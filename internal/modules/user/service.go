package user

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]UserResponse, error) {

	users, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Age:       user.Age,
		})
	}
	return response, nil
}

func (s *UserService) GetUser(id int) (UserResponse, error) {
	user, err := s.Repo.FindById(id)
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}, nil
}

func (s *UserService) CreateUser(input CreateUserRequest) (UserResponse, error) {

	user, err := s.Repo.Create(input)
	if err != nil {
		return UserResponse{}, err
	}
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}, nil
}

func (s *UserService) UpdateUser(id int, input UpdateUserRequest) (UserResponse, error) {
	user, err := s.Repo.Update(id, input)
	if err != nil {
		return UserResponse{}, err
	}
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.Delete(id)
}
