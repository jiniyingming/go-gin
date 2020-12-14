local typedefs = require "kong.db.schema.typedefs"

return {
    name = "custome-auth",
    fields = {
        { consumer = typedefs.no_consumer },
        { config = {
            type = "record",
            fields = {
                { Prefix = { type = "string", required = true, default = "fx:api:token:" }, },
                { Host = { type = "string", required = true, default = "172.20.0.4" }, },
                { Port = { type = "string", default = "6379" }, },
                { Password = { type = "string", required = false, }, },
                { Db = { type = "string", required = false }, },
                { dyeingUrl = { type = "string", required = false }, },
                { dyeingVersion = { type = "string", required = true, default = "default" }, },
                { dyeingPrefix = { type = "string", required = true, default = "dyeing:" }, },
                { dyeingTTL = { type = "string", required = true, default = "864000" }, },
            }, }, },
    },
}
