package cos

import (
	"github.com/aiagt/aiagt/pkg/logerr"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var cli *cos.Client

func Cli() *cos.Client {
	return cli
}

const (
	AvatarDir        = "avatar"
	AppLogoDir       = "app_logo"
	PluginLogoDir    = "plugin_logo"
	KnowledgeLogoDir = "knowledge_logo"
	ChatFile         = "chat_file"
	KnowledgeDir     = "knowledge"
)

func InitCos(cosURL, secretID, secretKey string) {
	u, err := url.Parse(cosURL)
	logerr.Fatal(err)

	bucket := &cos.BaseURL{BucketURL: u}
	cli = cos.NewClient(bucket, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
}

func WithContentTypeByExt(ext string) *cos.ObjectPutOptions {
	return &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/plain",
		},
	}
}
