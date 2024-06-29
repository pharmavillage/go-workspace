package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Message_read_users 
type Message_read_users struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display_name property
    display_name *string
    // The read_on property
    read_on *string
    // The user_id property
    user_id *int32
}
// NewMessage_read_users instantiates a new Message_read_users and sets the default values.
func NewMessage_read_users()(*Message_read_users) {
    m := &Message_read_users{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessage_read_usersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessage_read_usersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage_read_users(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Message_read_users) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the display_name property value. The display_name property
func (m *Message_read_users) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message_read_users) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["read_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOn(val)
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
// GetReadOn gets the read_on property value. The read_on property
func (m *Message_read_users) GetReadOn()(*string) {
    return m.read_on
}
// GetUserId gets the user_id property value. The user_id property
func (m *Message_read_users) GetUserId()(*int32) {
    return m.user_id
}
// Serialize serializes information the current object
func (m *Message_read_users) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("read_on", m.GetReadOn())
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
func (m *Message_read_users) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *Message_read_users) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetReadOn sets the read_on property value. The read_on property
func (m *Message_read_users) SetReadOn(value *string)() {
    m.read_on = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *Message_read_users) SetUserId(value *int32)() {
    m.user_id = value
}
// Message_read_usersable 
type Message_read_usersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetReadOn()(*string)
    GetUserId()(*int32)
    SetDisplayName(value *string)()
    SetReadOn(value *string)()
    SetUserId(value *int32)()
}
