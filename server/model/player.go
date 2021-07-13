// 自动生成模板Player
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Player struct {
      global.GVA_MODEL
      Pid  int `json:"pid" form:"pid" gorm:"column:pid;comment:;type:bigint;size:20;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(100);size:100;"`
      Curlevel  int `json:"curlevel" form:"curlevel" gorm:"column:curlevel;comment:;type:int;size:10;"`
      Exp  int `json:"exp" form:"exp" gorm:"column:exp;comment:;type:int;size:10;"`
      Physicalstrength  int `json:"physicalstrength" form:"physicalstrength" gorm:"column:physicalstrength;comment:;type:int;size:10;"`
      Rolemodelframe  int `json:"rolemodelframe" form:"rolemodelframe" gorm:"column:rolemodelframe;comment:;type:int;size:10;"`
}


func (Player) TableName() string {
  return "player"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type PlayerWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Player   `json:"business"`
// }

// func (Player) TableName() string {
// 	return "player"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["player"] = func() model.GVA_Workflow {
//   return new(model.PlayerWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["player"] = func() interface{} {
// 	return new(model.Player)
// }
