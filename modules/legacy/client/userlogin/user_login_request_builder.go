package userlogin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserLoginRequestBuilder builds and executes requests for operations under \user.login
type UserLoginRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserLoginRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserLoginRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserLoginRequestBuilderInternal instantiates a new UserLoginRequestBuilder and sets the default values.
func NewUserLoginRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserLoginRequestBuilder) {
    m := &UserLoginRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.login", pathParameters),
    }
    return m
}
// NewUserLoginRequestBuilder instantiates a new UserLoginRequestBuilder and sets the default values.
func NewUserLoginRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserLoginRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserLoginRequestBuilderInternal(urlParams, requestAdapter)
}
// Post login user
// Deprecated: This method is obsolete. Use PostAsLoginPostResponse instead.
func (m *UserLoginRequestBuilder) Post(ctx context.Context, body LoginPostRequestBodyable, requestConfiguration *UserLoginRequestBuilderPostRequestConfiguration)(LoginResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateLogin400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLoginResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LoginResponseable), nil
}
// PostAsLoginPostResponse login user
func (m *UserLoginRequestBuilder) PostAsLoginPostResponse(ctx context.Context, body LoginPostRequestBodyable, requestConfiguration *UserLoginRequestBuilderPostRequestConfiguration)(LoginPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateLogin400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLoginPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LoginPostResponseable), nil
}
// ToPostRequestInformation login user
func (m *UserLoginRequestBuilder) ToPostRequestInformation(ctx context.Context, body LoginPostRequestBodyable, requestConfiguration *UserLoginRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserLoginRequestBuilder) WithUrl(rawUrl string)(*UserLoginRequestBuilder) {
    return NewUserLoginRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
