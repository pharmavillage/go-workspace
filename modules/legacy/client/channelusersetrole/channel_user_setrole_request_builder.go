package channelusersetrole

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelUserSetroleRequestBuilder builds and executes requests for operations under \channel.user.setrole
type ChannelUserSetroleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelUserSetroleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelUserSetroleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelUserSetroleRequestBuilderInternal instantiates a new ChannelUserSetroleRequestBuilder and sets the default values.
func NewChannelUserSetroleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelUserSetroleRequestBuilder) {
    m := &ChannelUserSetroleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.user.setrole", pathParameters),
    }
    return m
}
// NewChannelUserSetroleRequestBuilder instantiates a new ChannelUserSetroleRequestBuilder and sets the default values.
func NewChannelUserSetroleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelUserSetroleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelUserSetroleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set role for a member. Only owner or admins set role for channel members. The role can be: 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
// Deprecated: This method is obsolete. Use PostAsSetrolePostResponse instead.
func (m *ChannelUserSetroleRequestBuilder) Post(ctx context.Context, body SetrolePostRequestBodyable, requestConfiguration *ChannelUserSetroleRequestBuilderPostRequestConfiguration)(SetroleResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSetrole400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetroleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetroleResponseable), nil
}
// PostAsSetrolePostResponse set role for a member. Only owner or admins set role for channel members. The role can be: 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *ChannelUserSetroleRequestBuilder) PostAsSetrolePostResponse(ctx context.Context, body SetrolePostRequestBodyable, requestConfiguration *ChannelUserSetroleRequestBuilderPostRequestConfiguration)(SetrolePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSetrole400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetrolePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetrolePostResponseable), nil
}
// ToPostRequestInformation set role for a member. Only owner or admins set role for channel members. The role can be: 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *ChannelUserSetroleRequestBuilder) ToPostRequestInformation(ctx context.Context, body SetrolePostRequestBodyable, requestConfiguration *ChannelUserSetroleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelUserSetroleRequestBuilder) WithUrl(rawUrl string)(*ChannelUserSetroleRequestBuilder) {
    return NewChannelUserSetroleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
