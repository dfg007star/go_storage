package storage

import (
	"errors"
	"github.com/dfg007star/go_storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	file, error := file.NewFile(filename, blob)
	if error != nil {
		return nil, error
	}

	s.files[file.ID] = file

	return file, nil
}

func (s *Storage) GetById(fileID uuid.UUID) (*file.File, error) {
	file, ok := s.files[fileID]
	if !ok {
		return nil, errors.New("file not found")
	}

	return file, nil
}
