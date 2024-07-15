package abstractfactory

type Widget interface {
	Render()
}

type LinuxButton struct{}

func (lb *LinuxButton) Render() {
	// Реализация для Linux кнопки
}

type MacOSButton struct{}

func (mb *MacOSButton) Render() {
	// Реализация для MacOS кнопки
}

type WindowsButton struct{}

func (wb *WindowsButton) Render() {
	// Реализация для Windows кнопки
}

type WidgetFactory interface {
	CreateButton() Widget
}

type LinuxFactory struct{}

func (lf *LinuxFactory) CreateButton() Widget {
	return &LinuxButton{}
}

type MacOSFactory struct{}

func (mf *MacOSFactory) CreateButton() Widget {
	return &MacOSButton{}
}

type WindowsFactory struct{}

func (mf *WindowsFactory) CreateButton() Widget {
	return &WindowsButton{}
}
