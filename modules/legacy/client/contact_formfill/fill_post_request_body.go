package contact_formfill

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FillPostRequestBody 
type FillPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email of the user that is filling the form
    form_filler_email *string
    // The message filled on the form
    form_filler_message *string
    // The name of the user that is filling the form
    form_filler_name *string
    // The form id to be filled
    form_id *int32
}
// NewFillPostRequestBody instantiates a new fillPostRequestBody and sets the default values.
func NewFillPostRequestBody()(*FillPostRequestBody) {
    m := &FillPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFillPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFillPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFillPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FillPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FillPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["form_filler_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormFillerEmail(val)
        }
        return nil
    }
    res["form_filler_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormFillerMessage(val)
        }
        return nil
    }
    res["form_filler_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormFillerName(val)
        }
        return nil
    }
    res["form_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormId(val)
        }
        return nil
    }
    return res
}
// GetFormFillerEmail gets the form_filler_email property value. The email of the user that is filling the form
func (m *FillPostRequestBody) GetFormFillerEmail()(*string) {
    return m.form_filler_email
}
// GetFormFillerMessage gets the form_filler_message property value. The message filled on the form
func (m *FillPostRequestBody) GetFormFillerMessage()(*string) {
    return m.form_filler_message
}
// GetFormFillerName gets the form_filler_name property value. The name of the user that is filling the form
func (m *FillPostRequestBody) GetFormFillerName()(*string) {
    return m.form_filler_name
}
// GetFormId gets the form_id property value. The form id to be filled
func (m *FillPostRequestBody) GetFormId()(*int32) {
    return m.form_id
}
// Serialize serializes information the current object
func (m *FillPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("form_filler_email", m.GetFormFillerEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("form_filler_message", m.GetFormFillerMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("form_filler_name", m.GetFormFillerName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("form_id", m.GetFormId())
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
func (m *FillPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFormFillerEmail sets the form_filler_email property value. The email of the user that is filling the form
func (m *FillPostRequestBody) SetFormFillerEmail(value *string)() {
    m.form_filler_email = value
}
// SetFormFillerMessage sets the form_filler_message property value. The message filled on the form
func (m *FillPostRequestBody) SetFormFillerMessage(value *string)() {
    m.form_filler_message = value
}
// SetFormFillerName sets the form_filler_name property value. The name of the user that is filling the form
func (m *FillPostRequestBody) SetFormFillerName(value *string)() {
    m.form_filler_name = value
}
// SetFormId sets the form_id property value. The form id to be filled
func (m *FillPostRequestBody) SetFormId(value *int32)() {
    m.form_id = value
}
// FillPostRequestBodyable 
type FillPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormFillerEmail()(*string)
    GetFormFillerMessage()(*string)
    GetFormFillerName()(*string)
    GetFormId()(*int32)
    SetFormFillerEmail(value *string)()
    SetFormFillerMessage(value *string)()
    SetFormFillerName(value *string)()
    SetFormId(value *int32)()
}
