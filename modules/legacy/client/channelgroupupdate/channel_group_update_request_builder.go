package channelgroupupdate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelGroupUpdateRequestBuilder builds and executes requests for operations under \channel.group.update
type ChannelGroupUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelGroupUpdateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelGroupUpdateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelGroupUpdateRequestBuilderInternal instantiates a new ChannelGroupUpdateRequestBuilder and sets the default values.
func NewChannelGroupUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelGroupUpdateRequestBuilder) {
    m := &ChannelGroupUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.group.update", pathParameters),
    }
    return m
}
// NewChannelGroupUpdateRequestBuilder instantiates a new ChannelGroupUpdateRequestBuilder and sets the default values.
func NewChannelGroupUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelGroupUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelGroupUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates a previously created channel group.
// Deprecated: This method is obsolete. Use PostAsUpdatePostResponse instead.
func (m *ChannelGroupUpdateRequestBuilder) Post(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ChannelGroupUpdateRequestBuilderPostRequestConfiguration)(UpdateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdateResponseable), nil
}
// PostAsUpdatePostResponse updates a previously created channel group.
func (m *ChannelGroupUpdateRequestBuilder) PostAsUpdatePostResponse(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ChannelGroupUpdateRequestBuilderPostRequestConfiguration)(UpdatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdatePostResponseable), nil
}
// ToPostRequestInformation updates a previously created channel group.
func (m *ChannelGroupUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ChannelGroupUpdateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelGroupUpdateRequestBuilder) WithUrl(rawUrl string)(*ChannelGroupUpdateRequestBuilder) {
    return NewChannelGroupUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
