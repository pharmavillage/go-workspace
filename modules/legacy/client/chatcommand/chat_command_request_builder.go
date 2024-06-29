package chatcommand

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChatCommandRequestBuilder builds and executes requests for operations under \chat.command
type ChatCommandRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChatCommandRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatCommandRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatCommandRequestBuilderInternal instantiates a new ChatCommandRequestBuilder and sets the default values.
func NewChatCommandRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatCommandRequestBuilder) {
    m := &ChatCommandRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chat.command", pathParameters),
    }
    return m
}
// NewChatCommandRequestBuilder instantiates a new ChatCommandRequestBuilder and sets the default values.
func NewChatCommandRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatCommandRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatCommandRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send a slash command inside a channel.The slash command signature is parsed and validated on the backend.
// Deprecated: This method is obsolete. Use PostAsCommandPostResponse instead.
func (m *ChatCommandRequestBuilder) Post(ctx context.Context, body CommandPostRequestBodyable, requestConfiguration *ChatCommandRequestBuilderPostRequestConfiguration)(CommandResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCommand400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCommandResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CommandResponseable), nil
}
// PostAsCommandPostResponse send a slash command inside a channel.The slash command signature is parsed and validated on the backend.
func (m *ChatCommandRequestBuilder) PostAsCommandPostResponse(ctx context.Context, body CommandPostRequestBodyable, requestConfiguration *ChatCommandRequestBuilderPostRequestConfiguration)(CommandPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCommand400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCommandPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CommandPostResponseable), nil
}
// ToPostRequestInformation send a slash command inside a channel.The slash command signature is parsed and validated on the backend.
func (m *ChatCommandRequestBuilder) ToPostRequestInformation(ctx context.Context, body CommandPostRequestBodyable, requestConfiguration *ChatCommandRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatCommandRequestBuilder) WithUrl(rawUrl string)(*ChatCommandRequestBuilder) {
    return NewChatCommandRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
