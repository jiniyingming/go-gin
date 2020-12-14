package process_limitflow

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joselee214/j7f/components/http/server"
	"go.uber.org/zap"
	"io"
	"io/ioutil"

	//"io/ioutil"
	"j7go/components"
	"j7go/utils"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
	"context"
)

const(
	StatusWait string = "wait"
	StatusAcceptWait string = "acceptwait"
	StatusErr string = "error"
	StatusStart string = "start"
	StatusRuning string = "runing"
	StatusTimeOut string = "timeout"
	StatusEnd string = "end"
	StatusKill string = "kill"
)

type ProcessLimitFlowController struct {
	server.Controller
}

type plfProcess struct {
	Cmd string	`json:"cmd"`
	Cmdid string	`json:"cmdid"`
	Cmdflagstr []string	`json:"cmdflagstr"`
	Cmdtype	string	`json:"cmdtype"`
	Cmdlfkey string	`json:"cmdlfkey"`
	Cmdoutput string	`json:"output"`
	Cmdpriority int `json:"riority"`
	Callback string `json:"callback"`
	Cmdlog *os.File
	CmdProcess *exec.Cmd
	StartTime int64	`json:"startTime"`
	HbreportTime uint	`json:"hbreportTime"`
	Uuid string	`json:"uuid"`
}

var torunCmds []*plfProcess
var runingCmds = make(map[string]*plfProcess)
var runingCmdsChannel = make(chan string)
var torunCmdsLock sync.RWMutex
var runingCmdsLock sync.RWMutex

var lfsettings = make(map[string]int)
var lfsettingsLock sync.RWMutex
var totalrunings = 0
var torunchan = make(chan int)


func Init(g *gin.Engine) {

	s := &ProcessLimitFlowController{}
	g.Routes()
	//g.GET("/favicon.ico",s.noop )//注册接口
	g.POST("/pm/newprocess",s.newprocess)
	g.Any("/pm/kill",s.kill)
	g.GET("/pm/stat",s.stat)
	g.GET("/pm/config",s.npconfig )//注册接口

	if Opts.LimitflowAll.Torunsave !="" {
		filePtr, err := os.Open(Opts.LimitflowAll.Torunsave)
		defer filePtr.Close()
		if err == nil {
			decoder := json.NewDecoder(filePtr)
			err = decoder.Decode(&torunCmds)
			components.L.Info("zzzz",zap.Any("zzz",torunCmds))
		}
	}
	go toDoAllCmds()
	go waitRuning()
	go checkRuning()
}


func (ctrl *ProcessLimitFlowController) noop(ctx *gin.Context)  {
}

func (ctrl *ProcessLimitFlowController) npconfig(ctx *gin.Context)  {
	var c struct{
		LimitflowAll     LimitflowAll	`json:"LimitflowAll"`
		Limitflowprocessdefault    LimitflowProcess	`json:"Limitflowprocessdefault"`
		LimitflowProcess    []LimitflowProcess	`json:"LimitflowProcess"`
	}

	c.LimitflowAll = Opts.LimitflowAll
	c.Limitflowprocessdefault = Opts.Limitflowprocessdefault
	c.LimitflowProcess = Opts.LimitflowProcess
	ctrl.Data = c
	ctrl.ResponseSuccess(ctx)
}

func (ctrl *ProcessLimitFlowController) stat(ctx *gin.Context)  {
	var ret = make(map[string]interface{})
	ret["to_runs"] = torunCmds
	ret["runings"] = runingCmds
	ret["totals"] = totalrunings
	ret["limitflows"] = lfsettings

	ctrl.Data = ret
	ctrl.ResponseSuccess(ctx)
}

func (ctrl *ProcessLimitFlowController) kill(ctx *gin.Context)  {
	uuid := ctx.PostForm("uuid")
	st := "no"
	if uuid!="" {
		//处理
		if len(torunCmds) > 0 {
			for idx,_ := range torunCmds {
				if torunCmds[idx].Uuid == uuid {
					torunCmdsLock.Lock()
					torunCmds = append(torunCmds[:idx], torunCmds[idx+1:]...)
					st = "remove"
					saveTorun()
					torunCmdsLock.Unlock()
					break
				}
			}
		}
		if len(runingCmds) > 0 {
			if _,ok := runingCmds[uuid]; ok {
				errk:=runingCmds[uuid].CmdProcess.Process.Kill()
				if errk==nil {
					go callBack(runingCmds[uuid],StatusKill)
					st = StatusKill
				}
			}
		}
	}
	ctrl.Data = st
	ctrl.ResponseSuccess(ctx)
}

func (ctrl *ProcessLimitFlowController) newprocess(ctx *gin.Context)  {
	if vv,ok := ctx.Get("json-param"); ok{
		body, err := ctx.GetRawData()
		if err == nil{
			var c plfProcess
			err := json.Unmarshal(body,&c)
			if err==nil{

				c.Cmd = replaceStr(c.Cmd)
				c.Cmdlfkey = replaceStr(c.Cmdlfkey)
				c.Cmdtype = replaceStr(c.Cmdtype)

				if c.Cmd != "" && len(c.Cmdflagstr)>0 {

					if c.Cmdlfkey == "" {
						c.Cmdlfkey = c.Cmdtype
					}
					if c.Cmdlfkey == "" {
						c.Cmdlfkey = c.Cmd
					}
					if c.Cmdtype=="" {
						c.Cmdtype = c.Cmdlfkey
					}
					if c.Cmdtype == "" {
						c.Cmdtype = c.Cmd
					}

					if c.Cmdoutput == "" {
						c.Cmdoutput = "/dev/null"
					}
					ctrl.Data = vv


					uuid,errr:=uuid.NewUUID()
					if errr == nil {
						c.Uuid = uuid.String()
					}

					torunCmdsLock.Lock()
					if c.Cmdpriority == 1 {
						torunCmds = append([]*plfProcess{&c},torunCmds...)
					} else {
						torunCmds = append(torunCmds,&c)
					}


					if Opts.LimitflowAll.AppendCmdLog != "" {
						filebytes, errf := json.MarshalIndent(c, "", "\t")
						if errf == nil {
							WriteAppendFile(Opts.LimitflowAll.AppendCmdLog, filebytes, 0666)
						}
					}

					fmt.Println("StatusAcceptWait")
					go callBack(&c,StatusAcceptWait)
					saveTorun()
					torunCmdsLock.Unlock()

					ctrl.ResponseSuccess(ctx)
					return

				} else {
					ctrl.Data = fmt.Sprintf("empty CMD or error flags !")
				}
			}
		}

	} else {
		ctrl.Data = fmt.Sprintf("nothing input")
	}
	ctrl.ResponseError(ctx,nil)
}

//判断是否限流
func checkLimitFlow(c *plfProcess) (bool) {
	lf := getLFsetting(c)
	if lf.LimitedFlow==0 {
		return true
	}

	if vv,ok := lfsettings[c.Cmdlfkey]; ok {
		//存在//
		if vv < lf.LimitedFlow {
			if (Opts.LimitflowAll.LimitedFlow >0) && (totalrunings < Opts.LimitflowAll.LimitedFlow) {
				return true
			}
		}
	} else {
		if (Opts.LimitflowAll.LimitedFlow >0) && (totalrunings < Opts.LimitflowAll.LimitedFlow) {
			return true
		}
	}
	return false
}

func addLimitFlow(c *plfProcess,incr int)  {
	lfsettingsLock.Lock()
	if _,ok := lfsettings[c.Cmdlfkey]; ok {
		//存在//
		lfsettings[c.Cmdlfkey] = lfsettings[c.Cmdlfkey] + incr
	} else {
		lfsettings[c.Cmdlfkey] = incr
	}
	if lfsettings[c.Cmdlfkey] <= 0 {
		delete(lfsettings,c.Cmdlfkey)
	}
	totalrunings = totalrunings + incr
	lfsettingsLock.Unlock()
}

func replaceStr(str string) string  {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return strings.Replace(str, "\t", "", -1)
}

//循环处理所有待执行的..数据
func toDoAllCmds() {
	for {
		select {
			case <- torunchan:
				//components.L.Info("toDoAllCommands from chan ")
				toDoAllCommands()
			case <- time.After( time.Duration(Opts.LimitflowAll.CheckTime) * time.Second ):
				//components.L.Info("toDoAllCommands from timeAfter ")
				toDoAllCommands()
		}
	}
}

func toDoAllCommands() {
	defer func() {
		if r := recover(); r != nil {
			components.L.Info("toDoAllCommands exception ", zap.Any("recover",r) )
		}
	}()

	if len(torunCmds) > 0 {
		var toRunNow = false
		//var toRunNowLF []LimitflowProcess
		var leftRunCmds []*plfProcess
		for idx, _ := range torunCmds {
			if ok := checkLimitFlow(torunCmds[idx]); ok {
				toRunNow = true
				runCmds(torunCmds[idx]) //新增执行
				//toRunNowLF = append(toRunNowLF,lmf)
			} else {
				leftRunCmds = append(leftRunCmds, torunCmds[idx])
			}
		}
		//如果判断有
		if toRunNow {
			torunCmdsLock.Lock()
			torunCmds = leftRunCmds //替换旧的
			saveTorun()
			torunCmdsLock.Unlock()
		}
	}
}

//检查正在执行的//
func checkRuning()  {
	for {
		time.Sleep( 1 * time.Second)
		checkRuningAndHb()
		checkWaitAndHb()
	}
}


func checkWaitAndHb(){
	if len(torunCmds) > 0 {
		torunCmdsLock.Lock()
		for idx, _ := range torunCmds {
			heartBeat(torunCmds[idx], StatusWait)
		}
		torunCmdsLock.Unlock()
	}
}

func checkRuningAndHb(){
	defer func() {
		if r := recover(); r != nil {
			components.L.Info("checkRuningAndHb exception ", zap.Any("recover",r) )
		}
	}()
	if len(runingCmds) > 0 {
		runingCmdsLock.Lock()
		for idx, _ := range runingCmds {
			//fmt.Println("checkRuning...runCmds...",idx, runingCmds[idx].CmdProcess )
			heartBeat(runingCmds[idx],StatusRuning)
		}
		runingCmdsLock.Unlock()
	}
}


func heartBeat(c *plfProcess,st string)  {
	lf := getLFsetting(c)

	//components.L.Info("heartBeat  " + st , zap.Any("lf.RuningHeartBeat",lf.RuningHeartBeat) )
	//components.L.Info("heartBeat  " + st , zap.Any("lf.RuningHeartBeat",c.HbreportTime) )

	if lf.RuningHeartBeat > 0 {
		timenow := utils.GetCurrentUnixTime()
		if  int(timenow - c.HbreportTime) > lf.RuningHeartBeat {
			c.HbreportTime = timenow
			go callBack(c,st)
		}
	}
}


//循环判断任务完成状态
func waitRuning(){
	for uukey := range runingCmdsChannel {
		endRunCmd(uukey)
	}
}

func endRunCmd(uukey string)  {
	defer func() {
		if r := recover(); r != nil {
			components.L.Info("endRunCmd exception ", zap.Any("recover",r) )
		}
	}()
	if _,ok := runingCmds[uukey]; ok {
		addLimitFlow(runingCmds[uukey],-1)

		runingCmdsLock.Lock()
		delete(runingCmds,uukey)
		runingCmdsLock.Unlock()

		torunchan <- 1
	}
}


//等待任务完成回调
func runCmdsWait(uukey string, c *plfProcess){

	defer func() {
		if r := recover(); r != nil {
			components.L.Info("runCmdsWait exception ", zap.Any("recover",r) )
		}
	}()

	c.CmdProcess.Wait()
	runingCmdsChannel <- uukey
	if c.Cmdoutput != "Stdout" {
		defer c.Cmdlog.Close()
	}

	if Opts.LimitflowAll.RuntimelessCmdLog != "" {
		nowt := utils.GetCurrentUnixNanoTime()
		if nowt - c.StartTime > Opts.LimitflowAll.RuntimelessNanoSecond {
			filebytes,errf := json.MarshalIndent(c,"","\t")
			if errf==nil {
				WriteAppendFile( Opts.LimitflowAll.RuntimelessCmdLog ,filebytes,0666)
			}
		}
	}

	go callBack(c,StatusEnd)
}

//批量执行任务
func runCmds(cs *plfProcess)  {
		runingCmdsLock.Lock()
		if _,ok := execProcess(cs); ok==nil{
			//fmt.Println("...runCmds..." )
			if cs.Uuid != "" {
				addLimitFlow(cs,1)
				go runCmdsWait(cs.Uuid,cs)
				runingCmds[cs.Uuid] = cs
			}
			go callBack(cs,StatusStart)
		} else {
			go callBack(cs,StatusErr)
		}
		runingCmdsLock.Unlock()
}




//进程状态通知
func callBack(c *plfProcess,st string){

	defer func() {
		if r := recover(); r != nil {
			components.L.Info("callback exception ", zap.Any("recover",r) )
		}
	}()

	callbackurl := c.Callback
	if c.Callback == "default" {
		lf := getLFsetting(c)
		callbackurl = lf.StatusCallBack
	}

	if callbackurl != "" {
		data := url.Values{}
		data.Set("cmd",c.Cmd)
		data.Set("cmdid",c.Cmdid)
		if len(c.Cmdflagstr) > 0 {
			for _,v := range c.Cmdflagstr{
				data.Add("cmdflagstr",v)
			}
		}
		data.Set("cmdtype",c.Cmdtype)
		data.Set("cmdlfkey",c.Cmdlfkey)
		data.Set("output",c.Cmdoutput)
		data.Set("uuid",c.Uuid)
		data.Set("riority",string(c.Cmdpriority))
		data.Set("statusstr",st)

		cc := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				Dial:            CallBackTimeoutDialer(),
			},
		}

		req,err := http.NewRequest(http.MethodPost , callbackurl , strings.NewReader(data.Encode()))
		if err==nil {
			req.Header = http.Header{}
			req.Header["Connection"] = []string{"Close"}
			req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=UTF-8"}
			req.Header["User-Agent"] = []string{"j7-http-client/1.1"}
			resp,err := cc.Do(req)
			if err == nil {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				components.L.Info("callback " + st, zap.String("response",string(body)) )
			}
		}
	}
}

//启动执行任务
func execProcess(c *plfProcess) (int,error) {
	lf := getLFsetting(c)
	var cmd *exec.Cmd
	if lf.MaxTimeOut > 0 {
		ctx, cancelFunc := context.WithCancel(context.Background())
		cmd = exec.CommandContext(ctx, c.Cmd,c.Cmdflagstr...)
		time.AfterFunc( time.Duration(lf.MaxTimeOut) * time.Second , func() {
			go callBack(c,StatusTimeOut)
			cancelFunc()
		})
	} else {
		cmd = exec.Command(c.Cmd,c.Cmdflagstr...)
	}


	if c.Cmdoutput == "Stdout" {
		c.Cmdlog = os.Stdout
	} else {
		file, err := os.OpenFile( c.Cmdoutput , os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			components.L.Info( "output file error ,use Stdout instead :",zap.Any("plfProcess",c))
			c.Cmdlog = os.Stdout
			c.Cmdoutput = "Stdout"
		} else {
			c.Cmdlog = file
		}
	}

	cmd.Stdout = c.Cmdlog // os.Stdout
	cmd.Stderr = c.Cmdlog // os.Stderr
	//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid:true}
	//fmt.Println("start",c.Cmd,c.Cmdflagstr)
	err := cmd.Start()
	c.StartTime = utils.GetCurrentUnixNanoTime()
	c.CmdProcess = cmd

	components.L.Info("start process",zap.Any("plfProcess",c))

	//fmt.Println("after start", &cmd )
	//fmt.Println("after start", c.startTime ," || pid:",cmd.Process )

	if err==nil {
		//fmt.Println("Wait",c.Cmd,c.Cmdflagstr,"pid:",cmd.Process )
		return cmd.Process.Pid,nil
	}
	return 0,err
}

//获取对应的配置
func getLFsetting(c *plfProcess) LimitflowProcess {
	var tp string
	tp = c.Cmdtype
	if tp == "" {
		tp = c.Cmdlfkey
	}
	if tp == "" {
		tp = c.Cmd
	}
	if _,ok := LfSettings[tp]; ok{
		return LfSettings[tp]
	}
	return Opts.Limitflowprocessdefault
}


func saveTorun() {
	if Opts.LimitflowAll.Torunsave !="" {
		filebytes,errf := json.MarshalIndent(torunCmds,"","\t")
		if errf==nil {
			errw := WriteFile( Opts.LimitflowAll.Torunsave ,filebytes,0666)
			if errw!=nil{
				components.L.Info("write torun File error ", zap.String("file",Opts.LimitflowAll.Torunsave) , zap.Error(errw))
			}
		}
	}
}


func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func WriteAppendFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}