package callstatus

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallStatusRequestBuilder builds and executes requests for operations under \call.status
type CallStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallStatusRequestBuilderGetQueryParameters send a channel update to all users of a channel - valid only for public channel
type CallStatusRequestBuilderGetQueryParameters struct {
    // Hash of the private call to send update
    Call_hash *string `uriparametername:"call_hash"`
    // Id of the channel to send the update
    Channel_id *string `uriparametername:"channel_id"`
}
// CallStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallStatusRequestBuilderGetQueryParameters
}
// NewCallStatusRequestBuilderInternal instantiates a new CallStatusRequestBuilder and sets the default values.
func NewCallStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallStatusRequestBuilder) {
    m := &CallStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/call.status?call_hash={call_hash}{&channel_id*}", pathParameters),
    }
    return m
}
// NewCallStatusRequestBuilder instantiates a new CallStatusRequestBuilder and sets the default values.
func NewCallStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get send a channel update to all users of a channel - valid only for public channel
// Deprecated: This method is obsolete. Use GetAsStatusGetResponse instead.
func (m *CallStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *CallStatusRequestBuilderGetRequestConfiguration)(StatusResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateStatus400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStatusResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StatusResponseable), nil
}
// GetAsStatusGetResponse send a channel update to all users of a channel - valid only for public channel
func (m *CallStatusRequestBuilder) GetAsStatusGetResponse(ctx context.Context, requestConfiguration *CallStatusRequestBuilderGetRequestConfiguration)(StatusGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateStatus400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStatusGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StatusGetResponseable), nil
}
// ToGetRequestInformation send a channel update to all users of a channel - valid only for public channel
func (m *CallStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CallStatusRequestBuilder) WithUrl(rawUrl string)(*CallStatusRequestBuilder) {
    return NewCallStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
