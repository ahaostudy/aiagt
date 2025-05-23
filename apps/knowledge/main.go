package main

import (
	"github.com/aiagt/aiagt/apps/knowledge/conf"
	"github.com/aiagt/aiagt/apps/knowledge/dal/db"
	"github.com/aiagt/aiagt/apps/knowledge/handler"
	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/apps/knowledge/pkg/rag"
	"github.com/aiagt/aiagt/common/cos"
	"github.com/aiagt/aiagt/common/kitex/ktserveroption"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	"github.com/aiagt/aiagt/common/logger"
	"github.com/aiagt/aiagt/common/observability"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc/knowledgeservice"
	"github.com/aiagt/aiagt/pkg/closer"
	"github.com/aiagt/aiagt/pkg/logerr"
	"github.com/aiagt/aiagt/rpc"
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktlog "github.com/aiagt/kitextool/option/server/log"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"log"
)

func main() {
	handle := handler.NewKnowledgeServiceImpl(
		db.NewKnowledgeDao(),
		db.NewKnowledgeDocumentDao(),
		db.NewKnowledgeChunkDao(),
	)

	config := conf.Conf()
	observability.InitMetrics(config.Server.Name, config.Metrics.Addr, config.Registry.Address[0])
	observability.InitTracing(config.Server.Name, config.Tracing.ExportAddr)

	svr := knowledgesvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			config,
			ktserveroption.WithLocalIpOption(),
			ktlog.WithLogger(logger.Logger()),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(config.GetServerConf(), rpc.UserCli)))

	logerr.Fatal(ktdb.DB().AutoMigrate(new(model.Knowledge), new(model.KnowledgeDocument), new(model.KnowledgeChunk)))
	logerr.Fatal(ktdb.DB().Use(tracing.NewPlugin(tracing.WithoutMetrics())))

	cos.InitCos(config.Cos.URL, config.Cos.SecretID, config.Cos.SecretKey)

	milvusClient, err := rag.NewMilvusClient()
	logerr.Fatal(err)
	defer closer.Close(milvusClient)

	handle.SetMilvusClient(milvusClient)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
