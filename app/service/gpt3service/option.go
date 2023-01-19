package gpt3service

import "time"

type Option func(s *service) error

// SetModel 設置模型
func SetModel(model Model) Option {
	return func(s *service) error {
		s.model = model
		return nil
	}
}

func SetTimeout(timeout time.Duration) Option {
	return func(s *service) error {
		s.httpClient.Timeout = timeout
		return nil
	}
}
