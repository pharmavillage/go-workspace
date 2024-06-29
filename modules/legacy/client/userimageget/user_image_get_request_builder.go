package userimageget

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserImageGetRequestBuilder builds and executes requests for operations under \user.image.get
type UserImageGetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserImageGetRequestBuilderGetQueryParameters get profile image of a  user.This endpoint can be accessed using a notification token (without been authenticated). The auth/token expirationis not verified for this endpoint.
type UserImageGetRequestBuilderGetQueryParameters struct {
    // This can be small, medium or full to retrieve the image
    Image_class *string `uriparametername:"image_class"`
    // The user id of the profile image to retrieve
    User_id *string `uriparametername:"user_id"`
}
// UserImageGetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserImageGetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserImageGetRequestBuilderGetQueryParameters
}
// NewUserImageGetRequestBuilderInternal instantiates a new UserImageGetRequestBuilder and sets the default values.
func NewUserImageGetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserImageGetRequestBuilder) {
    m := &UserImageGetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.image.get?image_class={image_class}&user_id={user_id}", pathParameters),
    }
    return m
}
// NewUserImageGetRequestBuilder instantiates a new UserImageGetRequestBuilder and sets the default values.
func NewUserImageGetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserImageGetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserImageGetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get profile image of a  user.This endpoint can be accessed using a notification token (without been authenticated). The auth/token expirationis not verified for this endpoint.
func (m *UserImageGetRequestBuilder) Get(ctx context.Context, requestConfiguration *UserImageGetRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateGet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get profile image of a  user.This endpoint can be accessed using a notification token (without been authenticated). The auth/token expirationis not verified for this endpoint.
func (m *UserImageGetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserImageGetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "image/*, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserImageGetRequestBuilder) WithUrl(rawUrl string)(*UserImageGetRequestBuilder) {
    return NewUserImageGetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
