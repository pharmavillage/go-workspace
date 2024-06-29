package channelclose

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelCloseRequestBuilder builds and executes requests for operations under \channel.close
type ChannelCloseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelCloseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelCloseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelCloseRequestBuilderInternal instantiates a new ChannelCloseRequestBuilder and sets the default values.
func NewChannelCloseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelCloseRequestBuilder) {
    m := &ChannelCloseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.close", pathParameters),
    }
    return m
}
// NewChannelCloseRequestBuilder instantiates a new ChannelCloseRequestBuilder and sets the default values.
func NewChannelCloseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelCloseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelCloseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post close a Channel. A closed channel will be moved out of active list. Any action such as posting message or adding new user will reactivate the channel.Only channel owner or Admin can close a channel
// Deprecated: This method is obsolete. Use PostAsClosePostResponse instead.
func (m *ChannelCloseRequestBuilder) Post(ctx context.Context, body ClosePostRequestBodyable, requestConfiguration *ChannelCloseRequestBuilderPostRequestConfiguration)(CloseResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateClose400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloseResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloseResponseable), nil
}
// PostAsClosePostResponse close a Channel. A closed channel will be moved out of active list. Any action such as posting message or adding new user will reactivate the channel.Only channel owner or Admin can close a channel
func (m *ChannelCloseRequestBuilder) PostAsClosePostResponse(ctx context.Context, body ClosePostRequestBodyable, requestConfiguration *ChannelCloseRequestBuilderPostRequestConfiguration)(ClosePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateClose400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateClosePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ClosePostResponseable), nil
}
// ToPostRequestInformation close a Channel. A closed channel will be moved out of active list. Any action such as posting message or adding new user will reactivate the channel.Only channel owner or Admin can close a channel
func (m *ChannelCloseRequestBuilder) ToPostRequestInformation(ctx context.Context, body ClosePostRequestBodyable, requestConfiguration *ChannelCloseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelCloseRequestBuilder) WithUrl(rawUrl string)(*ChannelCloseRequestBuilder) {
    return NewChannelCloseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
