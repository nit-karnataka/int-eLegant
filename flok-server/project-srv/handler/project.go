package handler

import (
	"context"
	"flok-server/lib"
	proto "flok-server/project-srv/projectproto"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "projects"

// ProjectServiceHandler implements the HouseService interface of proto
type ProjectServiceHandler struct {
	Store *lib.Store
}

// Create imlements proto
func (h *ProjectServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	req.Project.Id = bson.NewObjectId().Hex()

	err = db.DB("").C(collectionName).Insert(req.Project)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: req.Project.Id,
	}, nil
}

// Create imlements proto
func (h *ProjectServiceHandler) Read(ctx context.Context, req *proto.ReadRequest) (*proto.ReadResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	project := &proto.Project{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(project)
	if err != nil {
		return nil, err
	}

	return &proto.ReadResponse{
		Project: project,
	}, nil
}

// AddDocument imlements proto
func (h *ProjectServiceHandler) AddDocument(ctx context.Context, req *proto.AddDocumentRequest) (*proto.AddDocumentResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Id}, bson.M{
		"$addToSet": bson.M{
			"docs": req.Doc,
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.AddDocumentResponse{}, nil
}

// RemoveDocument imlements proto
func (h *ProjectServiceHandler) RemoveDocument(ctx context.Context, req *proto.RemoveDocumentRequest) (*proto.RemoveDocumentResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Id}, bson.M{
		"$pull": bson.M{
			"docs": req.Doc,
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.RemoveDocumentResponse{}, nil
}

// GetAll imlements proto
func (h *ProjectServiceHandler) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	projects := []*proto.Project{}

	err = db.DB("").C(collectionName).Find(nil).All(&projects)
	if err != nil {
		return nil, err
	}

	return &proto.GetAllResponse{
		Projects: projects,
	}, nil
}

// GetByStudent imlements proto
func (h *ProjectServiceHandler) GetByStudent(ctx context.Context, req *proto.GetByStudentRequest) (*proto.GetByStudentResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	projects := []*proto.Project{}

	err = db.DB("").C(collectionName).Find(bson.M{"student": req.Student}).All(&projects)
	if err != nil {
		return nil, err
	}

	return &proto.GetByStudentResponse{
		Projects: projects,
	}, nil
}

// SetIRBMeeting imlements proto
func (h *ProjectServiceHandler) SetIRBMeeting(ctx context.Context, req *proto.SetIRBMeetingRequest) (*proto.SetIRBMeetingResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Id}, bson.M{
		"$set": bson.M{
			"isIRBMeetingDone": req.Status,
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.SetIRBMeetingResponse{}, nil
}

// SetIsApproved imlements proto
func (h *ProjectServiceHandler) SetIsApproved(ctx context.Context, req *proto.SetIsApprovedRequest) (*proto.SetIsApprovedResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Id}, bson.M{
		"$set": bson.M{
			"isApproved": req.Status,
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.SetIsApprovedResponse{}, nil
}

// SetThesisDone imlements proto
func (h *ProjectServiceHandler) SetThesisDone(ctx context.Context, req *proto.SetThesisDoneRequest) (*proto.SetThesisDoneResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Id}, bson.M{
		"$set": bson.M{
			"isThesisDone": req.Status,
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.SetThesisDoneResponse{}, nil
}

/*
// Accept imlements proto
func (h *ProjectServiceHandler) Accept(ctx context.Context, req *proto.AcceptRequest) (*proto.AcceptResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	project := &proto.Project{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.ProjectID}).One(project)
	if err != nil {
		return nil, err
	}

	if project.PrincipalInvestigator == req.PrincipalInvestigator && project.Id == req.ProjectID {
		project.IsAcceptedByPI = true
		err = db.DB("").C(collectionName).Update(bson.M{"_id": project.Id}, project)
		if err != nil {
			return nil, err
		}
	}

	return nil, errors.New("project id and pi dont match")
}
*/
