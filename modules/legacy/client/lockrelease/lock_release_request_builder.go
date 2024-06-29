package lockrelease

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LockReleaseRequestBuilder builds and executes requests for operations under \lock.release
type LockReleaseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LockReleaseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LockReleaseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLockReleaseRequestBuilderInternal instantiates a new LockReleaseRequestBuilder and sets the default values.
func NewLockReleaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LockReleaseRequestBuilder) {
    m := &LockReleaseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lock.release", pathParameters),
    }
    return m
}
// NewLockReleaseRequestBuilder instantiates a new LockReleaseRequestBuilder and sets the default values.
func NewLockReleaseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LockReleaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLockReleaseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post release lock of the file locked earlier.
// Deprecated: This method is obsolete. Use PostAsReleasePostResponse instead.
func (m *LockReleaseRequestBuilder) Post(ctx context.Context, body ReleasePostRequestBodyable, requestConfiguration *LockReleaseRequestBuilderPostRequestConfiguration)(ReleaseResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRelease400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReleaseResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReleaseResponseable), nil
}
// PostAsReleasePostResponse release lock of the file locked earlier.
func (m *LockReleaseRequestBuilder) PostAsReleasePostResponse(ctx context.Context, body ReleasePostRequestBodyable, requestConfiguration *LockReleaseRequestBuilderPostRequestConfiguration)(ReleasePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRelease400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReleasePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReleasePostResponseable), nil
}
// ToPostRequestInformation release lock of the file locked earlier.
func (m *LockReleaseRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReleasePostRequestBodyable, requestConfiguration *LockReleaseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LockReleaseRequestBuilder) WithUrl(rawUrl string)(*LockReleaseRequestBuilder) {
    return NewLockReleaseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
