package observer

import "fmt"

type IObserver interface {
	Update()
}

type EmailNotificationObserver struct {
	email       string
	objservable IObservable
}

func NewEmailNotificationObserver(email string, stockObserver IObservable) *EmailNotificationObserver {
	return &EmailNotificationObserver{
		email:       email,
		objservable: stockObserver,
	}
}

func (e EmailNotificationObserver) Update() {
	e.SendEmail()
}

func (e EmailNotificationObserver) SendEmail() {
	fmt.Printf("Email on id %s, sent for observable %d\n", e.email, e.objservable.GetData())
}

type SMSNotificationObserver struct {
	mobile string
	objservable IObservable
}

func NewSMSNotificationObserver(mobile string, stockObserver IObservable) *SMSNotificationObserver{
	return &SMSNotificationObserver{
		mobile: mobile,
		objservable: stockObserver,
	}
}

func (s SMSNotificationObserver) Update() {
	fmt.Printf("SMS on id %s, sent for observable %d\n", s.mobile, s.objservable.GetData())
}