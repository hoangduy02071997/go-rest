package media

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest-api/common"
	"time"
)

func UploadHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dst := fmt.Sprintf("static/%d_%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternalServer(err))
		}

		img := common.Media{
			Id:     1,
			Url:    dst,
			Width:  200,
			Height: 200,
			Type:   "image",
			From:   "local",
			Ext:    "jpeg",
		}

		img.FullFill("http://localhost:8080")
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
