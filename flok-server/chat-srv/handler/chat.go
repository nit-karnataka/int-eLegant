package handler

import (
	"context"
	"errors"
	proto "flok-server/chat-srv/chatproto"
	"flok-server/lib"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/thoas/go-funk"
)

var collectionName = "chats"

// ChatServiceHandler implements the ChatService interface of proto
type ChatServiceHandler struct {
	Store *lib.Store
}

// Create imlements proto
func (h *ChatServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	chat := &proto.Chat{
		Id:             bson.NewObjectId().Hex(),
		PermittedUsers: req.Owner,
		IsFreeze:       req.IsFreeze,
	}

	err = db.DB("").C(collectionName).Insert(chat)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: chat.Id,
	}, nil
}

// InsertComment imlements proto
func (h *ChatServiceHandler) InsertComment(ctx context.Context, req *proto.InsertCommentRequest) (*proto.InsertCommentResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	chat := &proto.Chat{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.ChatID}).One(chat)
	if err != nil {
		return nil, err
	}

	if chat.IsFreeze {
		return &proto.InsertCommentResponse{
			Success: false,
		}, errors.New("chat is freezed")
	}

	log.Printf("Permitted Users ID: %+v, %s", chat.PermittedUsers[0], req.Owner)

	if !funk.ContainsString(chat.PermittedUsers, req.Owner) {
		return &proto.InsertCommentResponse{
			Success: false,
		}, errors.New("permission denied")
	}

	chat.Comments = append(chat.Comments, &proto.Comment{
		Content: req.Content,
		Owner:   req.Owner,
	})

	err = db.DB("").C(collectionName).Update(bson.M{"_id": chat.Id}, chat)
	if err != nil {
		return nil, err
	}

	return &proto.InsertCommentResponse{
		Success: true,
	}, nil
}

// View imlements proto
func (h *ChatServiceHandler) View(ctx context.Context, req *proto.ViewRequest) (*proto.ViewResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	chat := &proto.Chat{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.ChatID}).One(chat)
	if err != nil {
		return nil, err
	}

	if !funk.ContainsString(chat.PermittedUsers, req.UserID) {
		return nil, errors.New("permission denied")
	}

	return &proto.ViewResponse{
		Chat: chat,
	}, nil
}

// Freeze imlements proto
func (h *ChatServiceHandler) Freeze(ctx context.Context, req *proto.FreezeRequest) (*proto.FreezeResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	chat := &proto.Chat{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.ChatID}).One(chat)
	if err != nil {
		return nil, err
	}

	if !funk.ContainsString(chat.PermittedUsers, req.UserID) {
		return nil, errors.New("permission denied")
	}

	chat.IsFreeze = req.IsFreeze

	err = db.DB("").C(collectionName).Update(bson.M{"_id": chat.Id}, chat)
	if err != nil {
		return nil, err
	}

	return &proto.FreezeResponse{}, nil
}
