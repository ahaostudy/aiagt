package handler

import (
	"github.com/aiagt/aiagt/apps/chat/dal/db"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc/knowledgeservice"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct {
	conversationDao *db.ConversationDao
	messageDao      *db.MessageDao

	userCli        usersvc.Client
	appCli         appsvc.Client
	pluginCli      pluginsvc.Client
	modelCli       modelsvc.Client
	modelStreamCli modelsvc.StreamClient
	knowledgeCli   knowledgesvc.Client
}

func NewChatService(conversationDao *db.ConversationDao, messageDao *db.MessageDao, userCli usersvc.Client, appCli appsvc.Client, pluginCli pluginsvc.Client, modelCli modelsvc.Client, modelStreamCli modelsvc.StreamClient, knowledgeCli knowledgesvc.Client) *ChatServiceImpl {
	initServiceBusiness(4)

	return &ChatServiceImpl{conversationDao: conversationDao, messageDao: messageDao, userCli: userCli, appCli: appCli, pluginCli: pluginCli, modelCli: modelCli, modelStreamCli: modelStreamCli, knowledgeCli: knowledgeCli}
}
