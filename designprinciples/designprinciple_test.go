package designprinciples

import (
	"testing"

	d "go-program/designprinciples/deccorator"
	f "go-program/designprinciples/factory"
	o "go-program/designprinciples/observer"
	s "go-program/designprinciples/strategy"
	c "go-program/designprinciples/chainofresponsibility"
	co "go-program/designprinciples/composite"

	"github.com/stretchr/testify/assert"
)

func TestStretegyPattern1(t *testing.T) {
	normalStrategy := s.NormalDriveStrategy{}
	vehicle := s.Vehicle{
		Ty : "Normal",
		IDS : normalStrategy,
	}

	vehicle.Drive()
	assert.Equal(t, "Normal", vehicle.GetType())
	sportsDriveStretegy := s.SportDriveStrategy{}
	vehicle.IDS = sportsDriveStretegy
	vehicle.Ty = "Sport"

	sports := s.SportVehicle{
		Vehicle: vehicle,
	}
	sports.Vehicle.Drive()
	assert.Equal(t, "Sport", vehicle.GetType())

}

func TestObserverPattern1(t *testing.T) {
	stockObserver := o.NewStockObservable()
	emailNotification1 := o.NewEmailNotificationObserver("abc@xyz.com", stockObserver)
	emailNotification2 := o.NewEmailNotificationObserver("xyzwwe@xyz.com", stockObserver)
	emailNotification3 := o.NewEmailNotificationObserver("ktyop@xyz.com", stockObserver)
	emailNotification4 := o.NewEmailNotificationObserver("sdsfdfdf@xyz.com", stockObserver)
	smsNotification1 := o.NewSMSNotificationObserver("9898989898", stockObserver)
	smsNotification2 := o.NewSMSNotificationObserver("787382788", stockObserver)

	stockObserver.Add(emailNotification1)
	stockObserver.Add(emailNotification2)
	stockObserver.Add(emailNotification3)
	stockObserver.Add(emailNotification4)
	stockObserver.Add(smsNotification1)
	stockObserver.Add(smsNotification2)

	stockObserver.SetData(10)
	assert.Equal(t, 10, stockObserver.GetData())
	stockObserver.SetData(20)
	assert.Equal(t, 20, stockObserver.GetData())
	stockObserver.SetData(0)
	assert.Equal(t, 0, stockObserver.GetData())
	stockObserver.SetData(30)
	assert.Equal(t, 30, stockObserver.GetData())
}

func TestDecoratorPattern1(t *testing.T) {
	farmouse := d.NewFarmHouse("farmhouse")
	assert.Equal(t, 100, farmouse.Cost())

	margerita := d.NewMargrita("margarita")
	assert.Equal(t, 120, margerita.Cost())

	topping := d.Topping{
		Pizza: farmouse,
	}
	extraCheeze := d.CheezeTopping{
		Topping: topping,
	}
	assert.Equal(t, 110, extraCheeze.Cost())
	extraCheeze2 := d.CheezeTopping{
		Topping: d.Topping{
			Pizza: margerita,
		},
	}
	assert.Equal(t, 130, extraCheeze2.Cost())
}

func TestFactoryPattern1(t *testing.T) {
	factory := f.ShapeFactory{}
	shape := factory.GetShape("circle")
	shape.Draw()

	assert.Equal(t, "square", factory.GetShape("square").GetName())
	assert.Equal(t, "circle", factory.GetShape("circle").GetName())
}

func TestAbstractFactoryPattern1(t *testing.T) {
	abstractFactory := f.LuxuryVehicleFactory{}
	abstractFactory.GetVehicle("bmw").Drive()

	assert.Equal(t, "bmw", abstractFactory.GetVehicle("bmw").GetName())
	abstractFactory2 := f.OrdinaryVehicleFactory{}
	assert.Equal(t, "swift", abstractFactory2.GetVehicle("swift").GetName())
	abstractFactory2.GetVehicle("hundai").Drive()

}

func TestChainOfResponsibilityPattern1(t *testing.T) {
	logger := c.NewDebugLogger(c.NewInfoLogger(c.NewErrorLogger(&c.NOLogger{})))
	logger.Log(c.DEBUG, "I am debugging it")
	logger.Log(c.INFO, "I am printing it")
	logger.Log(c.ERROR, "I am Error printing it")
	logger.Log(5, "This should not be printing")
	assert.Equal(t, 1, int(c.DEBUG))
	assert.Equal(t, 2, int(c.INFO))
	assert.Equal(t, 3, int(c.ERROR))
}

func TestComposite(t *testing.T) {
	circle1 := co.Circle{Name: "Circle1"}
	circle1.Draw()
	rectangle1 := co.Rectangle{Name: "rectangle1"}
	rectangle1.Draw()
	circle2 := co.Circle{Name: "Circle2"}
	rectangle2 := co.Rectangle{Name: "rectangle2"}
	circle3 := co.Circle{Name: "Circle3"}
	rectangle3 := co.Rectangle{Name: "rectangle3"}
	circle4 := co.Circle{Name: "Circle4"}
	composite := co.Composite{Name: "Composit group"}
	composite.AddChild(&circle2)
	composite.AddChild(&rectangle2)
	composite.AddChild(&circle3)
	composite.AddChild(&rectangle3)
	composite.AddChild(&circle4)
	composite.Draw()
	assert.Equal(t, 5, composite.GetChildCount())
}