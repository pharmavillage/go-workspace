package actionupdate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActionUpdateRequestBuilder builds and executes requests for operations under \action.update
type ActionUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionUpdateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActionUpdateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewActionUpdateRequestBuilderInternal instantiates a new ActionUpdateRequestBuilder and sets the default values.
func NewActionUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionUpdateRequestBuilder) {
    m := &ActionUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/action.update", pathParameters),
    }
    return m
}
// NewActionUpdateRequestBuilder instantiates a new ActionUpdateRequestBuilder and sets the default values.
func NewActionUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActionUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an action
// Deprecated: This method is obsolete. Use PostAsUpdatePostResponse instead.
func (m *ActionUpdateRequestBuilder) Post(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ActionUpdateRequestBuilderPostRequestConfiguration)(UpdateResponseable, error) {
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
// PostAsUpdatePostResponse create an action
func (m *ActionUpdateRequestBuilder) PostAsUpdatePostResponse(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ActionUpdateRequestBuilderPostRequestConfiguration)(UpdatePostResponseable, error) {
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
// ToPostRequestInformation create an action
func (m *ActionUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body UpdatePostRequestBodyable, requestConfiguration *ActionUpdateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ActionUpdateRequestBuilder) WithUrl(rawUrl string)(*ActionUpdateRequestBuilder) {
    return NewActionUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
