package callupdate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallUpdateRequestBuilder builds and executes requests for operations under \call.update
type CallUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallUpdateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallUpdateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallUpdateRequestBuilderInternal instantiates a new CallUpdateRequestBuilder and sets the default values.
func NewCallUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallUpdateRequestBuilder) {
    m := &CallUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/call.update", pathParameters),
    }
    return m
}
// NewCallUpdateRequestBuilder instantiates a new CallUpdateRequestBuilder and sets the default values.
func NewCallUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update a  call.
// Deprecated: This method is obsolete. Use PostAsUpdatePostResponse instead.
func (m *CallUpdateRequestBuilder) Post(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *CallUpdateRequestBuilderPostRequestConfiguration)(UpdateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdateResponseable), nil
}
// PostAsUpdatePostResponse update a  call.
func (m *CallUpdateRequestBuilder) PostAsUpdatePostResponse(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *CallUpdateRequestBuilderPostRequestConfiguration)(UpdatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateUpdate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUpdatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UpdatePostResponseable), nil
}
// ToPostRequestInformation update a  call.
func (m *CallUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *CallUpdateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CallUpdateRequestBuilder) WithUrl(rawUrl string)(*CallUpdateRequestBuilder) {
    return NewCallUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
