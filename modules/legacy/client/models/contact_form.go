package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContactForm 
type ContactForm struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *string
    // The confirmation_message property
    confirmation_message *string
    // The copy_from_channel_id property
    copy_from_channel_id *int32
    // The enable_overlay property
    enable_overlay *bool
    // The enabled property
    enabled *bool
    // The form_hash property
    form_hash *string
    // The form_title property
    form_title *string
    // The id property
    id *int32
    // The owner_id property
    owner_id *int32
}
// NewContactForm instantiates a new ContactForm and sets the default values.
func NewContactForm()(*ContactForm) {
    m := &ContactForm{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContactFormFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContactFormFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContactForm(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContactForm) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *ContactForm) GetColor()(*string) {
    return m.color
}
// GetConfirmationMessage gets the confirmation_message property value. The confirmation_message property
func (m *ContactForm) GetConfirmationMessage()(*string) {
    return m.confirmation_message
}
// GetCopyFromChannelId gets the copy_from_channel_id property value. The copy_from_channel_id property
func (m *ContactForm) GetCopyFromChannelId()(*int32) {
    return m.copy_from_channel_id
}
// GetEnabled gets the enabled property value. The enabled property
func (m *ContactForm) GetEnabled()(*bool) {
    return m.enabled
}
// GetEnableOverlay gets the enable_overlay property value. The enable_overlay property
func (m *ContactForm) GetEnableOverlay()(*bool) {
    return m.enable_overlay
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContactForm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["form_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormHash(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["owner_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerId(val)
        }
        return nil
    }
    return res
}
// GetFormHash gets the form_hash property value. The form_hash property
func (m *ContactForm) GetFormHash()(*string) {
    return m.form_hash
}
// GetFormTitle gets the form_title property value. The form_title property
func (m *ContactForm) GetFormTitle()(*string) {
    return m.form_title
}
// GetId gets the id property value. The id property
func (m *ContactForm) GetId()(*int32) {
    return m.id
}
// GetOwnerId gets the owner_id property value. The owner_id property
func (m *ContactForm) GetOwnerId()(*int32) {
    return m.owner_id
}
// Serialize serializes information the current object
func (m *ContactForm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
        err := writer.WriteStringValue("form_hash", m.GetFormHash())
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
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("owner_id", m.GetOwnerId())
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
func (m *ContactForm) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *ContactForm) SetColor(value *string)() {
    m.color = value
}
// SetConfirmationMessage sets the confirmation_message property value. The confirmation_message property
func (m *ContactForm) SetConfirmationMessage(value *string)() {
    m.confirmation_message = value
}
// SetCopyFromChannelId sets the copy_from_channel_id property value. The copy_from_channel_id property
func (m *ContactForm) SetCopyFromChannelId(value *int32)() {
    m.copy_from_channel_id = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *ContactForm) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEnableOverlay sets the enable_overlay property value. The enable_overlay property
func (m *ContactForm) SetEnableOverlay(value *bool)() {
    m.enable_overlay = value
}
// SetFormHash sets the form_hash property value. The form_hash property
func (m *ContactForm) SetFormHash(value *string)() {
    m.form_hash = value
}
// SetFormTitle sets the form_title property value. The form_title property
func (m *ContactForm) SetFormTitle(value *string)() {
    m.form_title = value
}
// SetId sets the id property value. The id property
func (m *ContactForm) SetId(value *int32)() {
    m.id = value
}
// SetOwnerId sets the owner_id property value. The owner_id property
func (m *ContactForm) SetOwnerId(value *int32)() {
    m.owner_id = value
}
// ContactFormable 
type ContactFormable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetConfirmationMessage()(*string)
    GetCopyFromChannelId()(*int32)
    GetEnabled()(*bool)
    GetEnableOverlay()(*bool)
    GetFormHash()(*string)
    GetFormTitle()(*string)
    GetId()(*int32)
    GetOwnerId()(*int32)
    SetColor(value *string)()
    SetConfirmationMessage(value *string)()
    SetCopyFromChannelId(value *int32)()
    SetEnabled(value *bool)()
    SetEnableOverlay(value *bool)()
    SetFormHash(value *string)()
    SetFormTitle(value *string)()
    SetId(value *int32)()
    SetOwnerId(value *int32)()
}
