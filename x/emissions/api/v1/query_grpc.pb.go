// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: emissions/v1/query.proto

package emissionsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Params_FullMethodName                            = "/emissions.v1.Query/Params"
	Query_GetLastRewardsUpdate_FullMethodName              = "/emissions.v1.Query/GetLastRewardsUpdate"
	Query_GetNextTopicId_FullMethodName                    = "/emissions.v1.Query/GetNextTopicId"
	Query_GetTopic_FullMethodName                          = "/emissions.v1.Query/GetTopic"
	Query_GetActiveTopics_FullMethodName                   = "/emissions.v1.Query/GetActiveTopics"
	Query_GetAllTopics_FullMethodName                      = "/emissions.v1.Query/GetAllTopics"
	Query_GetExistingInferenceRequest_FullMethodName       = "/emissions.v1.Query/GetExistingInferenceRequest"
	Query_GetAllExistingInferenceRequests_FullMethodName   = "/emissions.v1.Query/GetAllExistingInferenceRequests"
	Query_GetTopicUnmetDemand_FullMethodName               = "/emissions.v1.Query/GetTopicUnmetDemand"
	Query_GetWorkerLatestInferenceByTopicId_FullMethodName = "/emissions.v1.Query/GetWorkerLatestInferenceByTopicId"
	Query_GetInferencesAtBlock_FullMethodName              = "/emissions.v1.Query/GetInferencesAtBlock"
	Query_GetForecastsAtBlock_FullMethodName               = "/emissions.v1.Query/GetForecastsAtBlock"
	Query_GetNetworkLossBundleAtBlock_FullMethodName       = "/emissions.v1.Query/GetNetworkLossBundleAtBlock"
	Query_GetTotalStake_FullMethodName                     = "/emissions.v1.Query/GetTotalStake"
	Query_GetReputerStakeList_FullMethodName               = "/emissions.v1.Query/GetReputerStakeList"
	Query_GetTopicStakeList_FullMethodName                 = "/emissions.v1.Query/GetTopicStakeList"
	Query_GetWorkerNodeRegistration_FullMethodName         = "/emissions.v1.Query/GetWorkerNodeRegistration"
	Query_GetWorkerAddressByP2PKey_FullMethodName          = "/emissions.v1.Query/GetWorkerAddressByP2PKey"
	Query_GetReputerAddressByP2PKey_FullMethodName         = "/emissions.v1.Query/GetReputerAddressByP2PKey"
	Query_GetRegisteredTopicIds_FullMethodName             = "/emissions.v1.Query/GetRegisteredTopicIds"
	Query_GetNetworkInferencesAtBlock_FullMethodName       = "/emissions.v1.Query/GetNetworkInferencesAtBlock"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	GetLastRewardsUpdate(ctx context.Context, in *QueryLastRewardsUpdateRequest, opts ...grpc.CallOption) (*QueryLastRewardsUpdateResponse, error)
	GetNextTopicId(ctx context.Context, in *QueryNextTopicIdRequest, opts ...grpc.CallOption) (*QueryNextTopicIdResponse, error)
	GetTopic(ctx context.Context, in *QueryTopicRequest, opts ...grpc.CallOption) (*QueryTopicResponse, error)
	GetActiveTopics(ctx context.Context, in *QueryActiveTopicsRequest, opts ...grpc.CallOption) (*QueryActiveTopicsResponse, error)
	GetAllTopics(ctx context.Context, in *QueryAllTopicsRequest, opts ...grpc.CallOption) (*QueryAllTopicsResponse, error)
	GetExistingInferenceRequest(ctx context.Context, in *QueryExistingInferenceRequest, opts ...grpc.CallOption) (*QueryExistingInferenceResponse, error)
	GetAllExistingInferenceRequests(ctx context.Context, in *QueryAllExistingInferenceRequest, opts ...grpc.CallOption) (*QueryAllExistingInferenceResponse, error)
	GetTopicUnmetDemand(ctx context.Context, in *QueryTopicUnmetDemandRequest, opts ...grpc.CallOption) (*QueryTopicUnmetDemandResponse, error)
	GetWorkerLatestInferenceByTopicId(ctx context.Context, in *QueryWorkerLatestInferenceRequest, opts ...grpc.CallOption) (*QueryWorkerLatestInferenceResponse, error)
	GetInferencesAtBlock(ctx context.Context, in *QueryInferencesAtBlockRequest, opts ...grpc.CallOption) (*QueryInferencesAtBlockResponse, error)
	GetForecastsAtBlock(ctx context.Context, in *QueryForecastsAtBlockRequest, opts ...grpc.CallOption) (*QueryForecastsAtBlockResponse, error)
	GetNetworkLossBundleAtBlock(ctx context.Context, in *QueryNetworkLossBundleAtBlockRequest, opts ...grpc.CallOption) (*QueryNetworkLossBundleAtBlockResponse, error)
	GetTotalStake(ctx context.Context, in *QueryTotalStakeRequest, opts ...grpc.CallOption) (*QueryTotalStakeResponse, error)
	GetReputerStakeList(ctx context.Context, in *QueryReputerStakeListRequest, opts ...grpc.CallOption) (*QueryReputerStakeListResponse, error)
	GetTopicStakeList(ctx context.Context, in *QueryTopicStakeListRequest, opts ...grpc.CallOption) (*QueryTopicStakeListResponse, error)
	GetWorkerNodeRegistration(ctx context.Context, in *QueryRegisteredWorkerNodesRequest, opts ...grpc.CallOption) (*QueryRegisteredWorkerNodesResponse, error)
	GetWorkerAddressByP2PKey(ctx context.Context, in *QueryWorkerAddressByP2PKeyRequest, opts ...grpc.CallOption) (*QueryWorkerAddressByP2PKeyResponse, error)
	GetReputerAddressByP2PKey(ctx context.Context, in *QueryReputerAddressByP2PKeyRequest, opts ...grpc.CallOption) (*QueryReputerAddressByP2PKeyResponse, error)
	GetRegisteredTopicIds(ctx context.Context, in *QueryRegisteredTopicIdsRequest, opts ...grpc.CallOption) (*QueryRegisteredTopicIdsResponse, error)
	GetNetworkInferencesAtBlock(ctx context.Context, in *QueryNetworkInferencesAtBlockRequest, opts ...grpc.CallOption) (*QueryNetworkInferencesAtBlockResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLastRewardsUpdate(ctx context.Context, in *QueryLastRewardsUpdateRequest, opts ...grpc.CallOption) (*QueryLastRewardsUpdateResponse, error) {
	out := new(QueryLastRewardsUpdateResponse)
	err := c.cc.Invoke(ctx, Query_GetLastRewardsUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetNextTopicId(ctx context.Context, in *QueryNextTopicIdRequest, opts ...grpc.CallOption) (*QueryNextTopicIdResponse, error) {
	out := new(QueryNextTopicIdResponse)
	err := c.cc.Invoke(ctx, Query_GetNextTopicId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTopic(ctx context.Context, in *QueryTopicRequest, opts ...grpc.CallOption) (*QueryTopicResponse, error) {
	out := new(QueryTopicResponse)
	err := c.cc.Invoke(ctx, Query_GetTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetActiveTopics(ctx context.Context, in *QueryActiveTopicsRequest, opts ...grpc.CallOption) (*QueryActiveTopicsResponse, error) {
	out := new(QueryActiveTopicsResponse)
	err := c.cc.Invoke(ctx, Query_GetActiveTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAllTopics(ctx context.Context, in *QueryAllTopicsRequest, opts ...grpc.CallOption) (*QueryAllTopicsResponse, error) {
	out := new(QueryAllTopicsResponse)
	err := c.cc.Invoke(ctx, Query_GetAllTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetExistingInferenceRequest(ctx context.Context, in *QueryExistingInferenceRequest, opts ...grpc.CallOption) (*QueryExistingInferenceResponse, error) {
	out := new(QueryExistingInferenceResponse)
	err := c.cc.Invoke(ctx, Query_GetExistingInferenceRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAllExistingInferenceRequests(ctx context.Context, in *QueryAllExistingInferenceRequest, opts ...grpc.CallOption) (*QueryAllExistingInferenceResponse, error) {
	out := new(QueryAllExistingInferenceResponse)
	err := c.cc.Invoke(ctx, Query_GetAllExistingInferenceRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTopicUnmetDemand(ctx context.Context, in *QueryTopicUnmetDemandRequest, opts ...grpc.CallOption) (*QueryTopicUnmetDemandResponse, error) {
	out := new(QueryTopicUnmetDemandResponse)
	err := c.cc.Invoke(ctx, Query_GetTopicUnmetDemand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetWorkerLatestInferenceByTopicId(ctx context.Context, in *QueryWorkerLatestInferenceRequest, opts ...grpc.CallOption) (*QueryWorkerLatestInferenceResponse, error) {
	out := new(QueryWorkerLatestInferenceResponse)
	err := c.cc.Invoke(ctx, Query_GetWorkerLatestInferenceByTopicId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetInferencesAtBlock(ctx context.Context, in *QueryInferencesAtBlockRequest, opts ...grpc.CallOption) (*QueryInferencesAtBlockResponse, error) {
	out := new(QueryInferencesAtBlockResponse)
	err := c.cc.Invoke(ctx, Query_GetInferencesAtBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetForecastsAtBlock(ctx context.Context, in *QueryForecastsAtBlockRequest, opts ...grpc.CallOption) (*QueryForecastsAtBlockResponse, error) {
	out := new(QueryForecastsAtBlockResponse)
	err := c.cc.Invoke(ctx, Query_GetForecastsAtBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetNetworkLossBundleAtBlock(ctx context.Context, in *QueryNetworkLossBundleAtBlockRequest, opts ...grpc.CallOption) (*QueryNetworkLossBundleAtBlockResponse, error) {
	out := new(QueryNetworkLossBundleAtBlockResponse)
	err := c.cc.Invoke(ctx, Query_GetNetworkLossBundleAtBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTotalStake(ctx context.Context, in *QueryTotalStakeRequest, opts ...grpc.CallOption) (*QueryTotalStakeResponse, error) {
	out := new(QueryTotalStakeResponse)
	err := c.cc.Invoke(ctx, Query_GetTotalStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetReputerStakeList(ctx context.Context, in *QueryReputerStakeListRequest, opts ...grpc.CallOption) (*QueryReputerStakeListResponse, error) {
	out := new(QueryReputerStakeListResponse)
	err := c.cc.Invoke(ctx, Query_GetReputerStakeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTopicStakeList(ctx context.Context, in *QueryTopicStakeListRequest, opts ...grpc.CallOption) (*QueryTopicStakeListResponse, error) {
	out := new(QueryTopicStakeListResponse)
	err := c.cc.Invoke(ctx, Query_GetTopicStakeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetWorkerNodeRegistration(ctx context.Context, in *QueryRegisteredWorkerNodesRequest, opts ...grpc.CallOption) (*QueryRegisteredWorkerNodesResponse, error) {
	out := new(QueryRegisteredWorkerNodesResponse)
	err := c.cc.Invoke(ctx, Query_GetWorkerNodeRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetWorkerAddressByP2PKey(ctx context.Context, in *QueryWorkerAddressByP2PKeyRequest, opts ...grpc.CallOption) (*QueryWorkerAddressByP2PKeyResponse, error) {
	out := new(QueryWorkerAddressByP2PKeyResponse)
	err := c.cc.Invoke(ctx, Query_GetWorkerAddressByP2PKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetReputerAddressByP2PKey(ctx context.Context, in *QueryReputerAddressByP2PKeyRequest, opts ...grpc.CallOption) (*QueryReputerAddressByP2PKeyResponse, error) {
	out := new(QueryReputerAddressByP2PKeyResponse)
	err := c.cc.Invoke(ctx, Query_GetReputerAddressByP2PKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetRegisteredTopicIds(ctx context.Context, in *QueryRegisteredTopicIdsRequest, opts ...grpc.CallOption) (*QueryRegisteredTopicIdsResponse, error) {
	out := new(QueryRegisteredTopicIdsResponse)
	err := c.cc.Invoke(ctx, Query_GetRegisteredTopicIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetNetworkInferencesAtBlock(ctx context.Context, in *QueryNetworkInferencesAtBlockRequest, opts ...grpc.CallOption) (*QueryNetworkInferencesAtBlockResponse, error) {
	out := new(QueryNetworkInferencesAtBlockResponse)
	err := c.cc.Invoke(ctx, Query_GetNetworkInferencesAtBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params returns the module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	GetLastRewardsUpdate(context.Context, *QueryLastRewardsUpdateRequest) (*QueryLastRewardsUpdateResponse, error)
	GetNextTopicId(context.Context, *QueryNextTopicIdRequest) (*QueryNextTopicIdResponse, error)
	GetTopic(context.Context, *QueryTopicRequest) (*QueryTopicResponse, error)
	GetActiveTopics(context.Context, *QueryActiveTopicsRequest) (*QueryActiveTopicsResponse, error)
	GetAllTopics(context.Context, *QueryAllTopicsRequest) (*QueryAllTopicsResponse, error)
	GetExistingInferenceRequest(context.Context, *QueryExistingInferenceRequest) (*QueryExistingInferenceResponse, error)
	GetAllExistingInferenceRequests(context.Context, *QueryAllExistingInferenceRequest) (*QueryAllExistingInferenceResponse, error)
	GetTopicUnmetDemand(context.Context, *QueryTopicUnmetDemandRequest) (*QueryTopicUnmetDemandResponse, error)
	GetWorkerLatestInferenceByTopicId(context.Context, *QueryWorkerLatestInferenceRequest) (*QueryWorkerLatestInferenceResponse, error)
	GetInferencesAtBlock(context.Context, *QueryInferencesAtBlockRequest) (*QueryInferencesAtBlockResponse, error)
	GetForecastsAtBlock(context.Context, *QueryForecastsAtBlockRequest) (*QueryForecastsAtBlockResponse, error)
	GetNetworkLossBundleAtBlock(context.Context, *QueryNetworkLossBundleAtBlockRequest) (*QueryNetworkLossBundleAtBlockResponse, error)
	GetTotalStake(context.Context, *QueryTotalStakeRequest) (*QueryTotalStakeResponse, error)
	GetReputerStakeList(context.Context, *QueryReputerStakeListRequest) (*QueryReputerStakeListResponse, error)
	GetTopicStakeList(context.Context, *QueryTopicStakeListRequest) (*QueryTopicStakeListResponse, error)
	GetWorkerNodeRegistration(context.Context, *QueryRegisteredWorkerNodesRequest) (*QueryRegisteredWorkerNodesResponse, error)
	GetWorkerAddressByP2PKey(context.Context, *QueryWorkerAddressByP2PKeyRequest) (*QueryWorkerAddressByP2PKeyResponse, error)
	GetReputerAddressByP2PKey(context.Context, *QueryReputerAddressByP2PKeyRequest) (*QueryReputerAddressByP2PKeyResponse, error)
	GetRegisteredTopicIds(context.Context, *QueryRegisteredTopicIdsRequest) (*QueryRegisteredTopicIdsResponse, error)
	GetNetworkInferencesAtBlock(context.Context, *QueryNetworkInferencesAtBlockRequest) (*QueryNetworkInferencesAtBlockResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetLastRewardsUpdate(context.Context, *QueryLastRewardsUpdateRequest) (*QueryLastRewardsUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastRewardsUpdate not implemented")
}
func (UnimplementedQueryServer) GetNextTopicId(context.Context, *QueryNextTopicIdRequest) (*QueryNextTopicIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextTopicId not implemented")
}
func (UnimplementedQueryServer) GetTopic(context.Context, *QueryTopicRequest) (*QueryTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedQueryServer) GetActiveTopics(context.Context, *QueryActiveTopicsRequest) (*QueryActiveTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveTopics not implemented")
}
func (UnimplementedQueryServer) GetAllTopics(context.Context, *QueryAllTopicsRequest) (*QueryAllTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTopics not implemented")
}
func (UnimplementedQueryServer) GetExistingInferenceRequest(context.Context, *QueryExistingInferenceRequest) (*QueryExistingInferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExistingInferenceRequest not implemented")
}
func (UnimplementedQueryServer) GetAllExistingInferenceRequests(context.Context, *QueryAllExistingInferenceRequest) (*QueryAllExistingInferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllExistingInferenceRequests not implemented")
}
func (UnimplementedQueryServer) GetTopicUnmetDemand(context.Context, *QueryTopicUnmetDemandRequest) (*QueryTopicUnmetDemandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicUnmetDemand not implemented")
}
func (UnimplementedQueryServer) GetWorkerLatestInferenceByTopicId(context.Context, *QueryWorkerLatestInferenceRequest) (*QueryWorkerLatestInferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerLatestInferenceByTopicId not implemented")
}
func (UnimplementedQueryServer) GetInferencesAtBlock(context.Context, *QueryInferencesAtBlockRequest) (*QueryInferencesAtBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInferencesAtBlock not implemented")
}
func (UnimplementedQueryServer) GetForecastsAtBlock(context.Context, *QueryForecastsAtBlockRequest) (*QueryForecastsAtBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForecastsAtBlock not implemented")
}
func (UnimplementedQueryServer) GetNetworkLossBundleAtBlock(context.Context, *QueryNetworkLossBundleAtBlockRequest) (*QueryNetworkLossBundleAtBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkLossBundleAtBlock not implemented")
}
func (UnimplementedQueryServer) GetTotalStake(context.Context, *QueryTotalStakeRequest) (*QueryTotalStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalStake not implemented")
}
func (UnimplementedQueryServer) GetReputerStakeList(context.Context, *QueryReputerStakeListRequest) (*QueryReputerStakeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReputerStakeList not implemented")
}
func (UnimplementedQueryServer) GetTopicStakeList(context.Context, *QueryTopicStakeListRequest) (*QueryTopicStakeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicStakeList not implemented")
}
func (UnimplementedQueryServer) GetWorkerNodeRegistration(context.Context, *QueryRegisteredWorkerNodesRequest) (*QueryRegisteredWorkerNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerNodeRegistration not implemented")
}
func (UnimplementedQueryServer) GetWorkerAddressByP2PKey(context.Context, *QueryWorkerAddressByP2PKeyRequest) (*QueryWorkerAddressByP2PKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerAddressByP2PKey not implemented")
}
func (UnimplementedQueryServer) GetReputerAddressByP2PKey(context.Context, *QueryReputerAddressByP2PKeyRequest) (*QueryReputerAddressByP2PKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReputerAddressByP2PKey not implemented")
}
func (UnimplementedQueryServer) GetRegisteredTopicIds(context.Context, *QueryRegisteredTopicIdsRequest) (*QueryRegisteredTopicIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredTopicIds not implemented")
}
func (UnimplementedQueryServer) GetNetworkInferencesAtBlock(context.Context, *QueryNetworkInferencesAtBlockRequest) (*QueryNetworkInferencesAtBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkInferencesAtBlock not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetLastRewardsUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastRewardsUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetLastRewardsUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetLastRewardsUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetLastRewardsUpdate(ctx, req.(*QueryLastRewardsUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetNextTopicId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNextTopicIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetNextTopicId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetNextTopicId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetNextTopicId(ctx, req.(*QueryNextTopicIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTopic(ctx, req.(*QueryTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetActiveTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActiveTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetActiveTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetActiveTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetActiveTopics(ctx, req.(*QueryActiveTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAllTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAllTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetAllTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAllTopics(ctx, req.(*QueryAllTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetExistingInferenceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExistingInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetExistingInferenceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetExistingInferenceRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetExistingInferenceRequest(ctx, req.(*QueryExistingInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAllExistingInferenceRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllExistingInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAllExistingInferenceRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetAllExistingInferenceRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAllExistingInferenceRequests(ctx, req.(*QueryAllExistingInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTopicUnmetDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTopicUnmetDemandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTopicUnmetDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTopicUnmetDemand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTopicUnmetDemand(ctx, req.(*QueryTopicUnmetDemandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetWorkerLatestInferenceByTopicId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWorkerLatestInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetWorkerLatestInferenceByTopicId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetWorkerLatestInferenceByTopicId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetWorkerLatestInferenceByTopicId(ctx, req.(*QueryWorkerLatestInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetInferencesAtBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInferencesAtBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetInferencesAtBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetInferencesAtBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetInferencesAtBlock(ctx, req.(*QueryInferencesAtBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetForecastsAtBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryForecastsAtBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetForecastsAtBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetForecastsAtBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetForecastsAtBlock(ctx, req.(*QueryForecastsAtBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetNetworkLossBundleAtBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNetworkLossBundleAtBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetNetworkLossBundleAtBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetNetworkLossBundleAtBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetNetworkLossBundleAtBlock(ctx, req.(*QueryNetworkLossBundleAtBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTotalStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalStakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTotalStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTotalStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTotalStake(ctx, req.(*QueryTotalStakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetReputerStakeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReputerStakeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetReputerStakeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetReputerStakeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetReputerStakeList(ctx, req.(*QueryReputerStakeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTopicStakeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTopicStakeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTopicStakeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTopicStakeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTopicStakeList(ctx, req.(*QueryTopicStakeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetWorkerNodeRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRegisteredWorkerNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetWorkerNodeRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetWorkerNodeRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetWorkerNodeRegistration(ctx, req.(*QueryRegisteredWorkerNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetWorkerAddressByP2PKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWorkerAddressByP2PKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetWorkerAddressByP2PKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetWorkerAddressByP2PKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetWorkerAddressByP2PKey(ctx, req.(*QueryWorkerAddressByP2PKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetReputerAddressByP2PKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReputerAddressByP2PKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetReputerAddressByP2PKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetReputerAddressByP2PKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetReputerAddressByP2PKey(ctx, req.(*QueryReputerAddressByP2PKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetRegisteredTopicIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRegisteredTopicIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetRegisteredTopicIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetRegisteredTopicIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetRegisteredTopicIds(ctx, req.(*QueryRegisteredTopicIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetNetworkInferencesAtBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNetworkInferencesAtBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetNetworkInferencesAtBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetNetworkInferencesAtBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetNetworkInferencesAtBlock(ctx, req.(*QueryNetworkInferencesAtBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emissions.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetLastRewardsUpdate",
			Handler:    _Query_GetLastRewardsUpdate_Handler,
		},
		{
			MethodName: "GetNextTopicId",
			Handler:    _Query_GetNextTopicId_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _Query_GetTopic_Handler,
		},
		{
			MethodName: "GetActiveTopics",
			Handler:    _Query_GetActiveTopics_Handler,
		},
		{
			MethodName: "GetAllTopics",
			Handler:    _Query_GetAllTopics_Handler,
		},
		{
			MethodName: "GetExistingInferenceRequest",
			Handler:    _Query_GetExistingInferenceRequest_Handler,
		},
		{
			MethodName: "GetAllExistingInferenceRequests",
			Handler:    _Query_GetAllExistingInferenceRequests_Handler,
		},
		{
			MethodName: "GetTopicUnmetDemand",
			Handler:    _Query_GetTopicUnmetDemand_Handler,
		},
		{
			MethodName: "GetWorkerLatestInferenceByTopicId",
			Handler:    _Query_GetWorkerLatestInferenceByTopicId_Handler,
		},
		{
			MethodName: "GetInferencesAtBlock",
			Handler:    _Query_GetInferencesAtBlock_Handler,
		},
		{
			MethodName: "GetForecastsAtBlock",
			Handler:    _Query_GetForecastsAtBlock_Handler,
		},
		{
			MethodName: "GetNetworkLossBundleAtBlock",
			Handler:    _Query_GetNetworkLossBundleAtBlock_Handler,
		},
		{
			MethodName: "GetTotalStake",
			Handler:    _Query_GetTotalStake_Handler,
		},
		{
			MethodName: "GetReputerStakeList",
			Handler:    _Query_GetReputerStakeList_Handler,
		},
		{
			MethodName: "GetTopicStakeList",
			Handler:    _Query_GetTopicStakeList_Handler,
		},
		{
			MethodName: "GetWorkerNodeRegistration",
			Handler:    _Query_GetWorkerNodeRegistration_Handler,
		},
		{
			MethodName: "GetWorkerAddressByP2PKey",
			Handler:    _Query_GetWorkerAddressByP2PKey_Handler,
		},
		{
			MethodName: "GetReputerAddressByP2PKey",
			Handler:    _Query_GetReputerAddressByP2PKey_Handler,
		},
		{
			MethodName: "GetRegisteredTopicIds",
			Handler:    _Query_GetRegisteredTopicIds_Handler,
		},
		{
			MethodName: "GetNetworkInferencesAtBlock",
			Handler:    _Query_GetNetworkInferencesAtBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emissions/v1/query.proto",
}
