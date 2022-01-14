package main

import (
	"flag"
	"fmt"
	"github.com/BaizeSec/Pusher/common"
	"github.com/BaizeSec/Pusher/crons"
	"github.com/BaizeSec/Pusher/lib"
	"github.com/BaizeSec/Pusher/pkg/logger"
	"github.com/robfig/cron"
)

func main() {
	flag.Parse()
	lib.Banner()
	lib.InitConfig()
	common.InitValues()
	c := cron.New()

	//crons.Daily()

	if lib.TestFlag {
		crons.Test()
	}

	//err := c.AddFunc("* * * * 1 ?", crons.Daily)
	logger.Info("Waiting for task scheduling..")
	err := c.AddFunc("0 30 9 * * ?", crons.Daily)

	//crons.Daily()

	if err != nil {
		fmt.Println(err)
	}

	c.Run()
}
