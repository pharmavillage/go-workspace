package passwordreset

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PasswordResetRequestBuilder builds and executes requests for operations under \password.reset
type PasswordResetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PasswordResetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PasswordResetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPasswordResetRequestBuilderInternal instantiates a new PasswordResetRequestBuilder and sets the default values.
func NewPasswordResetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PasswordResetRequestBuilder) {
    m := &PasswordResetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/password.reset", pathParameters),
    }
    return m
}
// NewPasswordResetRequestBuilder instantiates a new PasswordResetRequestBuilder and sets the default values.
func NewPasswordResetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PasswordResetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPasswordResetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset the password for the user
// Deprecated: This method is obsolete. Use PostAsResetPostResponse instead.
func (m *PasswordResetRequestBuilder) Post(ctx context.Context, body ResetPostRequestBodyable, requestConfiguration *PasswordResetRequestBuilderPostRequestConfiguration)(ResetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReset400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResetResponseable), nil
}
// PostAsResetPostResponse reset the password for the user
func (m *PasswordResetRequestBuilder) PostAsResetPostResponse(ctx context.Context, body ResetPostRequestBodyable, requestConfiguration *PasswordResetRequestBuilderPostRequestConfiguration)(ResetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateReset400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResetPostResponseable), nil
}
// ToPostRequestInformation reset the password for the user
func (m *PasswordResetRequestBuilder) ToPostRequestInformation(ctx context.Context, body ResetPostRequestBodyable, requestConfiguration *PasswordResetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PasswordResetRequestBuilder) WithUrl(rawUrl string)(*PasswordResetRequestBuilder) {
    return NewPasswordResetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
