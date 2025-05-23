// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dummy1.proto
package pb

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"go.saastack.io/protoc-gen-grpc-wrapper/middleware"

	"google.golang.org/grpc"
)

type wrapperDummy1 struct {
	client Dummy1Client
}

func NewDummy1HandlerWrapper(client Dummy1Client) Dummy1Client {
	return &wrapperDummy1{client: client}
}

func RegisterDummy1HandlerClientWrapper(ctx context.Context, mux *runtime.ServeMux, client Dummy1Client) error {
	return RegisterDummy1HandlerClient(ctx, mux, NewDummy1HandlerWrapper(client))
}

func (g *wrapperDummy1) GetDummy1Method1(ctx context.Context, in *GetDummy1Method1Request, opts ...grpc.CallOption) (*Dummy1Method1Object, error) {
	middleware.MaskFieldsFromRequest(in)
	var resp, err = g.client.GetDummy1Method1(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	middleware.MaskFieldsFromResponse(resp)
	return resp, nil
}

func (g *wrapperDummy1) GetDummy1Method2(ctx context.Context, in *GetDummy1Method2Request, opts ...grpc.CallOption) (*Dummy1Method2Object, error) {
	middleware.MaskFieldsFromRequest(in)
	var resp, err = g.client.GetDummy1Method2(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	middleware.MaskFieldsFromResponse(resp)
	return resp, nil
}
