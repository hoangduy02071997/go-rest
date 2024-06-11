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

func GetPosItems(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		var filters model.FilterItem
		if err := c.ShouldBind(&filters); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := postgres.NewSqlStorage(db)
		biz := business.NewGetListItemBusiness(store)
		results, err := biz.GetListItems(c.Request.Context(), &filters, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//c.JSON(http.StatusOK, gin.H{"data": todoItems, "paging": paging}) ~~> Chưa chuẩn hoá Response
		c.JSON(http.StatusOK, common.NewSuccessResponse(results, paging, filters))
	}
}
