// Code generated by Kitex v0.10.0. DO NOT EDIT.

package pluginservice

import (
	"context"
	"errors"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Create": kitex.NewMethodInfo(
		createHandler,
		newPluginServiceCreateArgs,
		newPluginServiceCreateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Update": kitex.NewMethodInfo(
		updateHandler,
		newPluginServiceUpdateArgs,
		newPluginServiceUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"List": kitex.NewMethodInfo(
		listHandler,
		newPluginServiceListArgs,
		newPluginServiceListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetByID": kitex.NewMethodInfo(
		getByIDHandler,
		newPluginServiceGetByIDArgs,
		newPluginServiceGetByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Delete": kitex.NewMethodInfo(
		deleteHandler,
		newPluginServiceDeleteArgs,
		newPluginServiceDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateTool": kitex.NewMethodInfo(
		createToolHandler,
		newPluginServiceCreateToolArgs,
		newPluginServiceCreateToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateTool": kitex.NewMethodInfo(
		updateToolHandler,
		newPluginServiceUpdateToolArgs,
		newPluginServiceUpdateToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListTool": kitex.NewMethodInfo(
		listToolHandler,
		newPluginServiceListToolArgs,
		newPluginServiceListToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetToolByID": kitex.NewMethodInfo(
		getToolByIDHandler,
		newPluginServiceGetToolByIDArgs,
		newPluginServiceGetToolByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteTool": kitex.NewMethodInfo(
		deleteToolHandler,
		newPluginServiceDeleteToolArgs,
		newPluginServiceDeleteToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CallPluginTool": kitex.NewMethodInfo(
		callPluginToolHandler,
		newPluginServiceCallPluginToolArgs,
		newPluginServiceCallPluginToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"TestPluginTool": kitex.NewMethodInfo(
		testPluginToolHandler,
		newPluginServiceTestPluginToolArgs,
		newPluginServiceTestPluginToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	pluginServiceServiceInfo                = NewServiceInfo()
	pluginServiceServiceInfoForClient       = NewServiceInfoForClient()
	pluginServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return pluginServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return pluginServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return pluginServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PluginService"
	handlerType := (*pluginsvc.PluginService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "pluginsvc",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.10.0",
		Extra:           extra,
	}
	return svcInfo
}

func createHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceCreateArgs)
	realResult := result.(*pluginsvc.PluginServiceCreateResult)
	success, err := handler.(pluginsvc.PluginService).Create(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceCreateArgs() interface{} {
	return pluginsvc.NewPluginServiceCreateArgs()
}

func newPluginServiceCreateResult() interface{} {
	return pluginsvc.NewPluginServiceCreateResult()
}

func updateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceUpdateArgs)
	realResult := result.(*pluginsvc.PluginServiceUpdateResult)
	success, err := handler.(pluginsvc.PluginService).Update(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceUpdateArgs() interface{} {
	return pluginsvc.NewPluginServiceUpdateArgs()
}

func newPluginServiceUpdateResult() interface{} {
	return pluginsvc.NewPluginServiceUpdateResult()
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceListArgs)
	realResult := result.(*pluginsvc.PluginServiceListResult)
	success, err := handler.(pluginsvc.PluginService).List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceListArgs() interface{} {
	return pluginsvc.NewPluginServiceListArgs()
}

func newPluginServiceListResult() interface{} {
	return pluginsvc.NewPluginServiceListResult()
}

func getByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceGetByIDArgs)
	realResult := result.(*pluginsvc.PluginServiceGetByIDResult)
	success, err := handler.(pluginsvc.PluginService).GetByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceGetByIDArgs() interface{} {
	return pluginsvc.NewPluginServiceGetByIDArgs()
}

func newPluginServiceGetByIDResult() interface{} {
	return pluginsvc.NewPluginServiceGetByIDResult()
}

func deleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceDeleteArgs)
	realResult := result.(*pluginsvc.PluginServiceDeleteResult)
	success, err := handler.(pluginsvc.PluginService).Delete(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceDeleteArgs() interface{} {
	return pluginsvc.NewPluginServiceDeleteArgs()
}

func newPluginServiceDeleteResult() interface{} {
	return pluginsvc.NewPluginServiceDeleteResult()
}

func createToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceCreateToolArgs)
	realResult := result.(*pluginsvc.PluginServiceCreateToolResult)
	success, err := handler.(pluginsvc.PluginService).CreateTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceCreateToolArgs() interface{} {
	return pluginsvc.NewPluginServiceCreateToolArgs()
}

func newPluginServiceCreateToolResult() interface{} {
	return pluginsvc.NewPluginServiceCreateToolResult()
}

func updateToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceUpdateToolArgs)
	realResult := result.(*pluginsvc.PluginServiceUpdateToolResult)
	success, err := handler.(pluginsvc.PluginService).UpdateTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceUpdateToolArgs() interface{} {
	return pluginsvc.NewPluginServiceUpdateToolArgs()
}

func newPluginServiceUpdateToolResult() interface{} {
	return pluginsvc.NewPluginServiceUpdateToolResult()
}

func listToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceListToolArgs)
	realResult := result.(*pluginsvc.PluginServiceListToolResult)
	success, err := handler.(pluginsvc.PluginService).ListTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceListToolArgs() interface{} {
	return pluginsvc.NewPluginServiceListToolArgs()
}

func newPluginServiceListToolResult() interface{} {
	return pluginsvc.NewPluginServiceListToolResult()
}

func getToolByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceGetToolByIDArgs)
	realResult := result.(*pluginsvc.PluginServiceGetToolByIDResult)
	success, err := handler.(pluginsvc.PluginService).GetToolByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceGetToolByIDArgs() interface{} {
	return pluginsvc.NewPluginServiceGetToolByIDArgs()
}

func newPluginServiceGetToolByIDResult() interface{} {
	return pluginsvc.NewPluginServiceGetToolByIDResult()
}

func deleteToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceDeleteToolArgs)
	realResult := result.(*pluginsvc.PluginServiceDeleteToolResult)
	success, err := handler.(pluginsvc.PluginService).DeleteTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceDeleteToolArgs() interface{} {
	return pluginsvc.NewPluginServiceDeleteToolArgs()
}

func newPluginServiceDeleteToolResult() interface{} {
	return pluginsvc.NewPluginServiceDeleteToolResult()
}

func callPluginToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceCallPluginToolArgs)
	realResult := result.(*pluginsvc.PluginServiceCallPluginToolResult)
	success, err := handler.(pluginsvc.PluginService).CallPluginTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceCallPluginToolArgs() interface{} {
	return pluginsvc.NewPluginServiceCallPluginToolArgs()
}

func newPluginServiceCallPluginToolResult() interface{} {
	return pluginsvc.NewPluginServiceCallPluginToolResult()
}

func testPluginToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceTestPluginToolArgs)
	realResult := result.(*pluginsvc.PluginServiceTestPluginToolResult)
	success, err := handler.(pluginsvc.PluginService).TestPluginTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceTestPluginToolArgs() interface{} {
	return pluginsvc.NewPluginServiceTestPluginToolArgs()
}

func newPluginServiceTestPluginToolResult() interface{} {
	return pluginsvc.NewPluginServiceTestPluginToolResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Create(ctx context.Context, req *pluginsvc.CreatePluginReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceCreateArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceCreateResult
	if err = p.c.Call(ctx, "Create", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Update(ctx context.Context, req *pluginsvc.UpdatePluginReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceUpdateArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceUpdateResult
	if err = p.c.Call(ctx, "Update", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, req *base.PaginationReq) (r *pluginsvc.ListPluginResp, err error) {
	var _args pluginsvc.PluginServiceListArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetByID(ctx context.Context, req *base.IDReq) (r *pluginsvc.Plugin, err error) {
	var _args pluginsvc.PluginServiceGetByIDArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceGetByIDResult
	if err = p.c.Call(ctx, "GetByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Delete(ctx context.Context, req *base.IDReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceDeleteArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceDeleteResult
	if err = p.c.Call(ctx, "Delete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceCreateToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceCreateToolResult
	if err = p.c.Call(ctx, "CreateTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceUpdateToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceUpdateToolResult
	if err = p.c.Call(ctx, "UpdateTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListTool(ctx context.Context, req *base.IDReq) (r *pluginsvc.ListPluginToolResp, err error) {
	var _args pluginsvc.PluginServiceListToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceListToolResult
	if err = p.c.Call(ctx, "ListTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetToolByID(ctx context.Context, req *base.IDReq) (r *pluginsvc.PluginTool, err error) {
	var _args pluginsvc.PluginServiceGetToolByIDArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceGetToolByIDResult
	if err = p.c.Call(ctx, "GetToolByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteTool(ctx context.Context, req *base.IDReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceDeleteToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceDeleteToolResult
	if err = p.c.Call(ctx, "DeleteTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (r *pluginsvc.CallPluginToolResp, err error) {
	var _args pluginsvc.PluginServiceCallPluginToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceCallPluginToolResult
	if err = p.c.Call(ctx, "CallPluginTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (r *pluginsvc.TestPluginToolResp, err error) {
	var _args pluginsvc.PluginServiceTestPluginToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceTestPluginToolResult
	if err = p.c.Call(ctx, "TestPluginTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
