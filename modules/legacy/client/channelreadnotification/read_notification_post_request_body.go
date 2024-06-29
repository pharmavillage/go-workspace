package channelreadnotification

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReadNotificationPostRequestBody 
type ReadNotificationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the channel to change the subscription level.
    channel_id *int32
    // Id of the message. If not supplied, all messages will be marked as read
    message_id *int32
}
// NewReadNotificationPostRequestBody instantiates a new readNotificationPostRequestBody and sets the default values.
func NewReadNotificationPostRequestBody()(*ReadNotificationPostRequestBody) {
    m := &ReadNotificationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReadNotificationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReadNotificationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReadNotificationPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReadNotificationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. Id of the channel to change the subscription level.
func (m *ReadNotificationPostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReadNotificationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["message_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the message_id property value. Id of the message. If not supplied, all messages will be marked as read
func (m *ReadNotificationPostRequestBody) GetMessageId()(*int32) {
    return m.message_id
}
// Serialize serializes information the current object
func (m *ReadNotificationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("message_id", m.GetMessageId())
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
func (m *ReadNotificationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. Id of the channel to change the subscription level.
func (m *ReadNotificationPostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetMessageId sets the message_id property value. Id of the message. If not supplied, all messages will be marked as read
func (m *ReadNotificationPostRequestBody) SetMessageId(value *int32)() {
    m.message_id = value
}
// ReadNotificationPostRequestBodyable 
type ReadNotificationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    GetMessageId()(*int32)
    SetChannelId(value *int32)()
    SetMessageId(value *int32)()
}
