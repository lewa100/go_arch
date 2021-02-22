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
	sync.Mutex
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

	done := make(chan bool, 1)
	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	var timer <-chan time.Time

	// Use flag -time
	if *fLimit == true {
		timer = time.After(time.Duration(*fCountLimit) * time.Second)
	}

	//flag count_flow
	for i := 0; i < *fCountFlows; i++ {
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle(done)
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
				done <- true
				close(jobChan)
				return
			case <-done:
				close(jobChan)
				return
			}
		}
	}()
	<-done
	wg.Wait()
	log.Printf("RPS: %s, AVG: %s", time.Duration(int(time.Since(startApp))/len(sliceTime)),
		time.Duration(sliceAvg()))
}

func (w *Worker) Handle(done chan bool) {
	defer w.wg.Done()
	for job := range w.jobChan {
		log.Printf("worker %d processing job with payload %s", w.num, string(job.payload))
		w.Lock()
		newDDosRequest(w, *job, done)
		w.Unlock()
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}

func newDDosRequest(w *Worker, job Job, done chan bool) {
	start := time.Now()
	resp, err := http.Get(*fUrl)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		timeTrack(start, resp, w, job, done)
	} else {
		log.Println("Argh! Broken")
	}
}

// timeTrack - func for print Log in realtime
func timeTrack(start time.Time, resp *http.Response, w *Worker, job Job, done chan bool) {
	elapsed := time.Since(start)
	sliceTime = append(sliceTime, elapsed)
	avg := sliceAvg()
	log.Printf("Flow id: %d Payload: %s HTTP Response Status: %d, %s, request time: %s", w.num, string(job.payload), resp.StatusCode, http.StatusText(resp.StatusCode), time.Duration(avg))
	//check status fLimit or fCountLimit
	if len(sliceTime) >= int(*fCountLimit) && *fLimit == false {
		log.Printf("Count limit: %d", len(sliceTime))
		done <- true
	}
}

func sliceAvg() int {
	var sum time.Duration
	for _, t := range sliceTime {
		sum += t
	}
	avg := int(sum) / len(sliceTime)
	return avg
}
