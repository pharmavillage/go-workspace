package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Meta 
type Meta struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *int32
    // The message property
    message *string
    // The ok property
    ok *bool
}
// NewMeta instantiates a new Meta and sets the default values.
func NewMeta()(*Meta) {
    m := &Meta{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMetaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMetaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeta(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Meta) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *Meta) GetCode()(*int32) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Meta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["ok"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOk(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *Meta) GetMessage()(*string) {
    return m.message
}
// GetOk gets the ok property value. The ok property
func (m *Meta) GetOk()(*bool) {
    return m.ok
}
// Serialize serializes information the current object
func (m *Meta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ok", m.GetOk())
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
func (m *Meta) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *Meta) SetCode(value *int32)() {
    m.code = value
}
// SetMessage sets the message property value. The message property
func (m *Meta) SetMessage(value *string)() {
    m.message = value
}
// SetOk sets the ok property value. The ok property
func (m *Meta) SetOk(value *bool)() {
    m.ok = value
}
// Metaable 
type Metaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*int32)
    GetMessage()(*string)
    GetOk()(*bool)
    SetCode(value *int32)()
    SetMessage(value *string)()
    SetOk(value *bool)()
}
