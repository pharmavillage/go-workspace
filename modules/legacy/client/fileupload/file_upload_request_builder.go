package fileupload

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FileUploadRequestBuilder builds and executes requests for operations under \file.upload
type FileUploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileUploadRequestBuilderPostQueryParameters upload a file
type FileUploadRequestBuilderPostQueryParameters struct {
    // (Optional) Indicates whether file is fully uploaded when set to 1. Default is 1 which means the file upload is complete.
    Complete *int32 `uriparametername:"complete"`
    // path to location
    Fspath *string `uriparametername:"fspath"`
    // (Optional) Optional offset from which file is being uploaded. Default is 0 which is the beginning of the file.When uploading very large files, you will want to send complete = 0 and then upload chunks one by one by moving the start value and complete the upload by sending complete=1 with the last chunk.Authentication is required to call this method
    Start *int32 `uriparametername:"start"`
}
// FileUploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileUploadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileUploadRequestBuilderPostQueryParameters
}
// NewFileUploadRequestBuilderInternal instantiates a new FileUploadRequestBuilder and sets the default values.
func NewFileUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileUploadRequestBuilder) {
    m := &FileUploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/file.upload?fspath={fspath}{&complete*,start*}", pathParameters),
    }
    return m
}
// NewFileUploadRequestBuilder instantiates a new FileUploadRequestBuilder and sets the default values.
func NewFileUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileUploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileUploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a file
// Deprecated: This method is obsolete. Use PostAsUploadPostResponse instead.
func (m *FileUploadRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *FileUploadRequestBuilderPostRequestConfiguration)(UploadResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpload400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUploadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UploadResponseable), nil
}
// PostAsUploadPostResponse upload a file
func (m *FileUploadRequestBuilder) PostAsUploadPostResponse(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *FileUploadRequestBuilderPostRequestConfiguration)(UploadPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpload400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUploadPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UploadPostResponseable), nil
}
// ToPostRequestInformation upload a file
func (m *FileUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *FileUploadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *FileUploadRequestBuilder) WithUrl(rawUrl string)(*FileUploadRequestBuilder) {
    return NewFileUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
