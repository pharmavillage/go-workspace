package usernotificationsreport

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserNotificationsReportRequestBuilder builds and executes requests for operations under \user.notifications.report
type UserNotificationsReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserNotificationsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserNotificationsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserNotificationsReportRequestBuilderInternal instantiates a new UserNotificationsReportRequestBuilder and sets the default values.
func NewUserNotificationsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserNotificationsReportRequestBuilder) {
    m := &UserNotificationsReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.notifications.report", pathParameters),
    }
    return m
}
// NewUserNotificationsReportRequestBuilder instantiates a new UserNotificationsReportRequestBuilder and sets the default values.
func NewUserNotificationsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserNotificationsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserNotificationsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a new abuse report for notifications.
// Deprecated: This method is obsolete. Use PostAsReportPostResponse instead.
func (m *UserNotificationsReportRequestBuilder) Post(ctx context.Context, body ReportPostRequestBodyable, requestConfiguration *UserNotificationsReportRequestBuilderPostRequestConfiguration)(ReportResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReport400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportResponseable), nil
}
// PostAsReportPostResponse creates a new abuse report for notifications.
func (m *UserNotificationsReportRequestBuilder) PostAsReportPostResponse(ctx context.Context, body ReportPostRequestBodyable, requestConfiguration *UserNotificationsReportRequestBuilderPostRequestConfiguration)(ReportPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReport400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportPostResponseable), nil
}
// ToPostRequestInformation creates a new abuse report for notifications.
func (m *UserNotificationsReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportPostRequestBodyable, requestConfiguration *UserNotificationsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserNotificationsReportRequestBuilder) WithUrl(rawUrl string)(*UserNotificationsReportRequestBuilder) {
    return NewUserNotificationsReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
