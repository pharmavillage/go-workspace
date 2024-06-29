package actionlist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActionListRequestBuilder builds and executes requests for operations under \action.list
type ActionListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionListRequestBuilderGetQueryParameters get list of all actions for a channel
type ActionListRequestBuilderGetQueryParameters struct {
    // Just return the actions from this channel.
    Channel_id *int32 `uriparametername:"channel_id"`
    // This value MUST be a valid action_id, base64 encoded. It's used by the infinite pagination, to define the point where to start to bring the results (the number of results are defined by the limit_after and limit_before params. The cursor action, by default is not included on the results, unless both limit_before and limit_after are defined.
    Cursor *string `uriparametername:"cursor"`
    // Is the max number of actions that will be returned after the defined cursor. If the cursor is not defined, it's the number of actions returned from the beggining. When no limit is defined (neither limit_after or limit_before), limit_after defaults to 30.
    Limit_after *int32 `uriparametername:"limit_after"`
    // Is the max number of actions that will be returned before the cursor. This parameter only makes sense when the cursor is defined (it's ignored otherwise). When both limit_before, and limit_after are defined, the cursor action is included on the results.
    Limit_before *int32 `uriparametername:"limit_before"`
    // A search term. Only actions that match this search term (on name or description), will be returned.
    Search *string `uriparametername:"search"`
    // Define the sorting method to return the list of actions. Possible values are: - default: The default sorting defined by the user (by drag-drop). This sorting schema is only available  when the channel_id filter is defined.- name: Sort by the action name - channel: Sort by the channel name - due_date: Sort by due date (considering sub-tasks).
    // Deprecated: This property is deprecated, use sort_byAsGetSort_byQueryParameterType instead
    Sort_by *string `uriparametername:"sort_by"`
    // Define the sorting method to return the list of actions. Possible values are: - default: The default sorting defined by the user (by drag-drop). This sorting schema is only available  when the channel_id filter is defined.- name: Sort by the action name - channel: Sort by the channel name - due_date: Sort by due date (considering sub-tasks).
    Sort_byAsGetSort_byQueryParameterType *GetSort_byQueryParameterType `uriparametername:"sort_by"`
    // Changes the sorting direction. By default all sorting schemas uses ascending direction, if this flag is set the sorting direction will be descending.
    Sort_desc *bool `uriparametername:"sort_desc"`
    // Just return the actions that have this status (0 for pending, 1 for complete)
    Status *int32 `uriparametername:"status"`
    // Just return the actions that are assigned to this user. MUST be a valid user id.
    User_id *int32 `uriparametername:"user_id"`
}
// ActionListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActionListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ActionListRequestBuilderGetQueryParameters
}
// NewActionListRequestBuilderInternal instantiates a new ActionListRequestBuilder and sets the default values.
func NewActionListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionListRequestBuilder) {
    m := &ActionListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/action.list?channel_id={channel_id}{&cursor*,limit_after*,limit_before*,search*,sort_by*,sort_desc*,status*,user_id*}", pathParameters),
    }
    return m
}
// NewActionListRequestBuilder instantiates a new ActionListRequestBuilder and sets the default values.
func NewActionListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActionListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get list of all actions for a channel
// Deprecated: This method is obsolete. Use GetAsListGetResponse instead.
func (m *ActionListRequestBuilder) Get(ctx context.Context, requestConfiguration *ActionListRequestBuilderGetRequestConfiguration)(ListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListResponseable), nil
}
// GetAsListGetResponse get list of all actions for a channel
func (m *ActionListRequestBuilder) GetAsListGetResponse(ctx context.Context, requestConfiguration *ActionListRequestBuilderGetRequestConfiguration)(ListGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateList400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateListGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ListGetResponseable), nil
}
// ToGetRequestInformation get list of all actions for a channel
func (m *ActionListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ActionListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ActionListRequestBuilder) WithUrl(rawUrl string)(*ActionListRequestBuilder) {
    return NewActionListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
