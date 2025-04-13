package service

import "github.com/sirupsen/logrus"

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
