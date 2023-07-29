package main

import (
	"BGP/src/Factory"
	"BGP/src/util"
	"flag"
)

// main
//
//	@Description: 主函数
func main() {
	defer func() {
		err := recover()
		if err != nil {
			util.Loglevel(util.Error, "main", util.Strval(err))
		}
	}()
	arg := flag.String("mode", "debug", "debug / release /test 环境")
	flag.Parse()
	Factory.LoadConfigFactory(arg)

	Factory.CloseFactory()
}
