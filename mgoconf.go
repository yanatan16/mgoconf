// Package mgoconf makes configuring mgo, the go mongo driver written by Canonical, easy to configure using a json file.
package mgoconf

import (
	"encoding/json"
	"io/ioutil"
	"labix.org/v2/mgo"
)

// A configuration for the Mongo connections
type Config struct {
	// Info for dialing
	Conn *mgo.DialInfo

	// Safety Characteristics
	Safety *mgo.Safe
}

// Create a new mongo configuration
func New() *Config {
	return &Config{
		Conn: &mgo.DialInfo{
			Database: "test",
			Addrs:    []string{"localhost"},
			Direct:   true,
		},
		Safety: &mgo.Safe{},
	}
}

func Read(fn string) (*Config, error) {
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

func (cfg *Config) Connect() (*mgo.Session, error) {
	sess, err := mgo.DialWithInfo(cfg.Conn)
	if err != nil {
		return nil, err
	}

	// Set Safety Parameters
	sess.SetSafe(cfg.Safety)

	return sess, nil
}
