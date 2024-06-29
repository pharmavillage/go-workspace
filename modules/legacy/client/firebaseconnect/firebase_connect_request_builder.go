package firebaseconnect

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FirebaseConnectRequestBuilder builds and executes requests for operations under \firebase.connect
type FirebaseConnectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FirebaseConnectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FirebaseConnectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFirebaseConnectRequestBuilderInternal instantiates a new FirebaseConnectRequestBuilder and sets the default values.
func NewFirebaseConnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FirebaseConnectRequestBuilder) {
    m := &FirebaseConnectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/firebase.connect", pathParameters),
    }
    return m
}
// NewFirebaseConnectRequestBuilder instantiates a new FirebaseConnectRequestBuilder and sets the default values.
func NewFirebaseConnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FirebaseConnectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFirebaseConnectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this call send the FCM token (a token that uniquely identify a device subscribed to Firebase Cloud Messaging) to the backend, linking it to the logged user. It allows the system to send push notification to all the devices where the logged user is connected. To refresh the token, just send a new connection to the same user. The invalid token will be removed automatically.
func (m *FirebaseConnectRequestBuilder) Post(ctx context.Context, body ConnectPostRequestBodyable, requestConfiguration *FirebaseConnectRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateConnect400ErrorFromDiscriminatorValue,
        "422": CreateConnect422ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation this call send the FCM token (a token that uniquely identify a device subscribed to Firebase Cloud Messaging) to the backend, linking it to the logged user. It allows the system to send push notification to all the devices where the logged user is connected. To refresh the token, just send a new connection to the same user. The invalid token will be removed automatically.
func (m *FirebaseConnectRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConnectPostRequestBodyable, requestConfiguration *FirebaseConnectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FirebaseConnectRequestBuilder) WithUrl(rawUrl string)(*FirebaseConnectRequestBuilder) {
    return NewFirebaseConnectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
