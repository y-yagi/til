package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Email string
}

func EncodeByGob() bytes.Buffer {
	buf := bytes.NewBuffer(nil)

	err := gob.NewEncoder(buf).Encode(buildUsers())
	if err != nil {
		panic(err)
	}

	return *buf
}

func DecodeByGob(buf bytes.Buffer) []User {
	users := []User{}
	err := gob.NewDecoder(&buf).Decode(&users)
	if err != nil {
		panic(err)
	}

	return users
}

func MarshalByJSON() []byte {
	result, err := json.Marshal(buildUsers())
	if err != nil {
		panic(err)
	}

	return result
}

func UnmarshalByJSON(buf []byte) []User {
	users := []User{}
	err := json.Unmarshal(buf, &users)
	if err != nil {
		panic(err)
	}

	return users
}

func buildUsers() []User {
	users := []User{}

	for i := 0; i < 1000; i++ {
		users = append(users, User{Name: fmt.Sprintf("Name%d", i), Email: fmt.Sprintf("Email%d@example.com", i)})
	}
	return users
}
