package process_limitflow

import (
	"github.com/fsnotify/fsnotify"
	"github.com/joselee214/j7f/components/config"
	"go.uber.org/zap"
	"j7go/components"
	"net"
	"os"
	"fmt"
	"time"
)

var Opts *Options

var LfSettings map[string]LimitflowProcess

type Options struct {
	Config           *config.Configer
	HotResetChan     chan int
	LimitflowAll     LimitflowAll              `mapstructure:"limitflowall"`
	Limitflowprocessdefault	LimitflowProcess	`mapstructure:"limitflowprocessdefault"`
	LimitflowProcess    []LimitflowProcess           `mapstructure:"limitflowprocess"`
}

type LimitflowAll struct {
	AppendCmdLog		string
	RuntimelessNanoSecond	int64
	RuntimelessCmdLog 	string
	Torunsave 		  	string
	LimitedFlow       	int
	CheckTime		  	int
}

type LimitflowProcess struct {
	TypeKey       string	`json:"TypeKey"`
	LimitedFlow     int		`json:"LimitedFlow"`
	RuningHeartBeat int		`json:"RuningHeartBeat"`
	StatusCallBack  string	`json:"StatusCallBack"`
	MaxTimeOut int			`json:"MaxTimeOut"`
}

func (o *Options) hotReset(e fsnotify.Event) {
	if e.Op == fsnotify.Write || e.Op == fsnotify.Create {
		o.HotResetChan <- 1
		err := o.Config.Unmarshal(o)
		if err != nil {
			components.L.Panic("faild unmarshal config", zap.Error(err))
		}

		if err != nil {
			components.L.Panic("faild init log config", zap.Error(err))
		}
		rebuildLfSettings()
		<- o.HotResetChan
	}
}

func init()  {
	env := os.Getenv("RUNTIME_ENV")
	if env == "" {
		env = "default"
	}
	cfgFile := fmt.Sprintf("./conf/%s/processlimitflow.yml", env)

	c, err := config.NewConfig()
	if err != nil {
		return
	}

	c.SetConfigFile(cfgFile)
	c.AutomaticEnv()

	err = c.ReadInConfig()
	if err != nil {
		components.L.Panic("faild init processlimitflow config", zap.Error(err))
		return
	}

	Opts = &Options{
		Config:       c,
		HotResetChan: make(chan int, 1),
	}

	err = c.Unmarshal(Opts)
	if err != nil {
		components.L.Panic("faild init processlimitflow config", zap.Error(err))
		return
	}

	rebuildLfSettings()
	c.WatchConfig()
	c.OnConfigChange(Opts.hotReset)
}

func rebuildLfSettings()  {
	LfSettings = make(map[string]LimitflowProcess)
	if len(Opts.LimitflowProcess)>0 {
		for _,v:= range Opts.LimitflowProcess {
			LfSettings[v.TypeKey] = v
		}
	}
}


func CallBackTimeoutDialer() func(net, addr string) (c net.Conn, err error) {

	cTimeout := 2 * time.Second
	rwTimeout := 500 * time.Millisecond

	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}