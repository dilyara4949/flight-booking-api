package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	pageDefault     = 1
	pageSizeDefault = 30
)

func GetPageInfo(c *gin.Context) (page int, pageSize int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = pageDefault
	}

	pageSize, err = strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = pageSizeDefault
	}
	return page, pageSize
}
