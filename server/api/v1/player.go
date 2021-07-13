package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Player
// @Summary 创建Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Player true "创建Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /player/createPlayer [post]
func CreatePlayer(c *gin.Context) {
	var player model.Player
	_ = c.ShouldBindJSON(&player)
	if err := service.CreatePlayer(player); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Player
// @Summary 删除Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Player true "删除Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /player/deletePlayer [delete]
func DeletePlayer(c *gin.Context) {
	var player model.Player
	_ = c.ShouldBindJSON(&player)
	if err := service.DeletePlayer(player); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Player
// @Summary 批量删除Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /player/deletePlayerByIds [delete]
func DeletePlayerByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePlayerByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Player
// @Summary 更新Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Player true "更新Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /player/updatePlayer [put]
func UpdatePlayer(c *gin.Context) {
	var player model.Player
	_ = c.ShouldBindJSON(&player)
	if err := service.UpdatePlayer(player); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Player
// @Summary 用id查询Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Player true "用id查询Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /player/findPlayer [get]
func FindPlayer(c *gin.Context) {
	var player model.Player
	_ = c.ShouldBindQuery(&player)
	if err, replayer := service.GetPlayer(player.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replayer": replayer}, c)
	}
}

// @Tags Player
// @Summary 分页获取Player列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PlayerSearch true "分页获取Player列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /player/getPlayerList [get]
func GetPlayerList(c *gin.Context) {
	var pageInfo request.PlayerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPlayerInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
