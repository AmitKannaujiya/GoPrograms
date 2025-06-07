package gorutines

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMinMaxOutput(t *testing.T) {
	data := [][]int{
		{9, 6, 19, 4, 90, 67, 0, 2}, {15, 17, 18, 4, 1, 78}, {89, 1, 7, 18, 67, 189, 89, 4, 56}, {89, 19, 17, 189, 56, 901, 90},
	}
	result := executeTheJob(data)
	fmt.Println(result)
	assert.Equal(t, []int{0, 90}, result[0])
}

func TestWorkerPool(t *testing.T) {
	data := [][]int{
		{9, 6, 19, 4, 90, 67, 0, 2}, {15, 17, 18, 4, 1, 78}, {89, 1, 7, 18, 67, 189, 89, 4, 56}, {89, 19, 17, 189, 56, 901, 90},
	}
	workerPool := NewWorkerPool(len(data))
	workerPool.Start()
	for i := 0; i < len(data); i++ {
		workerPool.Submit(func() (result interface{}, err error) {
			min, max := getMinMax(data[i])
			return []int{min, max}, nil
		})
	}
	//close(workerPool.taskQueue)

	for i := 0; i < len(data); i++ {
		result := workerPool.GetResult()
		fmt.Printf("worker id : %d, result : %v, error : %v\n", result.WorkerId, result.result, result.err)
	}
}

func TestProducerConsumer(t *testing.T) {
	ProducerConsumer()
}

func TestDoneChannelExample1(t *testing.T) {
	DoneChannelExample()
}

func Test_UrlTime(t *testing.T) {
	start := time.Now()

	urls := []string{
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
	}

	MultiURLTime(urls)

	duration := time.Since(start)
	log.Printf("%d URLs in %v", len(urls), duration)
}


func Test_ValidateSign(t *testing.T) {
	start := time.Now()
	files := []file {
		file{Signature: "dasdasassdafsadfasdfsa",
		Content: []byte("ssdsdsdfsfdsdsfdsfds"),
		Name: "abc.txt",
	},
	file{Signature: "sdcsdcdsccsdcsdvsdvsv",
		Content: []byte("dsvdsvsd"),
		Name: "bcd.txt",
	},
	file{Signature: "ascsadcwfrwfwdsc",
		Content: []byte("scdcdssdvsdv"),
		Name: "sdf.txt",
	},
	}
	ok, bad, _:= ValidateSigs(files)

	duration := time.Since(start)
	log.Printf("info: %d files in %v\n", len(ok)+len(bad), duration)
	log.Printf("ok: %v", ok)
	log.Printf("bad: %v", bad)
}

func Test_ContextTimeout(t *testing.T) {
	log.Printf("info: checking finish in time")
	ctx, cancel := context.WithTimeout(context.Background(), bmvTime*2)
	defer cancel()

	mOK := NextMovie(ctx, "ridley")
	log.Printf("info: got %+v", mOK)

	log.Printf("info: checking timeout")
	ctx, cancel = context.WithTimeout(context.Background(), bmvTime/2)
	defer cancel()

	mTimeout := NextMovie(ctx, "ridley")
	log.Printf("info: got %+v", mTimeout)
}

func Test_LimitGoroutine(t *testing.T) {
	start :=  time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	n :=  runtime.GOMAXPROCS(0)
	src :=[]string{"C:\\01Coding\\06Resume\\2024-10-28", "C:\\01Coding\\06Resume\\2024-12-24", "C:\\Users\\\amit_\\Downloads\\fwddocuments2"}
	err :=  CenterDir(ctx, src, "C:\\02Bills", n)

	duration := time.Since(start)
	log.Printf("info: finished in %v (err=%v)", duration, err)
}