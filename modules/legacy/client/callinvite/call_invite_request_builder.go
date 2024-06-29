package callinvite

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallInviteRequestBuilder builds and executes requests for operations under \call.invite
type CallInviteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallInviteRequestBuilderGetQueryParameters invite a user to join a private call. Issuing this for public call will be rejected. A websocket notification will be sent to the user on invite.
type CallInviteRequestBuilderGetQueryParameters struct {
    // Hash of the private call to send the invite for
    Call_hash *string `uriparametername:"call_hash"`
    // Id of the user to invite
    User_id *string `uriparametername:"user_id"`
}
// CallInviteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallInviteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallInviteRequestBuilderGetQueryParameters
}
// NewCallInviteRequestBuilderInternal instantiates a new CallInviteRequestBuilder and sets the default values.
func NewCallInviteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallInviteRequestBuilder) {
    m := &CallInviteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/call.invite?call_hash={call_hash}&user_id={user_id}", pathParameters),
    }
    return m
}
// NewCallInviteRequestBuilder instantiates a new CallInviteRequestBuilder and sets the default values.
func NewCallInviteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallInviteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallInviteRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invite a user to join a private call. Issuing this for public call will be rejected. A websocket notification will be sent to the user on invite.
// Deprecated: This method is obsolete. Use GetAsInviteGetResponse instead.
func (m *CallInviteRequestBuilder) Get(ctx context.Context, requestConfiguration *CallInviteRequestBuilderGetRequestConfiguration)(InviteResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInvite400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInviteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InviteResponseable), nil
}
// GetAsInviteGetResponse invite a user to join a private call. Issuing this for public call will be rejected. A websocket notification will be sent to the user on invite.
func (m *CallInviteRequestBuilder) GetAsInviteGetResponse(ctx context.Context, requestConfiguration *CallInviteRequestBuilderGetRequestConfiguration)(InviteGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInvite400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInviteGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InviteGetResponseable), nil
}
// ToGetRequestInformation invite a user to join a private call. Issuing this for public call will be rejected. A websocket notification will be sent to the user on invite.
func (m *CallInviteRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallInviteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CallInviteRequestBuilder) WithUrl(rawUrl string)(*CallInviteRequestBuilder) {
    return NewCallInviteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
