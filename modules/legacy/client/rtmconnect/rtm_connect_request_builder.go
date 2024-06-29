package rtmconnect

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RtmConnectRequestBuilder builds and executes requests for operations under \rtm.connect
type RtmConnectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RtmConnectRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RtmConnectRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRtmConnectRequestBuilderInternal instantiates a new RtmConnectRequestBuilder and sets the default values.
func NewRtmConnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RtmConnectRequestBuilder) {
    m := &RtmConnectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/rtm.connect", pathParameters),
    }
    return m
}
// NewRtmConnectRequestBuilder instantiates a new RtmConnectRequestBuilder and sets the default values.
func NewRtmConnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RtmConnectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRtmConnectRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this call is to get endpoint information for realtime engine and realtime auth token.
// Deprecated: This method is obsolete. Use GetAsConnectGetResponse instead.
func (m *RtmConnectRequestBuilder) Get(ctx context.Context, requestConfiguration *RtmConnectRequestBuilderGetRequestConfiguration)(ConnectResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateConnect400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConnectResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConnectResponseable), nil
}
// GetAsConnectGetResponse this call is to get endpoint information for realtime engine and realtime auth token.
func (m *RtmConnectRequestBuilder) GetAsConnectGetResponse(ctx context.Context, requestConfiguration *RtmConnectRequestBuilderGetRequestConfiguration)(ConnectGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateConnect400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConnectGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConnectGetResponseable), nil
}
// ToGetRequestInformation this call is to get endpoint information for realtime engine and realtime auth token.
func (m *RtmConnectRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RtmConnectRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *RtmConnectRequestBuilder) WithUrl(rawUrl string)(*RtmConnectRequestBuilder) {
    return NewRtmConnectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
