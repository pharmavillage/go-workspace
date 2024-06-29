package chatcommand

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommandPostRequestBody 
type CommandPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The channel_id property
    channel_id *int32
    // The text property
    text *string
}
// NewCommandPostRequestBody instantiates a new commandPostRequestBody and sets the default values.
func NewCommandPostRequestBody()(*CommandPostRequestBody) {
    m := &CommandPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommandPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommandPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommandPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommandPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. The channel_id property
func (m *CommandPostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommandPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channel_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetText gets the text property value. The text property
func (m *CommandPostRequestBody) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *CommandPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *CommandPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. The channel_id property
func (m *CommandPostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetText sets the text property value. The text property
func (m *CommandPostRequestBody) SetText(value *string)() {
    m.text = value
}
// CommandPostRequestBodyable 
type CommandPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    GetText()(*string)
    SetChannelId(value *int32)()
    SetText(value *string)()
}
