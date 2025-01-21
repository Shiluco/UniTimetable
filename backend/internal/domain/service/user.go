package service

import (
    "context"
    "github.com/Shiluco/UniTimetable/backend/ent"
)

type UserService struct {
    Client *ent.Client
}

func (s *UserService) GetUsers() ([]*ent.User, error) {
    return s.Client.User.Query().All(context.Background())
}
