package useralertack

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserAlertAckRequestBuilder builds and executes requests for operations under \user.alert.ack
type UserAlertAckRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserAlertAckRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserAlertAckRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserAlertAckRequestBuilderInternal instantiates a new UserAlertAckRequestBuilder and sets the default values.
func NewUserAlertAckRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserAlertAckRequestBuilder) {
    m := &UserAlertAckRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.alert.ack", pathParameters),
    }
    return m
}
// NewUserAlertAckRequestBuilder instantiates a new UserAlertAckRequestBuilder and sets the default values.
func NewUserAlertAckRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserAlertAckRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserAlertAckRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark an alert as read
// Deprecated: This method is obsolete. Use PostAsAckPostResponse instead.
func (m *UserAlertAckRequestBuilder) Post(ctx context.Context, body AckPostRequestBodyable, requestConfiguration *UserAlertAckRequestBuilderPostRequestConfiguration)(AckResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAck400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAckResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AckResponseable), nil
}
// PostAsAckPostResponse mark an alert as read
func (m *UserAlertAckRequestBuilder) PostAsAckPostResponse(ctx context.Context, body AckPostRequestBodyable, requestConfiguration *UserAlertAckRequestBuilderPostRequestConfiguration)(AckPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAck400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAckPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AckPostResponseable), nil
}
// ToPostRequestInformation mark an alert as read
func (m *UserAlertAckRequestBuilder) ToPostRequestInformation(ctx context.Context, body AckPostRequestBodyable, requestConfiguration *UserAlertAckRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
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
func (m *UserAlertAckRequestBuilder) WithUrl(rawUrl string)(*UserAlertAckRequestBuilder) {
    return NewUserAlertAckRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
