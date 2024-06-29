package userprofileset

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetPostRequestBody 
type SetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Email
    email *string
    // Display name
    name *string
    // Phone number
    phone *string
}
// NewSetPostRequestBody instantiates a new setPostRequestBody and sets the default values.
func NewSetPostRequestBody()(*SetPostRequestBody) {
    m := &SetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. Email
func (m *SetPostRequestBody) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Display name
func (m *SetPostRequestBody) GetName()(*string) {
    return m.name
}
// GetPhone gets the phone property value. Phone number
func (m *SetPostRequestBody) GetPhone()(*string) {
    return m.phone
}
// Serialize serializes information the current object
func (m *SetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteStringValue("phone", m.GetPhone())
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
func (m *SetPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. Email
func (m *SetPostRequestBody) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. Display name
func (m *SetPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPhone sets the phone property value. Phone number
func (m *SetPostRequestBody) SetPhone(value *string)() {
    m.phone = value
}
// SetPostRequestBodyable 
type SetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetName()(*string)
    GetPhone()(*string)
    SetEmail(value *string)()
    SetName(value *string)()
    SetPhone(value *string)()
}
