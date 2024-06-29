package channelupdate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelUpdateRequestBuilder builds and executes requests for operations under \channel.update
type ChannelUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelUpdateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelUpdateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelUpdateRequestBuilderInternal instantiates a new ChannelUpdateRequestBuilder and sets the default values.
func NewChannelUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelUpdateRequestBuilder) {
    m := &ChannelUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.update", pathParameters),
    }
    return m
}
// NewChannelUpdateRequestBuilder instantiates a new ChannelUpdateRequestBuilder and sets the default values.
func NewChannelUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update a channel. Only owner or admins set role for channel members. The role can be: 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
// Deprecated: This method is obsolete. Use PostAsUpdatePostResponse instead.
func (m *ChannelUpdateRequestBuilder) Post(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ChannelUpdateRequestBuilderPostRequestConfiguration)(UpdateResponseable, error) {
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
// PostAsUpdatePostResponse update a channel. Only owner or admins set role for channel members. The role can be: 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *ChannelUpdateRequestBuilder) PostAsUpdatePostResponse(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ChannelUpdateRequestBuilderPostRequestConfiguration)(UpdatePostResponseable, error) {
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
// ToPostRequestInformation update a channel. Only owner or admins set role for channel members. The role can be: 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *ChannelUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ChannelUpdateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelUpdateRequestBuilder) WithUrl(rawUrl string)(*ChannelUpdateRequestBuilder) {
    return NewChannelUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
