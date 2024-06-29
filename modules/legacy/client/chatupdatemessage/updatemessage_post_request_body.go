package chatupdatemessage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatemessagePostRequestBody 
type UpdatemessagePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The message_id property
    message_id *int32
    // The text property
    text *string
}
// NewUpdatemessagePostRequestBody instantiates a new updatemessagePostRequestBody and sets the default values.
func NewUpdatemessagePostRequestBody()(*UpdatemessagePostRequestBody) {
    m := &UpdatemessagePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdatemessagePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatemessagePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatemessagePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatemessagePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatemessagePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the message_id property value. The message_id property
func (m *UpdatemessagePostRequestBody) GetMessageId()(*int32) {
    return m.message_id
}
// GetText gets the text property value. The text property
func (m *UpdatemessagePostRequestBody) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *UpdatemessagePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("message_id", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *UpdatemessagePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessageId sets the message_id property value. The message_id property
func (m *UpdatemessagePostRequestBody) SetMessageId(value *int32)() {
    m.message_id = value
}
// SetText sets the text property value. The text property
func (m *UpdatemessagePostRequestBody) SetText(value *string)() {
    m.text = value
}
// UpdatemessagePostRequestBodyable 
type UpdatemessagePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessageId()(*int32)
    GetText()(*string)
    SetMessageId(value *int32)()
    SetText(value *string)()
}
