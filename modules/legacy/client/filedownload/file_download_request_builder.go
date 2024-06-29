package filedownload

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FileDownloadRequestBuilder builds and executes requests for operations under \file.download
type FileDownloadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileDownloadRequestBuilderGetQueryParameters download a file
type FileDownloadRequestBuilderGetQueryParameters struct {
    // path to location
    Fspath *string `uriparametername:"fspath"`
    // (Optional)Version identifier as returned by the file.list call. Can only be retrieved if you have write permissions to the resource.
    Versionid *string `uriparametername:"versionid"`
}
// FileDownloadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileDownloadRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileDownloadRequestBuilderGetQueryParameters
}
// NewFileDownloadRequestBuilderInternal instantiates a new FileDownloadRequestBuilder and sets the default values.
func NewFileDownloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileDownloadRequestBuilder) {
    m := &FileDownloadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/file.download?fspath={fspath}{&versionid*}", pathParameters),
    }
    return m
}
// NewFileDownloadRequestBuilder instantiates a new FileDownloadRequestBuilder and sets the default values.
func NewFileDownloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileDownloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileDownloadRequestBuilderInternal(urlParams, requestAdapter)
}
// Get download a file
func (m *FileDownloadRequestBuilder) Get(ctx context.Context, requestConfiguration *FileDownloadRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateDownload400ErrorFromDiscriminatorValue,
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
// ToGetRequestInformation download a file
func (m *FileDownloadRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileDownloadRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FileDownloadRequestBuilder) WithUrl(rawUrl string)(*FileDownloadRequestBuilder) {
    return NewFileDownloadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
