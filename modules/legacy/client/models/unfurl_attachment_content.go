package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnfurlAttachment_content 
type UnfurlAttachment_content struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The favicon property
    favicon *string
    // The images property
    images []string
    // The title property
    title *string
    // The url property
    url *string
}
// NewUnfurlAttachment_content instantiates a new UnfurlAttachment_content and sets the default values.
func NewUnfurlAttachment_content()(*UnfurlAttachment_content) {
    m := &UnfurlAttachment_content{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUnfurlAttachment_contentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnfurlAttachment_contentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnfurlAttachment_content(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnfurlAttachment_content) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
func (m *UnfurlAttachment_content) GetDescription()(*string) {
    return m.description
}
// GetFavicon gets the favicon property value. The favicon property
func (m *UnfurlAttachment_content) GetFavicon()(*string) {
    return m.favicon
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnfurlAttachment_content) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["favicon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFavicon(val)
        }
        return nil
    }
    res["images"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetImages(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetImages gets the images property value. The images property
func (m *UnfurlAttachment_content) GetImages()([]string) {
    return m.images
}
// GetTitle gets the title property value. The title property
func (m *UnfurlAttachment_content) GetTitle()(*string) {
    return m.title
}
// GetUrl gets the url property value. The url property
func (m *UnfurlAttachment_content) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *UnfurlAttachment_content) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("favicon", m.GetFavicon())
        if err != nil {
            return err
        }
    }
    if m.GetImages() != nil {
        err := writer.WriteCollectionOfStringValues("images", m.GetImages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *UnfurlAttachment_content) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *UnfurlAttachment_content) SetDescription(value *string)() {
    m.description = value
}
// SetFavicon sets the favicon property value. The favicon property
func (m *UnfurlAttachment_content) SetFavicon(value *string)() {
    m.favicon = value
}
// SetImages sets the images property value. The images property
func (m *UnfurlAttachment_content) SetImages(value []string)() {
    m.images = value
}
// SetTitle sets the title property value. The title property
func (m *UnfurlAttachment_content) SetTitle(value *string)() {
    m.title = value
}
// SetUrl sets the url property value. The url property
func (m *UnfurlAttachment_content) SetUrl(value *string)() {
    m.url = value
}
// UnfurlAttachment_contentable 
type UnfurlAttachment_contentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetFavicon()(*string)
    GetImages()([]string)
    GetTitle()(*string)
    GetUrl()(*string)
    SetDescription(value *string)()
    SetFavicon(value *string)()
    SetImages(value []string)()
    SetTitle(value *string)()
    SetUrl(value *string)()
}
