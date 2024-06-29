package channelrename

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelRenameRequestBuilder builds and executes requests for operations under \channel.rename
type ChannelRenameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelRenameRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelRenameRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelRenameRequestBuilderInternal instantiates a new ChannelRenameRequestBuilder and sets the default values.
func NewChannelRenameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelRenameRequestBuilder) {
    m := &ChannelRenameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.rename", pathParameters),
    }
    return m
}
// NewChannelRenameRequestBuilder instantiates a new ChannelRenameRequestBuilder and sets the default values.
func NewChannelRenameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelRenameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelRenameRequestBuilderInternal(urlParams, requestAdapter)
}
// Post rename an existing channel. Only channel owner or Admin can rename a channel. Deprecaded, since the `channel.update` endpoint is able to rename a channel.
// Deprecated: This method is obsolete. Use PostAsRenamePostResponse instead.
func (m *ChannelRenameRequestBuilder) Post(ctx context.Context, body RenamePostRequestBodyable, requestConfiguration *ChannelRenameRequestBuilderPostRequestConfiguration)(RenameResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRename400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRenameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RenameResponseable), nil
}
// PostAsRenamePostResponse rename an existing channel. Only channel owner or Admin can rename a channel. Deprecaded, since the `channel.update` endpoint is able to rename a channel.
// Deprecated: 
func (m *ChannelRenameRequestBuilder) PostAsRenamePostResponse(ctx context.Context, body RenamePostRequestBodyable, requestConfiguration *ChannelRenameRequestBuilderPostRequestConfiguration)(RenamePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRename400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRenamePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RenamePostResponseable), nil
}
// ToPostRequestInformation rename an existing channel. Only channel owner or Admin can rename a channel. Deprecaded, since the `channel.update` endpoint is able to rename a channel.
// Deprecated: 
func (m *ChannelRenameRequestBuilder) ToPostRequestInformation(ctx context.Context, body RenamePostRequestBodyable, requestConfiguration *ChannelRenameRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
func (m *ChannelRenameRequestBuilder) WithUrl(rawUrl string)(*ChannelRenameRequestBuilder) {
    return NewChannelRenameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
