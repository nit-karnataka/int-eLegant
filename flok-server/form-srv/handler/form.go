package handler

import (
	"context"
	proto "flok-server/form-srv/formproto"
	"flok-server/lib"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "forms"

// FormServiceHandler implements the FormService interface of proto
type FormServiceHandler struct {
	Store *lib.Store
}

// Create imlements proto
func (h *FormServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	req.Form.Id = bson.NewObjectId().Hex()

	err = db.DB("").C(collectionName).Insert(req.Form)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: req.Form.Id,
	}, nil
}

// View imlements proto
func (h *FormServiceHandler) View(ctx context.Context, req *proto.ViewRequest) (*proto.ViewResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	form := &proto.Form{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(form)
	if err != nil {
		return nil, err
	}

	return &proto.ViewResponse{
		Form: form,
	}, nil
}

// Update imlements proto
func (h *FormServiceHandler) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Form.Id}, req.Form)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateResponse{}, nil
}
