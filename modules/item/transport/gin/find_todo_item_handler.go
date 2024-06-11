package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest-api/common"
	"rest-api/modules/item/business"
	"rest-api/modules/item/storage/mysql"
	"strconv"
)

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get Id parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := mysql.NewSqlStorage(db)
		biz := business.NewGetTodoItemBusiness(store)

		item, err := biz.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		// c.JSON(http.StatusOK, gin.H{"data": todoItem}) ~~> Chưa chuẩn hoá Response
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(item))
	}
}
