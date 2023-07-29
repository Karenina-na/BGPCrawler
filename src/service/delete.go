package service

import (
	"BGP/src/Factory"
	"BGP/src/config"
	"BGP/src/exception"
	"BGP/src/pool"
	"BGP/src/util"
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

		// 等待四小时后执行，错开下载和转换时间
		time.Sleep(time.Hour * 4)

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
				util.Run("./script/delete.sh" +
					" " + year + month + day +
					" " + config.BGP.StoragePath +
					" " + config.BGP.ProcessPath)
				println("Delete-service end")
			}
		}

	}, func(message error) {
		exception.HandleException(message)
	})
}
