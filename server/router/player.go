package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPlayerRouter(Router *gin.RouterGroup) {
	PlayerRouter := Router.Group("player").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PlayerRouter.POST("createPlayer", v1.CreatePlayer)   // 新建Player
		PlayerRouter.DELETE("deletePlayer", v1.DeletePlayer) // 删除Player
		PlayerRouter.DELETE("deletePlayerByIds", v1.DeletePlayerByIds) // 批量删除Player
		PlayerRouter.PUT("updatePlayer", v1.UpdatePlayer)    // 更新Player
		PlayerRouter.GET("findPlayer", v1.FindPlayer)        // 根据ID获取Player
		PlayerRouter.GET("getPlayerList", v1.GetPlayerList)  // 获取Player列表
	}
}
