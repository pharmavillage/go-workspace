package actiondelete

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActionDeleteRequestBuilder builds and executes requests for operations under \action.delete
type ActionDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActionDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewActionDeleteRequestBuilderInternal instantiates a new ActionDeleteRequestBuilder and sets the default values.
func NewActionDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionDeleteRequestBuilder) {
    m := &ActionDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/action.delete", pathParameters),
    }
    return m
}
// NewActionDeleteRequestBuilder instantiates a new ActionDeleteRequestBuilder and sets the default values.
func NewActionDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActionDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete an action. Only owner can delete the action. A bot messageand realtime event will be raised for this operation.
// Deprecated: This method is obsolete. Use PostAsDeletePostResponse instead.
func (m *ActionDeleteRequestBuilder) Post(ctx context.Context, body DeletePostRequestBodyable, requestConfiguration *ActionDeleteRequestBuilderPostRequestConfiguration)(DeleteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateDelete400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteResponseable), nil
}
// PostAsDeletePostResponse delete an action. Only owner can delete the action. A bot messageand realtime event will be raised for this operation.
func (m *ActionDeleteRequestBuilder) PostAsDeletePostResponse(ctx context.Context, body DeletePostRequestBodyable, requestConfiguration *ActionDeleteRequestBuilderPostRequestConfiguration)(DeletePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateDelete400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletePostResponseable), nil
}
// ToPostRequestInformation delete an action. Only owner can delete the action. A bot messageand realtime event will be raised for this operation.
func (m *ActionDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeletePostRequestBodyable, requestConfiguration *ActionDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ActionDeleteRequestBuilder) WithUrl(rawUrl string)(*ActionDeleteRequestBuilder) {
    return NewActionDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
