package storage

import "errors"

type Storage interface {
	Save(id int) error
	GetAll() ([]int, error)
	Exists(id int) (bool, error)
}

var ErrNoSavedValues = errors.New("no saved values")
