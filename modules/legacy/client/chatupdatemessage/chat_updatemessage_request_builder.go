package chatupdatemessage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChatUpdatemessageRequestBuilder builds and executes requests for operations under \chat.updatemessage
type ChatUpdatemessageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChatUpdatemessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatUpdatemessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatUpdatemessageRequestBuilderInternal instantiates a new ChatUpdatemessageRequestBuilder and sets the default values.
func NewChatUpdatemessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatUpdatemessageRequestBuilder) {
    m := &ChatUpdatemessageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chat.updatemessage", pathParameters),
    }
    return m
}
// NewChatUpdatemessageRequestBuilder instantiates a new ChatUpdatemessageRequestBuilder and sets the default values.
func NewChatUpdatemessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatUpdatemessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatUpdatemessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update a message already posted. Only the creator of the message can update it.chat.updateMessage event will be sent to other users in realtime socket
// Deprecated: This method is obsolete. Use PostAsUpdatemessagePostResponse instead.
func (m *ChatUpdatemessageRequestBuilder) Post(ctx context.Context, body UpdatemessagePostRequestBodyable, requestConfiguration *ChatUpdatemessageRequestBuilderPostRequestConfiguration)(UpdatemessageResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdatemessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdatemessageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdatemessageResponseable), nil
}
// PostAsUpdatemessagePostResponse update a message already posted. Only the creator of the message can update it.chat.updateMessage event will be sent to other users in realtime socket
func (m *ChatUpdatemessageRequestBuilder) PostAsUpdatemessagePostResponse(ctx context.Context, body UpdatemessagePostRequestBodyable, requestConfiguration *ChatUpdatemessageRequestBuilderPostRequestConfiguration)(UpdatemessagePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdatemessage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdatemessagePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdatemessagePostResponseable), nil
}
// ToPostRequestInformation update a message already posted. Only the creator of the message can update it.chat.updateMessage event will be sent to other users in realtime socket
func (m *ChatUpdatemessageRequestBuilder) ToPostRequestInformation(ctx context.Context, body UpdatemessagePostRequestBodyable, requestConfiguration *ChatUpdatemessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatUpdatemessageRequestBuilder) WithUrl(rawUrl string)(*ChatUpdatemessageRequestBuilder) {
    return NewChatUpdatemessageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
