package userfinalize

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FinalizePostRequestBody 
type FinalizePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name
    name *string
    // Password for the user account
    password *string
    // Email or Phone of the user
    user *string
}
// NewFinalizePostRequestBody instantiates a new finalizePostRequestBody and sets the default values.
func NewFinalizePostRequestBody()(*FinalizePostRequestBody) {
    m := &FinalizePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFinalizePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFinalizePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFinalizePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FinalizePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FinalizePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetName gets the name property value. Display name
func (m *FinalizePostRequestBody) GetName()(*string) {
    return m.name
}
// GetPassword gets the password property value. Password for the user account
func (m *FinalizePostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetUser gets the user property value. Email or Phone of the user
func (m *FinalizePostRequestBody) GetUser()(*string) {
    return m.user
}
// Serialize serializes information the current object
func (m *FinalizePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
func (m *FinalizePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Display name
func (m *FinalizePostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPassword sets the password property value. Password for the user account
func (m *FinalizePostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetUser sets the user property value. Email or Phone of the user
func (m *FinalizePostRequestBody) SetUser(value *string)() {
    m.user = value
}
// FinalizePostRequestBodyable 
type FinalizePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetPassword()(*string)
    GetUser()(*string)
    SetName(value *string)()
    SetPassword(value *string)()
    SetUser(value *string)()
}
