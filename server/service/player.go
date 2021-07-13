package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePlayer
//@description: 创建Player记录
//@param: player model.Player
//@return: err error

func CreatePlayer(player model.Player) (err error) {
	err = global.GVA_DBMap["gamedata"].Create(&player).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePlayer
//@description: 删除Player记录
//@param: player model.Player
//@return: err error

func DeletePlayer(player model.Player) (err error) {
	err = global.GVA_DBMap["gamedata"].Delete(player).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePlayerByIds
//@description: 批量删除Player记录
//@param: ids request.IdsReq
//@return: err error

func DeletePlayerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBMap["gamedata"].Delete(&[]model.Player{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePlayer
//@description: 更新Player记录
//@param: player *model.Player
//@return: err error

func UpdatePlayer(player model.Player) (err error) {
	err = global.GVA_DBMap["gamedata"].Save(&player).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPlayer
//@description: 根据id获取Player记录
//@param: id uint
//@return: err error, player model.Player

func GetPlayer(id uint) (err error, player model.Player) {
	err = global.GVA_DBMap["gamedata"].Where("id = ?", id).First(&player).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPlayerInfoList
//@description: 分页获取Player记录
//@param: info request.PlayerSearch
//@return: err error, list interface{}, total int64

func GetPlayerInfoList(info request.PlayerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBMap["gamedata"].Model(&model.Player{})
	var players []model.Player
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&players).Error
	return err, players, total
}
