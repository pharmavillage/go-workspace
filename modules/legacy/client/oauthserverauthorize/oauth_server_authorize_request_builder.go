package oauthserverauthorize

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OauthServerAuthorizeRequestBuilder builds and executes requests for operations under \oauth.server.authorize
type OauthServerAuthorizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthorizeGetResponse composed type wrapper for classes authorizeGetResponseMember1, authorizeGetResponseMember2
type AuthorizeGetResponse struct {
    // Composed type representation for type authorizeGetResponseMember1
    authorizeGetResponseMember1 AuthorizeGetResponseMember1able
    // Composed type representation for type authorizeGetResponseMember2
    authorizeGetResponseMember2 AuthorizeGetResponseMember2able
}
// NewAuthorizeGetResponse instantiates a new authorizeGetResponse and sets the default values.
func NewAuthorizeGetResponse()(*AuthorizeGetResponse) {
    m := &AuthorizeGetResponse{
    }
    return m
}
// CreateAuthorizeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAuthorizeGetResponse()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetAuthorizeGetResponseMember1 gets the authorizeGetResponseMember1 property value. Composed type representation for type authorizeGetResponseMember1
func (m *AuthorizeGetResponse) GetAuthorizeGetResponseMember1()(AuthorizeGetResponseMember1able) {
    return m.authorizeGetResponseMember1
}
// GetAuthorizeGetResponseMember2 gets the authorizeGetResponseMember2 property value. Composed type representation for type authorizeGetResponseMember2
func (m *AuthorizeGetResponse) GetAuthorizeGetResponseMember2()(AuthorizeGetResponseMember2able) {
    return m.authorizeGetResponseMember2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *AuthorizeGetResponse) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *AuthorizeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAuthorizeGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAuthorizeGetResponseMember1())
        if err != nil {
            return err
        }
    } else if m.GetAuthorizeGetResponseMember2() != nil {
        err := writer.WriteObjectValue("", m.GetAuthorizeGetResponseMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizeGetResponseMember1 sets the authorizeGetResponseMember1 property value. Composed type representation for type authorizeGetResponseMember1
func (m *AuthorizeGetResponse) SetAuthorizeGetResponseMember1(value AuthorizeGetResponseMember1able)() {
    m.authorizeGetResponseMember1 = value
}
// SetAuthorizeGetResponseMember2 sets the authorizeGetResponseMember2 property value. Composed type representation for type authorizeGetResponseMember2
func (m *AuthorizeGetResponse) SetAuthorizeGetResponseMember2(value AuthorizeGetResponseMember2able)() {
    m.authorizeGetResponseMember2 = value
}
// AuthorizeResponse composed type wrapper for classes authorizeGetResponseMember1, authorizeGetResponseMember2
type AuthorizeResponse struct {
    // Composed type representation for type authorizeGetResponseMember1
    authorizeGetResponseMember1 AuthorizeGetResponseMember1able
    // Composed type representation for type authorizeGetResponseMember2
    authorizeGetResponseMember2 AuthorizeGetResponseMember2able
}
// NewAuthorizeResponse instantiates a new authorizeResponse and sets the default values.
func NewAuthorizeResponse()(*AuthorizeResponse) {
    m := &AuthorizeResponse{
    }
    return m
}
// CreateAuthorizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAuthorizeResponse()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetAuthorizeGetResponseMember1 gets the authorizeGetResponseMember1 property value. Composed type representation for type authorizeGetResponseMember1
func (m *AuthorizeResponse) GetAuthorizeGetResponseMember1()(AuthorizeGetResponseMember1able) {
    return m.authorizeGetResponseMember1
}
// GetAuthorizeGetResponseMember2 gets the authorizeGetResponseMember2 property value. Composed type representation for type authorizeGetResponseMember2
func (m *AuthorizeResponse) GetAuthorizeGetResponseMember2()(AuthorizeGetResponseMember2able) {
    return m.authorizeGetResponseMember2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizeResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *AuthorizeResponse) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *AuthorizeResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAuthorizeGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAuthorizeGetResponseMember1())
        if err != nil {
            return err
        }
    } else if m.GetAuthorizeGetResponseMember2() != nil {
        err := writer.WriteObjectValue("", m.GetAuthorizeGetResponseMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizeGetResponseMember1 sets the authorizeGetResponseMember1 property value. Composed type representation for type authorizeGetResponseMember1
func (m *AuthorizeResponse) SetAuthorizeGetResponseMember1(value AuthorizeGetResponseMember1able)() {
    m.authorizeGetResponseMember1 = value
}
// SetAuthorizeGetResponseMember2 sets the authorizeGetResponseMember2 property value. Composed type representation for type authorizeGetResponseMember2
func (m *AuthorizeResponse) SetAuthorizeGetResponseMember2(value AuthorizeGetResponseMember2able)() {
    m.authorizeGetResponseMember2 = value
}
// OauthServerAuthorizeRequestBuilderGetQueryParameters this endpoint must be called by the UI to authorize a OAuth access request. To use this endpoint, the caller must be authenticated, and it must provide all the required oauth parameters. If the user don't have any token already generated for the same oauth app/scope, the endpoint return the application and scopes info, so the UI can show it to the user and ask for his authorization (calling `oauth.server.authorize` endpoint to authorize the request). If the user already have a valid token, a new authorization is not required, so it will redirect to the callback uri with the generated authorization code. If the parameter `redirect` is not provided, or set to true, the redirect is made using a 302 response. If this parameter is set to true, then a 200 response is returned, containing the redirect uri on the body.
type OauthServerAuthorizeRequestBuilderGetQueryParameters struct {
    // A valid OAuth client_id
    Client_id *string `uriparametername:"client_id"`
    // Defines if the request should be redirected (302) directly or just return the redirect_uri on the reponse body.
    Redirect *bool `uriparametername:"redirect"`
    // The URI to redirect on an successful authorization. This URI must be provided, even if the caller isn't redirecting. It's used for validation.
    Redirect_uri *string `uriparametername:"redirect_uri"`
    // The type of the authorization. Only `code` is supported for now
    // Deprecated: This property is deprecated, use response_typeAsGetResponse_typeQueryParameterType instead
    Response_type *string `uriparametername:"response_type"`
    // The type of the authorization. Only `code` is supported for now
    Response_typeAsGetResponse_typeQueryParameterType *GetResponse_typeQueryParameterType `uriparametername:"response_type"`
    // A space separated list of scopes that the client app is requiring.
    Scope *string `uriparametername:"scope"`
}
// OauthServerAuthorizeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthServerAuthorizeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OauthServerAuthorizeRequestBuilderGetQueryParameters
}
// AuthorizeGetResponseable 
type AuthorizeGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizeGetResponseMember1()(AuthorizeGetResponseMember1able)
    GetAuthorizeGetResponseMember2()(AuthorizeGetResponseMember2able)
    SetAuthorizeGetResponseMember1(value AuthorizeGetResponseMember1able)()
    SetAuthorizeGetResponseMember2(value AuthorizeGetResponseMember2able)()
}
// AuthorizeResponseable 
type AuthorizeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizeGetResponseMember1()(AuthorizeGetResponseMember1able)
    GetAuthorizeGetResponseMember2()(AuthorizeGetResponseMember2able)
    SetAuthorizeGetResponseMember1(value AuthorizeGetResponseMember1able)()
    SetAuthorizeGetResponseMember2(value AuthorizeGetResponseMember2able)()
}
// NewOauthServerAuthorizeRequestBuilderInternal instantiates a new OauthServerAuthorizeRequestBuilder and sets the default values.
func NewOauthServerAuthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerAuthorizeRequestBuilder) {
    m := &OauthServerAuthorizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/oauth.server.authorize?client_id={client_id}&redirect_uri={redirect_uri}&response_type={response_type}{&redirect*,scope*}", pathParameters),
    }
    return m
}
// NewOauthServerAuthorizeRequestBuilder instantiates a new OauthServerAuthorizeRequestBuilder and sets the default values.
func NewOauthServerAuthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OauthServerAuthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOauthServerAuthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this endpoint must be called by the UI to authorize a OAuth access request. To use this endpoint, the caller must be authenticated, and it must provide all the required oauth parameters. If the user don't have any token already generated for the same oauth app/scope, the endpoint return the application and scopes info, so the UI can show it to the user and ask for his authorization (calling `oauth.server.authorize` endpoint to authorize the request). If the user already have a valid token, a new authorization is not required, so it will redirect to the callback uri with the generated authorization code. If the parameter `redirect` is not provided, or set to true, the redirect is made using a 302 response. If this parameter is set to true, then a 200 response is returned, containing the redirect uri on the body.
// Deprecated: This method is obsolete. Use GetAsAuthorizeGetResponse instead.
func (m *OauthServerAuthorizeRequestBuilder) Get(ctx context.Context, requestConfiguration *OauthServerAuthorizeRequestBuilderGetRequestConfiguration)(AuthorizeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAuthorize400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthorizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthorizeResponseable), nil
}
// GetAsAuthorizeGetResponse this endpoint must be called by the UI to authorize a OAuth access request. To use this endpoint, the caller must be authenticated, and it must provide all the required oauth parameters. If the user don't have any token already generated for the same oauth app/scope, the endpoint return the application and scopes info, so the UI can show it to the user and ask for his authorization (calling `oauth.server.authorize` endpoint to authorize the request). If the user already have a valid token, a new authorization is not required, so it will redirect to the callback uri with the generated authorization code. If the parameter `redirect` is not provided, or set to true, the redirect is made using a 302 response. If this parameter is set to true, then a 200 response is returned, containing the redirect uri on the body.
func (m *OauthServerAuthorizeRequestBuilder) GetAsAuthorizeGetResponse(ctx context.Context, requestConfiguration *OauthServerAuthorizeRequestBuilderGetRequestConfiguration)(AuthorizeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateAuthorize400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthorizeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthorizeGetResponseable), nil
}
// ToGetRequestInformation this endpoint must be called by the UI to authorize a OAuth access request. To use this endpoint, the caller must be authenticated, and it must provide all the required oauth parameters. If the user don't have any token already generated for the same oauth app/scope, the endpoint return the application and scopes info, so the UI can show it to the user and ask for his authorization (calling `oauth.server.authorize` endpoint to authorize the request). If the user already have a valid token, a new authorization is not required, so it will redirect to the callback uri with the generated authorization code. If the parameter `redirect` is not provided, or set to true, the redirect is made using a 302 response. If this parameter is set to true, then a 200 response is returned, containing the redirect uri on the body.
func (m *OauthServerAuthorizeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OauthServerAuthorizeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OauthServerAuthorizeRequestBuilder) WithUrl(rawUrl string)(*OauthServerAuthorizeRequestBuilder) {
    return NewOauthServerAuthorizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
