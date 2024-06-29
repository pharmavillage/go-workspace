package wikiget

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithChannelresourceItemRequestBuilder builds and executes requests for operations under \wiki.get\{channelresource}
type WithChannelresourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithChannelresourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithChannelresourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWithChannelresourceItemRequestBuilderInternal instantiates a new WithChannelresourceItemRequestBuilder and sets the default values.
func NewWithChannelresourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithChannelresourceItemRequestBuilder) {
    m := &WithChannelresourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/wiki.get/{channelresource}", pathParameters),
    }
    return m
}
// NewWithChannelresourceItemRequestBuilder instantiates a new WithChannelresourceItemRequestBuilder and sets the default values.
func NewWithChannelresourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithChannelresourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithChannelresourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the wiki associated with the channel. Typically, the wiki starts with a file named 'index.md'So, start the wiki off by using a path like '/wiki.get/wf/10001/index.md'.You can retrieve the wiki root for the channel from the channel info API call.
func (m *WithChannelresourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithChannelresourceItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateItemWithChannelresource400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets the wiki associated with the channel. Typically, the wiki starts with a file named 'index.md'So, start the wiki off by using a path like '/wiki.get/wf/10001/index.md'.You can retrieve the wiki root for the channel from the channel info API call.
func (m *WithChannelresourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithChannelresourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/html, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WithChannelresourceItemRequestBuilder) WithUrl(rawUrl string)(*WithChannelresourceItemRequestBuilder) {
    return NewWithChannelresourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
