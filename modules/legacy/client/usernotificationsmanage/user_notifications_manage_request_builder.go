package usernotificationsmanage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserNotificationsManageRequestBuilder builds and executes requests for operations under \user.notifications.manage
type UserNotificationsManageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserNotificationsManageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserNotificationsManageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserNotificationsManageRequestBuilderInternal instantiates a new UserNotificationsManageRequestBuilder and sets the default values.
func NewUserNotificationsManageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserNotificationsManageRequestBuilder) {
    m := &UserNotificationsManageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.notifications.manage", pathParameters),
    }
    return m
}
// NewUserNotificationsManageRequestBuilder instantiates a new UserNotificationsManageRequestBuilder and sets the default values.
func NewUserNotificationsManageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserNotificationsManageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserNotificationsManageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the subscribing level for notifications on the user.
// Deprecated: This method is obsolete. Use PostAsManagePostResponse instead.
func (m *UserNotificationsManageRequestBuilder) Post(ctx context.Context, body ManagePostRequestBodyable, requestConfiguration *UserNotificationsManageRequestBuilderPostRequestConfiguration)(ManageResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateManage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageResponseable), nil
}
// PostAsManagePostResponse set the subscribing level for notifications on the user.
func (m *UserNotificationsManageRequestBuilder) PostAsManagePostResponse(ctx context.Context, body ManagePostRequestBodyable, requestConfiguration *UserNotificationsManageRequestBuilderPostRequestConfiguration)(ManagePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateManage400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagePostResponseable), nil
}
// ToPostRequestInformation set the subscribing level for notifications on the user.
func (m *UserNotificationsManageRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagePostRequestBodyable, requestConfiguration *UserNotificationsManageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserNotificationsManageRequestBuilder) WithUrl(rawUrl string)(*UserNotificationsManageRequestBuilder) {
    return NewUserNotificationsManageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
