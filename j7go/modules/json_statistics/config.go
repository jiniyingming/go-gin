package json_statistics

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/joselee214/j7f/components/config"
	"go.uber.org/zap"
	"j7go/components"
	"os"
)

var Opts *Options

//var LfSettings map[string]LimitflowProcess

type Options struct {
	Config          *config.Configer
	HotResetChan    chan int
	Mongodb         string `mapstructure:"mongodb"`
	MongoCollection string `mapstructure:"mongocollection"`
	IsSplit         int    `mapstructure:"issplit"`
}

func (o *Options) hotReset(e fsnotify.Event) {

	//fmt.Println("=================reset >>>>",e.Op,e.String())

	if e.Op == fsnotify.Write || e.Op == fsnotify.Create {
		o.HotResetChan <- 1
		err := o.Config.Unmarshal(o)
		if err != nil {
			components.L.Panic("faild unmarshal config", zap.Error(err))
		}

		if err != nil {
			components.L.Panic("faild init log config", zap.Error(err))
		}
		<-o.HotResetChan
	}
}

func init() {
	env := os.Getenv("RUNTIME_ENV")
	if env == "" {
		env = "default"
	}
	cfgFile := fmt.Sprintf("./conf/%s/jsonstatistics.yml", env)

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
	c.WatchConfig()
	c.OnConfigChange(Opts.hotReset)
}
