package callinviteaccept

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallInviteAcceptRequestBuilder builds and executes requests for operations under \call.invite.accept
type CallInviteAcceptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallInviteAcceptRequestBuilderGetQueryParameters called when user accepts or rejects a invite to join a call. A notification call.invite.accept will be sent to the user connections. This is useful to stop ringing if user responds
type CallInviteAcceptRequestBuilderGetQueryParameters struct {
    // True if user joins and false if the user rejects the call invite
    Accept *bool `uriparametername:"accept"`
    // Hash of the private call to send the invite for
    Call_hash *string `uriparametername:"call_hash"`
}
// CallInviteAcceptRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallInviteAcceptRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallInviteAcceptRequestBuilderGetQueryParameters
}
// NewCallInviteAcceptRequestBuilderInternal instantiates a new CallInviteAcceptRequestBuilder and sets the default values.
func NewCallInviteAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallInviteAcceptRequestBuilder) {
    m := &CallInviteAcceptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/call.invite.accept?accept={accept}&call_hash={call_hash}", pathParameters),
    }
    return m
}
// NewCallInviteAcceptRequestBuilder instantiates a new CallInviteAcceptRequestBuilder and sets the default values.
func NewCallInviteAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallInviteAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallInviteAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Get called when user accepts or rejects a invite to join a call. A notification call.invite.accept will be sent to the user connections. This is useful to stop ringing if user responds
// Deprecated: This method is obsolete. Use GetAsAcceptGetResponse instead.
func (m *CallInviteAcceptRequestBuilder) Get(ctx context.Context, requestConfiguration *CallInviteAcceptRequestBuilderGetRequestConfiguration)(AcceptResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAccept400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAcceptResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AcceptResponseable), nil
}
// GetAsAcceptGetResponse called when user accepts or rejects a invite to join a call. A notification call.invite.accept will be sent to the user connections. This is useful to stop ringing if user responds
func (m *CallInviteAcceptRequestBuilder) GetAsAcceptGetResponse(ctx context.Context, requestConfiguration *CallInviteAcceptRequestBuilderGetRequestConfiguration)(AcceptGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAccept400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAcceptGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AcceptGetResponseable), nil
}
// ToGetRequestInformation called when user accepts or rejects a invite to join a call. A notification call.invite.accept will be sent to the user connections. This is useful to stop ringing if user responds
func (m *CallInviteAcceptRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallInviteAcceptRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CallInviteAcceptRequestBuilder) WithUrl(rawUrl string)(*CallInviteAcceptRequestBuilder) {
    return NewCallInviteAcceptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
