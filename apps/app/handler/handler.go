package handler

import (
	"github.com/aiagt/aiagt/apps/app/dal/db"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc/knowledgeservice"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
)

// AppServiceImpl implements the last service interface defined in the IDL.
type AppServiceImpl struct {
	appDao   *db.AppDao
	labelDao *db.LabelDao

	userCli      usersvc.Client
	pluginCli    pluginsvc.Client
	knowledgeCli knowledgesvc.Client
}

func NewAppService(appDao *db.AppDao, labelDao *db.LabelDao, userCli usersvc.Client, pluginCli pluginsvc.Client, knowledgeCli knowledgesvc.Client) *AppServiceImpl {
	initServiceBusiness(3)

	return &AppServiceImpl{appDao: appDao, labelDao: labelDao, userCli: userCli, pluginCli: pluginCli, knowledgeCli: knowledgeCli}
}
