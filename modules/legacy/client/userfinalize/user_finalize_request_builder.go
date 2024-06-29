package userfinalize

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserFinalizeRequestBuilder builds and executes requests for operations under \user.finalize
type UserFinalizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserFinalizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserFinalizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserFinalizeRequestBuilderInternal instantiates a new UserFinalizeRequestBuilder and sets the default values.
func NewUserFinalizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserFinalizeRequestBuilder) {
    m := &UserFinalizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.finalize", pathParameters),
    }
    return m
}
// NewUserFinalizeRequestBuilder instantiates a new UserFinalizeRequestBuilder and sets the default values.
func NewUserFinalizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserFinalizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFinalizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this endpoint can be accessed using a notification token (without been authenticated). The auth/token expirationis not verified for this endpoint.
// Deprecated: This method is obsolete. Use PostAsFinalizePostResponse instead.
func (m *UserFinalizeRequestBuilder) Post(ctx context.Context, body FinalizePostRequestBodyable, requestConfiguration *UserFinalizeRequestBuilderPostRequestConfiguration)(FinalizeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateFinalize400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFinalizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FinalizeResponseable), nil
}
// PostAsFinalizePostResponse this endpoint can be accessed using a notification token (without been authenticated). The auth/token expirationis not verified for this endpoint.
func (m *UserFinalizeRequestBuilder) PostAsFinalizePostResponse(ctx context.Context, body FinalizePostRequestBodyable, requestConfiguration *UserFinalizeRequestBuilderPostRequestConfiguration)(FinalizePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateFinalize400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFinalizePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FinalizePostResponseable), nil
}
// ToPostRequestInformation this endpoint can be accessed using a notification token (without been authenticated). The auth/token expirationis not verified for this endpoint.
func (m *UserFinalizeRequestBuilder) ToPostRequestInformation(ctx context.Context, body FinalizePostRequestBodyable, requestConfiguration *UserFinalizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserFinalizeRequestBuilder) WithUrl(rawUrl string)(*UserFinalizeRequestBuilder) {
    return NewUserFinalizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
