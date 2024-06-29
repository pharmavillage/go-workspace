package oauthserverapprove

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthServerApproveRequestBuilder builds and executes requests for operations under \oauth.server.approve
type OauthServerApproveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OauthServerApproveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthServerApproveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOauthServerApproveRequestBuilderInternal instantiates a new OauthServerApproveRequestBuilder and sets the default values.
func NewOauthServerApproveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerApproveRequestBuilder) {
    m := &OauthServerApproveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/oauth.server.approve", pathParameters),
    }
    return m
}
// NewOauthServerApproveRequestBuilder instantiates a new OauthServerApproveRequestBuilder and sets the default values.
func NewOauthServerApproveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerApproveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOauthServerApproveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this endpoint directly redirects to the callback url defined on the request, with the authorization code if the access was authorized, or with the errors if something went wrong.
func (m *OauthServerApproveRequestBuilder) Post(ctx context.Context, requestConfiguration *OauthServerApproveRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateApprove400ErrorFromDiscriminatorValue,
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
// ToPostRequestInformation this endpoint directly redirects to the callback url defined on the request, with the authorization code if the access was authorized, or with the errors if something went wrong.
func (m *OauthServerApproveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *OauthServerApproveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *OauthServerApproveRequestBuilder) WithUrl(rawUrl string)(*OauthServerApproveRequestBuilder) {
    return NewOauthServerApproveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
