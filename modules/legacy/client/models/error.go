package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Error 
type Error struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *int32
    // The error property
    error *string
    // The ok property
    ok *bool
}
// NewError instantiates a new Error and sets the default values.
func NewError()(*Error) {
    m := &Error{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewError(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *Error) GetCode()(*int32) {
    return m.code
}
// GetError gets the error property value. The error property
func (m *Error) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
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
// GetOk gets the ok property value. The ok property
func (m *Error) GetOk()(*bool) {
    return m.ok
}
// Serialize serializes information the current object
func (m *Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error", m.GetError())
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
func (m *Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *Error) SetCode(value *int32)() {
    m.code = value
}
// SetError sets the error property value. The error property
func (m *Error) SetError(value *string)() {
    m.error = value
}
// SetOk sets the ok property value. The ok property
func (m *Error) SetOk(value *bool)() {
    m.ok = value
}
// Errorable 
type Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*int32)
    GetError()(*string)
    GetOk()(*bool)
    SetCode(value *int32)()
    SetError(value *string)()
    SetOk(value *bool)()
}
