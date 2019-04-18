package scheduler

import "Demo01/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request

}

func (s *SimpleScheduler)WorkerChan()chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler)WorkerReady(chan engine.Request)  {

}

func (s *SimpleScheduler)ConfigureMasterWorkerChan(c chan engine.Request)  {
	//panic("implement me")
	s.workerChan = c
}

func (s *SimpleScheduler)run()  {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler)Submit(request engine.Request)  {
	go func() {
		s.workerChan <- request
	}()
}
