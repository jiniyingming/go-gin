local BasePlugin    = require "kong.plugins.base_plugin"
-- 调用模块
local redisUtil     = require "kong.plugins.custome-auth.redisUtil"
local httpPost_     = require "kong.plugins.custome-auth.httpPost"

local SumHandler    = BasePlugin:extend()
local json = require "cjson"
local log          = ngx.log
local ERR          = ngx.ERR


SumHandler.VERSION  = "1.0.0"
SumHandler.PRIORITY = 10

function SumHandler:access(config)

  local redisConfig = {}
  redisConfig.host = config['Host']
  redisConfig.port = config['Port']
  redisConfig.password = config['Password']
  redisConfig.database = config['Db']


  local redis         = redisUtil.new(redisConfig)

  local token_str = kong.request.get_headers()["token"] or nil
  if token_str then
    local auth_header_token = config['Prefix'] ..  token_str

--     log(ERR, "token" .. auth_header_token)

    -- 获取 token 缓存信息

    local res = getRedisStr(auth_header_token,redis)
    if res == nil then
      responseErr("get token data err")
    end


    local tokenJsonStr = res[1] or  nil
    local rst      = {}

    -- 用户信息解析
    if type(tokenJsonStr) == "string" then
      local decodeRet = json_decode(tokenJsonStr)

      if not decodeRet then
        responseErr("JsonString decode failed! token:: " .. tokenJsonStr)
      end

      if table.kIn(decodeRet, 'uid') == false then
        responseErr("User not found!")
      end
      ngx.header["j7_authed_uid"] = decodeRet['uid']

      if config['dyeingUrl'] ~= "" and config['dyeingVersion']~="" and config['dyeingPrefix']~="" then
            -- 需要重置 token string 信息
          if table.kIn(decodeRet, 'dyeingVersion') == false or decodeRet['dyeingVersion'] ~= config['dyeingVersion'] then
                -- 获取 dyeing 信息
                local needresetdyeing = 1
                local errored = 0
                local dyeingtoken = config['dyeingPrefix'] ..  token_str
                local dyeingStr = getRedisStr(dyeingtoken,redis)
                if dyeingStr ~= nil and type(dyeingStr) == "string" then
                    local decodeDyeingRet = json_decode(dyeingStr)
                    if decodeDyeingRet then
                        if table.kIn(decodeDyeingRet, 'version') ~= false and decodeDyeingRet['version']==config['dyeingVersion'] then
                            needresetdyeing = 0
                        end
                    end
                end

                local ttl = getRedisTTL( auth_header_token,redis )
                if needresetdyeing == 1 then
                      -- dyeing AB测试染色版本判断 --
                      local storeDyeing,errddd = dyeingUser("userId="..decodeRet['uid'].."&dyeingVersion="..config['dyeingVersion'] ,token_str,config)
                      if errddd == nil then
                            local str = json_encode(storeDyeing)
                            if str ~= nil then
                                local dottl = ttl
                                if dottl == nil or dottl <= 0 then
                                    dottl = tonumber(config['dyeingTTL'])
                                end
                                -- 保存 dyeingtoken
                                setRedis( dyeingtoken,str,redis, dottl )
                            else
                                errored = 1
                                log(ERR," ====== ERROR : dyeing json_encode ：" )
                            end
                      else
                        errored = 1
                        log(ERR," ====== ERROR : dyeingUser ：" .. errddd )
                      end
                end

                if errored == 0 then
                    -- 保存主键主信息...
                    decodeRet['dyeingVersion'] = config['dyeingVersion']
                    local mainstr = json_encode(decodeRet)
                    if mainstr ~= nil then
                        setRedis( auth_header_token,mainstr,redis,ttl )
                    end
                end

          end
      end


    end
  end
end

function responseErr (args)

  local rst      = {}
  rst["code"] = 400
  rst["msg"]     = args

  return kong.response.exit(400, rst,
    {
      ["Content-Type"] = "application/json",
    })

end



function dyeingUser(poststr,token,config)
  local errmsg = nil
  local storeDyeing = {}
  if config['dyeingUrl'] ~= "" and config['dyeingVersion']~="" and config['dyeingPrefix']~="" then
        local body,err= send_payload(config['dyeingUrl'],poststr)
        if err == nil then
            local decodeDyeingRet = json_decode(body)
              if not decodeDyeingRet then
                errmsg = "dyeing data decode error @ body:: " .. body
              else
                if table.kIn(decodeDyeingRet, 'code') ~= false and decodeDyeingRet['code']==200 and table.kIn(decodeDyeingRet, 'data') ~= false then
                    storeDyeing.data = decodeDyeingRet['data']
                    storeDyeing.version = config['dyeingVersion']
                    -- savedata
                else
                    errmsg = "response not 200 or empty data ：" .. body
                end
              end
        else
            errmsg = "error dyeing data @ err:",err
        end
  else
    errmsg = "empty config"
  end
  return storeDyeing,errmsg
end

function getRedisTTL(key,redis)
    local ttl, ttlerr = redis:exec(
      function(red)
        red:init_pipeline()
        red:ttl(key)
        return red:commit_pipeline()
      end
    )
    if ttlerr == nil then
        return ttl[1]
    end
    return nil
end

function getRedisStr(key,redis)
    local str, strerr = redis:exec(
      function(red)
        red:init_pipeline()
        red:get(key)
        return red:commit_pipeline()
      end
    )
    if strerr == nil then
        return str[1]
    end
    return nil
end

function setRedis(key,value,redis,expireseconds)
    local valueStr=""
    if type(value) ~= "string" then
        valueStr = json_encode(storeDyeing)
        if valueStr == nil then
            return nil,"cannot json_encode value"
        end
    else
        valueStr = value
    end
    return redis:exec(
      function(red)
        red:init_pipeline()
        red:set(key,valueStr)
        if expireseconds~=nil and expireseconds > 0 then
            red:expire(key,expireseconds)
        end
        return red:commit_pipeline()
      end
    )
end


local function _json_encode(str)
  return json.encode(str)
end

function json_encode( str )
    local ok, t = pcall(_json_encode, str)
    if not ok then
      return nil
    end
    return t
end

local function _json_decode(str)
  return json.decode(str)
end

function json_decode( str )
    local ok, t = pcall(_json_decode, str)
    if not ok then
      return nil
    end
    return t
end

function table.kIn(tbl, key)
    if tbl == nil then
        return false
    end
    for k, v in pairs(tbl) do
        if k == key then
            return true
        end
    end
    return false
end

function table.size(t)
    local s = 0;
    for k, v in pairs(t) do
        if v ~= nil then s = s + 1 end
    end
    return s;
end


return SumHandler

