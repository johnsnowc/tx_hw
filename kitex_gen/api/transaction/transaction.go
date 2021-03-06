// Code generated by Kitex v0.0.8. DO NOT EDIT.

package transaction

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"tx_hw/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return transactionServiceInfo
}

var transactionServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "Transaction"
	handlerType := (*api.Transaction)(nil)
	methods := map[string]kitex.MethodInfo{
		"add":              kitex.NewMethodInfo(addHandler, newTransactionAddArgs, newTransactionAddResult, false),
		"deleteByID":       kitex.NewMethodInfo(deleteByIDHandler, newTransactionDeleteByIDArgs, newTransactionDeleteByIDResult, false),
		"paymentsByRoomID": kitex.NewMethodInfo(paymentsByRoomIDHandler, newTransactionPaymentsByRoomIDArgs, newTransactionPaymentsByRoomIDResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.0.8",
		Extra:           extra,
	}
	return svcInfo
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TransactionAddArgs)
	realResult := result.(*api.TransactionAddResult)
	success, err := handler.(api.Transaction).Add(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTransactionAddArgs() interface{} {
	return api.NewTransactionAddArgs()
}

func newTransactionAddResult() interface{} {
	return api.NewTransactionAddResult()
}

func deleteByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TransactionDeleteByIDArgs)
	realResult := result.(*api.TransactionDeleteByIDResult)
	success, err := handler.(api.Transaction).DeleteByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTransactionDeleteByIDArgs() interface{} {
	return api.NewTransactionDeleteByIDArgs()
}

func newTransactionDeleteByIDResult() interface{} {
	return api.NewTransactionDeleteByIDResult()
}

func paymentsByRoomIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TransactionPaymentsByRoomIDArgs)
	realResult := result.(*api.TransactionPaymentsByRoomIDResult)
	success, err := handler.(api.Transaction).PaymentsByRoomID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTransactionPaymentsByRoomIDArgs() interface{} {
	return api.NewTransactionPaymentsByRoomIDArgs()
}

func newTransactionPaymentsByRoomIDResult() interface{} {
	return api.NewTransactionPaymentsByRoomIDResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Add(ctx context.Context, req *api.AddRequest) (r *api.AddResponse, err error) {
	var _args api.TransactionAddArgs
	_args.Req = req
	var _result api.TransactionAddResult
	if err = p.c.Call(ctx, "add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteByID(ctx context.Context, req *api.DeleteByIDRequest) (r *api.DeleteByIDResponse, err error) {
	var _args api.TransactionDeleteByIDArgs
	_args.Req = req
	var _result api.TransactionDeleteByIDResult
	if err = p.c.Call(ctx, "deleteByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PaymentsByRoomID(ctx context.Context, req *api.PaymentsByRoomIDRequest) (r *api.PaymentsByRoomIDResponse, err error) {
	var _args api.TransactionPaymentsByRoomIDArgs
	_args.Req = req
	var _result api.TransactionPaymentsByRoomIDResult
	if err = p.c.Call(ctx, "paymentsByRoomID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
