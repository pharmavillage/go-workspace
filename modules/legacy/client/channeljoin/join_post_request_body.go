package channeljoin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JoinPostRequestBody 
type JoinPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the channel
    channel_id *int32
    // The public hash included on the channel public url
    public_hash *string
}
// NewJoinPostRequestBody instantiates a new joinPostRequestBody and sets the default values.
func NewJoinPostRequestBody()(*JoinPostRequestBody) {
    m := &JoinPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJoinPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. Id of the channel
func (m *JoinPostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["public_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicHash(val)
        }
        return nil
    }
    return res
}
// GetPublicHash gets the public_hash property value. The public hash included on the channel public url
func (m *JoinPostRequestBody) GetPublicHash()(*string) {
    return m.public_hash
}
// Serialize serializes information the current object
func (m *JoinPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_hash", m.GetPublicHash())
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
func (m *JoinPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. Id of the channel
func (m *JoinPostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetPublicHash sets the public_hash property value. The public hash included on the channel public url
func (m *JoinPostRequestBody) SetPublicHash(value *string)() {
    m.public_hash = value
}
// JoinPostRequestBodyable 
type JoinPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    GetPublicHash()(*string)
    SetChannelId(value *int32)()
    SetPublicHash(value *string)()
}
