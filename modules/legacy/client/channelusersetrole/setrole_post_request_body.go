package channelusersetrole

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetrolePostRequestBody 
type SetrolePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the channel
    channel_id *int32
    // Id of the user to set role
    user_id *int32
    // 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
    user_role *int32
}
// NewSetrolePostRequestBody instantiates a new setrolePostRequestBody and sets the default values.
func NewSetrolePostRequestBody()(*SetrolePostRequestBody) {
    m := &SetrolePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSetrolePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetrolePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetrolePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetrolePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. Id of the channel
func (m *SetrolePostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetrolePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["user_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["user_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRole(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the user_id property value. Id of the user to set role
func (m *SetrolePostRequestBody) GetUserId()(*int32) {
    return m.user_id
}
// GetUserRole gets the user_role property value. 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *SetrolePostRequestBody) GetUserRole()(*int32) {
    return m.user_role
}
// Serialize serializes information the current object
func (m *SetrolePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_id", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_role", m.GetUserRole())
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
func (m *SetrolePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. Id of the channel
func (m *SetrolePostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetUserId sets the user_id property value. Id of the user to set role
func (m *SetrolePostRequestBody) SetUserId(value *int32)() {
    m.user_id = value
}
// SetUserRole sets the user_role property value. 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *SetrolePostRequestBody) SetUserRole(value *int32)() {
    m.user_role = value
}
// SetrolePostRequestBodyable 
type SetrolePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    GetUserId()(*int32)
    GetUserRole()(*int32)
    SetChannelId(value *int32)()
    SetUserId(value *int32)()
    SetUserRole(value *int32)()
}
