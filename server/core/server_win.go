//go:build windows
// +build windows

package core

import (
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	router.MaxMultipartMemory = math.MaxInt64
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 30,
	}

}
