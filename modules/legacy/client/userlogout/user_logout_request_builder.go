package userlogout

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserLogoutRequestBuilder builds and executes requests for operations under \user.logout
type UserLogoutRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserLogoutRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserLogoutRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserLogoutRequestBuilderInternal instantiates a new UserLogoutRequestBuilder and sets the default values.
func NewUserLogoutRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserLogoutRequestBuilder) {
    m := &UserLogoutRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.logout", pathParameters),
    }
    return m
}
// NewUserLogoutRequestBuilder instantiates a new UserLogoutRequestBuilder and sets the default values.
func NewUserLogoutRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserLogoutRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserLogoutRequestBuilderInternal(urlParams, requestAdapter)
}
// Post revokes the token used as authentication for the request. Works for internal JWT token, and for Oauth tokens.
// Deprecated: This method is obsolete. Use PostAsLogoutPostResponse instead.
func (m *UserLogoutRequestBuilder) Post(ctx context.Context, requestConfiguration *UserLogoutRequestBuilderPostRequestConfiguration)(LogoutResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateLogout401ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLogoutResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LogoutResponseable), nil
}
// PostAsLogoutPostResponse revokes the token used as authentication for the request. Works for internal JWT token, and for Oauth tokens.
func (m *UserLogoutRequestBuilder) PostAsLogoutPostResponse(ctx context.Context, requestConfiguration *UserLogoutRequestBuilderPostRequestConfiguration)(LogoutPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateLogout401ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLogoutPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LogoutPostResponseable), nil
}
// ToPostRequestInformation revokes the token used as authentication for the request. Works for internal JWT token, and for Oauth tokens.
func (m *UserLogoutRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *UserLogoutRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserLogoutRequestBuilder) WithUrl(rawUrl string)(*UserLogoutRequestBuilder) {
    return NewUserLogoutRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
