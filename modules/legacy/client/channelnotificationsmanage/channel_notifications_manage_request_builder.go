package channelnotificationsmanage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelNotificationsManageRequestBuilder builds and executes requests for operations under \channel.notifications.manage
type ChannelNotificationsManageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelNotificationsManageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelNotificationsManageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelNotificationsManageRequestBuilderInternal instantiates a new ChannelNotificationsManageRequestBuilder and sets the default values.
func NewChannelNotificationsManageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelNotificationsManageRequestBuilder) {
    m := &ChannelNotificationsManageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.notifications.manage", pathParameters),
    }
    return m
}
// NewChannelNotificationsManageRequestBuilder instantiates a new ChannelNotificationsManageRequestBuilder and sets the default values.
func NewChannelNotificationsManageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelNotificationsManageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelNotificationsManageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the subscribing level for notifications on the channel, for a particular user.
// Deprecated: This method is obsolete. Use PostAsManagePostResponse instead.
func (m *ChannelNotificationsManageRequestBuilder) Post(ctx context.Context, body ManagePostRequestBodyable, requestConfiguration *ChannelNotificationsManageRequestBuilderPostRequestConfiguration)(ManageResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateManage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageResponseable), nil
}
// PostAsManagePostResponse set the subscribing level for notifications on the channel, for a particular user.
func (m *ChannelNotificationsManageRequestBuilder) PostAsManagePostResponse(ctx context.Context, body ManagePostRequestBodyable, requestConfiguration *ChannelNotificationsManageRequestBuilderPostRequestConfiguration)(ManagePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateManage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagePostResponseable), nil
}
// ToPostRequestInformation set the subscribing level for notifications on the channel, for a particular user.
func (m *ChannelNotificationsManageRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagePostRequestBodyable, requestConfiguration *ChannelNotificationsManageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelNotificationsManageRequestBuilder) WithUrl(rawUrl string)(*ChannelNotificationsManageRequestBuilder) {
    return NewChannelNotificationsManageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
