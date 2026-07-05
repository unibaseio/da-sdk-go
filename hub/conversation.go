package hub

import (
	"net/http"
	"strconv"

	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/gin-gonic/gin"
)

func (s *Server) addConversation(g *gin.RouterGroup) {
	g.Group("/").POST("/conversation", s.conversationByPost)
	g.Group("/").GET("/conversation", s.conversationByGet)

	g.Group("/").GET("/listConversation", s.listConversationByGet)
	g.Group("/").GET("/getConversation", s.getConversationByGet)
}

// conversationByPost godoc
//
//	@Summary		list conversations or get conversation messages by conversation id
//	@Description	list all conversations or get conversation messages by conversation id
//	@Tags			conversation
//	@Accept			json
//	@Produce		json
//	@Param			owner	formData	string	true	"owner/account address"
//	@Param			bucket	formData	string	false	"bucket name"  (empty means list all conversations of owner)
//	@Param			name	formData	string	false	"conversation ID (empty means list all conversation ids)"
//	@Param			offset	formData	int		false	"pagination offset" default(0)
//	@Param			length	formData	int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/conversation [post]
func (s *Server) conversationByPost(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.PostForm("name")
	if mn == "" {
		mn = c.PostForm("id")
	}
	addr, ok := ResolveOwnerForList(c, c.PostForm("owner"))
	if !ok {
		return
	}
	bucket := c.PostForm("bucket")

	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 1024
	}
	if mn == "" {
		res, err := s.listConversation(addr, bucket, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.getConversation(ctx, mn, addr, bucket, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

// conversationByGet godoc
//
//	@Summary		list conversations or get conversation messages by conversation id
//	@Description	list all conversations or get conversation messages by conversation id
//	@Tags			conversation
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	true	"owner/account address"
//	@Param			bucket	query		string	false	"bucket name" (empty means list all conversations of owner)
//	@Param			name	query		string	false	"conversation ID (empty means list all conversation ids)"
//	@Param			offset	query		int		false	"pagination offset" default(0)
//	@Param			length	query		int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/conversation [get]
func (s *Server) conversationByGet(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.Query("name")
	if mn == "" {
		mn = c.Query("id")
	}
	addr, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Query("bucket")
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 1024
	}
	if mn == "" {
		res, err := s.listConversation(addr, bucket, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.getConversation(ctx, mn, addr, bucket, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

// listConversationByGet godoc
//
//	@Summary		list conversations with details
//	@Description	list all conversations with detailed information for a specific owner and bucket
//	@Tags			conversation
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	true	"owner/account address"
//	@Param			bucket	query		string	false	"bucket name (empty means list all conversations of owner)"
//	@Param			offset	query		int		false	"pagination offset" default(0)
//	@Param			length	query		int		false	"number of items per page" default(32)
//	@Success		200		{object}	[]types.Conversation
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listConversation [get]
func (s *Server) listConversationByGet(c *gin.Context) {
	addr, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Query("bucket")
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 1024
	}
	res, err := s.listConversationDisplay(addr, bucket, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// getConversationByGet godoc
//
//	@Summary		get conversation details
//	@Description	get detailed information for a specific conversation
//	@Tags			conversation
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	true	"owner/account address"
//	@Param			bucket	query		string	false	"bucket name (empty means search in all buckets)"
//	@Param			name	query		string	true	"conversation ID"
//	@Success		200		{object}	[]types.Conversation
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/getConversation [get]
func (s *Server) getConversationByGet(c *gin.Context) {
	addr, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Query("bucket")
	mn := c.Query("name")
	res, err := s.getConversationDisplay(addr, bucket, mn)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, res)
}
