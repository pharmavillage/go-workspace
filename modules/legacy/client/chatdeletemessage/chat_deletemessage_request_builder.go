package chatdeletemessage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChatDeletemessageRequestBuilder builds and executes requests for operations under \chat.deletemessage
type ChatDeletemessageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChatDeletemessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatDeletemessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatDeletemessageRequestBuilderInternal instantiates a new ChatDeletemessageRequestBuilder and sets the default values.
func NewChatDeletemessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatDeletemessageRequestBuilder) {
    m := &ChatDeletemessageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chat.deletemessage", pathParameters),
    }
    return m
}
// NewChatDeletemessageRequestBuilder instantiates a new ChatDeletemessageRequestBuilder and sets the default values.
func NewChatDeletemessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatDeletemessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatDeletemessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete a previously posted message in a channel. Only the creator of the message can delete it.chat.deleteMessage event will be sent to other users in realtime socket
// Deprecated: This method is obsolete. Use PostAsDeletemessagePostResponse instead.
func (m *ChatDeletemessageRequestBuilder) Post(ctx context.Context, body DeletemessagePostRequestBodyable, requestConfiguration *ChatDeletemessageRequestBuilderPostRequestConfiguration)(DeletemessageResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateDeletemessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletemessageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletemessageResponseable), nil
}
// PostAsDeletemessagePostResponse delete a previously posted message in a channel. Only the creator of the message can delete it.chat.deleteMessage event will be sent to other users in realtime socket
func (m *ChatDeletemessageRequestBuilder) PostAsDeletemessagePostResponse(ctx context.Context, body DeletemessagePostRequestBodyable, requestConfiguration *ChatDeletemessageRequestBuilderPostRequestConfiguration)(DeletemessagePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateDeletemessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletemessagePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletemessagePostResponseable), nil
}
// ToPostRequestInformation delete a previously posted message in a channel. Only the creator of the message can delete it.chat.deleteMessage event will be sent to other users in realtime socket
func (m *ChatDeletemessageRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeletemessagePostRequestBodyable, requestConfiguration *ChatDeletemessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatDeletemessageRequestBuilder) WithUrl(rawUrl string)(*ChatDeletemessageRequestBuilder) {
    return NewChatDeletemessageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
