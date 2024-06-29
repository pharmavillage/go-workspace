package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Message_parent_message 
type Message_parent_message struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content *string
    // The created_on property
    created_on *string
    // The created_on_ts property
    created_on_ts *int32
    // The display_name property
    display_name *string
    // The message_id property
    message_id *int32
    // The user_id property
    user_id *int32
}
// NewMessage_parent_message instantiates a new Message_parent_message and sets the default values.
func NewMessage_parent_message()(*Message_parent_message) {
    m := &Message_parent_message{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessage_parent_messageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessage_parent_messageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage_parent_message(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Message_parent_message) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
func (m *Message_parent_message) GetContent()(*string) {
    return m.content
}
// GetCreatedOn gets the created_on property value. The created_on property
func (m *Message_parent_message) GetCreatedOn()(*string) {
    return m.created_on
}
// GetCreatedOnTs gets the created_on_ts property value. The created_on_ts property
func (m *Message_parent_message) GetCreatedOnTs()(*int32) {
    return m.created_on_ts
}
// GetDisplayName gets the display_name property value. The display_name property
func (m *Message_parent_message) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message_parent_message) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["created_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOn(val)
        }
        return nil
    }
    res["created_on_ts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnTs(val)
        }
        return nil
    }
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["user_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the message_id property value. The message_id property
func (m *Message_parent_message) GetMessageId()(*int32) {
    return m.message_id
}
// GetUserId gets the user_id property value. The user_id property
func (m *Message_parent_message) GetUserId()(*int32) {
    return m.user_id
}
// Serialize serializes information the current object
func (m *Message_parent_message) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_on", m.GetCreatedOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("created_on_ts", m.GetCreatedOnTs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
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
        err := writer.WriteInt32Value("user_id", m.GetUserId())
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
func (m *Message_parent_message) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *Message_parent_message) SetContent(value *string)() {
    m.content = value
}
// SetCreatedOn sets the created_on property value. The created_on property
func (m *Message_parent_message) SetCreatedOn(value *string)() {
    m.created_on = value
}
// SetCreatedOnTs sets the created_on_ts property value. The created_on_ts property
func (m *Message_parent_message) SetCreatedOnTs(value *int32)() {
    m.created_on_ts = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *Message_parent_message) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetMessageId sets the message_id property value. The message_id property
func (m *Message_parent_message) SetMessageId(value *int32)() {
    m.message_id = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *Message_parent_message) SetUserId(value *int32)() {
    m.user_id = value
}
// Message_parent_messageable 
type Message_parent_messageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetCreatedOn()(*string)
    GetCreatedOnTs()(*int32)
    GetDisplayName()(*string)
    GetMessageId()(*int32)
    GetUserId()(*int32)
    SetContent(value *string)()
    SetCreatedOn(value *string)()
    SetCreatedOnTs(value *int32)()
    SetDisplayName(value *string)()
    SetMessageId(value *int32)()
    SetUserId(value *int32)()
}
