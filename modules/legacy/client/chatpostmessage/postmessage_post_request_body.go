package chatpostmessage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostmessagePostRequestBody 
type PostmessagePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional - Array of JSON encoded strings for storing file attachment Each JSON record will result in an attachment To send a file attachment, First upload file(s) and send that as an array in attachments" field.
    attachments *string
    // The channel_id property
    channel_id *int32
    // Optional - For quoting another message.
    quote_message_id *string
    // Optional - Send a email message to all users of the channel in addition to posting in the channel.
    send_email *bool
    // The text property
    text *string
}
// NewPostmessagePostRequestBody instantiates a new postmessagePostRequestBody and sets the default values.
func NewPostmessagePostRequestBody()(*PostmessagePostRequestBody) {
    m := &PostmessagePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePostmessagePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePostmessagePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostmessagePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostmessagePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachments gets the attachments property value. Optional - Array of JSON encoded strings for storing file attachment Each JSON record will result in an attachment To send a file attachment, First upload file(s) and send that as an array in attachments" field.
func (m *PostmessagePostRequestBody) GetAttachments()(*string) {
    return m.attachments
}
// GetChannelId gets the channel_id property value. The channel_id property
func (m *PostmessagePostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PostmessagePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachments(val)
        }
        return nil
    }
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
    res["quote_message_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuoteMessageId(val)
        }
        return nil
    }
    res["send_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendEmail(val)
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
// GetQuoteMessageId gets the quote_message_id property value. Optional - For quoting another message.
func (m *PostmessagePostRequestBody) GetQuoteMessageId()(*string) {
    return m.quote_message_id
}
// GetSendEmail gets the send_email property value. Optional - Send a email message to all users of the channel in addition to posting in the channel.
func (m *PostmessagePostRequestBody) GetSendEmail()(*bool) {
    return m.send_email
}
// GetText gets the text property value. The text property
func (m *PostmessagePostRequestBody) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *PostmessagePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("attachments", m.GetAttachments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("quote_message_id", m.GetQuoteMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("send_email", m.GetSendEmail())
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
func (m *PostmessagePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachments sets the attachments property value. Optional - Array of JSON encoded strings for storing file attachment Each JSON record will result in an attachment To send a file attachment, First upload file(s) and send that as an array in attachments" field.
func (m *PostmessagePostRequestBody) SetAttachments(value *string)() {
    m.attachments = value
}
// SetChannelId sets the channel_id property value. The channel_id property
func (m *PostmessagePostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetQuoteMessageId sets the quote_message_id property value. Optional - For quoting another message.
func (m *PostmessagePostRequestBody) SetQuoteMessageId(value *string)() {
    m.quote_message_id = value
}
// SetSendEmail sets the send_email property value. Optional - Send a email message to all users of the channel in addition to posting in the channel.
func (m *PostmessagePostRequestBody) SetSendEmail(value *bool)() {
    m.send_email = value
}
// SetText sets the text property value. The text property
func (m *PostmessagePostRequestBody) SetText(value *string)() {
    m.text = value
}
// PostmessagePostRequestBodyable 
type PostmessagePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()(*string)
    GetChannelId()(*int32)
    GetQuoteMessageId()(*string)
    GetSendEmail()(*bool)
    GetText()(*string)
    SetAttachments(value *string)()
    SetChannelId(value *int32)()
    SetQuoteMessageId(value *string)()
    SetSendEmail(value *bool)()
    SetText(value *string)()
}
