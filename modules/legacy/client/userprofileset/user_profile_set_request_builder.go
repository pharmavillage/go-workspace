package userprofileset

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserProfileSetRequestBuilder builds and executes requests for operations under \user.profile.set
type UserProfileSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserProfileSetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserProfileSetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserProfileSetRequestBuilderInternal instantiates a new UserProfileSetRequestBuilder and sets the default values.
func NewUserProfileSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserProfileSetRequestBuilder) {
    m := &UserProfileSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.profile.set", pathParameters),
    }
    return m
}
// NewUserProfileSetRequestBuilder instantiates a new UserProfileSetRequestBuilder and sets the default values.
func NewUserProfileSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserProfileSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserProfileSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set a profile information of currently logged in user
// Deprecated: This method is obsolete. Use PostAsSetPostResponse instead.
func (m *UserProfileSetRequestBuilder) Post(ctx context.Context, body SetPostRequestBodyable, requestConfiguration *UserProfileSetRequestBuilderPostRequestConfiguration)(SetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetResponseable), nil
}
// PostAsSetPostResponse set a profile information of currently logged in user
func (m *UserProfileSetRequestBuilder) PostAsSetPostResponse(ctx context.Context, body SetPostRequestBodyable, requestConfiguration *UserProfileSetRequestBuilderPostRequestConfiguration)(SetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetPostResponseable), nil
}
// ToPostRequestInformation set a profile information of currently logged in user
func (m *UserProfileSetRequestBuilder) ToPostRequestInformation(ctx context.Context, body SetPostRequestBodyable, requestConfiguration *UserProfileSetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserProfileSetRequestBuilder) WithUrl(rawUrl string)(*UserProfileSetRequestBuilder) {
    return NewUserProfileSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
