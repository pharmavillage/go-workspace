package channelremove

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelRemoveRequestBuilder builds and executes requests for operations under \channel.remove
type ChannelRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelRemoveRequestBuilderInternal instantiates a new ChannelRemoveRequestBuilder and sets the default values.
func NewChannelRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelRemoveRequestBuilder) {
    m := &ChannelRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.remove", pathParameters),
    }
    return m
}
// NewChannelRemoveRequestBuilder instantiates a new ChannelRemoveRequestBuilder and sets the default values.
func NewChannelRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete a Channel. Channel and all its associated content will be deleted. Once a channel is deleted, it cannot be recovered. Only admin or owner can delete a channel.    NOTE: Channel delete can take a some time after this call returns success.
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
func (m *ChannelRemoveRequestBuilder) Post(ctx context.Context, body RemovePostRequestBodyable, requestConfiguration *ChannelRemoveRequestBuilderPostRequestConfiguration)(RemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRemove400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RemoveResponseable), nil
}
// PostAsRemovePostResponse delete a Channel. Channel and all its associated content will be deleted. Once a channel is deleted, it cannot be recovered. Only admin or owner can delete a channel.    NOTE: Channel delete can take a some time after this call returns success.
func (m *ChannelRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body RemovePostRequestBodyable, requestConfiguration *ChannelRemoveRequestBuilderPostRequestConfiguration)(RemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRemove400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RemovePostResponseable), nil
}
// ToPostRequestInformation delete a Channel. Channel and all its associated content will be deleted. Once a channel is deleted, it cannot be recovered. Only admin or owner can delete a channel.    NOTE: Channel delete can take a some time after this call returns success.
func (m *ChannelRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body RemovePostRequestBodyable, requestConfiguration *ChannelRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelRemoveRequestBuilder) WithUrl(rawUrl string)(*ChannelRemoveRequestBuilder) {
    return NewChannelRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
