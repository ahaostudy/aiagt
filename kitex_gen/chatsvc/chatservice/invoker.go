// Code generated by Kitex v0.10.0. DO NOT EDIT.

package chatservice

import (
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler chatsvc.ChatService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
