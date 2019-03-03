package handler

import (
	"context"
	proto "flok-server/auth-srv/authproto"
	"flok-server/auth-srv/crypto"
	"flok-server/lib"
	"log"

	"github.com/go-redis/redis"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "auths"

// AuthServiceHandler implements the HouseService interface of proto
type AuthServiceHandler struct {
	Store       *lib.Store
	Hash        *crypto.Hash
	CacheClient *redis.Client
}

// Create imlements proto
func (h *AuthServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	hash, err := h.Hash.Generate(req.User.Password)
	if err != nil {
		return nil, err
	}

	req.User.Password = hash

	err = db.DB("").C(collectionName).Insert(req.User)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// VerifyUser imlements proto
func (h *AuthServiceHandler) VerifyUser(ctx context.Context, req *proto.VerifyUserRequest) (*proto.VerifyUserResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	u := &proto.User{}

	err = db.DB("").C(collectionName).Find(bson.M{"email": req.Email}).One(u)
	if err != nil {
		return nil, err
	}

	log.Printf("Checking for user: %+v", u)

	err = h.Hash.Compare(u.Password, req.Password)
	if err != nil {
		return &proto.VerifyUserResponse{}, err
	}

	return &proto.VerifyUserResponse{
		User: u,
	}, nil
}

// Delete imlements proto
func (h *AuthServiceHandler) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	db, err := h.Store.GetMongoSession()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Remove(bson.M{"email": req.Id})
	if err != nil {
		return nil, err
	}

	return &proto.DeleteResponse{}, nil
}

// Update imlements proto
func (h *AuthServiceHandler) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	db, err := h.Store.GetMongoSession()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	hash, err := h.Hash.Generate(req.User.Password)
	if err != nil {
		return nil, err
	}

	req.User.Password = hash

	err = db.DB("").C(collectionName).Update(bson.M{"email": req.User.Email}, req.User)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateResponse{}, nil
}
