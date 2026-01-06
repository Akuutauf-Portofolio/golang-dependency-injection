package simple

// membuat interface sebagai a 
type SayHello interface {
	// fungsi bawaan dari interface a sebagai attribute
	Hello(name string) string
}

// membuat struct service sebagai b
type HelloService struct {
	// memiliki attribute say hello dengan tipe say hello
	SayHello SayHello // tipe nya adalah interface
}

// membuat struct implementation sebagai implementasi dari interface SayHello / a
type SayHelloImpl struct {

}

// mengimplementasikan kontrak dari struct SayHelloImpl
// dan function say hello
func(s *SayHelloImpl) Hello(name string) string {
	return "Hello" + name
}

// membuat provider / function constructor untuk struct HelloService
// disarankan untuk langsung deklarasi return berupa struct (SayHelloImpl)
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: sayHello,
	}
}

// membuat provider / function constructor untuk struct SayHelloImpl
// disarankan untuk langsung deklarasi return berupa struct (SayHelloImpl)
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}