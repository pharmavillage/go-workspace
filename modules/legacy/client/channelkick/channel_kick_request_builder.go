package channelkick

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelKickRequestBuilder builds and executes requests for operations under \channel.kick
type ChannelKickRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelKickRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelKickRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelKickRequestBuilderInternal instantiates a new ChannelKickRequestBuilder and sets the default values.
func NewChannelKickRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelKickRequestBuilder) {
    m := &ChannelKickRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.kick", pathParameters),
    }
    return m
}
// NewChannelKickRequestBuilder instantiates a new ChannelKickRequestBuilder and sets the default values.
func NewChannelKickRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelKickRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelKickRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove a member from a channel. Only owner or Admins can remove members. Channel owners cannot be kicked out of the channel.
// Deprecated: This method is obsolete. Use PostAsKickPostResponse instead.
func (m *ChannelKickRequestBuilder) Post(ctx context.Context, body KickPostRequestBodyable, requestConfiguration *ChannelKickRequestBuilderPostRequestConfiguration)(KickResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateKick400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateKickResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(KickResponseable), nil
}
// PostAsKickPostResponse remove a member from a channel. Only owner or Admins can remove members. Channel owners cannot be kicked out of the channel.
func (m *ChannelKickRequestBuilder) PostAsKickPostResponse(ctx context.Context, body KickPostRequestBodyable, requestConfiguration *ChannelKickRequestBuilderPostRequestConfiguration)(KickPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateKick400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateKickPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(KickPostResponseable), nil
}
// ToPostRequestInformation remove a member from a channel. Only owner or Admins can remove members. Channel owners cannot be kicked out of the channel.
func (m *ChannelKickRequestBuilder) ToPostRequestInformation(ctx context.Context, body KickPostRequestBodyable, requestConfiguration *ChannelKickRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelKickRequestBuilder) WithUrl(rawUrl string)(*ChannelKickRequestBuilder) {
    return NewChannelKickRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
