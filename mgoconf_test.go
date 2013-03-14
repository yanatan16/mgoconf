package mgoconf

import (
	"testing"
)

func TestReadAndConnect(t *testing.T) {
	cfg, err := Read("config.json")
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Safety.W != 1 {
		t.Error("config was not read correctly")
	}

	sess, err := cfg.Dial()
	if err != nil {
		t.Fatal(err, cfg.Conn)
	}

	err = sess.DB("test").C("test").Insert(map[string]int{"key": 1, "val": 2})
	if err != nil {
		t.Fatal(err)
	}

	m := make(map[string]int)
	sess.DB("test").C("test").Find(map[string]int{"key": 1}).One(&m)

	if m["val"] != 2 {
		t.Error("Expected m[val]=2", m)
	}
}