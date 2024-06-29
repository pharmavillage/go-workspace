package channelgroupupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePostRequestBody 
type UpdatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The channel group id that must be updated.
    channel_group_id *int32
    // New channel group name
    name *string
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
// GetChannelGroupId gets the channel_group_id property value. The channel group id that must be updated.
func (m *UpdatePostRequestBody) GetChannelGroupId()(*int32) {
    return m.channel_group_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channel_group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelGroupId(val)
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
    return res
}
// GetName gets the name property value. New channel group name
func (m *UpdatePostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_group_id", m.GetChannelGroupId())
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
// SetChannelGroupId sets the channel_group_id property value. The channel group id that must be updated.
func (m *UpdatePostRequestBody) SetChannelGroupId(value *int32)() {
    m.channel_group_id = value
}
// SetName sets the name property value. New channel group name
func (m *UpdatePostRequestBody) SetName(value *string)() {
    m.name = value
}
// UpdatePostRequestBodyable 
type UpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelGroupId()(*int32)
    GetName()(*string)
    SetChannelGroupId(value *int32)()
    SetName(value *string)()
}
