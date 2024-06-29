package userinfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserInfoRequestBuilder builds and executes requests for operations under \user.info
type UserInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInfoRequestBuilderGetQueryParameters get information about an user. user_id or email is required for this endpoint. This endpoint can be accessed using a notification token (without been authenticated). If none of the parameters are included (email, phone or user_id), the logged user info will be returned.
type UserInfoRequestBuilderGetQueryParameters struct {
    // The email id of the user to retrieve
    Email *string `uriparametername:"email"`
    // The phone of the user to retrieve
    Phone *string `uriparametername:"phone"`
    // The user id of the user to retrieve
    User_id *string `uriparametername:"user_id"`
}
// UserInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInfoRequestBuilderGetQueryParameters
}
// NewUserInfoRequestBuilderInternal instantiates a new UserInfoRequestBuilder and sets the default values.
func NewUserInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInfoRequestBuilder) {
    m := &UserInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user.info{?email*,phone*,user_id*}", pathParameters),
    }
    return m
}
// NewUserInfoRequestBuilder instantiates a new UserInfoRequestBuilder and sets the default values.
func NewUserInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get information about an user. user_id or email is required for this endpoint. This endpoint can be accessed using a notification token (without been authenticated). If none of the parameters are included (email, phone or user_id), the logged user info will be returned.
// Deprecated: This method is obsolete. Use GetAsInfoGetResponse instead.
func (m *UserInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInfoRequestBuilderGetRequestConfiguration)(InfoResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoResponseable), nil
}
// GetAsInfoGetResponse get information about an user. user_id or email is required for this endpoint. This endpoint can be accessed using a notification token (without been authenticated). If none of the parameters are included (email, phone or user_id), the logged user info will be returned.
func (m *UserInfoRequestBuilder) GetAsInfoGetResponse(ctx context.Context, requestConfiguration *UserInfoRequestBuilderGetRequestConfiguration)(InfoGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateInfo400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInfoGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InfoGetResponseable), nil
}
// ToGetRequestInformation get information about an user. user_id or email is required for this endpoint. This endpoint can be accessed using a notification token (without been authenticated). If none of the parameters are included (email, phone or user_id), the logged user info will be returned.
func (m *UserInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserInfoRequestBuilder) WithUrl(rawUrl string)(*UserInfoRequestBuilder) {
    return NewUserInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
