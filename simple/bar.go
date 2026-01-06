package simple

type BarRepository struct {

}

// membuat provider / function construct bar repository
func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

// membuat provider / function construct bar service
func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{
		BarRepository: barRepository,
	}
}