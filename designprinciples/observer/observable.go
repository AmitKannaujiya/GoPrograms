package observer

type IObservable interface {
	Add(IObserver)
	Remove(IObserver) bool
	Notify()
	GetData() int
	SetData(int)
}

type StockObserable struct {
	stock int
	observers []IObserver
}

func NewStockObservable() *StockObserable{
	return &StockObserable{
		stock: 0,
		observers: make([]IObserver, 0),
	}
}

func (s *StockObserable) Add(obj IObserver) {
	s.observers = append(s.observers, obj)
}

func (s StockObserable) GetData() int {
	return s.stock
}

func (s *StockObserable) SetData(stock int) {
	sendNotification := false
	if s.stock == 0 {
		sendNotification = true
	}
	s.stock = stock
	if sendNotification {
		s.Notify()
	}
}

func (s *StockObserable) Remove(obj IObserver) bool {
	index := -1
	for i, obj := range s.observers {
		if obj == obj {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}
	s.observers = append(s.observers[:index], s.observers[index+1:]...)
	return true
} 

func (s StockObserable) Notify() {
	for _, observer :=  range s.observers {
		observer.Update()
	}
}

