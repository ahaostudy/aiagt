// Code generated by Kitex v0.10.0. DO NOT EDIT.

package modelservice

import (
	"context"
	"errors"
	"fmt"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GenToken": kitex.NewMethodInfo(
		genTokenHandler,
		newModelServiceGenTokenArgs,
		newModelServiceGenTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Chat": kitex.NewMethodInfo(
		chatHandler,
		newModelServiceChatArgs,
		newModelServiceChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingServer),
	),
}

var (
	modelServiceServiceInfo                = NewServiceInfo()
	modelServiceServiceInfoForClient       = NewServiceInfoForClient()
	modelServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return modelServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return modelServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return modelServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(true, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ModelService"
	handlerType := (*modelsvc.ModelService)(nil)
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
		"PackageName": "modelsvc",
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

func genTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*modelsvc.ModelServiceGenTokenArgs)
	realResult := result.(*modelsvc.ModelServiceGenTokenResult)
	success, err := handler.(modelsvc.ModelService).GenToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newModelServiceGenTokenArgs() interface{} {
	return modelsvc.NewModelServiceGenTokenArgs()
}

func newModelServiceGenTokenResult() interface{} {
	return modelsvc.NewModelServiceGenTokenResult()
}

func chatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st, ok := arg.(*streaming.Args)
	if !ok {
		return errors.New("ModelService.Chat is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &modelServiceChatServer{st.Stream}
	req := new(modelsvc.ChatReq)
	if err := st.Stream.RecvMsg(req); err != nil {
		return err
	}
	return handler.(modelsvc.ModelService).Chat(req, stream)
}

type modelServiceChatClient struct {
	streaming.Stream
}

func (x *modelServiceChatClient) DoFinish(err error) {
	if finisher, ok := x.Stream.(streaming.WithDoFinish); ok {
		finisher.DoFinish(err)
	} else {
		panic(fmt.Sprintf("streaming.WithDoFinish is not implemented by %T", x.Stream))
	}
}
func (x *modelServiceChatClient) Recv() (*modelsvc.ChatResp, error) {
	m := new(modelsvc.ChatResp)
	return m, x.Stream.RecvMsg(m)
}

type modelServiceChatServer struct {
	streaming.Stream
}

func (x *modelServiceChatServer) Send(m *modelsvc.ChatResp) error {
	return x.Stream.SendMsg(m)
}

func newModelServiceChatArgs() interface{} {
	return modelsvc.NewModelServiceChatArgs()
}

func newModelServiceChatResult() interface{} {
	return modelsvc.NewModelServiceChatResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GenToken(ctx context.Context, req *modelsvc.GenTokenReq) (r *modelsvc.GenTokenResp, err error) {
	var _args modelsvc.ModelServiceGenTokenArgs
	_args.Req = req
	var _result modelsvc.ModelServiceGenTokenResult
	if err = p.c.Call(ctx, "GenToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Chat(ctx context.Context, req *modelsvc.ChatReq) (ModelService_ChatClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "Chat", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &modelServiceChatClient{res.Stream}

	if err := stream.Stream.SendMsg(req); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}
