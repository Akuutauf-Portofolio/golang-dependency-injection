package simple

type FooBarService struct {
	*FooService
	*BarService
}

// membuat provider / function construct
func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}