package useralerts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserAlertsRequestBuilder builds and executes requests for operations under \user.alerts
type UserAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserAlertsRequestBuilderInternal instantiates a new UserAlertsRequestBuilder and sets the default values.
func NewUserAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserAlertsRequestBuilder) {
    m := &UserAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.alerts", pathParameters),
    }
    return m
}
// NewUserAlertsRequestBuilder instantiates a new UserAlertsRequestBuilder and sets the default values.
func NewUserAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get return all alert record for a user
// Deprecated: This method is obsolete. Use GetAsAlertsGetResponse instead.
func (m *UserAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserAlertsRequestBuilderGetRequestConfiguration)(AlertsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAlerts400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsResponseable), nil
}
// GetAsAlertsGetResponse return all alert record for a user
func (m *UserAlertsRequestBuilder) GetAsAlertsGetResponse(ctx context.Context, requestConfiguration *UserAlertsRequestBuilderGetRequestConfiguration)(AlertsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAlerts400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsGetResponseable), nil
}
// ToGetRequestInformation return all alert record for a user
func (m *UserAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserAlertsRequestBuilder) WithUrl(rawUrl string)(*UserAlertsRequestBuilder) {
    return NewUserAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
