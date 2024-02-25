//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthEndPoint(t *testing.T) {
	fmt.Println("running e2e test for health check endpoint")

	client := resty.New()

	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fatal()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
