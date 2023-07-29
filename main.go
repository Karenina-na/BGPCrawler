package main

import (
	"BGP/src/Factory"
	"BGP/src/service"
	"BGP/src/util"
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

/*
~ Licensed to the Apache Software Foundation (ASF) under one or more
~ contributor license agreements.  See the NOTICE file distributed with
~ this work for additional information regarding copyright ownership.
~ The ASF licenses this file to You under the Apache License, Version 2.0
~ (the "License"); you may not use this file except in compliance with
~ the License.  You may obtain a copy of the License at
~
~     http://www.apache.org/licenses/LICENSE-2.0
~
~ Unless required by applicable law or agreed to in writing, software
~ distributed under the License is distributed on an "AS IS" BASIS,
~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
~ See the License for the specific language governing permissions and
~ limitations under the License.
*/

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
	// 启动下载服务
	service.DownLoad()
	// 启动转换服务
	service.Transfer()
	// quit
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	util.Loglevel(util.Info, "main", "BGP is exiting...")
	Factory.CloseFactory()
	runtime.GC()
	util.Loglevel(util.Info, "main", "BGP is exited")
	time.Sleep(time.Second * 3)

}
