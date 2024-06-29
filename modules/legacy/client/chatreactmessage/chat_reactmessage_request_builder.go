package chatreactmessage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChatReactmessageRequestBuilder builds and executes requests for operations under \chat.reactmessage
type ChatReactmessageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChatReactmessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatReactmessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatReactmessageRequestBuilderInternal instantiates a new ChatReactmessageRequestBuilder and sets the default values.
func NewChatReactmessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatReactmessageRequestBuilder) {
    m := &ChatReactmessageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chat.reactmessage", pathParameters),
    }
    return m
}
// NewChatReactmessageRequestBuilder instantiates a new ChatReactmessageRequestBuilder and sets the default values.
func NewChatReactmessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatReactmessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatReactmessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post react to or remove reaction from a message already posted.chat.updateMessage event will be sent to other users in realtime socket and alert will be generated to the owner of the message
// Deprecated: This method is obsolete. Use PostAsReactmessagePostResponse instead.
func (m *ChatReactmessageRequestBuilder) Post(ctx context.Context, body ReactmessagePostRequestBodyable, requestConfiguration *ChatReactmessageRequestBuilderPostRequestConfiguration)(ReactmessageResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReactmessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReactmessageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReactmessageResponseable), nil
}
// PostAsReactmessagePostResponse react to or remove reaction from a message already posted.chat.updateMessage event will be sent to other users in realtime socket and alert will be generated to the owner of the message
func (m *ChatReactmessageRequestBuilder) PostAsReactmessagePostResponse(ctx context.Context, body ReactmessagePostRequestBodyable, requestConfiguration *ChatReactmessageRequestBuilderPostRequestConfiguration)(ReactmessagePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReactmessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReactmessagePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReactmessagePostResponseable), nil
}
// ToPostRequestInformation react to or remove reaction from a message already posted.chat.updateMessage event will be sent to other users in realtime socket and alert will be generated to the owner of the message
func (m *ChatReactmessageRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReactmessagePostRequestBodyable, requestConfiguration *ChatReactmessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatReactmessageRequestBuilder) WithUrl(rawUrl string)(*ChatReactmessageRequestBuilder) {
    return NewChatReactmessageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
