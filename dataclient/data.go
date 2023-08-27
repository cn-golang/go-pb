// Code generated by goctl. DO NOT EDIT.
// Source: data.proto

package dataclient

import (
	"context"

	"demo6/data/data"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddBindingDoctorReq = data.AddBindingDoctorReq
	AddBindingDoctorRes = data.AddBindingDoctorRes
	DelBindingDoctorReq = data.DelBindingDoctorReq
	DelBindingDoctorRes = data.DelBindingDoctorRes

	Data interface {
		AddBindingDoctor(ctx context.Context, in *AddBindingDoctorReq, opts ...grpc.CallOption) (*AddBindingDoctorRes, error)
		DelBindingDoctor(ctx context.Context, in *DelBindingDoctorReq, opts ...grpc.CallOption) (*DelBindingDoctorRes, error)
	}

	defaultData struct {
		cli zrpc.Client
	}
)

func NewData(cli zrpc.Client) Data {
	return &defaultData{
		cli: cli,
	}
}

func (m *defaultData) AddBindingDoctor(ctx context.Context, in *AddBindingDoctorReq, opts ...grpc.CallOption) (*AddBindingDoctorRes, error) {
	client := data.NewDataClient(m.cli.Conn())
	return client.AddBindingDoctor(ctx, in, opts...)
}

func (m *defaultData) DelBindingDoctor(ctx context.Context, in *DelBindingDoctorReq, opts ...grpc.CallOption) (*DelBindingDoctorRes, error) {
	client := data.NewDataClient(m.cli.Conn())
	return client.DelBindingDoctor(ctx, in, opts...)
}