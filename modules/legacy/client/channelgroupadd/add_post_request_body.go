package channelgroupadd

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddPostRequestBody 
type AddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id of the group where the channel should be inserted
    channel_group_id *int32
    // The channel to be included on the group.
    channel_id *int32
}
// NewAddPostRequestBody instantiates a new addPostRequestBody and sets the default values.
func NewAddPostRequestBody()(*AddPostRequestBody) {
    m := &AddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelGroupId gets the channel_group_id property value. The id of the group where the channel should be inserted
func (m *AddPostRequestBody) GetChannelGroupId()(*int32) {
    return m.channel_group_id
}
// GetChannelId gets the channel_id property value. The channel to be included on the group.
func (m *AddPostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *AddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_group_id", m.GetChannelGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
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
func (m *AddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelGroupId sets the channel_group_id property value. The id of the group where the channel should be inserted
func (m *AddPostRequestBody) SetChannelGroupId(value *int32)() {
    m.channel_group_id = value
}
// SetChannelId sets the channel_id property value. The channel to be included on the group.
func (m *AddPostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// AddPostRequestBodyable 
type AddPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelGroupId()(*int32)
    GetChannelId()(*int32)
    SetChannelGroupId(value *int32)()
    SetChannelId(value *int32)()
}
