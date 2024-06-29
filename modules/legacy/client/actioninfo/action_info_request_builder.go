package actioninfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActionInfoRequestBuilder builds and executes requests for operations under \action.info
type ActionInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionInfoRequestBuilderGetQueryParameters get information about an action with assignees
type ActionInfoRequestBuilderGetQueryParameters struct {
    // Id of action to get information
    Action_id *int32 `uriparametername:"action_id"`
}
// ActionInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActionInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ActionInfoRequestBuilderGetQueryParameters
}
// NewActionInfoRequestBuilderInternal instantiates a new ActionInfoRequestBuilder and sets the default values.
func NewActionInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionInfoRequestBuilder) {
    m := &ActionInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/action.info?action_id={action_id}", pathParameters),
    }
    return m
}
// NewActionInfoRequestBuilder instantiates a new ActionInfoRequestBuilder and sets the default values.
func NewActionInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActionInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get information about an action with assignees
// Deprecated: This method is obsolete. Use GetAsInfoGetResponse instead.
func (m *ActionInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *ActionInfoRequestBuilderGetRequestConfiguration)(InfoResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoResponseable), nil
}
// GetAsInfoGetResponse get information about an action with assignees
func (m *ActionInfoRequestBuilder) GetAsInfoGetResponse(ctx context.Context, requestConfiguration *ActionInfoRequestBuilderGetRequestConfiguration)(InfoGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoGetResponseable), nil
}
// ToGetRequestInformation get information about an action with assignees
func (m *ActionInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ActionInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ActionInfoRequestBuilder) WithUrl(rawUrl string)(*ActionInfoRequestBuilder) {
    return NewActionInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
