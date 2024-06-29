package channelmembers

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelMembersRequestBuilder builds and executes requests for operations under \channel.members
type ChannelMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelMembersRequestBuilderGetQueryParameters list users in a channel
type ChannelMembersRequestBuilderGetQueryParameters struct {
    // Id of channel to get information
    Channel_id *int32 `uriparametername:"channel_id"`
}
// ChannelMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelMembersRequestBuilderGetQueryParameters
}
// NewChannelMembersRequestBuilderInternal instantiates a new ChannelMembersRequestBuilder and sets the default values.
func NewChannelMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelMembersRequestBuilder) {
    m := &ChannelMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.members?channel_id={channel_id}", pathParameters),
    }
    return m
}
// NewChannelMembersRequestBuilder instantiates a new ChannelMembersRequestBuilder and sets the default values.
func NewChannelMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list users in a channel
// Deprecated: This method is obsolete. Use GetAsMembersGetResponse instead.
func (m *ChannelMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelMembersRequestBuilderGetRequestConfiguration)(MembersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateMembers400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMembersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MembersResponseable), nil
}
// GetAsMembersGetResponse list users in a channel
func (m *ChannelMembersRequestBuilder) GetAsMembersGetResponse(ctx context.Context, requestConfiguration *ChannelMembersRequestBuilderGetRequestConfiguration)(MembersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateMembers400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMembersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MembersGetResponseable), nil
}
// ToGetRequestInformation list users in a channel
func (m *ChannelMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelMembersRequestBuilder) WithUrl(rawUrl string)(*ChannelMembersRequestBuilder) {
    return NewChannelMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
