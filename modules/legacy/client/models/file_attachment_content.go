package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileAttachment_content 
type FileAttachment_content struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The file property
    file *string
    // The path property
    path *string
    // The size property
    size *int32
}
// NewFileAttachment_content instantiates a new FileAttachment_content and sets the default values.
func NewFileAttachment_content()(*FileAttachment_content) {
    m := &FileAttachment_content{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileAttachment_contentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileAttachment_contentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileAttachment_content(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileAttachment_content) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileAttachment_content) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. The file property
func (m *FileAttachment_content) GetFile()(*string) {
    return m.file
}
// GetPath gets the path property value. The path property
func (m *FileAttachment_content) GetPath()(*string) {
    return m.path
}
// GetSize gets the size property value. The size property
func (m *FileAttachment_content) GetSize()(*int32) {
    return m.size
}
// Serialize serializes information the current object
func (m *FileAttachment_content) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
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
func (m *FileAttachment_content) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFile sets the file property value. The file property
func (m *FileAttachment_content) SetFile(value *string)() {
    m.file = value
}
// SetPath sets the path property value. The path property
func (m *FileAttachment_content) SetPath(value *string)() {
    m.path = value
}
// SetSize sets the size property value. The size property
func (m *FileAttachment_content) SetSize(value *int32)() {
    m.size = value
}
// FileAttachment_contentable 
type FileAttachment_contentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFile()(*string)
    GetPath()(*string)
    GetSize()(*int32)
    SetFile(value *string)()
    SetPath(value *string)()
    SetSize(value *int32)()
}
