package service

import (
	"BGPCrawler/src/Factory"
	"BGPCrawler/src/config"
	"BGPCrawler/src/exception"
	"BGPCrawler/src/pool"
	"BGPCrawler/src/util"
	"time"
)

// Delete	转换服务
//
//	@Description: 转换服务
func Delete() {
	defer func() {
		r := recover()
		if r != nil {
			exception.HandleException(exception.NewSystemError("Delete", util.Strval(r)))
		}
	}()
	// 创建转换协程
	pool.RoutinePool.CreateWork(func() (E error) {
		defer func() {
			r := recover()
			if r != nil {
				E = exception.NewDownloadError("Delete-service", util.Strval(r))
			}
		}()

		// 等待一小时后执行，错开下载和转换时间
		time.Sleep(time.Minute * 20)

		for {
			select {
			case <-Factory.ServiceCloseChan:
				util.Loglevel(util.Info, "Delete", "download-service exit")
				return nil
			case <-time.After(time.Hour * time.Duration(config.BGP.Frequency)): // TODO -----------------------------------
				println("Delete-service start")
				y, m, d := time.Now().AddDate(0, 0, 0).Date()
				year := util.Strval(y)
				month := util.Strval(m)
				if len(month) == 1 {
					month = "0" + month
				}
				day := util.Strval(d)
				err := util.Run("./script/delete.sh" +
					" " + year + month + day +
					" " + config.BGP.StoragePath +
					" " + config.BGP.ProcessPath)
				if err != nil {
					exception.HandleException(exception.NewDeleteError("Delete-service", err.Error()))
				}
				println("Delete-service end")
			}
		}

	}, func(message error) {
		exception.HandleException(message)
	})
}
