package fileinfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FileInfoRequestBuilder builds and executes requests for operations under \file.info
type FileInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileInfoRequestBuilderGetQueryParameters gets info related to a file path
type FileInfoRequestBuilderGetQueryParameters struct {
    // path to location
    Fspath *string `uriparametername:"fspath"`
}
// FileInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileInfoRequestBuilderGetQueryParameters
}
// NewFileInfoRequestBuilderInternal instantiates a new FileInfoRequestBuilder and sets the default values.
func NewFileInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileInfoRequestBuilder) {
    m := &FileInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/file.info?fspath={fspath}", pathParameters),
    }
    return m
}
// NewFileInfoRequestBuilder instantiates a new FileInfoRequestBuilder and sets the default values.
func NewFileInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets info related to a file path
// Deprecated: This method is obsolete. Use GetAsInfoGetResponse instead.
func (m *FileInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *FileInfoRequestBuilderGetRequestConfiguration)(InfoResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoResponseable), nil
}
// GetAsInfoGetResponse gets info related to a file path
func (m *FileInfoRequestBuilder) GetAsInfoGetResponse(ctx context.Context, requestConfiguration *FileInfoRequestBuilderGetRequestConfiguration)(InfoGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoGetResponseable), nil
}
// ToGetRequestInformation gets info related to a file path
func (m *FileInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FileInfoRequestBuilder) WithUrl(rawUrl string)(*FileInfoRequestBuilder) {
    return NewFileInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
