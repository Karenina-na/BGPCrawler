package pool

import (
	"BGPCrawler/src/config"
	"BGPCrawler/src/exception"
	"BGPCrawler/src/util"
)

// RoutinePool goroutine池
var RoutinePool *util.Pool

// InitRoutinePool
//
//	@Description: 初始化协程池
func InitRoutinePool() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("InitRoutinePool-pool", util.Strval(r))
		}
	}()
	//初始化协程池
	RoutinePool = util.CreatePool(config.Goroutine.CoreRoutineNum,
		config.Goroutine.MaxRoutineNum, config.Goroutine.RoutineTimeOut)
	RoutinePool.SetExceptionFunc(func(r any) {
		exception.HandleException(exception.NewSystemError("Pool池", util.Strval(r)))
	})
	return nil
}

// CloseRoutinePool
//
//	@Description: 关闭协程池
func CloseRoutinePool() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("CloseRoutinePool-pool", util.Strval(r))
		}
	}()
	RoutinePool.Close()
	return nil
}
