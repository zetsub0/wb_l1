package main

import "fmt"

// Target - интерфейс, который ожидается от клиентского кода
type Target interface {
	Request() string
}

// Adaptee - адаптируемый тип
type Adaptee struct{}

// SpecificRequest - метод, который нужно адаптировать
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter - адаптер, который преобразует метод SpecificRequest к интерфейсу Target
type Adapter struct {
	adaptee *Adaptee
}

// NewAdapter - создает новый экземпляр адаптера с заданным адаптируемым объектом
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Request преобразует вызов Request к вызову SpecificRequest адаптируемого объекта
func (adapter *Adapter) Request() string {
	return adapter.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)

	fmt.Println(adapter.Request())
}
