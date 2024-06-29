package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnfurlAttachment 
type UnfurlAttachment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content UnfurlAttachment_contentable
    // The ctp property
    ctp *string
}
// NewUnfurlAttachment instantiates a new UnfurlAttachment and sets the default values.
func NewUnfurlAttachment()(*UnfurlAttachment) {
    m := &UnfurlAttachment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUnfurlAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnfurlAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnfurlAttachment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnfurlAttachment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
func (m *UnfurlAttachment) GetContent()(UnfurlAttachment_contentable) {
    return m.content
}
// GetCtp gets the ctp property value. The ctp property
func (m *UnfurlAttachment) GetCtp()(*string) {
    return m.ctp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnfurlAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnfurlAttachment_contentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(UnfurlAttachment_contentable))
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
func (m *UnfurlAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UnfurlAttachment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *UnfurlAttachment) SetContent(value UnfurlAttachment_contentable)() {
    m.content = value
}
// SetCtp sets the ctp property value. The ctp property
func (m *UnfurlAttachment) SetCtp(value *string)() {
    m.ctp = value
}
// UnfurlAttachmentable 
type UnfurlAttachmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(UnfurlAttachment_contentable)
    GetCtp()(*string)
    SetContent(value UnfurlAttachment_contentable)()
    SetCtp(value *string)()
}
