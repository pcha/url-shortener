package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	switch os.Args[1] {
	case "encode":
		res, err := encode(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "decode":
		res, err := encode(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "send":
		body, err := encode(os.Args[2])
		if err != nil {
			panic(err)
		}
		resp, err := http.Post("http://localhost:8080/", "application/x-msgpack", bytes.NewReader(body))
		if err != nil {
			panic(err)
		}
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		decRes, err := decode(string(result))
		if err != nil {
			panic(err)
		}
		fmt.Println(string(decRes))
	}
}

func encode(msg string) ([]byte, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return msgpack.Marshal(m)
}

func decode(msg string) ([]byte, error) {
	m := make(map[string]interface{})
	err := msgpack.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return json.Marshal(m)
}
