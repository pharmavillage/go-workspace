package channelcreate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreatePostRequestBody 
type CreatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Only used when `allow_join` is set. This allows anyone that have the channel link to see its contents (without joining the channel).
    allow_external_read *bool
    // Defines if the channel will have access through a link. If set to true, the channel will have an access link. This link will allow users to self join to the channel, and in some cases to see channel contents.
    allow_join *bool
    // Automatically close after these many days of inactivity
    auto_close_days *int32
    // Small blurb that briefly describes the channel content.
    blurb *string
    // New channel name
    channel_name *string
    // Copy files and settings from this channel to the new channe. This channel must be owned by this user.
    copy_from_channel_id *int32
    // Defines if the channel have public access. If set to true, the channel will have a public access link, and anyone with access to this link will be able to see the channel content and join the channel. This property is deprecated. Use allow_join.
    // Deprecated: 
    is_public *bool
    // Defines if the user must wait for a manager to approve his join to the channel. Only used when `allow_join` is set.
    require_join_approval *bool
    // Comma separated email or phone list to add in addition to the user requesting this API.test1@test.com,+18197432105,test2@gmail.com
    users *string
}
// NewCreatePostRequestBody instantiates a new createPostRequestBody and sets the default values.
func NewCreatePostRequestBody()(*CreatePostRequestBody) {
    m := &CreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    allow_external_readValue := "0"
    m.SetAllowExternalRead(&allow_external_readValue)
    allow_joinValue := "0"
    m.SetAllowJoin(&allow_joinValue)
    is_publicValue := "0"
    m.SetIsPublic(&is_publicValue)
    require_join_approvalValue := "1"
    m.SetRequireJoinApproval(&require_join_approvalValue)
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
// GetAllowExternalRead gets the allow_external_read property value. Only used when `allow_join` is set. This allows anyone that have the channel link to see its contents (without joining the channel).
func (m *CreatePostRequestBody) GetAllowExternalRead()(*bool) {
    return m.allow_external_read
}
// GetAllowJoin gets the allow_join property value. Defines if the channel will have access through a link. If set to true, the channel will have an access link. This link will allow users to self join to the channel, and in some cases to see channel contents.
func (m *CreatePostRequestBody) GetAllowJoin()(*bool) {
    return m.allow_join
}
// GetAutoCloseDays gets the auto_close_days property value. Automatically close after these many days of inactivity
func (m *CreatePostRequestBody) GetAutoCloseDays()(*int32) {
    return m.auto_close_days
}
// GetBlurb gets the blurb property value. Small blurb that briefly describes the channel content.
func (m *CreatePostRequestBody) GetBlurb()(*string) {
    return m.blurb
}
// GetChannelName gets the channel_name property value. New channel name
func (m *CreatePostRequestBody) GetChannelName()(*string) {
    return m.channel_name
}
// GetCopyFromChannelId gets the copy_from_channel_id property value. Copy files and settings from this channel to the new channe. This channel must be owned by this user.
func (m *CreatePostRequestBody) GetCopyFromChannelId()(*int32) {
    return m.copy_from_channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_external_read"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalRead(val)
        }
        return nil
    }
    res["allow_join"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowJoin(val)
        }
        return nil
    }
    res["auto_close_days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoCloseDays(val)
        }
        return nil
    }
    res["blurb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlurb(val)
        }
        return nil
    }
    res["channel_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelName(val)
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
    res["is_public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublic(val)
        }
        return nil
    }
    res["require_join_approval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireJoinApproval(val)
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
// GetIsPublic gets the is_public property value. Defines if the channel have public access. If set to true, the channel will have a public access link, and anyone with access to this link will be able to see the channel content and join the channel. This property is deprecated. Use allow_join.
// Deprecated: 
func (m *CreatePostRequestBody) GetIsPublic()(*bool) {
    return m.is_public
}
// GetRequireJoinApproval gets the require_join_approval property value. Defines if the user must wait for a manager to approve his join to the channel. Only used when `allow_join` is set.
func (m *CreatePostRequestBody) GetRequireJoinApproval()(*bool) {
    return m.require_join_approval
}
// GetUsers gets the users property value. Comma separated email or phone list to add in addition to the user requesting this API.test1@test.com,+18197432105,test2@gmail.com
func (m *CreatePostRequestBody) GetUsers()(*string) {
    return m.users
}
// Serialize serializes information the current object
func (m *CreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_external_read", m.GetAllowExternalRead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_join", m.GetAllowJoin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("auto_close_days", m.GetAutoCloseDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blurb", m.GetBlurb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("channel_name", m.GetChannelName())
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
        err := writer.WriteBoolValue("is_public", m.GetIsPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("require_join_approval", m.GetRequireJoinApproval())
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
// SetAllowExternalRead sets the allow_external_read property value. Only used when `allow_join` is set. This allows anyone that have the channel link to see its contents (without joining the channel).
func (m *CreatePostRequestBody) SetAllowExternalRead(value *bool)() {
    m.allow_external_read = value
}
// SetAllowJoin sets the allow_join property value. Defines if the channel will have access through a link. If set to true, the channel will have an access link. This link will allow users to self join to the channel, and in some cases to see channel contents.
func (m *CreatePostRequestBody) SetAllowJoin(value *bool)() {
    m.allow_join = value
}
// SetAutoCloseDays sets the auto_close_days property value. Automatically close after these many days of inactivity
func (m *CreatePostRequestBody) SetAutoCloseDays(value *int32)() {
    m.auto_close_days = value
}
// SetBlurb sets the blurb property value. Small blurb that briefly describes the channel content.
func (m *CreatePostRequestBody) SetBlurb(value *string)() {
    m.blurb = value
}
// SetChannelName sets the channel_name property value. New channel name
func (m *CreatePostRequestBody) SetChannelName(value *string)() {
    m.channel_name = value
}
// SetCopyFromChannelId sets the copy_from_channel_id property value. Copy files and settings from this channel to the new channe. This channel must be owned by this user.
func (m *CreatePostRequestBody) SetCopyFromChannelId(value *int32)() {
    m.copy_from_channel_id = value
}
// SetIsPublic sets the is_public property value. Defines if the channel have public access. If set to true, the channel will have a public access link, and anyone with access to this link will be able to see the channel content and join the channel. This property is deprecated. Use allow_join.
// Deprecated: 
func (m *CreatePostRequestBody) SetIsPublic(value *bool)() {
    m.is_public = value
}
// SetRequireJoinApproval sets the require_join_approval property value. Defines if the user must wait for a manager to approve his join to the channel. Only used when `allow_join` is set.
func (m *CreatePostRequestBody) SetRequireJoinApproval(value *bool)() {
    m.require_join_approval = value
}
// SetUsers sets the users property value. Comma separated email or phone list to add in addition to the user requesting this API.test1@test.com,+18197432105,test2@gmail.com
func (m *CreatePostRequestBody) SetUsers(value *string)() {
    m.users = value
}
// CreatePostRequestBodyable 
type CreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowExternalRead()(*bool)
    GetAllowJoin()(*bool)
    GetAutoCloseDays()(*int32)
    GetBlurb()(*string)
    GetChannelName()(*string)
    GetCopyFromChannelId()(*int32)
    GetIsPublic()(*bool)
    GetRequireJoinApproval()(*bool)
    GetUsers()(*string)
    SetAllowExternalRead(value *bool)()
    SetAllowJoin(value *bool)()
    SetAutoCloseDays(value *int32)()
    SetBlurb(value *string)()
    SetChannelName(value *string)()
    SetCopyFromChannelId(value *int32)()
    SetIsPublic(value *bool)()
    SetRequireJoinApproval(value *bool)()
    SetUsers(value *string)()
}
