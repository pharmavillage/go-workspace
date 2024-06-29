package channelgrouplist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelGroupListRequestBuilder builds and executes requests for operations under \channel.group.list
type ChannelGroupListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelGroupListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelGroupListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChannelGroupListRequestBuilderInternal instantiates a new ChannelGroupListRequestBuilder and sets the default values.
func NewChannelGroupListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelGroupListRequestBuilder) {
    m := &ChannelGroupListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.group.list", pathParameters),
    }
    return m
}
// NewChannelGroupListRequestBuilder instantiates a new ChannelGroupListRequestBuilder and sets the default values.
func NewChannelGroupListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelGroupListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelGroupListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns an ordered list os the channel groups defined for the logged user. Includes the user defined groups, and the virtual groups (like all channels, favorites, direct messages).
// Deprecated: This method is obsolete. Use GetAsListGetResponse instead.
func (m *ChannelGroupListRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelGroupListRequestBuilderGetRequestConfiguration)(ListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListResponseable), nil
}
// GetAsListGetResponse returns an ordered list os the channel groups defined for the logged user. Includes the user defined groups, and the virtual groups (like all channels, favorites, direct messages).
func (m *ChannelGroupListRequestBuilder) GetAsListGetResponse(ctx context.Context, requestConfiguration *ChannelGroupListRequestBuilderGetRequestConfiguration)(ListGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListGetResponseable), nil
}
// ToGetRequestInformation returns an ordered list os the channel groups defined for the logged user. Includes the user defined groups, and the virtual groups (like all channels, favorites, direct messages).
func (m *ChannelGroupListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelGroupListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelGroupListRequestBuilder) WithUrl(rawUrl string)(*ChannelGroupListRequestBuilder) {
    return NewChannelGroupListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
