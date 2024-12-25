package adapter

// client ---->  adapter ---> has a adaptee ---> adaptee.requireMethod
// example := getWeightByKg() --> adapterForGetWeightByKG <--- adpterGetWeightByKgIMP -->      adaptee.getWeightByPound()

// Adaptee :- WeightMachin
type WeightMachine struct {
	wg int
}

func NewWeightMachine() *WeightMachine {
	return &WeightMachine{
	}
}

func (wm *WeightMachine) SetWeight(weight int) {
	wm.wg = weight
}

func (wg *WeightMachine) GetWeightInPound() int {
	return wg.wg
}

// Interface WeightConvertor

type IWeightConvertor interface {
	GetWeightInKg() float32
	GetWeightInTonne() float32
	SetWeight(int)
}

// adapter Implementation

type WeightConvertorImp struct {
	weightMachine *WeightMachine
}

func NewWeightConvertorImp() *WeightConvertorImp {
	return &WeightConvertorImp{
		weightMachine: NewWeightMachine(),
	}
}

func (wci *WeightConvertorImp) GetWeightInKg() float32 {
	weight :=  wci.weightMachine.GetWeightInPound()
	return float32(weight) * 0.45
}

func (wci *WeightConvertorImp) GetWeightInTonne() float32 {
	weightInKG := wci.GetWeightInKg()
	return weightInKG / 1000
}

func (wci *WeightConvertorImp) SetWeight(weight int) {
	wci.weightMachine.SetWeight(weight)
}

// Client 
type CalculatorClient struct {
	weight int
	weightConvertorImp  *WeightConvertorImp 
}
func NewCalculatorClient(weight int) *CalculatorClient{
	calculatorClient := &CalculatorClient{
		weightConvertorImp: NewWeightConvertorImp(),
	}
	calculatorClient.weightConvertorImp.SetWeight(weight)
	return calculatorClient
}

func(cc *CalculatorClient) GetWeightInKg() float32 {
	return cc.weightConvertorImp.GetWeightInKg()
}
func(cc *CalculatorClient) GetWeightInTonne() float32 {
	return cc.weightConvertorImp.GetWeightInTonne()
}
