package handler

import (
	"context"
	"errors"
	"flok-server/lib"
	proto "flok-server/portal-srv/portalproto"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "portals"
var rejectionCommentFormat = "Your submission was last reject by \" %s \" due to \" %s \""

const (
	piDisplayName  = "PI"
	dorDisplayName = "DOR"
	secDisplayName = "Sec."
)

// PortalServiceHandler implements the PortalService interface of proto
type PortalServiceHandler struct {
	Store *lib.Store
}

// Create imlements proto
func (h *PortalServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	portal := &proto.Portal{
		Id:            bson.NewObjectId().Hex(),
		PiID:          req.PiID,
		IsDORApproved: false,
		IsPIApproved:  false,
		IsSecApproved: false,
		StudentID:     req.StudentID,
		IsFreeze:      req.IsFreeze,
	}

	err = db.DB("").C(collectionName).Insert(portal)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: portal.Id,
	}, nil
}

// View imlements proto
func (h *PortalServiceHandler) View(ctx context.Context, req *proto.ViewRequest) (*proto.ViewResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	portal := &proto.Portal{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.PortalID}).One(portal)
	if err != nil {
		return nil, err
	}

	if portal.StudentID != req.UserID && portal.PiID != req.UserID {
		return nil, errors.New("permission denied")
	}

	return &proto.ViewResponse{
		Portal: portal,
	}, nil
}

// Freeze imlements proto
func (h *PortalServiceHandler) Freeze(ctx context.Context, req *proto.FreezeRequest) (*proto.FreezeResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	portal := &proto.Portal{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.PortalID}).One(portal)
	if err != nil {
		return nil, err
	}

	portal.IsFreeze = req.IsFreeze

	err = db.DB("").C(collectionName).Update(bson.M{"_id": portal.Id}, portal)
	if err != nil {
		return nil, err
	}

	return &proto.FreezeResponse{}, nil
}

// SetApprove imlements proto
func (h *PortalServiceHandler) SetApprove(ctx context.Context, req *proto.SetApproveRequest) (*proto.SetApproveResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	portal := &proto.Portal{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.PortalID}).One(portal)
	if err != nil {
		return nil, err
	}

	if portal.IsFreeze {
		return &proto.SetApproveResponse{
			Success: false,
		}, errors.New("portal is freezed")
	}

	if portal.StudentID == req.UserID {
		return &proto.SetApproveResponse{
			Success: false,
		}, errors.New("permission denied")
	}

	portal.IsDORApproved = false
	portal.IsPIApproved = false
	portal.IsSecApproved = false

	displayName := ""

	if !req.IsApproved {
		portal.IsSubmitted = false
	}

	if portal.PiID == req.UserID {
		if portal.IsSecApproved || portal.IsDORApproved {
			return &proto.SetApproveResponse{
				Success: false,
			}, errors.New("update denied")
		}
		if req.IsApproved {
			portal.IsPIApproved = true
			portal.RejectionComment = ""
		} else {
			displayName = piDisplayName
		}
	}

	switch req.Level {
	case 2:
		if !portal.IsPIApproved || portal.IsDORApproved {
			return &proto.SetApproveResponse{
				Success: false,
			}, errors.New("update denied")
		}
		if req.IsApproved {
			portal.IsPIApproved = true
			portal.IsSecApproved = true
			portal.RejectionComment = ""
		} else {
			displayName = secDisplayName
		}
	case 3:
		if !portal.IsSecApproved || !portal.IsPIApproved {
			return &proto.SetApproveResponse{
				Success: false,
			}, errors.New("update denied")
		}
		if req.IsApproved {
			portal.IsDORApproved = true
			portal.IsPIApproved = true
			portal.IsSecApproved = true
			portal.RejectionComment = ""
			portal.IsFreeze = true
		} else {
			displayName = dorDisplayName
		}
	}
	/*
		return &proto.SetApproveResponse{
			Success: false,
		}, errors.New("permission denied") */

	if !req.IsApproved {
		portal.RejectionComment = fmt.Sprintf(rejectionCommentFormat, displayName, req.RejectionComment)
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": portal.Id}, portal)
	if err != nil {
		return nil, err
	}

	return &proto.SetApproveResponse{
		Success: true,
	}, nil
}

// SetSubmitted imlements proto
func (h *PortalServiceHandler) SetSubmitted(ctx context.Context, req *proto.SetSubmittedRequest) (*proto.SetSubmittedResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	portal := &proto.Portal{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(portal)
	if err != nil {
		return nil, err
	}

	if portal.IsFreeze {
		return &proto.SetSubmittedResponse{}, errors.New("portal is freezed")
	}

	if portal.StudentID != req.User {
		return nil, errors.New("permission denied")
	}

	portal.IsSubmitted = req.IsSubmitted

	err = db.DB("").C(collectionName).Update(bson.M{"_id": portal.Id}, portal)
	if portal.StudentID != req.User {
		return nil, errors.New("permission denied")
	}

	return &proto.SetSubmittedResponse{}, nil
}
