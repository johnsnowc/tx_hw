// Code generated by Kitex v0.0.8. DO NOT EDIT.

package transaction

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"tx_hw/kitex_gen/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Add(ctx context.Context, req *api.AddRequest, callOptions ...callopt.Option) (r *api.AddResponse, err error)
	DeleteByID(ctx context.Context, req *api.DeleteByIDRequest, callOptions ...callopt.Option) (r *api.DeleteByIDResponse, err error)
	PaymentsByRoomID(ctx context.Context, req *api.PaymentsByRoomIDRequest, callOptions ...callopt.Option) (r *api.PaymentsByRoomIDResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTransactionClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTransactionClient struct {
	*kClient
}

func (p *kTransactionClient) Add(ctx context.Context, req *api.AddRequest, callOptions ...callopt.Option) (r *api.AddResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Add(ctx, req)
}

func (p *kTransactionClient) DeleteByID(ctx context.Context, req *api.DeleteByIDRequest, callOptions ...callopt.Option) (r *api.DeleteByIDResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteByID(ctx, req)
}

func (p *kTransactionClient) PaymentsByRoomID(ctx context.Context, req *api.PaymentsByRoomIDRequest, callOptions ...callopt.Option) (r *api.PaymentsByRoomIDResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PaymentsByRoomID(ctx, req)
}
