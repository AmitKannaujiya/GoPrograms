package gorutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// caller api to get details from multiple guys
// parallely start worker with no of jobs
// each worker is responsible for sending the
// response :
// first api will get results from each jobs
// club and sort them
// second api will get result in the form of array and then merged the arr to prepare big array

func executeTheJob(data [][]int) [][]int {
	var output [][]int
	size := len(data)
	jobs := make(chan int, size)
	result := make(chan []int, size)
	noOfWorkers := rand.Intn(size)

	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(data, jobs, result, &wg)
	}

	for i := 0; i < size; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < size; i++ {
		output = append(output, <-result)
	}
	close(result)
	wg.Wait()
	return output
}

func worker(data [][]int, jobs chan int, result chan []int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for job := range jobs {
		min, max := getMinMax(data[job])
		result <- []int{min, max}
	}
	//defer close(result)
}

func getMinMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}

func producer(ch chan <-int) {
	defer close(ch)
	for i:=0; i< 5 ; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
}

func consumer(ch <-chan int) {
	for d := range ch {
		fmt.Println(d)
	}
}

func ProducerConsumer() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}

func DoneChannelExample() {
	job := make(chan int, 5)
	done := make(chan int)
	//var wg sync.WaitGroup
	for i:=0; i < 5; i++ {
		go func(){
			workerJob(i, job, done)
		}()
	}
	for i:=0; i < 5; i++ {
		job <- i
	}
	close(job)

	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(500 * time.Millisecond)
}

func workerJob(id int, job chan int, done chan int) {
	for{
		select {
		case j, ok := <- job :
			if ok {
				fmt.Printf("worker : %d, job  :%d, started\n", id, j)
				return
			} else {
				fmt.Printf("All Job done")
			}
		case <- done :
			fmt.Println("Done signal received, exiting")
			return
		}
	}
}