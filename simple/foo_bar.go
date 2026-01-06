package simple

type Foo struct {

}

// membuat provider / function constructor
func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

// membuat provider / function constructor
func NewBar() *Bar {
	return &Bar{}
}

// membuat struct yang attribute nya berisi provider dari foo dan bar
type FooBar struct {
	*Foo
	*Bar
}