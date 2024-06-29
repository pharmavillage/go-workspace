package wikipreview

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WikiPreviewRequestBuilder builds and executes requests for operations under \wiki.preview
type WikiPreviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByChannelresource gets an item from the kiota_airsend/client.wikiPreview.item collection
func (m *WikiPreviewRequestBuilder) ByChannelresource(channelresource string)(*WithChannelresourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if channelresource != "" {
        urlTplParams["channelresource"] = channelresource
    }
    return NewWithChannelresourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWikiPreviewRequestBuilderInternal instantiates a new WikiPreviewRequestBuilder and sets the default values.
func NewWikiPreviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WikiPreviewRequestBuilder) {
    m := &WikiPreviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/wiki.preview", pathParameters),
    }
    return m
}
// NewWikiPreviewRequestBuilder instantiates a new WikiPreviewRequestBuilder and sets the default values.
func NewWikiPreviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WikiPreviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWikiPreviewRequestBuilderInternal(urlParams, requestAdapter)
}
