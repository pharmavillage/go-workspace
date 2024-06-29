package passwordrecover

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PasswordRecoverRequestBuilder builds and executes requests for operations under \password.recover
type PasswordRecoverRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PasswordRecoverRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PasswordRecoverRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPasswordRecoverRequestBuilderInternal instantiates a new PasswordRecoverRequestBuilder and sets the default values.
func NewPasswordRecoverRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PasswordRecoverRequestBuilder) {
    m := &PasswordRecoverRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/password.recover", pathParameters),
    }
    return m
}
// NewPasswordRecoverRequestBuilder instantiates a new PasswordRecoverRequestBuilder and sets the default values.
func NewPasswordRecoverRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PasswordRecoverRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPasswordRecoverRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set password recovery code. Email or Phone is required
// Deprecated: This method is obsolete. Use PostAsRecoverPostResponse instead.
func (m *PasswordRecoverRequestBuilder) Post(ctx context.Context, body RecoverPostRequestBodyable, requestConfiguration *PasswordRecoverRequestBuilderPostRequestConfiguration)(RecoverResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRecover400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRecoverResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RecoverResponseable), nil
}
// PostAsRecoverPostResponse set password recovery code. Email or Phone is required
func (m *PasswordRecoverRequestBuilder) PostAsRecoverPostResponse(ctx context.Context, body RecoverPostRequestBodyable, requestConfiguration *PasswordRecoverRequestBuilderPostRequestConfiguration)(RecoverPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateRecover400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRecoverPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RecoverPostResponseable), nil
}
// ToPostRequestInformation set password recovery code. Email or Phone is required
func (m *PasswordRecoverRequestBuilder) ToPostRequestInformation(ctx context.Context, body RecoverPostRequestBodyable, requestConfiguration *PasswordRecoverRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PasswordRecoverRequestBuilder) WithUrl(rawUrl string)(*PasswordRecoverRequestBuilder) {
    return NewPasswordRecoverRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
