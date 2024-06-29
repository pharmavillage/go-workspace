package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageSearchResult the search result for a message
type MessageSearchResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The channel_id property
    channel_id *int32
    // The channel_name property
    channel_name *string
    // The content property
    content *string
    // The message_id property
    message_id *int32
    // The user_email property
    user_email *string
    // The user_id property
    user_id *int32
    // The user_name property
    user_name *string
}
// NewMessageSearchResult instantiates a new MessageSearchResult and sets the default values.
func NewMessageSearchResult()(*MessageSearchResult) {
    m := &MessageSearchResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessageSearchResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageSearchResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageSearchResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageSearchResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. The channel_id property
func (m *MessageSearchResult) GetChannelId()(*int32) {
    return m.channel_id
}
// GetChannelName gets the channel_name property value. The channel_name property
func (m *MessageSearchResult) GetChannelName()(*string) {
    return m.channel_name
}
// GetContent gets the content property value. The content property
func (m *MessageSearchResult) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageSearchResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channel_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    res["channel_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelName(val)
        }
        return nil
    }
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
    res["user_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
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
    res["user_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the message_id property value. The message_id property
func (m *MessageSearchResult) GetMessageId()(*int32) {
    return m.message_id
}
// GetUserEmail gets the user_email property value. The user_email property
func (m *MessageSearchResult) GetUserEmail()(*string) {
    return m.user_email
}
// GetUserId gets the user_id property value. The user_id property
func (m *MessageSearchResult) GetUserId()(*int32) {
    return m.user_id
}
// GetUserName gets the user_name property value. The user_name property
func (m *MessageSearchResult) GetUserName()(*string) {
    return m.user_name
}
// Serialize serializes information the current object
func (m *MessageSearchResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("channel_name", m.GetChannelName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content", m.GetContent())
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
        err := writer.WriteStringValue("user_email", m.GetUserEmail())
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
        err := writer.WriteStringValue("user_name", m.GetUserName())
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
func (m *MessageSearchResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. The channel_id property
func (m *MessageSearchResult) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetChannelName sets the channel_name property value. The channel_name property
func (m *MessageSearchResult) SetChannelName(value *string)() {
    m.channel_name = value
}
// SetContent sets the content property value. The content property
func (m *MessageSearchResult) SetContent(value *string)() {
    m.content = value
}
// SetMessageId sets the message_id property value. The message_id property
func (m *MessageSearchResult) SetMessageId(value *int32)() {
    m.message_id = value
}
// SetUserEmail sets the user_email property value. The user_email property
func (m *MessageSearchResult) SetUserEmail(value *string)() {
    m.user_email = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *MessageSearchResult) SetUserId(value *int32)() {
    m.user_id = value
}
// SetUserName sets the user_name property value. The user_name property
func (m *MessageSearchResult) SetUserName(value *string)() {
    m.user_name = value
}
// MessageSearchResultable 
type MessageSearchResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    GetChannelName()(*string)
    GetContent()(*string)
    GetMessageId()(*int32)
    GetUserEmail()(*string)
    GetUserId()(*int32)
    GetUserName()(*string)
    SetChannelId(value *int32)()
    SetChannelName(value *string)()
    SetContent(value *string)()
    SetMessageId(value *int32)()
    SetUserEmail(value *string)()
    SetUserId(value *int32)()
    SetUserName(value *string)()
}
