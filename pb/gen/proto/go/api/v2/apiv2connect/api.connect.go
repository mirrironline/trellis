// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v2/api.proto

package apiv2connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v2 "github.com/mirrironline/trellis/pb/gen/proto/go/api/v2"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ObservabilityServiceName is the fully-qualified name of the ObservabilityService service.
	ObservabilityServiceName = "api.v2.ObservabilityService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ObservabilityServiceGetLogsProcedure is the fully-qualified name of the ObservabilityService's
	// GetLogs RPC.
	ObservabilityServiceGetLogsProcedure = "/api.v2.ObservabilityService/GetLogs"
	// ObservabilityServicePutLogsProcedure is the fully-qualified name of the ObservabilityService's
	// PutLogs RPC.
	ObservabilityServicePutLogsProcedure = "/api.v2.ObservabilityService/PutLogs"
	// ObservabilityServiceGetMetricsProcedure is the fully-qualified name of the ObservabilityService's
	// GetMetrics RPC.
	ObservabilityServiceGetMetricsProcedure = "/api.v2.ObservabilityService/GetMetrics"
)

// ObservabilityServiceClient is a client for the api.v2.ObservabilityService service.
type ObservabilityServiceClient interface {
	GetLogs(context.Context, *connect_go.Request[v2.GetLogsRequest]) (*connect_go.Response[v2.GetLogsResponse], error)
	PutLogs(context.Context, *connect_go.Request[v2.PutLogsRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetMetrics(context.Context, *connect_go.Request[v2.GetMetricsRequest]) (*connect_go.Response[v2.GetMetricsResponse], error)
}

// NewObservabilityServiceClient constructs a client for the api.v2.ObservabilityService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewObservabilityServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ObservabilityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &observabilityServiceClient{
		getLogs: connect_go.NewClient[v2.GetLogsRequest, v2.GetLogsResponse](
			httpClient,
			baseURL+ObservabilityServiceGetLogsProcedure,
			opts...,
		),
		putLogs: connect_go.NewClient[v2.PutLogsRequest, emptypb.Empty](
			httpClient,
			baseURL+ObservabilityServicePutLogsProcedure,
			opts...,
		),
		getMetrics: connect_go.NewClient[v2.GetMetricsRequest, v2.GetMetricsResponse](
			httpClient,
			baseURL+ObservabilityServiceGetMetricsProcedure,
			opts...,
		),
	}
}

// observabilityServiceClient implements ObservabilityServiceClient.
type observabilityServiceClient struct {
	getLogs    *connect_go.Client[v2.GetLogsRequest, v2.GetLogsResponse]
	putLogs    *connect_go.Client[v2.PutLogsRequest, emptypb.Empty]
	getMetrics *connect_go.Client[v2.GetMetricsRequest, v2.GetMetricsResponse]
}

// GetLogs calls api.v2.ObservabilityService.GetLogs.
func (c *observabilityServiceClient) GetLogs(ctx context.Context, req *connect_go.Request[v2.GetLogsRequest]) (*connect_go.Response[v2.GetLogsResponse], error) {
	return c.getLogs.CallUnary(ctx, req)
}

// PutLogs calls api.v2.ObservabilityService.PutLogs.
func (c *observabilityServiceClient) PutLogs(ctx context.Context, req *connect_go.Request[v2.PutLogsRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.putLogs.CallUnary(ctx, req)
}

// GetMetrics calls api.v2.ObservabilityService.GetMetrics.
func (c *observabilityServiceClient) GetMetrics(ctx context.Context, req *connect_go.Request[v2.GetMetricsRequest]) (*connect_go.Response[v2.GetMetricsResponse], error) {
	return c.getMetrics.CallUnary(ctx, req)
}

// ObservabilityServiceHandler is an implementation of the api.v2.ObservabilityService service.
type ObservabilityServiceHandler interface {
	GetLogs(context.Context, *connect_go.Request[v2.GetLogsRequest]) (*connect_go.Response[v2.GetLogsResponse], error)
	PutLogs(context.Context, *connect_go.Request[v2.PutLogsRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetMetrics(context.Context, *connect_go.Request[v2.GetMetricsRequest]) (*connect_go.Response[v2.GetMetricsResponse], error)
}

// NewObservabilityServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewObservabilityServiceHandler(svc ObservabilityServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ObservabilityServiceGetLogsProcedure, connect_go.NewUnaryHandler(
		ObservabilityServiceGetLogsProcedure,
		svc.GetLogs,
		opts...,
	))
	mux.Handle(ObservabilityServicePutLogsProcedure, connect_go.NewUnaryHandler(
		ObservabilityServicePutLogsProcedure,
		svc.PutLogs,
		opts...,
	))
	mux.Handle(ObservabilityServiceGetMetricsProcedure, connect_go.NewUnaryHandler(
		ObservabilityServiceGetMetricsProcedure,
		svc.GetMetrics,
		opts...,
	))
	return "/api.v2.ObservabilityService/", mux
}

// UnimplementedObservabilityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedObservabilityServiceHandler struct{}

func (UnimplementedObservabilityServiceHandler) GetLogs(context.Context, *connect_go.Request[v2.GetLogsRequest]) (*connect_go.Response[v2.GetLogsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v2.ObservabilityService.GetLogs is not implemented"))
}

func (UnimplementedObservabilityServiceHandler) PutLogs(context.Context, *connect_go.Request[v2.PutLogsRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v2.ObservabilityService.PutLogs is not implemented"))
}

func (UnimplementedObservabilityServiceHandler) GetMetrics(context.Context, *connect_go.Request[v2.GetMetricsRequest]) (*connect_go.Response[v2.GetMetricsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v2.ObservabilityService.GetMetrics is not implemented"))
}
