// Code generated by tools/init_service. DO NOT EDIT.

package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "model"

	bizCodeChat              = 0
	bizCodeCreateModel       = 1
	bizCodeDeleteModel       = 2
	bizCodeGenToken          = 3
	bizCodeGetApikeyByModel  = 4
	bizCodeGetApikeyBySource = 5
	bizCodeGetModelById      = 6
	bizCodeListModel         = 7
	bizCodeUpdateModel       = 8
)

var (
	bizChat              *bizerr.Biz
	bizCreateModel       *bizerr.Biz
	bizDeleteModel       *bizerr.Biz
	bizGenToken          *bizerr.Biz
	bizGetApikeyByModel  *bizerr.Biz
	bizGetApikeyBySource *bizerr.Biz
	bizGetModelById      *bizerr.Biz
	bizListModel         *bizerr.Biz
	bizUpdateModel       *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizChat = bizerr.NewBiz(ServiceName, "chat", baseCode+bizCodeChat)
	bizCreateModel = bizerr.NewBiz(ServiceName, "create_model", baseCode+bizCodeCreateModel)
	bizDeleteModel = bizerr.NewBiz(ServiceName, "delete_model", baseCode+bizCodeDeleteModel)
	bizGenToken = bizerr.NewBiz(ServiceName, "gen_token", baseCode+bizCodeGenToken)
	bizGetApikeyByModel = bizerr.NewBiz(ServiceName, "get_api_key_by_model", baseCode+bizCodeGetApikeyByModel)
	bizGetApikeyBySource = bizerr.NewBiz(ServiceName, "get_api_key_by_source", baseCode+bizCodeGetApikeyBySource)
	bizGetModelById = bizerr.NewBiz(ServiceName, "get_model_by_id", baseCode+bizCodeGetModelById)
	bizListModel = bizerr.NewBiz(ServiceName, "list_model", baseCode+bizCodeListModel)
	bizUpdateModel = bizerr.NewBiz(ServiceName, "update_model", baseCode+bizCodeUpdateModel)
}
