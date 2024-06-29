package actioncreate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActionCreateRequestBuilder builds and executes requests for operations under \action.create
type ActionCreateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionCreateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActionCreateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewActionCreateRequestBuilderInternal instantiates a new ActionCreateRequestBuilder and sets the default values.
func NewActionCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionCreateRequestBuilder) {
    m := &ActionCreateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/action.create", pathParameters),
    }
    return m
}
// NewActionCreateRequestBuilder instantiates a new ActionCreateRequestBuilder and sets the default values.
func NewActionCreateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionCreateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActionCreateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new action
// Deprecated: This method is obsolete. Use PostAsCreatePostResponse instead.
func (m *ActionCreateRequestBuilder) Post(ctx context.Context, body CreatePostRequestBodyable, requestConfiguration *ActionCreateRequestBuilderPostRequestConfiguration)(CreateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCreate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCreateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CreateResponseable), nil
}
// PostAsCreatePostResponse create a new action
func (m *ActionCreateRequestBuilder) PostAsCreatePostResponse(ctx context.Context, body CreatePostRequestBodyable, requestConfiguration *ActionCreateRequestBuilderPostRequestConfiguration)(CreatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCreate400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCreatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CreatePostResponseable), nil
}
// ToPostRequestInformation create a new action
func (m *ActionCreateRequestBuilder) ToPostRequestInformation(ctx context.Context, body CreatePostRequestBodyable, requestConfiguration *ActionCreateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ActionCreateRequestBuilder) WithUrl(rawUrl string)(*ActionCreateRequestBuilder) {
    return NewActionCreateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
