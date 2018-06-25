package scheduler

import "mytest04/crawler/gocrawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//func (s *SimpleScheduler) ConfigureMasterWorkChan(
//	c chan engine.Request) {
//	s.workerChan = c
//}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	// 为每个Request创建goroutine,由goroutine往统一的channel分发
	//发送完毕即结束该goroutine
	go func() { s.workerChan <- r }()

}
