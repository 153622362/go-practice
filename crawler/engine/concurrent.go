package engine

//并发版引擎结构
type ConcurrentEnine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

//调度器 它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
type Scheduler interface {
	ReadyNotifier
	Submit(request Request) //发送请求
	WorkerChan() chan Request //向调度器请求channel
	//ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

//多任务版并发爬虫
func (e *ConcurrentEnine) Run (seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run() //启动调度器

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) //生成worker
	}

	for _, r := range seeds {
		isfalse := isDuplicte(r.Url)
		if isfalse == false {
			e.Scheduler.Submit(r)
		}
	}

	//itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items {
			//log.Printf("Got item: #%d: %v", itemCount, item)
			//itemCount++

			go func() {
					e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			isfalse := isDuplicte(request.Url)
			if isfalse == false {
				e.Scheduler.Submit(request)
			}
		}
	}
}

func createWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifier)  {
		//in := make(chan Request)
		go func() {
			for  {
				// tell scheduler i'm ready
				ready.WorkerReady(in)
				request := <- in
				result, err := worker(request)
				if err != nil {
					continue
				}
				out <- result
			}
		}()
}

var visitedUrls = make(map[string]bool) //集合 string key_data_type 		bool value_data_type

func isDuplicte(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}