package userlogin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoginPostRequestBody 
type LoginPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The password property
    password *string
    // The remember_me property
    remember_me *bool
    // Browser timezone. This value will be ignored if the user already have a timezone set.
    timezone *string
    // The user property
    user *string
}
// NewLoginPostRequestBody instantiates a new loginPostRequestBody and sets the default values.
func NewLoginPostRequestBody()(*LoginPostRequestBody) {
    m := &LoginPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLoginPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoginPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoginPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoginPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["remember_me"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRememberMe(val)
        }
        return nil
    }
    res["timezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimezone(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
func (m *LoginPostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRememberMe gets the remember_me property value. The remember_me property
func (m *LoginPostRequestBody) GetRememberMe()(*bool) {
    return m.remember_me
}
// GetTimezone gets the timezone property value. Browser timezone. This value will be ignored if the user already have a timezone set.
func (m *LoginPostRequestBody) GetTimezone()(*string) {
    return m.timezone
}
// GetUser gets the user property value. The user property
func (m *LoginPostRequestBody) GetUser()(*string) {
    return m.user
}
// Serialize serializes information the current object
func (m *LoginPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("remember_me", m.GetRememberMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timezone", m.GetTimezone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *LoginPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPassword sets the password property value. The password property
func (m *LoginPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRememberMe sets the remember_me property value. The remember_me property
func (m *LoginPostRequestBody) SetRememberMe(value *bool)() {
    m.remember_me = value
}
// SetTimezone sets the timezone property value. Browser timezone. This value will be ignored if the user already have a timezone set.
func (m *LoginPostRequestBody) SetTimezone(value *string)() {
    m.timezone = value
}
// SetUser sets the user property value. The user property
func (m *LoginPostRequestBody) SetUser(value *string)() {
    m.user = value
}
// LoginPostRequestBodyable 
type LoginPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    GetRememberMe()(*bool)
    GetTimezone()(*string)
    GetUser()(*string)
    SetPassword(value *string)()
    SetRememberMe(value *bool)()
    SetTimezone(value *string)()
    SetUser(value *string)()
}
