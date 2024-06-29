package chatpostmessage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChatPostmessageRequestBuilder builds and executes requests for operations under \chat.postmessage
type ChatPostmessageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChatPostmessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatPostmessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatPostmessageRequestBuilderInternal instantiates a new ChatPostmessageRequestBuilder and sets the default values.
func NewChatPostmessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatPostmessageRequestBuilder) {
    m := &ChatPostmessageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chat.postmessage", pathParameters),
    }
    return m
}
// NewChatPostmessageRequestBuilder instantiates a new ChatPostmessageRequestBuilder and sets the default values.
func NewChatPostmessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatPostmessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatPostmessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post post a message to a channel. User must have write level permission in the channel to post.chat.postMessage event will be sent to other users in realtime socket
// Deprecated: This method is obsolete. Use PostAsPostmessagePostResponse instead.
func (m *ChatPostmessageRequestBuilder) Post(ctx context.Context, body PostmessagePostRequestBodyable, requestConfiguration *ChatPostmessageRequestBuilderPostRequestConfiguration)(PostmessageResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreatePostmessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePostmessageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PostmessageResponseable), nil
}
// PostAsPostmessagePostResponse post a message to a channel. User must have write level permission in the channel to post.chat.postMessage event will be sent to other users in realtime socket
func (m *ChatPostmessageRequestBuilder) PostAsPostmessagePostResponse(ctx context.Context, body PostmessagePostRequestBodyable, requestConfiguration *ChatPostmessageRequestBuilderPostRequestConfiguration)(PostmessagePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreatePostmessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePostmessagePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PostmessagePostResponseable), nil
}
// ToPostRequestInformation post a message to a channel. User must have write level permission in the channel to post.chat.postMessage event will be sent to other users in realtime socket
func (m *ChatPostmessageRequestBuilder) ToPostRequestInformation(ctx context.Context, body PostmessagePostRequestBodyable, requestConfiguration *ChatPostmessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatPostmessageRequestBuilder) WithUrl(rawUrl string)(*ChatPostmessageRequestBuilder) {
    return NewChatPostmessageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
