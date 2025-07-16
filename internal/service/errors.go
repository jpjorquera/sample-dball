package service

import "errors"

var (
	ErrNotFound    = errors.New("record not found")
	ErrExternalAPI = errors.New("failed to connect to upstream service")
	ErrDatabase    = errors.New("database failure")
)
