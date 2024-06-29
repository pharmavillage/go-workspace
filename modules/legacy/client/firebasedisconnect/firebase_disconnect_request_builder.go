package firebasedisconnect

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FirebaseDisconnectRequestBuilder builds and executes requests for operations under \firebase.disconnect
type FirebaseDisconnectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FirebaseDisconnectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FirebaseDisconnectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFirebaseDisconnectRequestBuilderInternal instantiates a new FirebaseDisconnectRequestBuilder and sets the default values.
func NewFirebaseDisconnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FirebaseDisconnectRequestBuilder) {
    m := &FirebaseDisconnectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/firebase.disconnect", pathParameters),
    }
    return m
}
// NewFirebaseDisconnectRequestBuilder instantiates a new FirebaseDisconnectRequestBuilder and sets the default values.
func NewFirebaseDisconnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FirebaseDisconnectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFirebaseDisconnectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this call send the current FCM token (a token that uniquely identify a device subscribed to Firebase Cloud Messaging) to the backend, deleting it, and preventint future push notifications to this token.
func (m *FirebaseDisconnectRequestBuilder) Post(ctx context.Context, body DisconnectPostRequestBodyable, requestConfiguration *FirebaseDisconnectRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateDisconnect400ErrorFromDiscriminatorValue,
        "422": CreateDisconnect422ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation this call send the current FCM token (a token that uniquely identify a device subscribed to Firebase Cloud Messaging) to the backend, deleting it, and preventint future push notifications to this token.
func (m *FirebaseDisconnectRequestBuilder) ToPostRequestInformation(ctx context.Context, body DisconnectPostRequestBodyable, requestConfiguration *FirebaseDisconnectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FirebaseDisconnectRequestBuilder) WithUrl(rawUrl string)(*FirebaseDisconnectRequestBuilder) {
    return NewFirebaseDisconnectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
