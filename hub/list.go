package hub

import (
	"net/http"
	"strconv"

	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/gin-gonic/gin"
)

func (s *Server) addList(g *gin.RouterGroup) {
	g.Group("/").POST("/listAccount", s.listAccountByPost)
	g.Group("/").GET("/listAccount", s.listAccountByGet)

	g.Group("/").POST("/listBucket", s.listBucketByPost)
	g.Group("/").GET("/listBucket", s.listBucketByGet)

	g.Group("/").POST("/listNeedle", s.listNeedleByPost)
	g.Group("/").GET("/listNeedle", s.listNeedleByGet)

	g.Group("/").POST("/listVolume", s.listVolumeByPost)
	g.Group("/").GET("/listVolume", s.listVolumeByGet)

	g.Group("/").GET("/getAccount", s.getAccountByGet)
	g.Group("/").GET("/getBucket", s.getBucketByGet)
	g.Group("/").GET("/getNeedle", s.getNeedleByGet)
	g.Group("/").GET("/getVolume", s.getVolumeByGet)
}

// getAccountByGet godoc
//
//	@Summary		get account information
//	@Description	get account information for a specific owner
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"account owner address"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/getAccount [get]
func (s *Server) getAccountByGet(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	res, err := s.getAccount(owner)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// listAccountByGet godoc
//
//	@Summary		list accounts
//	@Description	get a list of accounts with pagination support
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int		false	"pagination offset" default(0)
//	@Param			length	query		int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listAccount [get]
func (s *Server) listAccountByGet(c *gin.Context) {
	if !RequireAuthenticated(c) { // account registry = global enumeration, not anonymous
		return
	}
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listAccount(offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// listAccountByPost godoc
//
//	@Summary		list accounts
//	@Description	get a list of accounts with pagination support
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			offset	formData	int		false	"pagination offset" default(0)
//	@Param			length	formData	int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listAccount [post]
func (s *Server) listAccountByPost(c *gin.Context) {
	if !RequireAuthenticated(c) { // account registry = global enumeration, not anonymous
		return
	}
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listAccount(offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// getBucketByGet godoc
//
//	@Summary		Get bucket information
//	@Description	Retrieve detailed information for a specific bucket including metadata, storage statistics, and configuration settings. Returns bucket display information with owner details, creation time, and current status.
//	@Tags			bucket
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"Account owner address (wallet address or account identifier)"
//	@Param			bucket	query		string	true	"Bucket name (unique identifier within the owner's account)"
//	@Success		200		{object}	map[string]interface{}	"Bucket information including metadata, statistics, and configuration"
//	@Failure		400		{object}	lerror.APIError			"Bad request - missing or invalid parameters"
//	@Failure		404		{object}	lerror.APIError			"Bucket not found"
//	@Failure		599		{object}	lerror.APIError			"Internal server error"
//	@Router			/api/getBucket [get]
func (s *Server) getBucketByGet(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Query("bucket")
	if bucket == "" {
		bucket = c.Query("name")
	}

	// result is []types.BucketDisplay containing bucket metadata, description, state, transport, type, and last update time
	res, err := s.getBucket(owner, bucket)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// listBucketByGet godoc
//
//	@Summary		List buckets with pagination
//	@Description	Retrieve a paginated list of buckets. If owner is specified, returns buckets owned by that account. If owner is empty, returns all buckets in the system. Supports pagination with configurable offset and page size. Returns bucket display information including names, creation times, and basic statistics.
//	@Tags			bucket
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"Owner address to filter buckets (leave empty to list all buckets)" default("")
//	@Param			offset	query		int		false	"Pagination offset - number of items to skip" default(0) minimum(0)
//	@Param			length	query		int		false	"Number of items per page (max 100)" default(32) minimum(1) maximum(100)
//	@Success		200		{object}	map[string]interface{}	"Paginated list of buckets with metadata and statistics"
//	@Failure		400		{object}	lerror.APIError			"Bad request - invalid pagination parameters"
//	@Failure		599		{object}	lerror.APIError			"Internal server error"
//	@Router			/api/listBucket [get]
func (s *Server) listBucketByGet(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	// kind: optional scenario filter (memory/model/dataset); empty = all kinds.
	res, err := s.listBucket(owner, c.Query("kind"), offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// listBucketByPost godoc
//
//	@Summary		List buckets with pagination (POST)
//	@Description	Retrieve a paginated list of buckets using POST method. If owner is specified, returns buckets owned by that account. If owner is empty, returns all buckets in the system. Supports pagination with configurable offset and page size. Returns bucket display information including names, creation times, and basic statistics. This POST variant is useful for complex filtering scenarios or when URL length limits are a concern.
//	@Tags			bucket
//	@Accept			application/x-www-form-urlencoded
//	@Produce		json
//	@Param			owner	formData	string	false	"Owner address to filter buckets (leave empty to list all buckets)" default("")
//	@Param			offset	formData	int		false	"Pagination offset - number of items to skip" default(0) minimum(0)
//	@Param			length	formData	int		false	"Number of items per page (max 100)" default(32) minimum(1) maximum(100)
//	@Success		200		{object}	map[string]interface{}	"Paginated list of buckets with metadata and statistics"
//	@Failure		400		{object}	lerror.APIError			"Bad request - invalid pagination parameters"
//	@Failure		599		{object}	lerror.APIError			"Internal server error"
//	@Router			/api/listBucket [post]
func (s *Server) listBucketByPost(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.PostForm("owner"))
	if !ok {
		return
	}
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	// kind: optional scenario filter (memory/model/dataset); empty = all kinds.
	res, err := s.listBucket(owner, c.PostForm("kind"), offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// getNeedleByGet godoc
//
//	@Summary		get needle information
//	@Description	get needle information for a specific owner
//	@Tags			needle
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"owner address" default("")
//	@Param			bucket	query		string	false	"bucket name" default("")
//	@Param			name	query		string	false	"needle name" default("")
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/getNeedle [get]
func (s *Server) getNeedleByGet(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Query("bucket")
	name := c.Query("name")
	res, err := s.getNeedleDisplay(owner, bucket, name)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// listNeedleByGet godoc
//
//	@Summary		list needles
//	@Description	get a list of needles for a specific owner with pagination support
//	@Tags			needle
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"owner address" default("")
//	@Param			bucket	query		string	false	"bucket name" default("")
//	@Param			offset	query		int		false	"pagination offset" default(0)
//	@Param			length	query		int		false		"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listNeedle [get]
func (s *Server) listNeedleByGet(c *gin.Context) {
	addr, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Query("bucket")
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}

	conversation := c.Query("conversation")
	if conversation != "" {
		res, err := s.listNeedleDisplayByConversation(addr, bucket, conversation, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.listNeedleDisplay(addr, bucket, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}

}

// listNeedleByPost godoc
//
//	@Summary		list needles
//	@Description	get a list of needles for a specific owner with pagination support
//	@Tags			needle
//	@Accept			json
//	@Produce		json
//	@Param			owner	formData	string	false	"owner address" default("")
//	@Param			bucket	formData	string	false	"bucket name" default("")
//	@Param			offset	formData	int		false	"pagination offset" default(0)
//	@Param			length	formData	int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listNeedle [post]
func (s *Server) listNeedleByPost(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.PostForm("owner"))
	if !ok {
		return
	}
	bucket := c.PostForm("bucket")
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	conversation := c.PostForm("conversation")
	if conversation != "" {
		res, err := s.listNeedleDisplayByConversation(owner, bucket, conversation, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.listNeedleDisplay(owner, bucket, offset, length)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

// listVolumeByGet godoc
//
//	@Summary		list volumes
//	@Description	get a list of volumes for a specific owner with pagination support
//	@Tags			volume
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"owner address"
//	@Param			offset	query		int		false	"pagination offset" default(0)
//	@Param			length	query		int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listVolume [get]
func (s *Server) listVolumeByGet(c *gin.Context) {
	addr, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listVolume(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// listVolumeByPost godoc
//
//	@Summary		list volumes
//	@Description	get a list of volumes for a specific owner with pagination support
//	@Tags			volume
//	@Accept			json
//	@Produce		json
//	@Param			owner	formData	string	false	"owner address"
//	@Param			offset	formData	int		false	"pagination offset" default(0)
//	@Param			length	formData	int		false	"number of items per page" default(32)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/listVolume [post]
func (s *Server) listVolumeByPost(c *gin.Context) {
	addr, ok := RequireOwnerForList(c, c.PostForm("owner"))
	if !ok {
		return
	}
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listVolume(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// getVolumeByGet godoc
//
//	@Summary		get volume information
//	@Description	get volume information for a specific owner
//	@Tags			volume
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"owner address"
//	@Param			file	query		int		false	"file name"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/getVolume [get]
func (s *Server) getVolumeByGet(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	fid, _ := strconv.Atoi(c.Query("file"))
	res, err := s.getVolume(owner, uint64(fid))
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
