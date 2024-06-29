package channeloneonone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelOneOnOneRequestBuilder builds and executes requests for operations under \channel.one-on-one
type ChannelOneOnOneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelOneOnOneRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelOneOnOneRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelOneOnOneRequestBuilderInternal instantiates a new ChannelOneOnOneRequestBuilder and sets the default values.
func NewChannelOneOnOneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelOneOnOneRequestBuilder) {
    m := &ChannelOneOnOneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.one-on-one", pathParameters),
    }
    return m
}
// NewChannelOneOnOneRequestBuilder instantiates a new ChannelOneOnOneRequestBuilder and sets the default values.
func NewChannelOneOnOneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelOneOnOneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelOneOnOneRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return the 1x1 channel for the provided user with the logged user. If it don't exists, creates it.
// Deprecated: This method is obsolete. Use PostAsOneOnOnePostResponse instead.
func (m *ChannelOneOnOneRequestBuilder) Post(ctx context.Context, body OneOnOnePostRequestBodyable, requestConfiguration *ChannelOneOnOneRequestBuilderPostRequestConfiguration)(OneOnOneResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateOneOnOne400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOneOnOneResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OneOnOneResponseable), nil
}
// PostAsOneOnOnePostResponse return the 1x1 channel for the provided user with the logged user. If it don't exists, creates it.
func (m *ChannelOneOnOneRequestBuilder) PostAsOneOnOnePostResponse(ctx context.Context, body OneOnOnePostRequestBodyable, requestConfiguration *ChannelOneOnOneRequestBuilderPostRequestConfiguration)(OneOnOnePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateOneOnOne400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOneOnOnePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OneOnOnePostResponseable), nil
}
// ToPostRequestInformation return the 1x1 channel for the provided user with the logged user. If it don't exists, creates it.
func (m *ChannelOneOnOneRequestBuilder) ToPostRequestInformation(ctx context.Context, body OneOnOnePostRequestBodyable, requestConfiguration *ChannelOneOnOneRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelOneOnOneRequestBuilder) WithUrl(rawUrl string)(*ChannelOneOnOneRequestBuilder) {
    return NewChannelOneOnOneRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
