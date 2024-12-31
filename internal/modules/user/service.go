package user

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]UserResponse, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetUser(id int) (UserResponse, error) {
	return s.Repo.FindById(id)
}

func (s *UserService) CreateUser(input CreateUserRequest) (UserResponse, error) {
	return s.Repo.Create(User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Age:       input.Age,
		Password:  input.Password,
	})
}
