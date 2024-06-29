package userimageset

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserImageSetRequestBuilder builds and executes requests for operations under \user.image.set
type UserImageSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserImageSetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserImageSetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserImageSetRequestBuilderInternal instantiates a new UserImageSetRequestBuilder and sets the default values.
func NewUserImageSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserImageSetRequestBuilder) {
    m := &UserImageSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.image.set", pathParameters),
    }
    return m
}
// NewUserImageSetRequestBuilder instantiates a new UserImageSetRequestBuilder and sets the default values.
func NewUserImageSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserImageSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserImageSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set a profile image of currently logged in user. Three versionsof the image will be automatically stored and can be retrievedusing user.image.get api
// Deprecated: This method is obsolete. Use PostAsSetPostResponse instead.
func (m *UserImageSetRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *UserImageSetRequestBuilderPostRequestConfiguration)(SetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetResponseable), nil
}
// PostAsSetPostResponse set a profile image of currently logged in user. Three versionsof the image will be automatically stored and can be retrievedusing user.image.get api
func (m *UserImageSetRequestBuilder) PostAsSetPostResponse(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *UserImageSetRequestBuilderPostRequestConfiguration)(SetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateSet400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SetPostResponseable), nil
}
// ToPostRequestInformation set a profile image of currently logged in user. Three versionsof the image will be automatically stored and can be retrievedusing user.image.get api
func (m *UserImageSetRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *UserImageSetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserImageSetRequestBuilder) WithUrl(rawUrl string)(*UserImageSetRequestBuilder) {
    return NewUserImageSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
