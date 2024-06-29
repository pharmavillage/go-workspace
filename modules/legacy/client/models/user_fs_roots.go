package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// User_fs_roots 
type User_fs_roots struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The location property
    location *string
    // The type property
    typeEscaped *string
}
// NewUser_fs_roots instantiates a new user_fs_roots and sets the default values.
func NewUser_fs_roots()(*User_fs_roots) {
    m := &User_fs_roots{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUser_fs_rootsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUser_fs_rootsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser_fs_roots(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *User_fs_roots) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *User_fs_roots) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. The location property
func (m *User_fs_roots) GetLocation()(*string) {
    return m.location
}
// GetTypeEscaped gets the type property value. The type property
func (m *User_fs_roots) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *User_fs_roots) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *User_fs_roots) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocation sets the location property value. The location property
func (m *User_fs_roots) SetLocation(value *string)() {
    m.location = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *User_fs_roots) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// User_fs_rootsable 
type User_fs_rootsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(*string)
    GetTypeEscaped()(*string)
    SetLocation(value *string)()
    SetTypeEscaped(value *string)()
}
