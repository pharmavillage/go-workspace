package contact_formcreate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreatePostRequestBody 
type CreatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Color used to render the form. Can be on css hex formats: FFF, #FFF, FFFFFF, #FFFFFF. It's case insensitive.
    color *string
    // The confirmation message. Can include a %CHANNEL_LINK% variable.
    confirmation_message *string
    // The id of the channel to be used as a template when creating new channels
    copy_from_channel_id *int32
    // Defines if the generated HTML code should open on a modal window, or be directly embeded to the code.
    enable_overlay *bool
    // The form title
    form_title *string
}
// NewCreatePostRequestBody instantiates a new createPostRequestBody and sets the default values.
func NewCreatePostRequestBody()(*CreatePostRequestBody) {
    m := &CreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. Color used to render the form. Can be on css hex formats: FFF, #FFF, FFFFFF, #FFFFFF. It's case insensitive.
func (m *CreatePostRequestBody) GetColor()(*string) {
    return m.color
}
// GetConfirmationMessage gets the confirmation_message property value. The confirmation message. Can include a %CHANNEL_LINK% variable.
func (m *CreatePostRequestBody) GetConfirmationMessage()(*string) {
    return m.confirmation_message
}
// GetCopyFromChannelId gets the copy_from_channel_id property value. The id of the channel to be used as a template when creating new channels
func (m *CreatePostRequestBody) GetCopyFromChannelId()(*int32) {
    return m.copy_from_channel_id
}
// GetEnableOverlay gets the enable_overlay property value. Defines if the generated HTML code should open on a modal window, or be directly embeded to the code.
func (m *CreatePostRequestBody) GetEnableOverlay()(*bool) {
    return m.enable_overlay
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["confirmation_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfirmationMessage(val)
        }
        return nil
    }
    res["copy_from_channel_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyFromChannelId(val)
        }
        return nil
    }
    res["enable_overlay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOverlay(val)
        }
        return nil
    }
    res["form_title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormTitle(val)
        }
        return nil
    }
    return res
}
// GetFormTitle gets the form_title property value. The form title
func (m *CreatePostRequestBody) GetFormTitle()(*string) {
    return m.form_title
}
// Serialize serializes information the current object
func (m *CreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("confirmation_message", m.GetConfirmationMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copy_from_channel_id", m.GetCopyFromChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enable_overlay", m.GetEnableOverlay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("form_title", m.GetFormTitle())
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
func (m *CreatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. Color used to render the form. Can be on css hex formats: FFF, #FFF, FFFFFF, #FFFFFF. It's case insensitive.
func (m *CreatePostRequestBody) SetColor(value *string)() {
    m.color = value
}
// SetConfirmationMessage sets the confirmation_message property value. The confirmation message. Can include a %CHANNEL_LINK% variable.
func (m *CreatePostRequestBody) SetConfirmationMessage(value *string)() {
    m.confirmation_message = value
}
// SetCopyFromChannelId sets the copy_from_channel_id property value. The id of the channel to be used as a template when creating new channels
func (m *CreatePostRequestBody) SetCopyFromChannelId(value *int32)() {
    m.copy_from_channel_id = value
}
// SetEnableOverlay sets the enable_overlay property value. Defines if the generated HTML code should open on a modal window, or be directly embeded to the code.
func (m *CreatePostRequestBody) SetEnableOverlay(value *bool)() {
    m.enable_overlay = value
}
// SetFormTitle sets the form_title property value. The form title
func (m *CreatePostRequestBody) SetFormTitle(value *string)() {
    m.form_title = value
}
// CreatePostRequestBodyable 
type CreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetConfirmationMessage()(*string)
    GetCopyFromChannelId()(*int32)
    GetEnableOverlay()(*bool)
    GetFormTitle()(*string)
    SetColor(value *string)()
    SetConfirmationMessage(value *string)()
    SetCopyFromChannelId(value *int32)()
    SetEnableOverlay(value *bool)()
    SetFormTitle(value *string)()
}
