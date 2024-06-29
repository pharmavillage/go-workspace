package oauthserveraccess_token

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthServerAccess_tokenRequestBuilder builds and executes requests for operations under \oauth.server.access_token
type OauthServerAccess_tokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OauthServerAccess_tokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthServerAccess_tokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOauthServerAccess_tokenRequestBuilderInternal instantiates a new OauthServerAccess_tokenRequestBuilder and sets the default values.
func NewOauthServerAccess_tokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerAccess_tokenRequestBuilder) {
    m := &OauthServerAccess_tokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/oauth.server.access_token", pathParameters),
    }
    return m
}
// NewOauthServerAccess_tokenRequestBuilder instantiates a new OauthServerAccess_tokenRequestBuilder and sets the default values.
func NewOauthServerAccess_tokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerAccess_tokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOauthServerAccess_tokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this endpoint can be called using one of the supported grant types: client_credentials, authorization_code or refresh token to retrieve an access token.
func (m *OauthServerAccess_tokenRequestBuilder) Post(ctx context.Context, requestConfiguration *OauthServerAccess_tokenRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAccess_token400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation this endpoint can be called using one of the supported grant types: client_credentials, authorization_code or refresh token to retrieve an access token.
func (m *OauthServerAccess_tokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *OauthServerAccess_tokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *OauthServerAccess_tokenRequestBuilder) WithUrl(rawUrl string)(*OauthServerAccess_tokenRequestBuilder) {
    return NewOauthServerAccess_tokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
