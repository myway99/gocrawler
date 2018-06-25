package engine

type ConcurrentEngine struct {
	Scheduler 	Scheduler
	WorkerCount int
	ItemChan  	chan  Item
	RequestProcessor Processor
}

// 定义Processor函数类型
type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkChan(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}


func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//e.createWorker(e.Scheduler.workerChan(), out, e.Scheduler)
		e.createWorker(e.Scheduler.WorkerChan(),
			out, e.Scheduler)
	}

	for _, r := range seeds {

		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: " +
			//	"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
				//log.Printf("Got item #%d: %v",
				//	itemCount, item)
				//itemCount++

				//go save(item)
				go func(i Item) {
					e.ItemChan <- i
				}(item)
		}

		// URL dedup
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request: " +
				//	"%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan  ParseResult, ready ReadyNotifier)  {
	//in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			//s.WorkerReady(in)
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(
				request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false

}
