package wikiget

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WikiGetRequestBuilder builds and executes requests for operations under \wiki.get
type WikiGetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByChannelresource gets an item from the kiota_airsend/client.wikiGet.item collection
func (m *WikiGetRequestBuilder) ByChannelresource(channelresource string)(*WithChannelresourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if channelresource != "" {
        urlTplParams["channelresource"] = channelresource
    }
    return NewWithChannelresourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWikiGetRequestBuilderInternal instantiates a new WikiGetRequestBuilder and sets the default values.
func NewWikiGetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WikiGetRequestBuilder) {
    m := &WikiGetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/wiki.get", pathParameters),
    }
    return m
}
// NewWikiGetRequestBuilder instantiates a new WikiGetRequestBuilder and sets the default values.
func NewWikiGetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WikiGetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWikiGetRequestBuilderInternal(urlParams, requestAdapter)
}
