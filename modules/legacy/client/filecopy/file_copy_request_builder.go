package filecopy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FileCopyRequestBuilder builds and executes requests for operations under \file.copy
type FileCopyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileCopyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileCopyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileCopyRequestBuilderInternal instantiates a new FileCopyRequestBuilder and sets the default values.
func NewFileCopyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileCopyRequestBuilder) {
    m := &FileCopyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/file.copy", pathParameters),
    }
    return m
}
// NewFileCopyRequestBuilder instantiates a new FileCopyRequestBuilder and sets the default values.
func NewFileCopyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileCopyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileCopyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post copy file or folder from one path to another
// Deprecated: This method is obsolete. Use PostAsCopyPostResponse instead.
func (m *FileCopyRequestBuilder) Post(ctx context.Context, body CopyPostRequestBodyable, requestConfiguration *FileCopyRequestBuilderPostRequestConfiguration)(CopyResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCopy400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCopyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CopyResponseable), nil
}
// PostAsCopyPostResponse copy file or folder from one path to another
func (m *FileCopyRequestBuilder) PostAsCopyPostResponse(ctx context.Context, body CopyPostRequestBodyable, requestConfiguration *FileCopyRequestBuilderPostRequestConfiguration)(CopyPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCopy400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCopyPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CopyPostResponseable), nil
}
// ToPostRequestInformation copy file or folder from one path to another
func (m *FileCopyRequestBuilder) ToPostRequestInformation(ctx context.Context, body CopyPostRequestBodyable, requestConfiguration *FileCopyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *FileCopyRequestBuilder) WithUrl(rawUrl string)(*FileCopyRequestBuilder) {
    return NewFileCopyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
