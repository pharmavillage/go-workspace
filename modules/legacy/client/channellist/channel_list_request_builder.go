package channellist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelListRequestBuilder builds and executes requests for operations under \channel.list
type ChannelListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelListRequestBuilderGetQueryParameters list all channel for a user
type ChannelListRequestBuilderGetQueryParameters struct {
    // Include or exclude closed channels in the list response. If not supplied, it is considered to be true
    Exclude_closed *bool `uriparametername:"exclude_closed"`
}
// ChannelListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelListRequestBuilderGetQueryParameters
}
// NewChannelListRequestBuilderInternal instantiates a new ChannelListRequestBuilder and sets the default values.
func NewChannelListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelListRequestBuilder) {
    m := &ChannelListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.list{?exclude_closed*}", pathParameters),
    }
    return m
}
// NewChannelListRequestBuilder instantiates a new ChannelListRequestBuilder and sets the default values.
func NewChannelListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all channel for a user
// Deprecated: This method is obsolete. Use GetAsListGetResponse instead.
func (m *ChannelListRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelListRequestBuilderGetRequestConfiguration)(ListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListResponseable), nil
}
// GetAsListGetResponse list all channel for a user
func (m *ChannelListRequestBuilder) GetAsListGetResponse(ctx context.Context, requestConfiguration *ChannelListRequestBuilderGetRequestConfiguration)(ListGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListGetResponseable), nil
}
// ToGetRequestInformation list all channel for a user
func (m *ChannelListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelListRequestBuilder) WithUrl(rawUrl string)(*ChannelListRequestBuilder) {
    return NewChannelListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
