package oauthserverclientcreate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthServerClientCreateRequestBuilder builds and executes requests for operations under \oauth.server.client.create
type OauthServerClientCreateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OauthServerClientCreateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthServerClientCreateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOauthServerClientCreateRequestBuilderInternal instantiates a new OauthServerClientCreateRequestBuilder and sets the default values.
func NewOauthServerClientCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerClientCreateRequestBuilder) {
    m := &OauthServerClientCreateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/oauth.server.client.create", pathParameters),
    }
    return m
}
// NewOauthServerClientCreateRequestBuilder instantiates a new OauthServerClientCreateRequestBuilder and sets the default values.
func NewOauthServerClientCreateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerClientCreateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOauthServerClientCreateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a new OAuth Client
func (m *OauthServerClientCreateRequestBuilder) Post(ctx context.Context, requestConfiguration *OauthServerClientCreateRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateCreate400ErrorFromDiscriminatorValue,
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
// ToPostRequestInformation creates a new OAuth Client
func (m *OauthServerClientCreateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *OauthServerClientCreateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *OauthServerClientCreateRequestBuilder) WithUrl(rawUrl string)(*OauthServerClientCreateRequestBuilder) {
    return NewOauthServerClientCreateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
