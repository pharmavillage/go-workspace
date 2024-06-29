package systemping

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SystemPingRequestBuilder builds and executes requests for operations under \system.ping
type SystemPingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SystemPingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SystemPingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSystemPingRequestBuilderInternal instantiates a new SystemPingRequestBuilder and sets the default values.
func NewSystemPingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemPingRequestBuilder) {
    m := &SystemPingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system.ping", pathParameters),
    }
    return m
}
// NewSystemPingRequestBuilder instantiates a new SystemPingRequestBuilder and sets the default values.
func NewSystemPingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemPingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSystemPingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send a ping to HTTP end point and receive it in WS to validate backend routes
// Deprecated: This method is obsolete. Use PostAsPingPostResponse instead.
func (m *SystemPingRequestBuilder) Post(ctx context.Context, body PingPostRequestBodyable, requestConfiguration *SystemPingRequestBuilderPostRequestConfiguration)(PingResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreatePing400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePingResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PingResponseable), nil
}
// PostAsPingPostResponse send a ping to HTTP end point and receive it in WS to validate backend routes
func (m *SystemPingRequestBuilder) PostAsPingPostResponse(ctx context.Context, body PingPostRequestBodyable, requestConfiguration *SystemPingRequestBuilderPostRequestConfiguration)(PingPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreatePing400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePingPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PingPostResponseable), nil
}
// ToPostRequestInformation send a ping to HTTP end point and receive it in WS to validate backend routes
func (m *SystemPingRequestBuilder) ToPostRequestInformation(ctx context.Context, body PingPostRequestBodyable, requestConfiguration *SystemPingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SystemPingRequestBuilder) WithUrl(rawUrl string)(*SystemPingRequestBuilder) {
    return NewSystemPingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
