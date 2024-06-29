package channelinvite

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InvitePostRequestBody 
type InvitePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id of the channel to add
    channel_id *int32
    // Role (100 - Owner, 50 - Manager, 20 - Collaborator, 10 - View) for the users to add
    user_role *int32
    // Comma separated email or phone list to add in addition to the user requesting this API.test1@test.com,+19023762345,test2@gmail.com
    users *string
}
// NewInvitePostRequestBody instantiates a new invitePostRequestBody and sets the default values.
func NewInvitePostRequestBody()(*InvitePostRequestBody) {
    m := &InvitePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInvitePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvitePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvitePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InvitePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. The id of the channel to add
func (m *InvitePostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvitePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsers(val)
        }
        return nil
    }
    return res
}
// GetUserRole gets the user_role property value. Role (100 - Owner, 50 - Manager, 20 - Collaborator, 10 - View) for the users to add
func (m *InvitePostRequestBody) GetUserRole()(*int32) {
    return m.user_role
}
// GetUsers gets the users property value. Comma separated email or phone list to add in addition to the user requesting this API.test1@test.com,+19023762345,test2@gmail.com
func (m *InvitePostRequestBody) GetUsers()(*string) {
    return m.users
}
// Serialize serializes information the current object
func (m *InvitePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("users", m.GetUsers())
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
func (m *InvitePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. The id of the channel to add
func (m *InvitePostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetUserRole sets the user_role property value. Role (100 - Owner, 50 - Manager, 20 - Collaborator, 10 - View) for the users to add
func (m *InvitePostRequestBody) SetUserRole(value *int32)() {
    m.user_role = value
}
// SetUsers sets the users property value. Comma separated email or phone list to add in addition to the user requesting this API.test1@test.com,+19023762345,test2@gmail.com
func (m *InvitePostRequestBody) SetUsers(value *string)() {
    m.users = value
}
// InvitePostRequestBodyable 
type InvitePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    GetUserRole()(*int32)
    GetUsers()(*string)
    SetChannelId(value *int32)()
    SetUserRole(value *int32)()
    SetUsers(value *string)()
}
