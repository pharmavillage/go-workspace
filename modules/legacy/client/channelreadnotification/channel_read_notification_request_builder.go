package channelreadnotification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelReadNotificationRequestBuilder builds and executes requests for operations under \channel.read-notification
type ChannelReadNotificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelReadNotificationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelReadNotificationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelReadNotificationRequestBuilderInternal instantiates a new ChannelReadNotificationRequestBuilder and sets the default values.
func NewChannelReadNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelReadNotificationRequestBuilder) {
    m := &ChannelReadNotificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.read-notification", pathParameters),
    }
    return m
}
// NewChannelReadNotificationRequestBuilder instantiates a new ChannelReadNotificationRequestBuilder and sets the default values.
func NewChannelReadNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelReadNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelReadNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark one or all message as read in the channel.
// Deprecated: This method is obsolete. Use PostAsReadNotificationPostResponse instead.
func (m *ChannelReadNotificationRequestBuilder) Post(ctx context.Context, body ReadNotificationPostRequestBodyable, requestConfiguration *ChannelReadNotificationRequestBuilderPostRequestConfiguration)(ReadNotificationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReadNotification400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReadNotificationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReadNotificationResponseable), nil
}
// PostAsReadNotificationPostResponse mark one or all message as read in the channel.
func (m *ChannelReadNotificationRequestBuilder) PostAsReadNotificationPostResponse(ctx context.Context, body ReadNotificationPostRequestBodyable, requestConfiguration *ChannelReadNotificationRequestBuilderPostRequestConfiguration)(ReadNotificationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReadNotification400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReadNotificationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReadNotificationPostResponseable), nil
}
// ToPostRequestInformation mark one or all message as read in the channel.
func (m *ChannelReadNotificationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReadNotificationPostRequestBodyable, requestConfiguration *ChannelReadNotificationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelReadNotificationRequestBuilder) WithUrl(rawUrl string)(*ChannelReadNotificationRequestBuilder) {
    return NewChannelReadNotificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
