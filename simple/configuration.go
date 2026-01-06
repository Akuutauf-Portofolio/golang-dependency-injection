package simple

type Configuration struct {
	Name string
}

type Application struct {
	// attribute nya berisi struct Configuration
	*Configuration
}

// membuat provider / function construct untuk application
func NewApplication() *Application {
	return &Application{
		// kita ingin menjadi configuration di bawah ini sebagai provider
		Configuration: &Configuration{ 
			Name: "Belajar Golang",
		},
	}
}