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

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get Id parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//	return
		//}
		store := mysql.NewSqlStorage(db)
		biz := business.NewDeleteTodoItemBusiness(store)
		if err := biz.DeleteItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//c.JSON(http.StatusOK, gin.H{"data": true})
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
