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
// 1. Attempt to read env variables from the system.
// 2. If the above fails, fall back to reading env from file.
//        See AutoPath
// If this function fails, it will panic if AutoPanic is set to true.
//        See AutoPanic
func Auto(e interface{}) {
	err := UnmarshalFromSystem(e)
	if err != nil {
		if AutoVerbose {
			log.Printf("envie: WARNING: failed to read environment variables from system\n")
			log.Printf("envie: fallback to file: %v\n", AutoPath)
		}
		err = UnmarshalFromFile(AutoPath, e)
		if err != nil {
			log.Printf("envie: failed to read environment variables:\n")
			empty := empty(e)
			for _, v := range empty {
				fmt.Printf("\t%v", v)
			}
			if AutoPanic {
				panic("envie: panic.")
			}
		}
	}
}

// UnmarshalFromSystem reads an entire struct of env variables. Returns an error
// if any of those variables does not exist in the environment.
func UnmarshalFromSystem(e interface{}) error {
	t := reflect.TypeOf(e).Elem()
	v := reflect.ValueOf(e).Elem()
	errors := []string{}
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		env := tag.Get("envie")
		val := os.Getenv(env)
		if len(val) <= 0 {
			errors = append(errors, env)
			continue
		}
		if AutoVerbose {
			log.Printf("export %v=%v", env, val)
		}
		v.Field(i).SetString(val)
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

// UnmarshalFromFile attempts to read a struct from an existing env file. It
// will ignore any values not annotated as `envie="VAR_NAME"`. It returns an
// error if it fails.
func UnmarshalFromFile(path string, e interface{}) error {
	if AutoVerbose {
		log.Printf("envie: unmarshaling env from %v", path)
	}
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
		if !ok || len(val) == 0 {
			errors = append(errors, env)
			continue
		}
		if AutoVerbose {
			log.Printf("export %v=%v", env, val)
		}
		v.Field(i).SetString(val)
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

// Empty returns a slice of empty environment variables.
func empty(e interface{}) []string {
	v := reflect.ValueOf(e).Elem()
	t := reflect.TypeOf(e).Elem()
	o := []string{}
	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag
		env := tag.Get("envie")
		if len(v.Field(i).String()) == 0 {
			o = append(o, env)
		}
	}
	return o
}
