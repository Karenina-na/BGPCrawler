package Factory

import (
	"BGPCrawler/src/exception"
	"BGPCrawler/src/mapper"
	"BGPCrawler/src/pool"
	"BGPCrawler/src/util"
)

// CloseFactory	关闭工厂
//
//	@Description: 关闭工厂
func CloseFactory() {
	err := pool.CloseRoutinePool()
	util.Loglevel(util.Debug, "CloseFactory", "关闭协程池")
	if err != nil {
		exception.HandleException(err)
	}

	err = mapper.CloseConnect()
	util.Loglevel(util.Debug, "CloseFactory", "关闭mongodb")
	if err != nil {
		exception.HandleException(err)
	}

	close(ServiceCloseChan)
}
