package gorutines

import (
	"fmt"
	"testing"

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
			min, max:=getMinMax(data[i])
			return []int{min,max}, nil
		})
	}
	//close(workerPool.taskQueue)

	for i := 0; i < len(data); i++ {
		result := workerPool.GetResult()
		fmt.Printf("worker id : %d, result : %v, error : %v\n",result.WorkerId, result.result, result.err)
	}
}

func TestProducerConsumer(t *testing.T) {
	ProducerConsumer()
}
