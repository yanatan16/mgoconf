mgoconf
=======

Small configuration file reading for mgo

LICENSE
-------

Copyright (c) 2012 Jon Eisen

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

API
---

### PACKAGE

		package mgoconf

Package mgoconf makes configuring mgo, the go mongo driver written by
Canonical, easy to configure using a json file.

### TYPES

		type Config struct {
		    // Info for dialing
		    Conn *mgo.DialInfo

		    // Safety Characteristics
		    Safety *mgo.Safe
		}

A configuration for the Mongo connections

		func New() *Config

Create a new mgo configuration

		func Read(fn string) (*Config, error)

Read the mgo configuration from a json file.

		func (cfg *Config) Connect() (*mgo.Session, error)

Connect to the mongo instance.

Example
-------

```json
{
	"Conn": {
		"Database" : "ua",
		"Addrs": [
			"localhost:27017"
		],
		"Direct": false,

		"_TimeoutComment": "nanoseconds",
		"Timeout": 100000,

		"_commented": {
			"_comment": "Use these for authentication",

			"Username": "something",
			"Password": "something"
		}
	},

	"Safety": { 
		"_reference": "http://go.pkgdoc.org/labix.org/v2/mgo#Session.SetSafe",
		"W": 1,
		"Fsync": false
	}
}
```

```go
package main

import (
	"github.com/yanatan16/mgoconf"
)

func main() {
	cfg, err := mgoconf.Read("config.json")
	if err != nil {
		panic(err)
	}

	session, err := cfg.Connect()
	if err != nil {
		panic(err)
	}

	session.DB("test").C("test").find() // etc..
}
```
