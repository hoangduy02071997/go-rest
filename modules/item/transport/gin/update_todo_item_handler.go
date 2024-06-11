package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest-api/common"
	"rest-api/modules/item/business"
	"rest-api/modules/item/model"
	"rest-api/modules/item/storage/mysql"
	"strconv"
)

func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var item model.TodoItemUpdate
		// Get Id parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := mysql.NewSqlStorage(db)
		biz := business.NewUpdateTodoItemBusiness(store)
		if err := biz.UpdateItemById(c.Request.Context(), &item, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//c.JSON(http.StatusOK, gin.H{"data": true}) ~~> Chưa chuẩn hoá response
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
