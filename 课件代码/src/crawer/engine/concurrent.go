package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)

	WorkerChan() chan Request
	Run()
	//ConfigureMasterWorkerChan(chan Request)

}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	//in := make( chan Request)
	out := make( chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i:=0;i<e.WorkerCount; i++{
		//createWorker(in,out)
		createWorker(e.Scheduler.WorkerChan() ,out,e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount :=0
	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v",itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier)  {
	//in := make(chan Request)
	go func() {
		for  {
			// tell scheduler I'm ready
			ready.WorkerReady(in)
			request := <-in
			result, e := worker(request)
			if e != nil {
				continue
			}
			out <-result
		}
	}()
	
}