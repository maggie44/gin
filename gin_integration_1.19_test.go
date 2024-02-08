// Copyright 2023 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build go1.19

package gin

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunQUIC(t *testing.T) {
	router := New()

	certPath := "./testdata/certificate/cert.pem"
	keyPath := "./testdata/certificate/key.pem"

	path := "/example"
	port := ":8443"

	address := "https://localhost" + port + path

	responseBody := "it worked"
	responseStatusCode := http.StatusOK

	go func() {
		router.GET("/example", func(c *Context) { c.String(responseStatusCode, responseBody) })

		assert.NoError(t, router.RunQUIC(port, certPath, keyPath))
	}()

	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

	assert.Error(t, router.RunQUIC(port, certPath, keyPath))
	testQuicRequest(t, responseBody, responseStatusCode, address, certPath)
}
