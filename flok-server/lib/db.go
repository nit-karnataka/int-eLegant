package lib

import (
	"errors"
	"time"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

func init() {
	bson.SetJSONTagFallback(true)
}

// ErrSessionNil means session is nil
var ErrSessionNil = errors.New("session is nil")

// Store is a connecton to MongoDB
type Store struct {
	mgoSession *mgo.Session
}

// GetMongoSession returns the copy of mongo session
func (s *Store) GetMongoSession() (*mgo.Session, error) {
	if s.mgoSession == nil {
		return nil, ErrSessionNil
	}
	return s.mgoSession.Clone(), nil
}

// Connect connects the mgoSession
func (s *Store) Connect(address, username, password, database string) error {
	dbConfig := &mgo.DialInfo{
		Addrs:    []string{address},
		Timeout:  600 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	sess, err := mgo.DialWithInfo(dbConfig)
	if err != nil {
		s.mgoSession = nil
		return err
	}

	if err := sess.Ping(); err != nil {
		s.mgoSession = nil
		return err
	}

	s.mgoSession = sess
	return nil
}

// Close the database connection
func (s *Store) Close() {
	if s.mgoSession != nil {
		s.mgoSession.Close()
	}
}
