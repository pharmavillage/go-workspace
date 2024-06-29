package channelhistory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChannelHistoryRequestBuilder builds and executes requests for operations under \channel.history
type ChannelHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChannelHistoryRequestBuilderGetQueryParameters get messages for a channel.This endpoint can be accessed using a notification token (without been authenticated)
type ChannelHistoryRequestBuilderGetQueryParameters struct {
    // Id of channel to get information
    Channel_id *int32 `uriparametername:"channel_id"`
    // Paginate through collections of data by setting the cursor parameter to a next_cursor attribute returned by a previous request's response_metadata.Optional. If this is not supplied, Newest message in the system is used as the cursor.This will NOT INCLUDE the message pointed by this cursor in the result set when retrieving OLDER messages . This will INCLUDE the message pointed by this cursor in the result set when retrieving NEWER messages (limit_newer is supplied)
    Cursor *string `uriparametername:"cursor"`
    // Number of older messages to retrieve. Optional. (Default is 10 unless limit_newer is supplied in which case default is 0)
    Limit *int32 `uriparametername:"limit"`
    // Number of newer messages to retrieve. Optional. Default is 0
    Limit_newer *int32 `uriparametername:"limit_newer"`
}
// ChannelHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelHistoryRequestBuilderGetQueryParameters
}
// NewChannelHistoryRequestBuilderInternal instantiates a new ChannelHistoryRequestBuilder and sets the default values.
func NewChannelHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelHistoryRequestBuilder) {
    m := &ChannelHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/channel.history?channel_id={channel_id}{&cursor*,limit*,limit_newer*}", pathParameters),
    }
    return m
}
// NewChannelHistoryRequestBuilder instantiates a new ChannelHistoryRequestBuilder and sets the default values.
func NewChannelHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get messages for a channel.This endpoint can be accessed using a notification token (without been authenticated)
// Deprecated: This method is obsolete. Use GetAsHistoryGetResponse instead.
func (m *ChannelHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelHistoryRequestBuilderGetRequestConfiguration)(HistoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateHistory400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHistoryResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(HistoryResponseable), nil
}
// GetAsHistoryGetResponse get messages for a channel.This endpoint can be accessed using a notification token (without been authenticated)
func (m *ChannelHistoryRequestBuilder) GetAsHistoryGetResponse(ctx context.Context, requestConfiguration *ChannelHistoryRequestBuilderGetRequestConfiguration)(HistoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateHistory400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHistoryGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(HistoryGetResponseable), nil
}
// ToGetRequestInformation get messages for a channel.This endpoint can be accessed using a notification token (without been authenticated)
func (m *ChannelHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChannelHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelHistoryRequestBuilder) WithUrl(rawUrl string)(*ChannelHistoryRequestBuilder) {
    return NewChannelHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
