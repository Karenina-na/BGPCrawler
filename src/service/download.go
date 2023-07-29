package service

import (
	"BGP/src/Factory"
	"BGP/src/config"
	"BGP/src/exception"
	"BGP/src/pool"
	"BGP/src/util"
	"time"
)

// DownLoad	下载服务
//
//	@Description: 下载服务
func DownLoad() {
	defer func() {
		r := recover()
		if r != nil {
			exception.HandleException(exception.NewSystemError("DownLoad", util.Strval(r)))
		}
	}()
	// 创建下载协程
	pool.RoutinePool.CreateWork(func() (E error) {
		defer func() {
			r := recover()
			if r != nil {
				E = exception.NewDownloadError("Persistence-service", util.Strval(r))
			}
		}()

		util.Loglevel(util.Info, "DownLoad", "download-service start")
		if err := callPython(); err != nil {
			return err
		}
		util.Loglevel(util.Info, "DownLoad", "download-service end")

		for {
			select {
			case <-Factory.ServiceCloseChan:
				util.Loglevel(util.Info, "DownLoad", "download-service exit")
				return nil
			case <-time.After(time.Hour * time.Duration(config.BGP.Frequency)): // TODO -----------------------------------
				util.Loglevel(util.Info, "DownLoad", "download-service start")
				if err := callPython(); err != nil {
					return err
				}
				util.Loglevel(util.Info, "DownLoad", "download-service end")
			}
		}

	}, func(message error) {
		exception.HandleException(message)
	})
}

// callPython
//
//	@Description: 调用python脚本
//	@return E	error
func callPython() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewDownloadError("callPython-service", util.Strval(r))
		}
	}()
	y, m, _ := time.Now().AddDate(0, 0, 0).Date()
	year := util.Strval(y)
	month := util.Strval(m)
	if len(month) == 1 {
		month = "0" + month
	}
	url := "https://archive.routeviews.org/bgpdata/" + year + "." + month + "/RIBS/"
	targetFolder := config.BGP.StoragePath
	days := config.BGP.StorageTime

	err := util.Run("python " + "./script/download.py" +
		" " + url +
		" " + targetFolder +
		" " + util.Strval(days))
	if err != nil {
		exception.HandleException(exception.NewDownloadError("callPython-service", err.Error()))
	}
	return nil
}
