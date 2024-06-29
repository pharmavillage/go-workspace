package channelexport

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelExportRequestBuilder builds and executes requests for operations under \channel.export
type ChannelExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelExportRequestBuilderGetQueryParameters this API will generate a zip file with all files, CSV file of message and actions. The user must be the owner of the channel to perform this operation
type ChannelExportRequestBuilderGetQueryParameters struct {
    // ID of the channel to export
    Channel_id *int32 `uriparametername:"channel_id"`
}
// ChannelExportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelExportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelExportRequestBuilderGetQueryParameters
}
// NewChannelExportRequestBuilderInternal instantiates a new ChannelExportRequestBuilder and sets the default values.
func NewChannelExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelExportRequestBuilder) {
    m := &ChannelExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.export{?channel_id*}", pathParameters),
    }
    return m
}
// NewChannelExportRequestBuilder instantiates a new ChannelExportRequestBuilder and sets the default values.
func NewChannelExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this API will generate a zip file with all files, CSV file of message and actions. The user must be the owner of the channel to perform this operation
func (m *ChannelExportRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelExportRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateExport400ErrorFromDiscriminatorValue,
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
// ToGetRequestInformation this API will generate a zip file with all files, CSV file of message and actions. The user must be the owner of the channel to perform this operation
func (m *ChannelExportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelExportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ChannelExportRequestBuilder) WithUrl(rawUrl string)(*ChannelExportRequestBuilder) {
    return NewChannelExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
