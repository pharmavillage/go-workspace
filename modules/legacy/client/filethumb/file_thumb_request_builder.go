package filethumb

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FileThumbRequestBuilder builds and executes requests for operations under \file.thumb
type FileThumbRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileThumbRequestBuilderGetQueryParameters get the thumb image for a file
type FileThumbRequestBuilderGetQueryParameters struct {
    // path to location
    Fspath *string `uriparametername:"fspath"`
    // Height of the thumbnail in pixels
    Height *int32 `uriparametername:"height"`
    // Width of the thumbnail in pixels
    Width *int32 `uriparametername:"width"`
}
// FileThumbRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileThumbRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileThumbRequestBuilderGetQueryParameters
}
// NewFileThumbRequestBuilderInternal instantiates a new FileThumbRequestBuilder and sets the default values.
func NewFileThumbRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileThumbRequestBuilder) {
    m := &FileThumbRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/file.thumb?fspath={fspath}&height={height}&width={width}", pathParameters),
    }
    return m
}
// NewFileThumbRequestBuilder instantiates a new FileThumbRequestBuilder and sets the default values.
func NewFileThumbRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileThumbRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileThumbRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the thumb image for a file
func (m *FileThumbRequestBuilder) Get(ctx context.Context, requestConfiguration *FileThumbRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateThumb400ErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get the thumb image for a file
func (m *FileThumbRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileThumbRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FileThumbRequestBuilder) WithUrl(rawUrl string)(*FileThumbRequestBuilder) {
    return NewFileThumbRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
