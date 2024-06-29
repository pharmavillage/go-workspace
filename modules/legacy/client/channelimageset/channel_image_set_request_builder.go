package channelimageset

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelImageSetRequestBuilder builds and executes requests for operations under \channel.image.set
type ChannelImageSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelImageSetRequestBuilderPostQueryParameters set logo or background image of a channel. Only owner or admin can set the channel asset.
type ChannelImageSetRequestBuilderPostQueryParameters struct {
    // Has to be logo or background
    Channel_asset_type *string `uriparametername:"channel_asset_type"`
    // The channel id image asset to store
    Channel_id *int32 `uriparametername:"channel_id"`
    // Clear the specified asset type if set. If specified, any files sent with this request will be ignored
    Clear_asset *bool `uriparametername:"clear_asset"`
}
// ChannelImageSetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelImageSetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelImageSetRequestBuilderPostQueryParameters
}
// NewChannelImageSetRequestBuilderInternal instantiates a new ChannelImageSetRequestBuilder and sets the default values.
func NewChannelImageSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelImageSetRequestBuilder) {
    m := &ChannelImageSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.image.set?channel_asset_type={channel_asset_type}&channel_id={channel_id}{&clear_asset*}", pathParameters),
    }
    return m
}
// NewChannelImageSetRequestBuilder instantiates a new ChannelImageSetRequestBuilder and sets the default values.
func NewChannelImageSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelImageSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelImageSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set logo or background image of a channel. Only owner or admin can set the channel asset.
// Deprecated: This method is obsolete. Use PostAsSetPostResponse instead.
func (m *ChannelImageSetRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *ChannelImageSetRequestBuilderPostRequestConfiguration)(SetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetResponseable), nil
}
// PostAsSetPostResponse set logo or background image of a channel. Only owner or admin can set the channel asset.
func (m *ChannelImageSetRequestBuilder) PostAsSetPostResponse(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *ChannelImageSetRequestBuilderPostRequestConfiguration)(SetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetPostResponseable), nil
}
// ToPostRequestInformation set logo or background image of a channel. Only owner or admin can set the channel asset.
func (m *ChannelImageSetRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *ChannelImageSetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelImageSetRequestBuilder) WithUrl(rawUrl string)(*ChannelImageSetRequestBuilder) {
    return NewChannelImageSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
