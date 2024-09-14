package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	port = "4001"
	path = "http://localhost:4001"
)

var serverSingleton *Server

func getServer() *Server {
	if serverSingleton == nil {
		serverSingleton = NewServer(port)
		go serverSingleton.Start()
		time.Sleep(1000 * time.Millisecond)
	}
	return serverSingleton
}

func TestServerCreation(t *testing.T) {
	s := getServer()
	assert.NotNil(t, s)
}

type responseForm struct {
	Result string
	Error  string
}

func addBook(title, author string, duplicate bool) error {
	data := url.Values{
		"title":  []string{title},
		"author": []string{author},
	}

	resp, err := http.DefaultClient.PostForm(path+"/book", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status code not OK: %d", resp.StatusCode))
	}

	var rf responseForm

	err = json.NewDecoder(resp.Body).Decode(&rf)
	if err != nil {
		return err
	}

	if rf.Result != fmt.Sprintf("added book %s by %s", strings.ToLower(title), strings.ToLower(author)) && !duplicate {
		return errors.New(fmt.Sprintf("result message is incorrect: %s", rf.Result))
	}

	if rf.Result == fmt.Sprintf("this book is already in the library") && duplicate {
		return errors.New(rf.Result)
	}

	return nil
}

func TestAddBook(t *testing.T) {
	err := addBook("alice in wonderland", "JC", false)
	assert.Nil(t, err)
}

type testBook struct {
	Title  string
	Author string
}

func getBook(title, author string) (testBook, error) {
	title = url.QueryEscape(title)
	author = url.QueryEscape(author)

	resp, err := http.DefaultClient.Get(fmt.Sprintf("%s/book?title=%s&author=%s", path, title, author))
	if err != nil {
		return testBook{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return testBook{}, errors.New(fmt.Sprintf("invalid status code %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return testBook{}, err
	}

	var result testBook
	err = json.Unmarshal(body, &result)
	if err != nil {
		return testBook{}, err
	}

	return result, nil
}

func TestGetBook(t *testing.T) {
	book, err := getBook("alice in wonderland", "JC")
	assert.Nil(t, err)
	assert.Equal(t, "alice in wonderland", book.Title)
	assert.Equal(t, "jc", book.Author)
}
