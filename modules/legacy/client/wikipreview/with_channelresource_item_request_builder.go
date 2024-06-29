package wikipreview

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithChannelresourceItemRequestBuilder builds and executes requests for operations under \wiki.preview\{channelresource}
type WithChannelresourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithChannelresourceItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithChannelresourceItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWithChannelresourceItemRequestBuilderInternal instantiates a new WithChannelresourceItemRequestBuilder and sets the default values.
func NewWithChannelresourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithChannelresourceItemRequestBuilder) {
    m := &WithChannelresourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/wiki.preview/{channelresource}", pathParameters),
    }
    return m
}
// NewWithChannelresourceItemRequestBuilder instantiates a new WithChannelresourceItemRequestBuilder and sets the default values.
func NewWithChannelresourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithChannelresourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithChannelresourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post previews the wiki markdown associated with the wiki resource path. POST the wiki markdown in the body of the HTTP Request, set Content-Type to text/markdownSo, use a path like '/wiki.get/wf/10001/index.md'.You can retrieve the wiki root for the channel from the channel info API call.
func (m *WithChannelresourceItemRequestBuilder) Post(ctx context.Context, body []byte, requestConfiguration *WithChannelresourceItemRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation previews the wiki markdown associated with the wiki resource path. POST the wiki markdown in the body of the HTTP Request, set Content-Type to text/markdownSo, use a path like '/wiki.get/wf/10001/index.md'.You can retrieve the wiki root for the channel from the channel info API call.
func (m *WithChannelresourceItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body []byte, requestConfiguration *WithChannelresourceItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/html, application/json")
    requestInfo.SetStreamContentAndContentType(body, "text/markdown")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WithChannelresourceItemRequestBuilder) WithUrl(rawUrl string)(*WithChannelresourceItemRequestBuilder) {
    return NewWithChannelresourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
