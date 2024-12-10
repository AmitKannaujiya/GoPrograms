package gorutines

type Task func() (result interface{}, err error) 

type Result struct {
	WorkerId int
	result interface{}
	err error
}

type Worker struct {
	id          int
	taskQueue   <- chan Task  //receiver
	result 		chan <- Result
}

func(w *Worker) Start() {
	go func (){
		for task := range w.taskQueue {
			result, err := task()
			w.result <- Result{WorkerId:  w.id, result:  result, err: err}
		}
	}()
}
type WorkerPool struct {
	taskQueue		chan Task
	resultChan		chan Result
	workerCount		int
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskQueue: make(chan Task),
		resultChan:  make(chan Result),
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i:=0; i < wp.workerCount; i++ {
		worker := Worker{id: i, taskQueue: wp.taskQueue, result: wp.resultChan}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(task Task) {
	wp.taskQueue <- task
}

func (wp *WorkerPool) GetResult() Result {
	return <- wp.resultChan
}