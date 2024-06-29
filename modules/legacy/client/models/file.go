package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// File 
type File struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The creation property
    creation *string
    // The ext property
    ext *string
    // The modification property
    modification *string
    // The name property
    name *string
    // The type property
    typeEscaped *string
}
// NewFile instantiates a new File and sets the default values.
func NewFile()(*File) {
    m := &File{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *File) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreation gets the creation property value. The creation property
func (m *File) GetCreation()(*string) {
    return m.creation
}
// GetExt gets the ext property value. The ext property
func (m *File) GetExt()(*string) {
    return m.ext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *File) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["creation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreation(val)
        }
        return nil
    }
    res["ext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExt(val)
        }
        return nil
    }
    res["modification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModification(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetModification gets the modification property value. The modification property
func (m *File) GetModification()(*string) {
    return m.modification
}
// GetName gets the name property value. The name property
func (m *File) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. The type property
func (m *File) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *File) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("creation", m.GetCreation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ext", m.GetExt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("modification", m.GetModification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *File) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreation sets the creation property value. The creation property
func (m *File) SetCreation(value *string)() {
    m.creation = value
}
// SetExt sets the ext property value. The ext property
func (m *File) SetExt(value *string)() {
    m.ext = value
}
// SetModification sets the modification property value. The modification property
func (m *File) SetModification(value *string)() {
    m.modification = value
}
// SetName sets the name property value. The name property
func (m *File) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *File) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// Fileable 
type Fileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreation()(*string)
    GetExt()(*string)
    GetModification()(*string)
    GetName()(*string)
    GetTypeEscaped()(*string)
    SetCreation(value *string)()
    SetExt(value *string)()
    SetModification(value *string)()
    SetName(value *string)()
    SetTypeEscaped(value *string)()
}
