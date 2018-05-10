package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

/*
	RequestHandler adalah tipe data yang merupakan
	fungsi yang menerima tipe interface sebagai
	parameternya dan return nothing
*/
type RequestHandler func(interface{})

// membuat Request baru
func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {

	return Request
				{
					Data: "Hello", 
						Handler: func(i interface{})
						{
							defer wg.Done()

							// casting tipe interface ke String	
							s, ok := i.(string)

							if !ok{
								log.Fatal("Invalid casting to string")
							}
							
							fmt.Println(s)
						}
				}
}

/*
	Handle Request

	Implement LaunchWorker yang parameternya
	adalah berupa Channel terkait Request yg
	nantinya menjadi titik masuk ke pipeline

	Tiap-tiap request diciptakan menjadi channel
	untuk dilempar ke pipeline yang akan berjalan
*/
type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

/*
	Handling Workers

	Eksekusi Tiap-Tiap Reqeust secara paralel dan
	handling semua channel yang akan dijalankan

	Implement tiap-tiap method.
*/
type Dispatcher interface {
	/*
		A Dispatcher interface can launch an injected WorkerLaunchers type in its own
		LaunchWorker method. The Dispatcher interface must use the LaunchWorker method
		of any of the WorkerLauncher types to initialize a pipeline. This way we can 
		reuse the Dispatcher interface to launch many types of WorkerLaunchers

		eksekusi WorkerLaunchers yang dikirimkan melalui parameter 
	*/
	LaunchWorker(w WorkerLauncher)

	// inject request baru ke WorkersPool
	MakeRequest(Request)

	// menghentikan ketika Go Routine telah selesai. 
	Stop()
}

/*
	Concrete Class Interface Dispacther
*/
type dispatcher struct {
	inCh chan Request
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh:make(chan Request, b),
	}
}

func (d *dispatcher) MakeRequest(r Request) {
	d.inCh <- r
}

func (d *dispatcher) LaunchWorker(id int, w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop(){
	select {
		case d.inCh <- r:
		case <-time.After(time.Second * 5):
			return
	}
}

/*
	The Pipeline


	That's all for workers! We simply pass the returning channels to the next steps in the
	Pipeline, as we did in the previous chapter. Remember that the pipeline is executed from
	inside to outside of the calls. So, what's the order of execution of any incoming data 
	to the pipeline?
	
	1. The data enters the pipeline through the Goroutine launched in the uppercase method.
	2. Then, it goes to the Goroutine launched in append.
	3. Finally, in enters the Goroutine launched in prefix method, which doesn't return 
	   anything but executes the handler after prefixing the incoming string with more data.

	Now we have a full pipeline and a dispatcher of pipelines. The dispatcher will launch as
	many instances of the pipelines as we want to route the incoming requests to any available
	worker.

	If none of the workers takes the request within 5 seconds, the request is lost.
	
	Let's use this library in a small app.
*/

// Concrete Class dari Interface WorkerLauncher
type PreffixSuffixWorker struct {
	id int
	prefixS string
	suffixS string
}

func (w *PreffixSuffixWorker) LaunchWorker(i int, in chan Request) {}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request 
{
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
				if !ok {
					msg.handler(nil)
					continue
				}
			
			msg.Data = strings.ToUpper(s)
			out <- msg
		}

		close(out)
	}()

	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	
	out := make(chan Request)

	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			// append Data in Incoming Channel
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}

		close(out)
	}()

	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			
			msg.handler(fmt.Sprintf("%s%s", w.prefixS, uppercasedStringWithSuffix))
		}
	}()
}

func main() {
	
	bufferSize := 100
	var dispatcher dDispatcher = NewDispatcher(bufferSize)

	workers := 3

	for i := 0; i < workers; i++ {

		var w WorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:i,
		}

		dispatcher.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	/*
		We will make 10 requests. We also need a WaitGroup to properly synchronize the app so
		that it doesn't exit too early. You can find yourself using WaitGroups quite a lot when
		dealing with concurrent applications. For 10 requests, we'll need to wait for 10 calls 
		to the Done() method, so we call the Add() method with a delta of 10. It's called delta 
		because you can also pass a -5 later to leave it in five requests
	*/

	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()
}
