// Code generated by tools/init_service. DO NOT EDIT.

package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "chat"

	bizCodeChat                = 0
	bizCodeDeleteConversation  = 1
	bizCodeDeleteMessage       = 2
	bizCodeGetConversationByID = 3
	bizCodeInitDevelop         = 4
	bizCodeListConversation    = 5
	bizCodeListMessage         = 6
	bizCodeUpdateConversation  = 7
	bizCodeUpdateMessage       = 8
)

var (
	bizChat                *bizerr.Biz
	bizDeleteConversation  *bizerr.Biz
	bizDeleteMessage       *bizerr.Biz
	bizGetConversationByID *bizerr.Biz
	bizInitDevelop         *bizerr.Biz
	bizListConversation    *bizerr.Biz
	bizListMessage         *bizerr.Biz
	bizUpdateConversation  *bizerr.Biz
	bizUpdateMessage       *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizChat = bizerr.NewBiz(ServiceName, "chat", baseCode+bizCodeChat)
	bizDeleteConversation = bizerr.NewBiz(ServiceName, "delete_conversation", baseCode+bizCodeDeleteConversation)
	bizDeleteMessage = bizerr.NewBiz(ServiceName, "delete_message", baseCode+bizCodeDeleteMessage)
	bizGetConversationByID = bizerr.NewBiz(ServiceName, "get_conversation_by_id", baseCode+bizCodeGetConversationByID)
	bizInitDevelop = bizerr.NewBiz(ServiceName, "init_develop", baseCode+bizCodeInitDevelop)
	bizListConversation = bizerr.NewBiz(ServiceName, "list_conversation", baseCode+bizCodeListConversation)
	bizListMessage = bizerr.NewBiz(ServiceName, "list_message", baseCode+bizCodeListMessage)
	bizUpdateConversation = bizerr.NewBiz(ServiceName, "update_conversation", baseCode+bizCodeUpdateConversation)
	bizUpdateMessage = bizerr.NewBiz(ServiceName, "update_message", baseCode+bizCodeUpdateMessage)
}
