package handler

import (
	"context"
	"flok-server/lib"
	proto "flok-server/user-srv/userproto"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "users"

// UserServiceHandler implements the HouseService interface of proto
type UserServiceHandler struct {
	Store *lib.Store
}

/* // Create imlements proto
func (h *UserServiceHandler) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	req.Profile.Id = bson.NewObjectId().Hex()
	req.Profile.Created = time.Now().Unix()
	req.Profile.Updated = time.Now().Unix()

	err = db.DB("").C(collectionName).Insert(req.Profile)
	if err != nil {
		return nil, err
	}

	log.Printf("New profile %+v created", req.Profile)

	return &proto.CreateResponse{}, nil
}

// Read imlements proto
func (h *UserServiceHandler) Read(ctx context.Context, req *proto.ReadRequest) (*proto.ReadResponse, error) {
	p := &proto.Profile{}
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(p)
	if err != nil {
		return nil, err
	}

	return &proto.ReadResponse{
		Profile: p,
	}, nil
}

// Delete imlements proto
func (h *UserServiceHandler) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	db, err := h.Store.GetMongoSession()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Remove(bson.M{"_id": req.Id})
	if err != nil {
		return nil, err
	}

	return &proto.DeleteResponse{}, nil
}

// Update imlements proto
func (h *UserServiceHandler) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	db, err := h.Store.GetMongoSession()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	t := req.Profile
	p := &proto.Profile{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": t.Id}).One(p)
	if err != nil {
		return nil, err
	}

	if t.Address != "" {
		p.Address = t.Address
	}

	if t.DisplayName != "" {
		p.DisplayName = t.DisplayName
	}

	if t.Email != "" {
		p.Email = t.Email
	}

	if t.Name != "" {
		p.Name = t.Name
	}

	if t.PhoneNumber != "" {
		p.PhoneNumber = t.PhoneNumber
	}

	err = db.DB("").C(collectionName).Update(bson.M{"_id": t.Id}, p)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateResponse{}, nil
} */

// CreateStudent implements proto
func (h *UserServiceHandler) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Student)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadStudent implements proto
func (h *UserServiceHandler) ReadStudent(ctx context.Context, req *proto.ReadRequest) (*proto.ReadStudentResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	s := &proto.Student{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(s)
	if err != nil {
		return nil, err
	}

	return &proto.ReadStudentResponse{
		Student: s,
	}, nil
}

// UpdateStudent implements proto
func (h *UserServiceHandler) UpdateStudent(ctx context.Context, req *proto.UpdateStudentRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeleteStudent implements proto
func (h *UserServiceHandler) DeleteStudent(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

// CreateDOR implements proto
func (h *UserServiceHandler) CreateDOR(ctx context.Context, req *proto.CreateDORRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Dor)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadDOR implements proto
func (h *UserServiceHandler) ReadDOR(ctx context.Context, req *proto.ReadRequest) (*proto.ReadDORResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	d := &proto.DOR{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(d)
	if err != nil {
		return nil, err
	}

	return &proto.ReadDORResponse{
		Dor: d,
	}, nil
}

// UpdateDOR implements proto
func (h *UserServiceHandler) UpdateDOR(ctx context.Context, req *proto.UpdateDORRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeleteDOR implements proto
func (h *UserServiceHandler) DeleteDOR(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

// CreateSecretariat implements proto
func (h *UserServiceHandler) CreateSecretariat(ctx context.Context, req *proto.CreateSecretariatRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Sec)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadSecretariat implements proto
func (h *UserServiceHandler) ReadSecretariat(ctx context.Context, req *proto.ReadRequest) (*proto.ReadSecretariatResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	s := &proto.Secretariat{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(s)
	if err != nil {
		return nil, err
	}

	return &proto.ReadSecretariatResponse{
		Sec: s,
	}, nil
}

// UpdateSecretariat implements proto
func (h *UserServiceHandler) UpdateSecretariat(ctx context.Context, req *proto.UpdateSecretariatRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeleteSecretariat implements proto
func (h *UserServiceHandler) DeleteSecretariat(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

// CreateEC implements proto
func (h *UserServiceHandler) CreateEC(ctx context.Context, req *proto.CreateECRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Ec)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadEC implements proto
func (h *UserServiceHandler) ReadEC(ctx context.Context, req *proto.ReadRequest) (*proto.ReadECResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ec := &proto.EC{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(ec)
	if err != nil {
		return nil, err
	}

	return &proto.ReadECResponse{
		Ec: ec,
	}, nil
}

// UpdateEC implements proto
func (h *UserServiceHandler) UpdateEC(ctx context.Context, req *proto.UpdateECRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeleteEC implements proto
func (h *UserServiceHandler) DeleteEC(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

// GetAllEC implements proto
func (h *UserServiceHandler) GetAllEC(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllECResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ecs := []*proto.EC{}

	err = db.DB("").C(collectionName).Find(nil).All(ecs)
	if err != nil {
		return nil, err
	}

	return &proto.GetAllECResponse{
		Ecs: ecs,
	}, nil
}

// CreatePI implements proto
func (h *UserServiceHandler) CreatePI(ctx context.Context, req *proto.CreatePIRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Pi)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadPI implements proto
func (h *UserServiceHandler) ReadPI(ctx context.Context, req *proto.ReadRequest) (*proto.ReadPIResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	pi := &proto.PI{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(pi)
	if err != nil {
		return nil, err
	}

	return &proto.ReadPIResponse{
		Pi: pi,
	}, nil
}

// UpdatePI implements proto
func (h *UserServiceHandler) UpdatePI(ctx context.Context, req *proto.UpdatePIRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeletePI implements proto
func (h *UserServiceHandler) DeletePI(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

// CreateCoI implements proto
func (h *UserServiceHandler) CreateCoI(ctx context.Context, req *proto.CreateCoIRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Coi)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadCoI implements proto
func (h *UserServiceHandler) ReadCoI(ctx context.Context, req *proto.ReadRequest) (*proto.ReadCoIResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	c := &proto.CoI{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(c)
	if err != nil {
		return nil, err
	}

	return &proto.ReadCoIResponse{
		Coi: c,
	}, nil
}

// UpdateCoI implements proto
func (h *UserServiceHandler) UpdateCoI(ctx context.Context, req *proto.UpdateCoIRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeleteCoI implements proto
func (h *UserServiceHandler) DeleteCoI(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

// CreateHR implements proto
func (h *UserServiceHandler) CreateHR(ctx context.Context, req *proto.CreateHRRequest) (*proto.CreateResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.DB("").C(collectionName).Insert(req.Hr)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{}, nil
}

// ReadHR implements proto
func (h *UserServiceHandler) ReadHR(ctx context.Context, req *proto.ReadRequest) (*proto.ReadHRResponse, error) {
	db, err := h.Store.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	hr := &proto.HR{}

	err = db.DB("").C(collectionName).Find(bson.M{"_id": req.Id}).One(hr)
	if err != nil {
		return nil, err
	}

	return &proto.ReadHRResponse{
		Hr: hr,
	}, nil
}

// UpdateHR implements proto
func (h *UserServiceHandler) UpdateHR(ctx context.Context, req *proto.UpdateHRRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

// DeleteHR implements proto
func (h *UserServiceHandler) DeleteHR(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
