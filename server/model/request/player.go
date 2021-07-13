package request

import "gin-vue-admin/model"

type PlayerSearch struct{
    model.Player
    PageInfo
}