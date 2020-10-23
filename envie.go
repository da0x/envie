//
//   envie.go
//   envie
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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/da0x/envie/props"
)

var AutoPath = ".env"
var AutoPanic = true
var AutoVerbose = true

// Auto reads an struct of enironment variables from any of the following:
// 1. From a properties file from the current directory.
//        See AutoPath
// 2. From the operating system's environment variable.
// If this function fails, it will panic if AutoPanic is set to true.
//        See AutoPanic
func Auto(e interface{}) {
	err1 := UnmarshalFromEnvFile(AutoPath, e)
	err2 := UnmarshalFromEnv(e)
	if err1 != nil && err2 != nil {
		if AutoPanic {
			log.Fatalf("envie:\n%v\n%v\n", err1, err2)
		}
		if AutoVerbose {
			log.Printf("envie:\n%v\n%v\n", err1, err2)
		}
	}
	if AutoVerbose {
		log.Println("envie: settings:")
		log.Printf("%+v", e)
	}
}

// UnmarshalFromEnv reads an entire struct of env variables. Returns an error
// if any of those variables does not exist in the environment.
func UnmarshalFromEnv(e interface{}) error {
	t := reflect.TypeOf(e).Elem()
	v := reflect.ValueOf(e).Elem()
	errors := []string{}
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		env := tag.Get("envie")
		val := os.Getenv(env)
		v.Field(i).SetString(val)
		if len(val) <= 0 {
			errors = append(errors, env)
		}
	}
	if len(errors) != 0 {
		str := "environment variables not found:\n"
		for _, err := range errors {
			str += "\t" + err + "\n"
		}
		return fmt.Errorf(str)
	}
	return nil
}

// Properties loads an environment file and returns it as a map[string]string.
// This is the best option for a lightweight env reader.
func Properties(path string) (map[string]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	p, err := props.Read(bytes.NewBufferString(string(content[:])))
	if err != nil {
		return nil, err
	}
	o := make(map[string]string)
	names := p.Names()
	for _, key := range names {
		o[key] = p.Get(key)
	}
	return o, nil
}

// UnmarshalFromEnvFile attempts to read a struct from an existing env file. It
// will ignore any values not annotated as `envie="VAR_NAME"`. It returns an
// error if it fails.
func UnmarshalFromEnvFile(path string, e interface{}) error {
	props, err := Properties(path)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(e).Elem()
	v := reflect.ValueOf(e).Elem()
	errors := []string{}
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		env := tag.Get("envie")
		val, ok := props[env]
		if ok {
			v.Field(i).SetString(val)
			if len(val) <= 0 {
				errors = append(errors, env)
			}
		}
	}
	if len(errors) != 0 {
		str := "environment variable(s) not found:\n"
		for _, err := range errors {
			str += "\t" + err + "\n"
		}
		return fmt.Errorf(str)
	}
	return nil
}
