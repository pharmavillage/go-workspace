package userloginrefresh

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserLoginRefreshRequestBuilder builds and executes requests for operations under \user.login.refresh
type UserLoginRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserLoginRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserLoginRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserLoginRefreshRequestBuilderInternal instantiates a new UserLoginRefreshRequestBuilder and sets the default values.
func NewUserLoginRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserLoginRefreshRequestBuilder) {
    m := &UserLoginRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.login.refresh", pathParameters),
    }
    return m
}
// NewUserLoginRefreshRequestBuilder instantiates a new UserLoginRefreshRequestBuilder and sets the default values.
func NewUserLoginRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserLoginRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserLoginRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post refreshes auth token from a authenticated user.This endpoint don't check for auth expiration, but verifies if a valid refresh token (rtk claim on jwt) is set. So it's possible to use this endpoint to refresh tokens, even when expired. The refresh token will only be available if the user set remember_me on the login process. A refresh token can only be used one time.As a part of this refresh process, these steps are needed if you use RTM service.1. Perform user.login.refresh in HTTP backend before expiration of your current token2. Perform rtm.connect to get get a new RTM token3. send the auth command in websocket to refresh the session with the new tokenType a message
// Deprecated: This method is obsolete. Use PostAsRefreshPostResponse instead.
func (m *UserLoginRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *UserLoginRefreshRequestBuilderPostRequestConfiguration)(RefreshResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateRefresh401ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRefreshResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RefreshResponseable), nil
}
// PostAsRefreshPostResponse refreshes auth token from a authenticated user.This endpoint don't check for auth expiration, but verifies if a valid refresh token (rtk claim on jwt) is set. So it's possible to use this endpoint to refresh tokens, even when expired. The refresh token will only be available if the user set remember_me on the login process. A refresh token can only be used one time.As a part of this refresh process, these steps are needed if you use RTM service.1. Perform user.login.refresh in HTTP backend before expiration of your current token2. Perform rtm.connect to get get a new RTM token3. send the auth command in websocket to refresh the session with the new tokenType a message
func (m *UserLoginRefreshRequestBuilder) PostAsRefreshPostResponse(ctx context.Context, requestConfiguration *UserLoginRefreshRequestBuilderPostRequestConfiguration)(RefreshPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateRefresh401ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRefreshPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RefreshPostResponseable), nil
}
// ToPostRequestInformation refreshes auth token from a authenticated user.This endpoint don't check for auth expiration, but verifies if a valid refresh token (rtk claim on jwt) is set. So it's possible to use this endpoint to refresh tokens, even when expired. The refresh token will only be available if the user set remember_me on the login process. A refresh token can only be used one time.As a part of this refresh process, these steps are needed if you use RTM service.1. Perform user.login.refresh in HTTP backend before expiration of your current token2. Perform rtm.connect to get get a new RTM token3. send the auth command in websocket to refresh the session with the new tokenType a message
func (m *UserLoginRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *UserLoginRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserLoginRefreshRequestBuilder) WithUrl(rawUrl string)(*UserLoginRefreshRequestBuilder) {
    return NewUserLoginRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
