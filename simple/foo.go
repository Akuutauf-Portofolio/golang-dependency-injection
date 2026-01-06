package simple

type FooRepository struct {

}

// membuat construct foo repository
func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

// membuat construct foo service
func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{
		FooRepository: fooRepository,
	}
}