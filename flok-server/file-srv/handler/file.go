package handler

import (
	"context"
	proto "flok-server/file-srv/fileproto"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/uuid"
)

// FileServiceHandler implements the fileService interface of proto
type FileServiceHandler struct {
}

// Create imlements proto
func (h *FileServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	t, _ := uuid.NewRandom()
	s := t.String()
	s = strings.Replace(s, "-", "", -1)
	req.Name += "-"
	req.Name += s
	req.Name += req.Extension

	if err := ioutil.WriteFile("../data/"+req.Name, req.Data, 0644); err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: req.Name,
	}, nil
}

// Read imlements proto
func (h *FileServiceHandler) Read(ctx context.Context, req *proto.ReadRequest) (*proto.ReadResponse, error) {
	data, err := ioutil.ReadFile("../data/" + req.Name)
	if err != nil {
		return nil, err
	}

	return &proto.ReadResponse{
		File: &proto.File{
			Data: data,
			Name: req.Name,
		},
	}, nil
}

// Delete imlements proto
func (h *FileServiceHandler) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	if err := os.Remove("../data/" + req.Name); err != nil {
		return nil, err
	}
	return &proto.DeleteResponse{}, nil
}
