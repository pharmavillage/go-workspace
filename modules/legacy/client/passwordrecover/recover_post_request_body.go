package passwordrecover

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecoverPostRequestBody 
type RecoverPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The opt_email property
    opt_email *string
    // The opt_phone property
    opt_phone *string
}
// NewRecoverPostRequestBody instantiates a new recoverPostRequestBody and sets the default values.
func NewRecoverPostRequestBody()(*RecoverPostRequestBody) {
    m := &RecoverPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecoverPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecoverPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecoverPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecoverPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecoverPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["opt_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptEmail(val)
        }
        return nil
    }
    res["opt_phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptPhone(val)
        }
        return nil
    }
    return res
}
// GetOptEmail gets the opt_email property value. The opt_email property
func (m *RecoverPostRequestBody) GetOptEmail()(*string) {
    return m.opt_email
}
// GetOptPhone gets the opt_phone property value. The opt_phone property
func (m *RecoverPostRequestBody) GetOptPhone()(*string) {
    return m.opt_phone
}
// Serialize serializes information the current object
func (m *RecoverPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("opt_email", m.GetOptEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("opt_phone", m.GetOptPhone())
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
func (m *RecoverPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOptEmail sets the opt_email property value. The opt_email property
func (m *RecoverPostRequestBody) SetOptEmail(value *string)() {
    m.opt_email = value
}
// SetOptPhone sets the opt_phone property value. The opt_phone property
func (m *RecoverPostRequestBody) SetOptPhone(value *string)() {
    m.opt_phone = value
}
// RecoverPostRequestBodyable 
type RecoverPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOptEmail()(*string)
    GetOptPhone()(*string)
    SetOptEmail(value *string)()
    SetOptPhone(value *string)()
}
