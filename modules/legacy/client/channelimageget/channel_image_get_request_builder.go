package channelimageget

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelImageGetRequestBuilder builds and executes requests for operations under \channel.image.get
type ChannelImageGetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelImageGetRequestBuilderGetQueryParameters get logo or background image asset for a channel
type ChannelImageGetRequestBuilderGetQueryParameters struct {
    // Has to be logo or background
    Channel_asset_type *string `uriparametername:"channel_asset_type"`
    // The channel id image asset to retrieve
    Channel_id *int32 `uriparametername:"channel_id"`
}
// ChannelImageGetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelImageGetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelImageGetRequestBuilderGetQueryParameters
}
// NewChannelImageGetRequestBuilderInternal instantiates a new ChannelImageGetRequestBuilder and sets the default values.
func NewChannelImageGetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelImageGetRequestBuilder) {
    m := &ChannelImageGetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.image.get?channel_asset_type={channel_asset_type}&channel_id={channel_id}", pathParameters),
    }
    return m
}
// NewChannelImageGetRequestBuilder instantiates a new ChannelImageGetRequestBuilder and sets the default values.
func NewChannelImageGetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelImageGetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelImageGetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get logo or background image asset for a channel
func (m *ChannelImageGetRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelImageGetRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateGet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get logo or background image asset for a channel
func (m *ChannelImageGetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelImageGetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "image/*, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelImageGetRequestBuilder) WithUrl(rawUrl string)(*ChannelImageGetRequestBuilder) {
    return NewChannelImageGetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
