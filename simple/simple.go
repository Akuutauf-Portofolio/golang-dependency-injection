package simple

import "errors"

// membuat struct awal untuk repository
type SimpleRepository struct {
	// menambahkan attribute error
	Error bool
}

// membuat struct awal untuk service
type SimpleService struct {
	// service membutuhkan repository
	*SimpleRepository
}

// membuat Provider (google wire) / function constructor
// biasanya provider menggunakan awalan 'New' untuk penamaan function nya

// membuat provider untuk repository
// dengan menambahkan parameter yang sekiranya nanti diperlukan pada injector
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

// membuat provider untuk service yang membutuhkan (depend) repository
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	// melakukan pengecekan, jika repositorynya error nya bernilai true
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}

}
