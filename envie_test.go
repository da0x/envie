//
//   envie_test.go
//   olog
//
//   Copyright 2020 Daher Alfawares
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2
//
//   Unless required by applicable law or agreed to in writing,
//   distributed under the License is distributed on an "AS IS" BASIS
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
//   See the License for the specific language governing permissions
//   limitations under the License

package envie

import (
	"os"
	"testing"
)

type entity struct {
	V1 string `envie:"TEST_VARIABLE_1"`
	V2 string `envie:"TEST_VARIABLE_2"`
}

func TestUnmarshalFromEnvFile(t *testing.T) {
	hello := "hello"
	world := "world"
	var e entity
	err := UnmarshalFromEnvFile(".env", &e)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if e.V1 != hello {
		t.Errorf("envie: incorrect environment variable:\nexpected:%v\nfound:%v", hello, e.V1)
	}
	if e.V2 != world {
		t.Errorf("envie: incorrect environment variable:\nexpected:%v\nfound:%v", world, e.V2)
	}
}

func TestAuto(t *testing.T) {
	hello := "Hello"
	world := "world"
	os.Setenv("TEST_VARIABLE_1", hello)
	var e entity
	Auto(&e)
	if e.V1 != hello {
		t.Errorf("envie: incorrect environment variable:\nexpected:%v\nfound:%v", hello, e.V1)
	}
	if e.V2 != world {
		t.Errorf("envie: incorrect environment variable:\nexpected:%v\nfound:%v", world, e.V2)
	}
}

func TestUnmarshalFromEnv(t *testing.T) {
	hello := "hello"
	world := "world"
	os.Setenv("TEST_VARIABLE_1", hello)
	os.Setenv("TEST_VARIABLE_2", world)
	var e entity
	err := UnmarshalFromEnv(&e)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if e.V1 != hello {
		t.Errorf("envie: incorrect environment variable:\nexpected:%v\nfound:%v", hello, e.V1)
	}
	if e.V2 != world {
		t.Errorf("envie: incorrect environment variable:\nexpected:%v\nfound:%v", world, e.V2)
	}
}
