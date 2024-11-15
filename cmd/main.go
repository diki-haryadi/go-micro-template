package main

import (
	"fmt"
	"github.com/diki-haryadi/go-micro-template/app"
	"github.com/diki-haryadi/ztools/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		fmt.Println(err)
		logger.Zap.Sugar().Fatal(err)
	}
}
