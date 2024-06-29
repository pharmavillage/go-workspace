package oauthserverclientlist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthServerClientListRequestBuilder builds and executes requests for operations under \oauth.server.client.list
type OauthServerClientListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OauthServerClientListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthServerClientListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOauthServerClientListRequestBuilderInternal instantiates a new OauthServerClientListRequestBuilder and sets the default values.
func NewOauthServerClientListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerClientListRequestBuilder) {
    m := &OauthServerClientListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/oauth.server.client.list", pathParameters),
    }
    return m
}
// NewOauthServerClientListRequestBuilder instantiates a new OauthServerClientListRequestBuilder and sets the default values.
func NewOauthServerClientListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerClientListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOauthServerClientListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the OAuth clients owned by the logged user
func (m *OauthServerClientListRequestBuilder) Get(ctx context.Context, requestConfiguration *OauthServerClientListRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
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
// ToGetRequestInformation list the OAuth clients owned by the logged user
func (m *OauthServerClientListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OauthServerClientListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *OauthServerClientListRequestBuilder) WithUrl(rawUrl string)(*OauthServerClientListRequestBuilder) {
    return NewOauthServerClientListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
