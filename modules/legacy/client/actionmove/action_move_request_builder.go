package actionmove

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActionMoveRequestBuilder builds and executes requests for operations under \action.move
type ActionMoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionMoveRequestBuilderPostQueryParameters changes an action position
type ActionMoveRequestBuilderPostQueryParameters struct {
    // Id of action to move
    Action_id *int32 `uriparametername:"action_id"`
}
// ActionMoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActionMoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ActionMoveRequestBuilderPostQueryParameters
}
// NewActionMoveRequestBuilderInternal instantiates a new ActionMoveRequestBuilder and sets the default values.
func NewActionMoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionMoveRequestBuilder) {
    m := &ActionMoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/action.move?action_id={action_id}", pathParameters),
    }
    return m
}
// NewActionMoveRequestBuilder instantiates a new ActionMoveRequestBuilder and sets the default values.
func NewActionMoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActionMoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActionMoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post changes an action position
// Deprecated: This method is obsolete. Use PostAsMovePostResponse instead.
func (m *ActionMoveRequestBuilder) Post(ctx context.Context, body MovePostRequestBodyable, requestConfiguration *ActionMoveRequestBuilderPostRequestConfiguration)(MoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateMove400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MoveResponseable), nil
}
// PostAsMovePostResponse changes an action position
func (m *ActionMoveRequestBuilder) PostAsMovePostResponse(ctx context.Context, body MovePostRequestBodyable, requestConfiguration *ActionMoveRequestBuilderPostRequestConfiguration)(MovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateMove400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MovePostResponseable), nil
}
// ToPostRequestInformation changes an action position
func (m *ActionMoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body MovePostRequestBodyable, requestConfiguration *ActionMoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ActionMoveRequestBuilder) WithUrl(rawUrl string)(*ActionMoveRequestBuilder) {
    return NewActionMoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
