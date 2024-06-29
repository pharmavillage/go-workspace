package channelactivate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelActivateRequestBuilder builds and executes requests for operations under \channel.activate
type ChannelActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelActivateRequestBuilderInternal instantiates a new ChannelActivateRequestBuilder and sets the default values.
func NewChannelActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelActivateRequestBuilder) {
    m := &ChannelActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.activate", pathParameters),
    }
    return m
}
// NewChannelActivateRequestBuilder instantiates a new ChannelActivateRequestBuilder and sets the default values.
func NewChannelActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a channel that was previously closed. Only owner or admins activate a channel.
// Deprecated: This method is obsolete. Use PostAsActivatePostResponse instead.
func (m *ChannelActivateRequestBuilder) Post(ctx context.Context, body ActivatePostRequestBodyable, requestConfiguration *ChannelActivateRequestBuilderPostRequestConfiguration)(ActivateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateActivate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateActivateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ActivateResponseable), nil
}
// PostAsActivatePostResponse activate a channel that was previously closed. Only owner or admins activate a channel.
func (m *ChannelActivateRequestBuilder) PostAsActivatePostResponse(ctx context.Context, body ActivatePostRequestBodyable, requestConfiguration *ChannelActivateRequestBuilderPostRequestConfiguration)(ActivatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateActivate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateActivatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ActivatePostResponseable), nil
}
// ToPostRequestInformation activate a channel that was previously closed. Only owner or admins activate a channel.
func (m *ChannelActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ActivatePostRequestBodyable, requestConfiguration *ChannelActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelActivateRequestBuilder) WithUrl(rawUrl string)(*ChannelActivateRequestBuilder) {
    return NewChannelActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
