// hub-mock is a tiny standalone server that mounts ONLY the three /api
// middlewares (MaxBodySize, AuthMiddleware, RateLimit) plus dummy handlers,
// so you can exercise the auth/rate/size rejection paths with curl without
// running a real hub (no chain, no LogFS, no SQLite).
//
// Usage:
//
//	CHAIN_TYPE=bnb-testnet-dao go run ./cmd/hub-mock
//	# default listen 127.0.0.1:8086, override with -addr
//
// Then curl as you wish, see auth_test.go for canned payloads.
package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unibaseio/da-sdk-go/hub"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8086", "listen address")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	g := r.Group("/api")
	g.Use(hub.MaxBodySize())
	g.Use(hub.AuthMiddleware())
	g.Use(hub.RateLimit())

	// bypassed
	g.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "endpoint": "info"})
	})

	// write-style endpoints: owner must equal signer
	upload := func(c *gin.Context) {
		var body struct {
			Owner   string `json:"owner"`
			ID      string `json:"id"`
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(599, gin.H{"err": err.Error()})
			return
		}
		if !hub.RequireOwnerMatch(c, body.Owner) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "endpoint": "upload", "owner": body.Owner, "id": body.ID})
	}
	g.POST("/upload", upload)

	uploadData := func(c *gin.Context) {
		owner := c.PostForm("owner")
		if !hub.RequireOwnerMatch(c, owner) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "endpoint": "uploadData", "owner": owner})
	}
	g.POST("/uploadData", uploadData)

	// read-style endpoints: owner defaults to signer when empty
	download := func(c *gin.Context) {
		o := c.PostForm("owner")
		if o == "" {
			o = c.Query("owner")
		}
		owner, ok := hub.ResolveOwnerForList(c, o)
		if !ok {
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "endpoint": "download", "owner": owner})
	}
	g.POST("/download", download)
	g.GET("/download", download)

	r.GET("/_health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok\n")
	})

	if err := r.Run(*addr); err != nil {
		panic(err)
	}
}
