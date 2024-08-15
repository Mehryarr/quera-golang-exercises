package main

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	port = "4001"
	path = "http://127.0.0.1:4001"
)

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

var testServer *Server

func getServer() *Server {
	if testServer == nil {
		testServer = NewServer(port)
		go testServer.Start()
	}
	time.Sleep(100 * time.Millisecond)
	return testServer
}

func TestSampleCreation(t *testing.T) {
	s := getServer()
	assert.NotNil(t, s)
}

func TestSampleServerStart(t *testing.T) {
	getServer()
	conn, err := net.Dial("tcp", "localhost:"+port)
	assert.Nil(t, err)
	defer conn.Close()
}

func TestSampleAddHandler(t *testing.T) {
	getServer()
	resp, err := http.DefaultClient.Get(path + "/add?numbers=1,2,3")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
}

func TestSampleSubtraction(t *testing.T) {
	getServer()
	resp, err := http.DefaultClient.Get(path + "/sub?numbers=104,204")
	assert.Nil(t, err)
	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	var response Response
	err = json.Unmarshal(s, &response)
	assert.Nil(t, err)
	assert.Equal(t, "The result of your query is: -100", response.Result)
}
