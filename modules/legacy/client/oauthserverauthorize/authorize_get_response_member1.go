package oauthserverauthorize

import (
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizeGetResponseMember1 
type AuthorizeGetResponseMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The application description
    description *string
    // The meta property
    meta i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable
    // The name of the application that is requesting access
    name *string
    // A value that must be passed on the approve call to validate the authorization
    request_key *string
    // The scopes property
    scopes []string
}
// NewAuthorizeGetResponseMember1 instantiates a new authorizeGetResponseMember1 and sets the default values.
func NewAuthorizeGetResponseMember1()(*AuthorizeGetResponseMember1) {
    m := &AuthorizeGetResponseMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthorizeGetResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizeGetResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorizeGetResponseMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthorizeGetResponseMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The application description
func (m *AuthorizeGetResponseMember1) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizeGetResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["meta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateMetaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["request_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestKey(val)
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetScopes(res)
        }
        return nil
    }
    return res
}
// GetMeta gets the meta property value. The meta property
func (m *AuthorizeGetResponseMember1) GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable) {
    return m.meta
}
// GetName gets the name property value. The name of the application that is requesting access
func (m *AuthorizeGetResponseMember1) GetName()(*string) {
    return m.name
}
// GetRequestKey gets the request_key property value. A value that must be passed on the approve call to validate the authorization
func (m *AuthorizeGetResponseMember1) GetRequestKey()(*string) {
    return m.request_key
}
// GetScopes gets the scopes property value. The scopes property
func (m *AuthorizeGetResponseMember1) GetScopes()([]string) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *AuthorizeGetResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("meta", m.GetMeta())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("request_key", m.GetRequestKey())
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        err := writer.WriteCollectionOfStringValues("scopes", m.GetScopes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthorizeGetResponseMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The application description
func (m *AuthorizeGetResponseMember1) SetDescription(value *string)() {
    m.description = value
}
// SetMeta sets the meta property value. The meta property
func (m *AuthorizeGetResponseMember1) SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)() {
    m.meta = value
}
// SetName sets the name property value. The name of the application that is requesting access
func (m *AuthorizeGetResponseMember1) SetName(value *string)() {
    m.name = value
}
// SetRequestKey sets the request_key property value. A value that must be passed on the approve call to validate the authorization
func (m *AuthorizeGetResponseMember1) SetRequestKey(value *string)() {
    m.request_key = value
}
// SetScopes sets the scopes property value. The scopes property
func (m *AuthorizeGetResponseMember1) SetScopes(value []string)() {
    m.scopes = value
}
// AuthorizeGetResponseMember1able 
type AuthorizeGetResponseMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)
    GetName()(*string)
    GetRequestKey()(*string)
    GetScopes()([]string)
    SetDescription(value *string)()
    SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)()
    SetName(value *string)()
    SetRequestKey(value *string)()
    SetScopes(value []string)()
}
