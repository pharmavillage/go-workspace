package lockacquire

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LockAcquireRequestBuilder builds and executes requests for operations under \lock.acquire
type LockAcquireRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LockAcquireRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LockAcquireRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLockAcquireRequestBuilderInternal instantiates a new LockAcquireRequestBuilder and sets the default values.
func NewLockAcquireRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LockAcquireRequestBuilder) {
    m := &LockAcquireRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lock.acquire", pathParameters),
    }
    return m
}
// NewLockAcquireRequestBuilder instantiates a new LockAcquireRequestBuilder and sets the default values.
func NewLockAcquireRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LockAcquireRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLockAcquireRequestBuilderInternal(urlParams, requestAdapter)
}
// Post takes lock of the supplied file. The user must have access to the file tolock the file. If the lock is granted a push event via websocket will be sent toall users of the channel if the file is associated with a channel.
// Deprecated: This method is obsolete. Use PostAsAcquirePostResponse instead.
func (m *LockAcquireRequestBuilder) Post(ctx context.Context, body AcquirePostRequestBodyable, requestConfiguration *LockAcquireRequestBuilderPostRequestConfiguration)(AcquireResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAcquire400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAcquireResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AcquireResponseable), nil
}
// PostAsAcquirePostResponse takes lock of the supplied file. The user must have access to the file tolock the file. If the lock is granted a push event via websocket will be sent toall users of the channel if the file is associated with a channel.
func (m *LockAcquireRequestBuilder) PostAsAcquirePostResponse(ctx context.Context, body AcquirePostRequestBodyable, requestConfiguration *LockAcquireRequestBuilderPostRequestConfiguration)(AcquirePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAcquire400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAcquirePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AcquirePostResponseable), nil
}
// ToPostRequestInformation takes lock of the supplied file. The user must have access to the file tolock the file. If the lock is granted a push event via websocket will be sent toall users of the channel if the file is associated with a channel.
func (m *LockAcquireRequestBuilder) ToPostRequestInformation(ctx context.Context, body AcquirePostRequestBodyable, requestConfiguration *LockAcquireRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LockAcquireRequestBuilder) WithUrl(rawUrl string)(*LockAcquireRequestBuilder) {
    return NewLockAcquireRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
