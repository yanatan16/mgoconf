package mgoconf

import (
	"labix.org/v2/mgo"
	"io/ioutil"
	"encoding/json"
)

// A configuration for the Mongo connections
type MongoConfig struct {
	// Info for dialing
	Conn *mgo.DialInfo

	// Safety Characteristics
	Safety *mgo.Safe
}

// Create a new mongo configuration
func New() (*MongoConfig) {
	return &MongoConfig{
		Conn: &mgo.DialInfo{	
			Database     : "test",
			Addrs  : []string{"localhost"},
			Direct : true,
		},
		Safety : &mgo.Safe{},
	}
}

func Read(fn string) (*MongoConfig, error) {
	cfg := New()

	file, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *MongoConfig) Dial() (*mgo.Session, error) {
	sess, err := mgo.DialWithInfo(cfg.Conn)
	if err != nil {
		return nil, err
	}

	// Set Safety Parameters
	sess.SetSafe(cfg.Safety)

	return sess, nil
}