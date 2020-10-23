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
	V1 string `envie:"951f83e8-c682-405f-9b49-49b498a41613"`
	V2 string `envie:"99217878-c9c3-4eaf-90ec-2b54d4da396b"`
}

func TestUnmarshalFromEnv(t *testing.T) {
	hello := "hello"
	world := "world"
	os.Setenv("951f83e8-c682-405f-9b49-49b498a41613", hello)
	os.Setenv("99217878-c9c3-4eaf-90ec-2b54d4da396b", world)
	var e entity
	err := UnmarshalFromEnv(&e)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if e.V1 != hello || e.V2 != world {
		t.Errorf("envie: incorrect environment variables")
	}
}

func TestUnmarshalFromEnvFile(t *testing.T) {
	hello := "hello"
	world := "world"
	os.Setenv("951f83e8-c682-405f-9b49-49b498a41613", hello)
	os.Setenv("99217878-c9c3-4eaf-90ec-2b54d4da396b", world)
	var e entity
	err := UnmarshalFromEnvFile(".env", &e)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if e.V1 != hello || e.V2 != world {
		t.Errorf("envie: incorrect environment variables")
	}
}

func TestAuto(t *testing.T) {
	hello := "Hello"
	world := "world"
	os.Setenv("951f83e8-c682-405f-9b49-49b498a41613", hello)
	var e entity
	Auto(&e)
	if e.V1 != hello || e.V2 != world {
		t.Errorf("envie: incorrect environment variables")
	}
}
