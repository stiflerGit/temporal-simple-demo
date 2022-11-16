package store

import "time"

type Storage struct {
	userToLastSignUpTime map[string]time.Time
}

func (s *Storage) GetUserLastSignUpTime(username string) (time.Time, error) {
	t := s.userToLastSignUpTime[username]
	return t, nil
}
