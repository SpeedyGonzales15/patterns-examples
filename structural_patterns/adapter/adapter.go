package adapter

// Target представляет интерфейс, к которому мы хотим адаптировать другой класс.
type Target interface {
	Request() string
}

// Adaptee содержит полезную функциональность, но его интерфейс несовместим с существующим клиентским кодом.
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "specific request"
}

// Adapter делает интерфейс Adaptee совместимым с Target интерфейсом.
type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
