package util

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Pool is a goroutine pool
type Pool struct {
	lock          sync.RWMutex       // 锁
	goChan        chan func()        // 任务队列
	coreNum       int                // 核心协程数
	maxNum        int                // 最大协程数
	activeNum     int                // 活跃协程数
	jobNum        int                // 任务数
	timeout       int                // 超时时间
	exceptionFunc func(r any)        // 异常处理
	ctx           context.Context    // 上下文
	cancelFunc    context.CancelFunc // 取消函数
}

// CreatePool
// @Description: 创建一个协程池
// @param        coreNum 核心协程数
// @param        maxNum  最大协程数
// @param        timeout 超时时间
// @return       *Pool   协程池
func CreatePool(coreNum int, maxNum int, timeout int) *Pool {
	ctx, cancelFunc := context.WithCancel(context.Background())
	P := &Pool{
		lock:       sync.RWMutex{},
		goChan:     make(chan func(), 5*maxNum),
		coreNum:    coreNum,
		maxNum:     maxNum,
		timeout:    timeout,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}

	// 初始化协程池核心协程
	for i := 0; i < coreNum; i++ {
		go P.work()
	}
	P.lock.Lock()
	P.activeNum = coreNum
	P.lock.Unlock()
	return P
}

// CheckStatus
// @Description: 检查协程池状态
// @receiver     P         协程池
// @return       activeNum 活跃协程数
// @return       jobNum    任务数
func (P *Pool) CheckStatus() (coreNum int, maxNum int, activeNum int, jobNum int) {
	P.lock.RLock()
	defer P.lock.RUnlock()
	return P.coreNum, P.maxNum, P.activeNum, P.jobNum
}

// CreateWork
// @Description: 创建一个任务
// @receiver     P             协程池
// @param        f             任务函数
// @param        exceptionFunc 异常处理函数
func (P *Pool) CreateWork(f func() (E error), exceptionFunc func(Message error)) {

	// 包装任务函数
	F := func() {
		if err := f(); err != nil {
			exceptionFunc(err)
			return
		}
	}

	// 阻塞等待任务队列有空闲
	select {
	case P.goChan <- F:
		P.lock.Lock()
		P.jobNum++
		P.lock.Unlock()
	case <-time.After(time.Duration(P.timeout) * time.Second):
		P.exceptionFunc(errors.New("goroutine队列溢出，超时"))
		return
	}

	// 动态创建协程
	P.lock.Lock()
	if P.activeNum < P.maxNum && P.jobNum > P.activeNum {
		P.activeNum++
		go P.work()
	}
	P.lock.Unlock()
}

// work
// @Description: 协程池工作函数
// @receiver     P 协程池
func (P *Pool) work() {
	defer func() {
		r := recover()
		if r != nil {
			P.exceptionFunc(r)
		}
	}()

	// 从任务队列中获取任务
	for {
		select {
		case <-P.ctx.Done():
			P.lock.Lock()
			P.activeNum--
			P.lock.Unlock()
			return
		case f := <-P.goChan:
			f()

			// 任务完成，任务数减一
			P.lock.Lock()
			P.jobNum--
			if P.activeNum > P.coreNum && P.jobNum < P.activeNum {
				P.activeNum--
				P.lock.Unlock()
				return
			}
			P.lock.Unlock()
		case <-time.After(time.Duration(P.timeout) * time.Second):
		}
	}
}

// SetExceptionFunc
// @Description: 设置异常处理函数
// @receiver     P 协程池
// @param        f 异常处理函数
func (P *Pool) SetExceptionFunc(f func(r any)) {
	P.exceptionFunc = f
}

// Close
// @Description: 关闭协程池
// @receiver     P 协程池
func (P *Pool) Close() {
	P.lock.Lock()
	P.maxNum = 0
	P.coreNum = 0
	P.cancelFunc()
	P.lock.Unlock()
	time.Sleep(time.Second * time.Duration(P.timeout))
}
