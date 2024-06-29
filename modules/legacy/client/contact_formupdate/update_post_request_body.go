package contact_formupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePostRequestBody 
type UpdatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Color used to render the form. Can be on css hex formats: FFF, #FFF, FFFFFF, #FFFFFF. It's case insensitive.If you pass the "0" value to this field, the current color is removed, and the field gets back to null.
    color *string
    // The new confirmation message.
    confirmation_message *string
    // The id of the channel to be used as a template when creating new channels. It's optional (the valueis kept if it's not provided). To remove an existing template, it must be set to -1.
    copy_from_channel_id *int32
    // Defines if the generated HTML code should open on a modal window, or be directly embeded to the code.
    enable_overlay *bool
    // The form id to be updated
    form_id *int32
    // The new form title.
    form_title *string
}
// NewUpdatePostRequestBody instantiates a new updatePostRequestBody and sets the default values.
func NewUpdatePostRequestBody()(*UpdatePostRequestBody) {
    m := &UpdatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. Color used to render the form. Can be on css hex formats: FFF, #FFF, FFFFFF, #FFFFFF. It's case insensitive.If you pass the "0" value to this field, the current color is removed, and the field gets back to null.
func (m *UpdatePostRequestBody) GetColor()(*string) {
    return m.color
}
// GetConfirmationMessage gets the confirmation_message property value. The new confirmation message.
func (m *UpdatePostRequestBody) GetConfirmationMessage()(*string) {
    return m.confirmation_message
}
// GetCopyFromChannelId gets the copy_from_channel_id property value. The id of the channel to be used as a template when creating new channels. It's optional (the valueis kept if it's not provided). To remove an existing template, it must be set to -1.
func (m *UpdatePostRequestBody) GetCopyFromChannelId()(*int32) {
    return m.copy_from_channel_id
}
// GetEnableOverlay gets the enable_overlay property value. Defines if the generated HTML code should open on a modal window, or be directly embeded to the code.
func (m *UpdatePostRequestBody) GetEnableOverlay()(*bool) {
    return m.enable_overlay
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["form_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormId(val)
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
// GetFormId gets the form_id property value. The form id to be updated
func (m *UpdatePostRequestBody) GetFormId()(*int32) {
    return m.form_id
}
// GetFormTitle gets the form_title property value. The new form title.
func (m *UpdatePostRequestBody) GetFormTitle()(*string) {
    return m.form_title
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("form_id", m.GetFormId())
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
func (m *UpdatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. Color used to render the form. Can be on css hex formats: FFF, #FFF, FFFFFF, #FFFFFF. It's case insensitive.If you pass the "0" value to this field, the current color is removed, and the field gets back to null.
func (m *UpdatePostRequestBody) SetColor(value *string)() {
    m.color = value
}
// SetConfirmationMessage sets the confirmation_message property value. The new confirmation message.
func (m *UpdatePostRequestBody) SetConfirmationMessage(value *string)() {
    m.confirmation_message = value
}
// SetCopyFromChannelId sets the copy_from_channel_id property value. The id of the channel to be used as a template when creating new channels. It's optional (the valueis kept if it's not provided). To remove an existing template, it must be set to -1.
func (m *UpdatePostRequestBody) SetCopyFromChannelId(value *int32)() {
    m.copy_from_channel_id = value
}
// SetEnableOverlay sets the enable_overlay property value. Defines if the generated HTML code should open on a modal window, or be directly embeded to the code.
func (m *UpdatePostRequestBody) SetEnableOverlay(value *bool)() {
    m.enable_overlay = value
}
// SetFormId sets the form_id property value. The form id to be updated
func (m *UpdatePostRequestBody) SetFormId(value *int32)() {
    m.form_id = value
}
// SetFormTitle sets the form_title property value. The new form title.
func (m *UpdatePostRequestBody) SetFormTitle(value *string)() {
    m.form_title = value
}
// UpdatePostRequestBodyable 
type UpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetConfirmationMessage()(*string)
    GetCopyFromChannelId()(*int32)
    GetEnableOverlay()(*bool)
    GetFormId()(*int32)
    GetFormTitle()(*string)
    SetColor(value *string)()
    SetConfirmationMessage(value *string)()
    SetCopyFromChannelId(value *int32)()
    SetEnableOverlay(value *bool)()
    SetFormId(value *int32)()
    SetFormTitle(value *string)()
}
