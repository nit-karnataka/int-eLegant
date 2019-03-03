package handler

import (
	"context"
	"errors"
	"flok-server/lib"
	proto "flok-server/meeting-srv/meetingproto"
	"fmt"

	"github.com/thoas/go-funk"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "meetings"
var rejectionCommentFormat = "Your submission was last reject by \" %s \" due to \" %s \""

const (
	piDisplayName  = "PI"
	dorDisplayName = "DOR"
	secDisplayName = "Sec."
)

// MeetingServiceHandler implements the MeetingService interface of proto
type MeetingServiceHandler struct {
	Store *lib.Store
}

// Create imlements proto
func (h *MeetingServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	meeting := proto.Meeting{
		Id:          bson.NewObjectId().Hex(),
		ChairPerson: req.ChairPerson,
	}

	err = db.DB("").C(collectionName).Insert(meeting)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: meeting.Id,
	}, nil
}

// View imlements proto
func (h *MeetingServiceHandler) View(ctx context.Context, req *proto.ViewRequest) (*proto.ViewResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	meeting := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(meeting)
	if err != nil {
		return nil, err
	}

	return &proto.ViewResponse{
		Meeting: meeting,
	}, nil
}

// Freeze imlements proto
func (h *MeetingServiceHandler) Freeze(ctx context.Context, req *proto.FreezeRequest) (*proto.FreezeResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	meeting := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(meeting)
	if err != nil {
		return nil, err
	}

	meeting.IsFreeze = req.IsFreeze

	err = db.DB("").C(collectionName).Update(bson.M{"_id": meeting.Id}, meeting)
	if err != nil {
		return nil, err
	}

	return &proto.FreezeResponse{}, nil
}

// AddInitialMembers imlements proto
func (h *MeetingServiceHandler) AddInitialMembers(ctx context.Context, req *proto.AddInitialMembersRequest) (*proto.AddInitialMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	member := []*proto.InitialMember{}

	for _, mem := range req.Members {
		member = append(member, &proto.InitialMember{
			Member:       mem,
			ReceiptDocID: "",
			Status:       false,
		})
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"initialMembers": bson.M{"$each": member}}})
	if err != nil {
		return nil, err
	}

	return &proto.AddInitialMembersResponse{}, nil
}

// RemoveInitialMembers imlements proto
func (h *MeetingServiceHandler) RemoveInitialMembers(ctx context.Context, req *proto.RemoveInitialMembersRequest) (*proto.RemoveInitialMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	member := []*proto.InitialMember{}

	for _, mem := range m.InitialMembers {
		if funk.ContainsString(req.Members, mem.Member) {
			member = append(member, mem)
		}
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$pullAll": bson.M{"initialMembers": member}})
	if err != nil {
		return nil, err
	}

	return &proto.RemoveInitialMembersResponse{}, nil
}

// SetAgenda imlements proto
func (h *MeetingServiceHandler) SetAgenda(ctx context.Context, req *proto.SetAgendaRequest) (*proto.SetAgendaResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.AgendaDocID = req.Doc
	m.IsAgendaApproved = false

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.SetAgendaResponse{}, nil
}

// SetMinutes imlements proto
func (h *MeetingServiceHandler) SetMinutes(ctx context.Context, req *proto.SetMinutesRequest) (*proto.SetMinutesResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.MinutesDocID = req.Doc
	m.IsMinutesAccepted = false

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.SetMinutesResponse{}, nil
}

// ApproveMinute imlements proto
func (h *MeetingServiceHandler) ApproveMinute(ctx context.Context, req *proto.ApproveMinuteRequest) (*proto.ApproveMinuteResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsMinutesAccepted = req.Approve
	m.MinutesRejectionComment = ""
	if !req.Approve {
		m.MinutesRejectionComment = fmt.Sprintf(rejectionCommentFormat, dorDisplayName, req.Comment)
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.ApproveMinuteResponse{}, nil
}

// ApproveAgenda imlements proto
func (h *MeetingServiceHandler) ApproveAgenda(ctx context.Context, req *proto.ApproveAgendaRequest) (*proto.ApproveAgendaResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsAgendaApproved = req.Approve
	m.AgendaRejectionComment = ""
	if !req.Approve {
		m.AgendaRejectionComment = fmt.Sprintf(rejectionCommentFormat, dorDisplayName, req.Comment)
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.ApproveAgendaResponse{}, nil
}

// AddPresentMembers imlements proto
func (h *MeetingServiceHandler) AddPresentMembers(ctx context.Context, req *proto.AddPresentMembersRequest) (*proto.AddPresentMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"presentMembers": bson.M{"$each": req.Members}}})
	if err != nil {
		return nil, err
	}

	return &proto.AddPresentMembersResponse{}, nil
}

// RemovePresentMembers imlements proto
func (h *MeetingServiceHandler) RemovePresentMembers(ctx context.Context, req *proto.RemovePresentMembersRequest) (*proto.RemovePresentMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$pullAll": bson.M{"presentMembers": req.Members}})
	if err != nil {
		return nil, err
	}

	return &proto.RemovePresentMembersResponse{}, nil
}

// AddQLMembers imlements proto
func (h *MeetingServiceHandler) AddQLMembers(ctx context.Context, req *proto.AddQLMembersRequest) (*proto.AddQLMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	member := []*proto.ApprovalMember{}

	for _, mem := range req.Members {
		member = append(member, &proto.ApprovalMember{
			Member:           mem,
			RejectionComment: "",
			Status:           false,
		})
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"approvalMembers": bson.M{"$each": member}}})
	if err != nil {
		return nil, err
	}

	return &proto.AddQLMembersResponse{}, nil
}

// RemoveQLMembers imlements proto
func (h *MeetingServiceHandler) RemoveQLMembers(ctx context.Context, req *proto.RemoveQLMembersRequest) (*proto.RemoveQLMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	member := []*proto.ApprovalMember{}

	for _, mem := range m.ApprovalMembers {
		if funk.ContainsString(req.Members, mem.Member) {
			member = append(member, mem)
		}
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$pullAll": bson.M{"approvalMembers": member}})
	if err != nil {
		return nil, err
	}

	return &proto.RemoveQLMembersResponse{}, nil
}

// AddFAMembers imlements proto
func (h *MeetingServiceHandler) AddFAMembers(ctx context.Context, req *proto.AddFAMembersRequest) (*proto.AddFAMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	member := []*proto.ApprovalMember{}

	for _, mem := range req.Members {
		member = append(member, &proto.ApprovalMember{
			Member:           mem,
			RejectionComment: "",
			Status:           false,
		})
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"approvalMembers": bson.M{"$each": member}}})
	if err != nil {
		return nil, err
	}

	return &proto.AddFAMembersResponse{}, nil
}

// RemoveFAMembers imlements proto
func (h *MeetingServiceHandler) RemoveFAMembers(ctx context.Context, req *proto.RemoveFAMembersRequest) (*proto.RemoveFAMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	member := []*proto.ApprovalMember{}

	for _, mem := range m.ApprovalMembers {
		if funk.ContainsString(req.Members, mem.Member) {
			member = append(member, mem)
		}
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$pullAll": bson.M{"approvalMembers": member}})
	if err != nil {
		return nil, err
	}

	return &proto.RemoveFAMembersResponse{}, nil
}

// SetTranslation imlements proto
func (h *MeetingServiceHandler) SetTranslation(ctx context.Context, req *proto.SetTranslationRequest) (*proto.SetTranslationResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsTranslationApproved = req.T

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.SetTranslationResponse{}, nil
}

// AddQL imlements proto
func (h *MeetingServiceHandler) AddQL(ctx context.Context, req *proto.AddQLRequest) (*proto.AddQLResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"queryLetter": req.Ql}})

	return &proto.AddQLResponse{}, nil
}

// SetChairPerson imlements proto
func (h *MeetingServiceHandler) SetChairPerson(ctx context.Context, req *proto.SetChairPersonRequest) (*proto.SetChairPersonResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{"chairPerson": req.User}})

	return &proto.SetChairPersonResponse{}, nil
}

// SetQLSubmitted imlements proto
func (h *MeetingServiceHandler) SetQLSubmitted(ctx context.Context, req *proto.SetQLSubmittedRequest) (*proto.SetQLSubmittedResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsQLSubmitted = req.IsQLSubmitted

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.SetQLSubmittedResponse{}, nil
}

// AddQLR imlements proto
func (h *MeetingServiceHandler) AddQLR(ctx context.Context, req *proto.AddQLRRequest) (*proto.AddQLRResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"queryLetterReplies": req.Qlr}})
	if err != nil {
		return nil, err
	}

	return &proto.AddQLRResponse{}, nil
}

// RemoveQLR imlements proto
func (h *MeetingServiceHandler) RemoveQLR(ctx context.Context, req *proto.RemoveQLRRequest) (*proto.RemoveQLRResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$pull": bson.M{"queryLetterReplies": req.Qlr}})
	if err != nil {
		return nil, err
	}

	return &proto.RemoveQLRResponse{}, nil
}

// ApproveQLRPI imlements proto
func (h *MeetingServiceHandler) ApproveQLRPI(ctx context.Context, req *proto.ApproveQLRPIRequest) (*proto.ApproveQLRPIResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsQLPIApproved = req.IsApproved

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.ApproveQLRPIResponse{}, nil
}

// ApproveQLRSec imlements proto
func (h *MeetingServiceHandler) ApproveQLRSec(ctx context.Context, req *proto.ApproveQLRSecRequest) (*proto.ApproveQLRSecResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsQLSecApproved = req.IsApproved

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.ApproveQLRSecResponse{}, nil
}

// AddTranslation imlements proto
func (h *MeetingServiceHandler) AddTranslation(ctx context.Context, req *proto.AddTranslationRequest) (*proto.AddTranslationResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$addToSet": bson.M{"translations": req.Translation}})
	if err != nil {
		return nil, err
	}

	return &proto.AddTranslationResponse{}, nil
}

// RemoveTranslation imlements proto
func (h *MeetingServiceHandler) RemoveTranslation(ctx context.Context, req *proto.RemoveTranslationRequest) (*proto.RemoveTranslationResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$pull": bson.M{"translations": req.Translation}})
	if err != nil {
		return nil, err
	}

	return &proto.RemoveTranslationResponse{}, nil
}

// ApproveTranslationPI imlements proto
func (h *MeetingServiceHandler) ApproveTranslationPI(ctx context.Context, req *proto.ApproveTranslationPIRequest) (*proto.ApproveTranslationPIResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsTPIApproved = req.IsApproved

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.ApproveTranslationPIResponse{}, nil
}

// ApproveTranslationSec imlements proto
func (h *MeetingServiceHandler) ApproveTranslationSec(ctx context.Context, req *proto.ApproveTranslationSecRequest) (*proto.ApproveTranslationSecResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	m.IsTSecApproved = req.IsApproved

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, m)
	if err != nil {
		return nil, err
	}

	return &proto.ApproveTranslationSecResponse{}, nil
}

// SetQLReview imlements proto
func (h *MeetingServiceHandler) SetQLReview(ctx context.Context, req *proto.SetQLReviewRequest) (*proto.SetQLReviewResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	if m.ChairPerson == req.Member && req.Status {
		err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
			"isQLApproved": true,
		}})
		if err != nil {
			return nil, err
		}
		_, err = h.RemoveAllApprovalMembers(ctx, &proto.RemoveAllApprovalMembersRequest{
			Id: req.Id,
		})
		if err != nil {
			return nil, err
		}
		return &proto.SetQLReviewResponse{
			IsAccepted: true,
		}, nil
	}
	comment := ""
	if !req.Status {
		comment = fmt.Sprintf(rejectionCommentFormat, "IRB Member", req.RejectionComment)
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id, "approvalMembers.member": req.Member}, bson.M{
		"$set": bson.M{
			"aprrovalMembers.$.status":           req.Status,
			"aprrovalMembers.$.rejectionComment": comment,
		},
	})
	if err != nil {
		return nil, err
	}

	if req.Status {
		cnt := 0
		for _, mem := range m.ApprovalMembers {
			if mem.Status {
				cnt++
			}
		}

		if int32(cnt+1) >= m.NomForApproval {
			err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
				"isQLApproved": true,
			}})
			if err != nil {
				return nil, err
			}
			_, err = h.RemoveAllApprovalMembers(ctx, &proto.RemoveAllApprovalMembersRequest{
				Id: req.Id,
			})
			if err != nil {
				return nil, err
			}
			return &proto.SetQLReviewResponse{
				IsAccepted: true,
			}, nil
		}
	}

	return &proto.SetQLReviewResponse{
		IsAccepted: false,
	}, nil
}

// SetTranslationReview imlements proto
func (h *MeetingServiceHandler) SetTranslationReview(ctx context.Context, req *proto.SetTranslationReviewRequest) (*proto.SetTranslationReviewResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	if m.ChairPerson == req.Member && req.Status {
		err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
			"isTranslationpproved": true,
		}})
		if err != nil {
			return nil, err
		}
		_, err = h.RemoveAllApprovalMembers(ctx, &proto.RemoveAllApprovalMembersRequest{
			Id: req.Id,
		})
		if err != nil {
			return nil, err
		}
		return &proto.SetTranslationReviewResponse{
			IsAccepted: true,
		}, nil
	}
	comment := ""
	if !req.Status {
		comment = fmt.Sprintf(rejectionCommentFormat, "IRB Member", req.RejectionComment)
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id, "approvalMembers.member": req.Member}, bson.M{
		"$set": bson.M{
			"aprrovalMembers.$.status":           req.Status,
			"aprrovalMembers.$.rejectionComment": comment,
		},
	})
	if err != nil {
		return nil, err
	}

	if req.Status {
		cnt := 0
		for _, mem := range m.ApprovalMembers {
			if mem.Status {
				cnt++
			}
		}

		if int32(cnt+1) >= m.NomForApproval {
			err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
				"isTranslationApproved": true,
			}})
			if err != nil {
				return nil, err
			}
			_, err = h.RemoveAllApprovalMembers(ctx, &proto.RemoveAllApprovalMembersRequest{
				Id: req.Id,
			})
			if err != nil {
				return nil, err
			}
			return &proto.SetTranslationReviewResponse{
				IsAccepted: true,
			}, nil
		}
	}

	return &proto.SetTranslationReviewResponse{
		IsAccepted: false,
	}, nil
}

// RemoveAllApprovalMembers imlements proto
func (h *MeetingServiceHandler) RemoveAllApprovalMembers(ctx context.Context, req *proto.RemoveAllApprovalMembersRequest) (*proto.RemoveAllApprovalMembersResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}
	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
		"approvalMembers": []*proto.ApprovalMember{},
	}})
	if err != nil {
		return nil, err
	}
	return &proto.RemoveAllApprovalMembersResponse{}, nil
}

// SetFA imlements proto
func (h *MeetingServiceHandler) SetFA(ctx context.Context, req *proto.SetFARequest) (*proto.SetFAResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(m)
	if err != nil {
		return nil, err
	}

	if m.IsFreeze {
		return nil, errors.New("portal is freezed")
	}

	if m.ChairPerson == req.Member && req.Status {
		err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
			"isApproved": true,
		}})
		if err != nil {
			return nil, err
		}
		return &proto.SetFAResponse{
			IsApproved: true,
		}, nil
	}
	comment := ""

	err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id, "approvalMembers.member": req.Member}, bson.M{
		"$set": bson.M{
			"aprrovalMembers.$.status":           req.Status,
			"aprrovalMembers.$.rejectionComment": comment,
		},
	})
	if err != nil {
		return nil, err
	}

	if req.Status {
		cnt := 0
		for _, mem := range m.ApprovalMembers {
			if mem.Status {
				cnt++
			}
		}

		if int32(cnt+1) >= m.NomForApproval {
			err = db.DB("").C(collectionName).Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
				"isApproved": true,
			}})
			if err != nil {
				return nil, err
			}
			return &proto.SetFAResponse{
				IsApproved: true,
			}, nil
		}
	}

	return &proto.SetFAResponse{
		IsApproved: false,
	}, nil
}

// SetAgendaSubmitted imlements proto
func (h *MeetingServiceHandler) SetAgendaSubmitted(ctx context.Context, req *proto.SetAgendaSubmittedRequest) (*proto.SetAgendaSubmittedResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	meeting := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(meeting)
	if err != nil {
		return nil, err
	}

	meeting.IsAgendaSubmitted = req.Status

	err = db.DB("").C(collectionName).Update(bson.M{"_id": meeting.Id}, meeting)
	if err != nil {
		return nil, err
	}

	return &proto.SetAgendaSubmittedResponse{}, nil
}

// SetMinutesSubmitted imlements proto
func (h *MeetingServiceHandler) SetMinutesSubmitted(ctx context.Context, req *proto.SetMinutesSubmittedRequest) (*proto.SetMinutesSubmittedResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	meeting := &proto.Meeting{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(meeting)
	if err != nil {
		return nil, err
	}

	meeting.IsMinutesSubmitted = req.Status

	err = db.DB("").C(collectionName).Update(bson.M{"_id": meeting.Id}, meeting)
	if err != nil {
		return nil, err
	}

	return &proto.SetMinutesSubmittedResponse{}, nil
}
