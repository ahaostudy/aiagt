package test

import (
	"fmt"
	"github.com/aiagt/aiagt/apps/knowledge/conf"
	"github.com/aiagt/aiagt/apps/knowledge/pkg/rag"
	"github.com/aiagt/aiagt/common/confutil"
	"github.com/aiagt/aiagt/common/tests"
	"github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/aiagt/aiagt/rpc"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"github.com/stretchr/testify/require"
	"testing"
)

var ctx = tests.InitTesting()

func init() {
	confutil.LoadConf(conf.Conf(), "../conf")
}

func TestSaveKnowledge(t *testing.T) {
	tests.RpcCallWrap(rpc.KnowledgeCli.SaveKnowledge(ctx, &knowledgesvc.SaveKnowledgeReq{
		Id:             utils.PtrOf(int64(1922279058809294848)),
		Name:           "测试知识库",
		Description:    "测试知识库",
		IsPrivate:      true,
		EmbedModelId:   1,
		TopK:           5,
		ScoreThreshold: utils.PtrOf(0.0),
	}))
}

func TestUploadDocument(t *testing.T) {
	tests.RpcCallWrap(rpc.KnowledgeCli.UploadDocument(ctx, &knowledgesvc.UploadDocumentReq{
		KnowledgeId: 1922279058809294848,
		Name:        "特朗普第二任期对华关税.txt",
		Content: []byte(`2025年2月1日，美国总统特朗普签署行政命令，对所有进口自中国的商品，美国将在现有关税基础上加征10%的关税；[34]2月4日该行政命令正式生效。作为回应，中方宣布一系列反制措施，包括对煤炭、液化天然气、原油、农业机械、大排量汽车、皮卡等美国进口商品加征10-15%关税[421]、对钨、碲、铋、钼、铟相关物项实施出口管制[422]、对谷歌公司展开反垄断调查[423]、以及将PVH集团和因美纳公司列入不可靠实体清单[424]。

2025年3月3日，特朗普签署行政命令，将进入美国的中国商品关税从此前的10%提高至20%。[425]同日，中华人民共和国商务部、国务院关税税则委员会对美国执行多项反制措施，其中商务部将莱多斯（Leidos）等15家美国实体列入出口管制管控名单，将特科姆公司（TCOM,Limited Partnership）等10家美国企业列入不可靠实体清单，对早前列入不可靠实体清单的因美纳公司采取禁止其向中国出口基因测序仪的措施；国务院宣布自3月10日起对美国进口鸡肉、小麦、玉米、棉花加征15%关税，高粱、大豆、猪肉、牛肉、水产品、水果、蔬菜、乳制品加征10%关税[426]。

2025年4月2日，特朗普宣布对美国100多个贸易伙伴征收关税，其中对中国商品额外加征34%关税。[427]特朗普还在当天签署行政令，决定自5月2日起取消中国大陆与香港的小额包裹免税措施。[428]

4月4日，中国公布多项反制措施：对原产于美国的所有进口商品，在现行适用关税税率基础上加征34%关税；暂停一家美国企业高粱输华资质，三家美国企业禽肉骨粉输华资质，并暂停两家美国企业禽肉产品输华；将斯凯迪欧无人机公司等11家美国实体，列入不可靠实体清单；将高点航空技术公司（High Point Aerotechnologies LLC）等16家美国实体，列入出口管制管控名单；对原产于美国、印度的进口相关医用CT球管发起反倾销立案调查；对钐、钆、铽、镝、镥、钪、钇等七类中重稀土相关物项实施出口管制措施。在世贸组织争端解决机制发起起诉。[429]同日，市场监管总局宣布对杜邦中国公司进行反垄断调查。[430]

4月7日，特朗普要求中国在4月9日前取消反制措施，否则会再加征50%关税[431][432]。由于中国未取消反制措施，美国4月9日对中国再征收50%的额外关税，令特朗普在第二任期中对中国所有商品加征的新关税税率就此达到104%[433]。作为回应，中国也宣布自2025年4月10日12时01分起，对原产于美国的进口商品加征84%关税措施 [434]；中国商务部将护盾人工智能公司等6家美国企业列入不可靠实体清单[435]。

4月9日，特朗普在社交媒体宣布，由于中国的反制措施，将中国商品的关税税率加至125%[436]。同时他下令，由于各国展现出谈判诚意，且并未采取报复性加税政策，因此对除了中国以外的75个国家给予90日的关税暂缓期。并且在此期间，美国政府同意统一降低各国的点对点关税税额至10％[437]。

4月10日，白宫指出，由于特朗普先前曾下令就中国违法走私鸦片类药物芬太尼进入美国的争议向中方课征20％关税，故美国政府对中国征收的关税累计总额已达145％[438]。

4月11日，中华人民共和国财政部、国务院关税税则委员会宣布将加征关税税率由84%提高至125%，并不再理会美方的关税数字游戏[439][440]。

4月15日，白宫公布了一份有关工业矿产与相关产品进口措施说明的《事实清单》（英语：Fact Sheet: President Donald J. Trump Ensures National Security and Economic Resilience Through Section 232 Actions on Processed Critical Minerals and Derivative Products[441]。），其中列出由于中国对美采取报复性措施，因此美国政府再次调高对中关税至245％[442]。对此，中华人民共和国外交部则称“中方已多次阐明在关税问题上的严正立场，这场关税战是美方发起的，中方采取必要的反制措施，是为了维护自身的正当权益和国际公平正义，完全合理合法。关税战、贸易战没有赢家，中方不愿打，但也绝不怕打。[443]”

4月16日，中国突然宣布更换贸易谈判代表。前驻世界贸易组织（WTO）代表李成钢接替2022年起担任该职的王受文，并兼任商务部副部长。在2月日内瓦的一次世贸组织会议上，李成钢曾猛烈批评美方“任意地对包括中方在内的贸易伙伴加征或威胁加征关税”，并警告此举将给全球带来“关税冲击”。[444]

4月17日，特朗普周四在白宫记者会上说表示：“我不要关税再升高，因为升到了某个程度，人们就不愿意再买东西了。所以，我可能不会调高关税，甚至可能不会让它升到那个水平。我可能会降低关税，因为你希望人们买东西。”这可能预示中美关税战将会降温。[445]

4月22日，特朗普受访时表态，接下来将以“非常友好”的态度跟中方谈判，预计在双方达成协议后，美国对中国进口产品的关税将大幅下降，但不会降至零。但同时特朗普也指出，若中国不同意达成贸易协议，美国将自行设定条件。特朗普重申“和中国的关系很好”，虽然双方早前在关税问题上交锋，但他不会在贸易谈判中对中国领导人习近平采取强硬态度[446]。

5月6日，中美两国宣布各派代表前往瑞士进行贸易谈判，美国财政部长斯科特·贝森特、美国贸易代表贾米森·格里尔和中共中央政治局委员、中国国务院副总理、中共中央财经委员会办公室主任何立峰出席此次谈判[447]。

5月12日，中美两国公布瑞士会谈联合声明，美国宣布将修改2025年4月2日第14257号行政命令中规定的对中国商品（包括香港特别行政区和澳门特别行政区商品）加征的从价关税，其中，24％的关税在初始的90天内暂停实施，同时保留按该行政命令规定对这些商品加征的剩余10％关税；取消根据2025年4月8日第14259号行政命令和2025年4月9日第14266号行政命令对这些商品加征的关税。中国将相应修改税委会公告2025年第4号中规定的对美国商品加征的从价关税，其中，24％的关税在初始的90天内暂停实施，同时保留对这些商品加征的剩余10％关税，并取消根据税委会公告2025年第5号和第6号对这些商品加征的关税；采取必要措施，暂停或取消自2025年4月2日起针对美国的非关税反制措施[448]。`),
		Type:        ".txt",
		ChunkSize:   1024,
		OverlapSize: 100,
		Separators:  nil,
	}))
}

func TestRetrieval(t *testing.T) {
	resp, err := rpc.KnowledgeCli.Retrieval(ctx, &knowledgesvc.RetrievalReq{
		KnowledgeIds: []int64{1922279058809294848},
		Query:        "违法走私鸦片类药物芬太尼进",
	})
	require.NoError(t, err)

	fmt.Println(utils.PrettyIndent(resp, "", "  "))
}

func TestMilvusSearch(t *testing.T) {
	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address:  conf.Conf().Milvus.Address,
		Username: conf.Conf().Milvus.Username,
		Password: conf.Conf().Milvus.Password,
	})
	require.NoError(t, err)

	collection := conf.Conf().Milvus.Collection

	result, err := cli.Query(ctx, milvusclient.NewQueryOption(collection).
		WithFilter("knowledge_id == '1922279058809294848'").
		WithOutputFields("id", "knowledge_id", "document_id", "content", "vector").
		WithLimit(5),
	)
	require.NoError(t, err)

	t.Log(result.ResultCount)

	var records []*rag.KnowledgeSchame
	err = result.Fields.Unmarshal(&records)
	require.NoError(t, err)

	t.Log(utils.Pretty(records, 1<<16))
}

func TestMilvusDelete(t *testing.T) {
	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address:  conf.Conf().Milvus.Address,
		Username: conf.Conf().Milvus.Username,
		Password: conf.Conf().Milvus.Password,
	})
	require.NoError(t, err)

	collection := conf.Conf().Milvus.Collection

	_, err = cli.Delete(ctx, milvusclient.NewDeleteOption(collection).WithExpr("knowledge_id == '1922279058809294848'"))
	require.NoError(t, err)
}
