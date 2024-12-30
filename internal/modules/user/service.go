package user

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) CreateUser(input CreateUserRequest) (User, error) {
	return s.Repo.Create(User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Age:       input.Age,
		Password:  input.Password,
	})
}
