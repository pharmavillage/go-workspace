package channelinfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelInfoRequestBuilder builds and executes requests for operations under \channel.info
type ChannelInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelInfoRequestBuilderGetQueryParameters get information about a channel.This endpoint can be accessed using a notification token (without been authenticated).
type ChannelInfoRequestBuilderGetQueryParameters struct {
    // Id of channel to get information
    Channel_id *int32 `uriparametername:"channel_id"`
}
// ChannelInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelInfoRequestBuilderGetQueryParameters
}
// NewChannelInfoRequestBuilderInternal instantiates a new ChannelInfoRequestBuilder and sets the default values.
func NewChannelInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelInfoRequestBuilder) {
    m := &ChannelInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.info?channel_id={channel_id}", pathParameters),
    }
    return m
}
// NewChannelInfoRequestBuilder instantiates a new ChannelInfoRequestBuilder and sets the default values.
func NewChannelInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get information about a channel.This endpoint can be accessed using a notification token (without been authenticated).
// Deprecated: This method is obsolete. Use GetAsInfoGetResponse instead.
func (m *ChannelInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelInfoRequestBuilderGetRequestConfiguration)(InfoResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoResponseable), nil
}
// GetAsInfoGetResponse get information about a channel.This endpoint can be accessed using a notification token (without been authenticated).
func (m *ChannelInfoRequestBuilder) GetAsInfoGetResponse(ctx context.Context, requestConfiguration *ChannelInfoRequestBuilderGetRequestConfiguration)(InfoGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoGetResponseable), nil
}
// ToGetRequestInformation get information about a channel.This endpoint can be accessed using a notification token (without been authenticated).
func (m *ChannelInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelInfoRequestBuilder) WithUrl(rawUrl string)(*ChannelInfoRequestBuilder) {
    return NewChannelInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
