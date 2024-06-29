package userverify

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserVerifyRequestBuilder builds and executes requests for operations under \user.verify
type UserVerifyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserVerifyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserVerifyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserVerifyRequestBuilderInternal instantiates a new UserVerifyRequestBuilder and sets the default values.
func NewUserVerifyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserVerifyRequestBuilder) {
    m := &UserVerifyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.verify", pathParameters),
    }
    return m
}
// NewUserVerifyRequestBuilder instantiates a new UserVerifyRequestBuilder and sets the default values.
func NewUserVerifyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserVerifyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserVerifyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post verify the email of phone for the user account.
// Deprecated: This method is obsolete. Use PostAsVerifyPostResponse instead.
func (m *UserVerifyRequestBuilder) Post(ctx context.Context, body VerifyPostRequestBodyable, requestConfiguration *UserVerifyRequestBuilderPostRequestConfiguration)(VerifyResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateVerify400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVerifyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VerifyResponseable), nil
}
// PostAsVerifyPostResponse verify the email of phone for the user account.
func (m *UserVerifyRequestBuilder) PostAsVerifyPostResponse(ctx context.Context, body VerifyPostRequestBodyable, requestConfiguration *UserVerifyRequestBuilderPostRequestConfiguration)(VerifyPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateVerify400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVerifyPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VerifyPostResponseable), nil
}
// ToPostRequestInformation verify the email of phone for the user account.
func (m *UserVerifyRequestBuilder) ToPostRequestInformation(ctx context.Context, body VerifyPostRequestBodyable, requestConfiguration *UserVerifyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserVerifyRequestBuilder) WithUrl(rawUrl string)(*UserVerifyRequestBuilder) {
    return NewUserVerifyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
