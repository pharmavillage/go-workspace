package filelist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FileListRequestBuilder builds and executes requests for operations under \file.list
type FileListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileListRequestBuilderGetQueryParameters get list of files for a given path
type FileListRequestBuilderGetQueryParameters struct {
    // path to location. If an empty path is provided, the user level top folders are provided.
    Fspath *string `uriparametername:"fspath"`
    // (Optional) The max number of items to retrieve. Default value is 100. To get all items, specify -1
    Limit *int32 `uriparametername:"limit"`
    // (Optional) Offset from which the listing items should be retrieved. Default value is 0.
    Start *int32 `uriparametername:"start"`
}
// FileListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileListRequestBuilderGetQueryParameters
}
// NewFileListRequestBuilderInternal instantiates a new FileListRequestBuilder and sets the default values.
func NewFileListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileListRequestBuilder) {
    m := &FileListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/file.list?fspath={fspath}{&limit*,start*}", pathParameters),
    }
    return m
}
// NewFileListRequestBuilder instantiates a new FileListRequestBuilder and sets the default values.
func NewFileListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get list of files for a given path
// Deprecated: This method is obsolete. Use GetAsListGetResponse instead.
func (m *FileListRequestBuilder) Get(ctx context.Context, requestConfiguration *FileListRequestBuilderGetRequestConfiguration)(ListResponseable, error) {
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
// GetAsListGetResponse get list of files for a given path
func (m *FileListRequestBuilder) GetAsListGetResponse(ctx context.Context, requestConfiguration *FileListRequestBuilderGetRequestConfiguration)(ListGetResponseable, error) {
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
// ToGetRequestInformation get list of files for a given path
func (m *FileListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FileListRequestBuilder) WithUrl(rawUrl string)(*FileListRequestBuilder) {
    return NewFileListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
