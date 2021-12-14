package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaxwu/him/conf/log"
	"github.com/jiaxwu/him/service/common"
)

// Recovery 异常恢复
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		log.WithField("err", err).Error("a panic captured")
		common.Failure(c, common.ErrCodeInternalError)
	})
}
