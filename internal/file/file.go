package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, data []byte) (*File, error) {
	fileId, error := uuid.NewRandom()

	if error != nil {
		return nil, error
	}

	return &File{
		ID:   fileId,
		Name: name,
		Data: data,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File{ID: %s, Name: %s, Data: %s}", f.ID, f.Name, f.Data)
}
