package services

import "time"

type TimeProvider struct{}

func NewTimeProvider() *TimeProvider {
	return &TimeProvider{}
}

func (tp *TimeProvider) Now() string {
	return time.Now().Format("15:04:05")
}
