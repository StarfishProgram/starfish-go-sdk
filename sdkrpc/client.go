package sdkrpc

import (
	reflect "reflect"

	"github.com/StarfishProgram/starfish-go-sdk/sdk"
	"github.com/StarfishProgram/starfish-go-sdk/sdkcodes"
	"github.com/StarfishProgram/starfish-go-sdk/sdklog"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

var clientIns map[string]*_Client

func init() {
	clientIns = make(map[string]*_Client)
}

type _Client struct {
	client GRPCServiceClient
}

func InitClient(url string, key ...string) {
	InitClientWithOptions(url, []grpc.DialOption{}, key...)
}

func InitClientWithOptions(url string, opts []grpc.DialOption, key ...string) {
	dialOpts := make([]grpc.DialOption, 0, len(opts)+1)
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dialOpts = append(dialOpts, opts...)
	conn, err := grpc.Dial(
		url,
		dialOpts...,
	)
	if err != nil {
		sdklog.AddCallerSkip(1).Panic(err)
	}
	client := NewGRPCServiceClient(conn)
	ins := _Client{client: client}
	if len(key) == 0 {
		clientIns[""] = &ins
	} else {
		clientIns[key[0]] = &ins
	}
}

func Client(key ...string) *_Client {
	if len(key) == 0 {
		return clientIns[""]
	} else {
		return clientIns[key[0]]
	}
}

type CallResult[D protoreflect.ProtoMessage] struct {
	Code *Code
	Data D
}

// AssertCode 错误断言
func (cr *CallResult[D]) AssertCode() {
	if cr.Code != nil {
		code := sdkcodes.New(
			cr.Code.Code,
			cr.Code.Msg,
			cr.Code.I18N,
			cr.Code.Meta,
		)
		panic(code)
	}
}

func Call[P, R protoreflect.ProtoMessage](client *_Client, param P) CallResult[R] {
	var r R
	anyParam, err := anypb.New(param)
	if err != nil {
		return CallResult[R]{
			Code: &Code{
				Code: sdkcodes.Internal.Code(),
				Msg:  err.Error(),
				I18N: sdkcodes.Internal.I18n(),
				Meta: sdkcodes.Internal.Meta(),
			},
			Data: r,
		}
	}
	result, err := client.client.Call(sdk.Context(), anyParam)
	if err != nil {
		sdklog.AddCallerSkip(1).Error(err)
		return CallResult[R]{
			Code: &Code{
				Code: sdkcodes.Internal.Code(),
				Msg:  err.Error(),
				I18N: sdkcodes.Internal.I18n(),
				Meta: sdkcodes.Internal.Meta(),
			},
			Data: r,
		}
	}
	if result.Code != nil {
		return CallResult[R]{
			Code: &Code{
				Code: result.Code.Code,
				Msg:  result.Code.Msg,
				I18N: result.Code.I18N,
				Meta: result.Code.Meta,
			},
			Data: r,
		}
	}
	realData := reflect.New(reflect.TypeOf(r).Elem()).Interface().(R)
	if err := result.Data.UnmarshalTo(realData); err != nil {
		return CallResult[R]{
			Code: &Code{
				Code: sdkcodes.Internal.Code(),
				Msg:  err.Error(),
				I18N: sdkcodes.Internal.I18n(),
				Meta: sdkcodes.Internal.Meta(),
			},
			Data: r,
		}
	}
	return CallResult[R]{Data: realData}
}
