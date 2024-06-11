package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest-api/common"
	"rest-api/modules/item/business"
	"rest-api/modules/item/model"
	"rest-api/modules/item/storage/postgres"
)

func CreatePosItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var item model.TodoItemCreation
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := postgres.NewSqlStorage(db)
		biz := business.NewCreateTodoItemBusiness(store)
		if err := biz.CreateNewTodoItemBusiness(c.Request.Context(), &item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// c.JSON(http.StatusOK, gin.H{"data": item.Id}) ~~> Chưa chuẩn hoá Response
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(item.Id))
	}
}
