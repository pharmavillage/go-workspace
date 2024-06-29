package channelleave

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelLeaveRequestBuilder builds and executes requests for operations under \channel.leave
type ChannelLeaveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelLeaveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelLeaveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelLeaveRequestBuilderInternal instantiates a new ChannelLeaveRequestBuilder and sets the default values.
func NewChannelLeaveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelLeaveRequestBuilder) {
    m := &ChannelLeaveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.leave", pathParameters),
    }
    return m
}
// NewChannelLeaveRequestBuilder instantiates a new ChannelLeaveRequestBuilder and sets the default values.
func NewChannelLeaveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelLeaveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelLeaveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post leave from a Channel. Channel owner cannot leave the channel.
// Deprecated: This method is obsolete. Use PostAsLeavePostResponse instead.
func (m *ChannelLeaveRequestBuilder) Post(ctx context.Context, body LeavePostRequestBodyable, requestConfiguration *ChannelLeaveRequestBuilderPostRequestConfiguration)(LeaveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateLeave400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLeaveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LeaveResponseable), nil
}
// PostAsLeavePostResponse leave from a Channel. Channel owner cannot leave the channel.
func (m *ChannelLeaveRequestBuilder) PostAsLeavePostResponse(ctx context.Context, body LeavePostRequestBodyable, requestConfiguration *ChannelLeaveRequestBuilderPostRequestConfiguration)(LeavePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateLeave400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLeavePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LeavePostResponseable), nil
}
// ToPostRequestInformation leave from a Channel. Channel owner cannot leave the channel.
func (m *ChannelLeaveRequestBuilder) ToPostRequestInformation(ctx context.Context, body LeavePostRequestBodyable, requestConfiguration *ChannelLeaveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelLeaveRequestBuilder) WithUrl(rawUrl string)(*ChannelLeaveRequestBuilder) {
    return NewChannelLeaveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
