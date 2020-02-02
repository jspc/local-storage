package main

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/j-and-j-global/storage-service"
	"github.com/pkg/errors"
)

type Server struct {
	ID        string
	Directory string
	Mode      os.FileMode
}

func (s Server) NewFile(context.Context, *empty.Empty) (f *storage.File, err error) {
	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	f = &storage.File{
		Id: id.String(),
	}

	return
}

func (s Server) Upload(stream storage.Storage_UploadServer) error {
	var f *os.File

	for {
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				goto END
			}

			return errors.Wrap(err,
				"failed unexpectadely while reading chunks from stream")
		}

		if f == nil {
			// first chunk
			f, err = os.OpenFile(s.path(chunk.File.Id), os.O_RDWR|os.O_CREATE, s.Mode)
			if err != nil {
				return err
			}

			defer f.Close()
		}

		_, err = f.Write(chunk.Data)
		if err != nil {
			return err
		}
	}

END:

	return stream.SendAndClose(nil)
}

func (s Server) Delete(ctx context.Context, f *storage.File) (*empty.Empty, error) {
	return nil, os.Remove(s.path(f.Id))
}

func (s Server) Status(context.Context, *empty.Empty) (*storage.ServerStatus, error) {
	return &storage.ServerStatus{
		Type:  "local-storage",
		Id:    s.ID,
		Ready: true,
	}, nil
}

func (s Server) path(s1 string) string {
	return filepath.Join(s.Directory, s1)
}
