//go:build e2e
// +build e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComments(t *testing.T) {
	client := resty.New()

	resp, err := client.R().Get(BASE_URL + "/api/comments")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()

	resp, err := client.R().SetBody(`{"slug": "/","body": "This is API","author": "J.K"}`).Post(BASE_URL + "/api/comment")

	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPutComment(t *testing.T) {
	client := resty.New()

	resp, err := client.R().SetBody(`{"slug":"/new","body":"nobody cares","author": "Khabib"}`).Put(BASE_URL + "/api/comment/9")
	assert.NoError(t, err)

	assert.Equal(t, resp.StatusCode(), 200)
}

func TestDeleteComment(t *testing.T) {
	client := resty.New()

	resp, err := client.R().Delete(BASE_URL + "/api/comment/1")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
