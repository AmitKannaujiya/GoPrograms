package facade

import "fmt"

type memory struct{}
type hardDisk struct{}
type cpu struct{}

type computer struct {
	cpu      *cpu
	memory   *memory
	hardDisk *hardDisk
}

func (c *cpu) execute() {
	fmt.Println("cpu is executed")
}

func (c *cpu) jump() {
	fmt.Println("cpu is switched")
}

func (c *cpu) freeze() {
	fmt.Println("cpu is freeze")
}

func (m *memory) load() {
	fmt.Println("memory is loaded")
}

func (h *hardDisk) read() {
	fmt.Println("harddisk is readed")
}

func NewComputer()*computer {
	return &computer{
		cpu: &cpu{},
		memory: &memory{},
		hardDisk: &hardDisk{},
	}
}

func (c *computer) Start() bool {
	c.cpu.freeze()
	c.memory.load()
	c.cpu.jump()
	c.hardDisk.read()
	return true
}