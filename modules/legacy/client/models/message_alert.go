package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageAlert 
type MessageAlert struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The alert_id property
    alert_id *int32
    // The alert_text property
    alert_text *string
    // The alert_type property
    alert_type *int32
    // The channel_id property
    channel_id *int32
    // The from property
    from []User_Abbrable
    // The is_read property
    is_read *bool
    // The message_id property
    message_id *int32
}
// NewMessageAlert instantiates a new MessageAlert and sets the default values.
func NewMessageAlert()(*MessageAlert) {
    m := &MessageAlert{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessageAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageAlert(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageAlert) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlertId gets the alert_id property value. The alert_id property
func (m *MessageAlert) GetAlertId()(*int32) {
    return m.alert_id
}
// GetAlertText gets the alert_text property value. The alert_text property
func (m *MessageAlert) GetAlertText()(*string) {
    return m.alert_text
}
// GetAlertType gets the alert_type property value. The alert_type property
func (m *MessageAlert) GetAlertType()(*int32) {
    return m.alert_type
}
// GetChannelId gets the channel_id property value. The channel_id property
func (m *MessageAlert) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alert_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertId(val)
        }
        return nil
    }
    res["alert_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertText(val)
        }
        return nil
    }
    res["alert_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertType(val)
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
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUser_AbbrFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]User_Abbrable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(User_Abbrable)
                }
            }
            m.SetFrom(res)
        }
        return nil
    }
    res["is_read"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRead(val)
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
    return res
}
// GetFrom gets the from property value. The from property
func (m *MessageAlert) GetFrom()([]User_Abbrable) {
    return m.from
}
// GetIsRead gets the is_read property value. The is_read property
func (m *MessageAlert) GetIsRead()(*bool) {
    return m.is_read
}
// GetMessageId gets the message_id property value. The message_id property
func (m *MessageAlert) GetMessageId()(*int32) {
    return m.message_id
}
// Serialize serializes information the current object
func (m *MessageAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("alert_id", m.GetAlertId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alert_text", m.GetAlertText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("alert_type", m.GetAlertType())
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
    if m.GetFrom() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFrom()))
        for i, v := range m.GetFrom() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("from", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_read", m.GetIsRead())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageAlert) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlertId sets the alert_id property value. The alert_id property
func (m *MessageAlert) SetAlertId(value *int32)() {
    m.alert_id = value
}
// SetAlertText sets the alert_text property value. The alert_text property
func (m *MessageAlert) SetAlertText(value *string)() {
    m.alert_text = value
}
// SetAlertType sets the alert_type property value. The alert_type property
func (m *MessageAlert) SetAlertType(value *int32)() {
    m.alert_type = value
}
// SetChannelId sets the channel_id property value. The channel_id property
func (m *MessageAlert) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetFrom sets the from property value. The from property
func (m *MessageAlert) SetFrom(value []User_Abbrable)() {
    m.from = value
}
// SetIsRead sets the is_read property value. The is_read property
func (m *MessageAlert) SetIsRead(value *bool)() {
    m.is_read = value
}
// SetMessageId sets the message_id property value. The message_id property
func (m *MessageAlert) SetMessageId(value *int32)() {
    m.message_id = value
}
// MessageAlertable 
type MessageAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertId()(*int32)
    GetAlertText()(*string)
    GetAlertType()(*int32)
    GetChannelId()(*int32)
    GetFrom()([]User_Abbrable)
    GetIsRead()(*bool)
    GetMessageId()(*int32)
    SetAlertId(value *int32)()
    SetAlertText(value *string)()
    SetAlertType(value *int32)()
    SetChannelId(value *int32)()
    SetFrom(value []User_Abbrable)()
    SetIsRead(value *bool)()
    SetMessageId(value *int32)()
}
