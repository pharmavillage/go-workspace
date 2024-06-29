package chatreactmessage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReactmessagePostRequestBody 
type ReactmessagePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The emoji_value property
    emoji_value *string
    // The message_id property
    message_id *int32
    // The remove property
    remove *bool
}
// NewReactmessagePostRequestBody instantiates a new reactmessagePostRequestBody and sets the default values.
func NewReactmessagePostRequestBody()(*ReactmessagePostRequestBody) {
    m := &ReactmessagePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReactmessagePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReactmessagePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReactmessagePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReactmessagePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmojiValue gets the emoji_value property value. The emoji_value property
func (m *ReactmessagePostRequestBody) GetEmojiValue()(*string) {
    return m.emoji_value
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReactmessagePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emoji_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmojiValue(val)
        }
        return nil
    }
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
    res["remove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemove(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the message_id property value. The message_id property
func (m *ReactmessagePostRequestBody) GetMessageId()(*int32) {
    return m.message_id
}
// GetRemove gets the remove property value. The remove property
func (m *ReactmessagePostRequestBody) GetRemove()(*bool) {
    return m.remove
}
// Serialize serializes information the current object
func (m *ReactmessagePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emoji_value", m.GetEmojiValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("message_id", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("remove", m.GetRemove())
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
func (m *ReactmessagePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmojiValue sets the emoji_value property value. The emoji_value property
func (m *ReactmessagePostRequestBody) SetEmojiValue(value *string)() {
    m.emoji_value = value
}
// SetMessageId sets the message_id property value. The message_id property
func (m *ReactmessagePostRequestBody) SetMessageId(value *int32)() {
    m.message_id = value
}
// SetRemove sets the remove property value. The remove property
func (m *ReactmessagePostRequestBody) SetRemove(value *bool)() {
    m.remove = value
}
// ReactmessagePostRequestBodyable 
type ReactmessagePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmojiValue()(*string)
    GetMessageId()(*int32)
    GetRemove()(*bool)
    SetEmojiValue(value *string)()
    SetMessageId(value *int32)()
    SetRemove(value *bool)()
}
