package userverify

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifyPostRequestBody 
type VerifyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email or phone of the user to verify
    user *string
    // The verify code received by the user for verification
    verify_code *string
}
// NewVerifyPostRequestBody instantiates a new verifyPostRequestBody and sets the default values.
func NewVerifyPostRequestBody()(*VerifyPostRequestBody) {
    m := &VerifyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVerifyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifyPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["verify_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifyCode(val)
        }
        return nil
    }
    return res
}
// GetUser gets the user property value. The email or phone of the user to verify
func (m *VerifyPostRequestBody) GetUser()(*string) {
    return m.user
}
// GetVerifyCode gets the verify_code property value. The verify code received by the user for verification
func (m *VerifyPostRequestBody) GetVerifyCode()(*string) {
    return m.verify_code
}
// Serialize serializes information the current object
func (m *VerifyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("verify_code", m.GetVerifyCode())
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
func (m *VerifyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUser sets the user property value. The email or phone of the user to verify
func (m *VerifyPostRequestBody) SetUser(value *string)() {
    m.user = value
}
// SetVerifyCode sets the verify_code property value. The verify code received by the user for verification
func (m *VerifyPostRequestBody) SetVerifyCode(value *string)() {
    m.verify_code = value
}
// VerifyPostRequestBodyable 
type VerifyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUser()(*string)
    GetVerifyCode()(*string)
    SetUser(value *string)()
    SetVerifyCode(value *string)()
}
