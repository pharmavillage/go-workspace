package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileAttachment 
type FileAttachment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content FileAttachment_contentable
    // The ctp property
    ctp *string
}
// NewFileAttachment instantiates a new FileAttachment and sets the default values.
func NewFileAttachment()(*FileAttachment) {
    m := &FileAttachment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileAttachment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileAttachment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
func (m *FileAttachment) GetContent()(FileAttachment_contentable) {
    return m.content
}
// GetCtp gets the ctp property value. The ctp property
func (m *FileAttachment) GetCtp()(*string) {
    return m.ctp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileAttachment_contentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(FileAttachment_contentable))
        }
        return nil
    }
    res["ctp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCtp(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FileAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ctp", m.GetCtp())
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
func (m *FileAttachment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *FileAttachment) SetContent(value FileAttachment_contentable)() {
    m.content = value
}
// SetCtp sets the ctp property value. The ctp property
func (m *FileAttachment) SetCtp(value *string)() {
    m.ctp = value
}
// FileAttachmentable 
type FileAttachmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(FileAttachment_contentable)
    GetCtp()(*string)
    SetContent(value FileAttachment_contentable)()
    SetCtp(value *string)()
}
