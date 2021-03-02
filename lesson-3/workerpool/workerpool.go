package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	payload []byte
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

var (
	fUrl        *string
	fCountFlows *int
	fCountLimit *int64
	fLimit      *bool
	sliceTime   []time.Duration
)

func initFlag() {
	fUrl = flag.String("url", "http://localhost:8081/item", "Url for DDoS")
	fCountFlows = flag.Int("count_flow", 1, "Count flows")
	fLimit = flag.Bool("time", false, "False - Count Limit, True - time limit")
	fCountLimit = flag.Int64("weight", 0, "Count Limit / Time limit in seconds")

	flag.Parse()

	//fmt.Println("u:", *fUrl)
	//fmt.Println("c:", *fCountFlows)
	//fmt.Println("f:", *fLimit)
	//fmt.Println("w:", *fCountLimit)
	//fmt.Println("tail:", flag.Args())

	if *fCountLimit == 0 || len(flag.Args()) != 0 {
		log.Fatalf("Flags error")
	}
}

func main() {
	startApp := time.Now()
	initFlag()

	timer := make(<-chan time.Time)
	done := make(chan bool, 1)
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	jobChan := make(chan *Job)

	// Use flag -time
	//if *fLimit == true {
	timer = time.After(time.Duration(*fCountLimit) * time.Second)
	//}

	//flag count_flow
	for i := 0; i < *fCountFlows; i++ {
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle(done, m)
	}

	go func() {
		for i := 0; ; i++ {
			select {
			default:
				jobChan <- &Job{
					payload: []byte(fmt.Sprintf("Job %d", i)),
				}
			case <-timer:
				log.Printf("Time out %s", time.Duration(*fCountLimit)*time.Second)
				close(jobChan)
				done <- true
				return
			case <-done:
				//close(jobChan)
				return
			}
		}
	}()
	<-done
	wg.Wait()
	log.Printf("RPS: %s, AVG: %s", time.Duration(int(time.Since(startApp))/len(sliceTime)),
		time.Duration(sliceAvg()))
}

func (w *Worker) Handle(done chan bool, m *sync.Mutex) {
	defer w.wg.Done()
	for job := range w.jobChan {
		start := time.Now()
		resp, err := http.Get(*fUrl)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			w.timeTrack(start, resp, job, done, m)
		} else {
			log.Println("Argh! Broken")
		}
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}

// timeTrack - func for print Log in realtime
func (w *Worker) timeTrack(start time.Time, resp *http.Response, job *Job, done chan bool, m *sync.Mutex) {
	elapsed := time.Since(start)
	m.Lock()
	sliceTime = append(sliceTime, elapsed)
	avg := sliceAvg()
	m.Unlock()
	log.Printf("Flow id: %d Payload: %s HTTP Response Status: %d, %s, request time: %s", w.num, string(job.payload), resp.StatusCode, http.StatusText(resp.StatusCode), time.Duration(avg))
	//check status fLimit or fCountLimit
	//if len(sliceTime) >= int(*fCountLimit) && *fLimit == false {
	//	log.Printf("Count limit: %d", len(sliceTime))
	//	done <- true
	//}
}

func sliceAvg() int {
	var sum time.Duration
	for _, t := range sliceTime {
		sum += t
	}
	avg := int(sum) / len(sliceTime)
	return avg
}

//func main() {
//	//startApp := time.Now()
//	initFlag()
//	ctx, cancel := context.WithCancel(context.Background())
//	stopped := make(chan struct{})
//
//	go run(ctx, stopped)
//
//	cancel()
//	<-stopped
//	log.Warn("Stopped")
//}
//
//func run(ctx context.Context, stopped chan struct{}) {
//
//	waiter := new(sync.WaitGroup)
//	waiter.Add(1)
//	go sniffer.startFlow(ctx, waiter)
//
//	waiter.Wait()
//	stopped <- struct{}{}
//}
