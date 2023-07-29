package service

import (
	"BGP/src/Factory"
	"BGP/src/config"
	"BGP/src/exception"
	"BGP/src/pool"
	"BGP/src/util"
	"time"
)

// Transfer	转换服务
//
//	@Description: 转换服务
func Transfer() {
	defer func() {
		r := recover()
		if r != nil {
			exception.HandleException(exception.NewSystemError("Transfer", util.Strval(r)))
		}
	}()
	// 创建转换协程
	pool.RoutinePool.CreateWork(func() (E error) {
		defer func() {
			r := recover()
			if r != nil {
				E = exception.NewDownloadError("Transfer-service", util.Strval(r))
			}
		}()

		// 等待二小时后执行，错开下载时间
		time.Sleep(time.Hour * 2)

		for {
			select {
			case <-Factory.ServiceCloseChan:
				util.Loglevel(util.Info, "Transfer", "download-service exit")
				return nil
			case <-time.After(time.Hour * time.Duration(config.BGP.Frequency)): // TODO -----------------------------------
				println("Transfer-service start")
				err := util.Run("./script/transfer.sh" +
					" " + config.BGP.StoragePath +
					" " + config.BGP.ProcessPath)
				if err != nil {
					exception.HandleException(exception.NewTransferError("Transfer", err.Error()))
				}
				println("Transfer-service end")
			}
		}

	}, func(message error) {
		exception.HandleException(message)
	})
}
