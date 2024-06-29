package calljoin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallJoinRequestBuilder builds and executes requests for operations under \call.join
type CallJoinRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallJoinRequestBuilderGetQueryParameters call to allow if a client is allowed to join a call or not
type CallJoinRequestBuilderGetQueryParameters struct {
    // Hash of the call the user is attempting to join
    Call_hash *string `uriparametername:"call_hash"`
    // RTM token of the user attempting to join a private call
    Rtm_token *string `uriparametername:"rtm_token"`
}
// CallJoinRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallJoinRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallJoinRequestBuilderGetQueryParameters
}
// NewCallJoinRequestBuilderInternal instantiates a new CallJoinRequestBuilder and sets the default values.
func NewCallJoinRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallJoinRequestBuilder) {
    m := &CallJoinRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/call.join?call_hash={call_hash}{&rtm_token*}", pathParameters),
    }
    return m
}
// NewCallJoinRequestBuilder instantiates a new CallJoinRequestBuilder and sets the default values.
func NewCallJoinRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallJoinRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallJoinRequestBuilderInternal(urlParams, requestAdapter)
}
// Get call to allow if a client is allowed to join a call or not
// Deprecated: This method is obsolete. Use GetAsJoinGetResponse instead.
func (m *CallJoinRequestBuilder) Get(ctx context.Context, requestConfiguration *CallJoinRequestBuilderGetRequestConfiguration)(JoinResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateJoin400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateJoinResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(JoinResponseable), nil
}
// GetAsJoinGetResponse call to allow if a client is allowed to join a call or not
func (m *CallJoinRequestBuilder) GetAsJoinGetResponse(ctx context.Context, requestConfiguration *CallJoinRequestBuilderGetRequestConfiguration)(JoinGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateJoin400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateJoinGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(JoinGetResponseable), nil
}
// ToGetRequestInformation call to allow if a client is allowed to join a call or not
func (m *CallJoinRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallJoinRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CallJoinRequestBuilder) WithUrl(rawUrl string)(*CallJoinRequestBuilder) {
    return NewCallJoinRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
