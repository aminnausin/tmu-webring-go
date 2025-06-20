package user

import (
	"tmu-webring-go/internal/mailer"
	"tmu-webring-go/internal/storage"
)

type UserService struct {
	mailer  mailer.Client
	storage storage.Storage
}

func NewUserService(storage storage.Storage, mailer mailer.Client) *UserService {
	return &UserService{storage: storage, mailer: mailer}
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *UserService) Create(req *CreateUserRequest) error {
	user := &storage.User{Name: req.Name, Email: req.Email}

	// if err := s.storage.CreateUser(user); err != nil {
	// 	return err

	// }

	s.mailer.Send(mailer.UserWelcomeTemplate, user.Name, user.Email, user, false)

	return nil
}
