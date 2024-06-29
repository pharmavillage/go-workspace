package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Message_emoticons 
type Message_emoticons struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dn property
    dn *string
    // The ev property
    ev *string
    // The uid property
    uid *int32
}
// NewMessage_emoticons instantiates a new Message_emoticons and sets the default values.
func NewMessage_emoticons()(*Message_emoticons) {
    m := &Message_emoticons{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessage_emoticonsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessage_emoticonsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage_emoticons(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Message_emoticons) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDn gets the dn property value. The dn property
func (m *Message_emoticons) GetDn()(*string) {
    return m.dn
}
// GetEv gets the ev property value. The ev property
func (m *Message_emoticons) GetEv()(*string) {
    return m.ev
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message_emoticons) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDn(val)
        }
        return nil
    }
    res["ev"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEv(val)
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    return res
}
// GetUid gets the uid property value. The uid property
func (m *Message_emoticons) GetUid()(*int32) {
    return m.uid
}
// Serialize serializes information the current object
func (m *Message_emoticons) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dn", m.GetDn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ev", m.GetEv())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("uid", m.GetUid())
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
func (m *Message_emoticons) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDn sets the dn property value. The dn property
func (m *Message_emoticons) SetDn(value *string)() {
    m.dn = value
}
// SetEv sets the ev property value. The ev property
func (m *Message_emoticons) SetEv(value *string)() {
    m.ev = value
}
// SetUid sets the uid property value. The uid property
func (m *Message_emoticons) SetUid(value *int32)() {
    m.uid = value
}
// Message_emoticonsable 
type Message_emoticonsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDn()(*string)
    GetEv()(*string)
    GetUid()(*int32)
    SetDn(value *string)()
    SetEv(value *string)()
    SetUid(value *int32)()
}
