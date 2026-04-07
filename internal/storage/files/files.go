package files

import (
	"encoding/gob"
	"errors"
	"log"
	"os"
	"path/filepath"

	"stanBot/internal/storage"
)

type Storage struct {
	filePath string
}

func New(filePath string) Storage {
	return Storage{filePath: filePath}
}

func (s Storage) Save(id int) error {
	log.Printf("storage wants to save chatID = %v", id)
	exists, err := s.Exists(id)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	values, err := s.GetAll()
	if err != nil && !errors.Is(err, storage.ErrNoSavedValues) {
		return err
	}

	values = append(values, id)

	return s.saveAll(values)
}

func (s Storage) GetAll() ([]int, error) {
	file, err := os.Open(s.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, storage.ErrNoSavedValues
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var values []int

	if err := gob.NewDecoder(file).Decode(&values); err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, storage.ErrNoSavedValues
	}

	log.Printf("succesfully retrived chatIDs: %v", values)

	return values, nil
}

func (s Storage) Exists(id int) (bool, error) {
	values, err := s.GetAll()
	if err != nil {
		if errors.Is(err, storage.ErrNoSavedValues) {
			return false, nil
		}
		return false, err
	}

	for _, v := range values {
		if v == id {
			return true, nil
		}
	}

	return false, nil
}

func (s Storage) saveAll(values []int) error {
	if err := os.MkdirAll(filepath.Dir(s.filePath), 0774); err != nil {
		return err
	}

	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	return gob.NewEncoder(file).Encode(values)
}
