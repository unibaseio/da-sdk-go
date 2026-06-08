package hub

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/unibaseio/da-sdk-go/build"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/gin-gonic/gin"
)

func (s *Server) addDownload(g *gin.RouterGroup) {
	g.Group("/").POST("/download", s.downloadByPOST)
	g.Group("/").GET("/download", s.downloadByGET)
}

func (s *Server) downloadByGET(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.Query("name")
	if mn == "" {
		mn = c.Query("id")
	}
	addr, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}

	head := fmt.Sprintf("attachment; filename=\"%s\"", mn)
	extraHeaders := map[string]string{
		"Content-Disposition": head,
	}

	var w bytes.Buffer
	size, err := s.download(ctx, mn, addr, &w)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.DataFromReader(http.StatusOK, size, "text/plain", &w, extraHeaders)
}

func (s *Server) downloadByPOST(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.PostForm("name")
	if mn == "" {
		mn = c.PostForm("id")
	}
	addr, ok := ResolveOwnerForList(c, c.PostForm("owner"))
	if !ok {
		return
	}

	head := fmt.Sprintf("attachment; filename=\"%s\"", mn)
	extraHeaders := map[string]string{
		"Content-Disposition": head,
	}

	var w bytes.Buffer
	size, err := s.download(ctx, mn, addr, &w)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.DataFromReader(http.StatusOK, size, "text/plain", &w, extraHeaders)
}

func (s *Server) download(ctx context.Context, name, owner string, w io.Writer) (int64, error) {
	size, err := s.logFSRead(owner, name, w)
	if err == nil {
		return size, nil
	}

	fr, err := sdk.GetFileReceipt(build.ServerURL, s.auth, name)
	if err == nil {
		err = sdk.Download(build.ServerURL, s.auth, name, s.ps, w)
		if err == nil {
			return fr.Size, nil
		}
	}

	pr, err := sdk.GetPieceReceipt(build.ServerURL, s.auth, name)
	if err == nil {
		err = sdk.DownloadPieceAndSave(build.ServerURL, s.auth, name, s.ps)
		if err != nil {
			return 0, err
		}

		_, err := s.ps.GetPiece(ctx, name, w, types.Options{})
		if err == nil {
			return pr.Size, nil
		}
	}

	res, err := s.ps.GetReplica(ctx, name, w, types.Options{})
	if err == nil {
		return res.Size, nil
	}

	return 0, fmt.Errorf("no such file: %s at %s", name, owner)
}
