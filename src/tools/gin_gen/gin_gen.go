package gin_gen

import (
	"tgsms/internal/entity/global"

	"github.com/gin-gonic/gin"
)

func HandleError(gctx *gin.Context, err error) {
	gctx.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
}
