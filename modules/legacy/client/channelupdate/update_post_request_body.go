package channelupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePostRequestBody 
type UpdatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Only used when `allow_join` is set. This allows anyone that have the channel link to see its contents (without joining the channel). If not provided, previous status is kept.
    allow_external_read *bool
    // Defines if the channel will have access through a link. If set to true, the channel will have an access link. This link will allow users to self join to the channel, and in some cases to see channel contents. If not provided, previous status is kept.
    allow_join *bool
    // The new blurb for the channel. If not provided, the blurb is not changed
    blurb *string
    // Id of the channel
    channel_id *int32
    // The new name of the channel. If not provided, the name is not changed
    channel_name *string
    // true for favorite, false otherwise
    is_favorite *bool
    // Defines if the channel have public access. If set to true, the channel will have a public access link, and anyone with access to this link will be able to see the channel content and join the channel. If not provided, previous status is kept. This property is deprecated. Use allow_join.
    // Deprecated: 
    is_public *bool
    // Defines if the user must wait for a manager to approve his join to the channel. Only used when `allow_join` is set. If not provided, previous status is kept.
    require_join_approval *bool
}
// NewUpdatePostRequestBody instantiates a new updatePostRequestBody and sets the default values.
func NewUpdatePostRequestBody()(*UpdatePostRequestBody) {
    m := &UpdatePostRequestBody{
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
// CreateUpdatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowExternalRead gets the allow_external_read property value. Only used when `allow_join` is set. This allows anyone that have the channel link to see its contents (without joining the channel). If not provided, previous status is kept.
func (m *UpdatePostRequestBody) GetAllowExternalRead()(*bool) {
    return m.allow_external_read
}
// GetAllowJoin gets the allow_join property value. Defines if the channel will have access through a link. If set to true, the channel will have an access link. This link will allow users to self join to the channel, and in some cases to see channel contents. If not provided, previous status is kept.
func (m *UpdatePostRequestBody) GetAllowJoin()(*bool) {
    return m.allow_join
}
// GetBlurb gets the blurb property value. The new blurb for the channel. If not provided, the blurb is not changed
func (m *UpdatePostRequestBody) GetBlurb()(*string) {
    return m.blurb
}
// GetChannelId gets the channel_id property value. Id of the channel
func (m *UpdatePostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetChannelName gets the channel_name property value. The new name of the channel. If not provided, the name is not changed
func (m *UpdatePostRequestBody) GetChannelName()(*string) {
    return m.channel_name
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["is_favorite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavorite(val)
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
    return res
}
// GetIsFavorite gets the is_favorite property value. true for favorite, false otherwise
func (m *UpdatePostRequestBody) GetIsFavorite()(*bool) {
    return m.is_favorite
}
// GetIsPublic gets the is_public property value. Defines if the channel have public access. If set to true, the channel will have a public access link, and anyone with access to this link will be able to see the channel content and join the channel. If not provided, previous status is kept. This property is deprecated. Use allow_join.
// Deprecated: 
func (m *UpdatePostRequestBody) GetIsPublic()(*bool) {
    return m.is_public
}
// GetRequireJoinApproval gets the require_join_approval property value. Defines if the user must wait for a manager to approve his join to the channel. Only used when `allow_join` is set. If not provided, previous status is kept.
func (m *UpdatePostRequestBody) GetRequireJoinApproval()(*bool) {
    return m.require_join_approval
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("blurb", m.GetBlurb())
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
        err := writer.WriteStringValue("channel_name", m.GetChannelName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_favorite", m.GetIsFavorite())
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
// SetAllowExternalRead sets the allow_external_read property value. Only used when `allow_join` is set. This allows anyone that have the channel link to see its contents (without joining the channel). If not provided, previous status is kept.
func (m *UpdatePostRequestBody) SetAllowExternalRead(value *bool)() {
    m.allow_external_read = value
}
// SetAllowJoin sets the allow_join property value. Defines if the channel will have access through a link. If set to true, the channel will have an access link. This link will allow users to self join to the channel, and in some cases to see channel contents. If not provided, previous status is kept.
func (m *UpdatePostRequestBody) SetAllowJoin(value *bool)() {
    m.allow_join = value
}
// SetBlurb sets the blurb property value. The new blurb for the channel. If not provided, the blurb is not changed
func (m *UpdatePostRequestBody) SetBlurb(value *string)() {
    m.blurb = value
}
// SetChannelId sets the channel_id property value. Id of the channel
func (m *UpdatePostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetChannelName sets the channel_name property value. The new name of the channel. If not provided, the name is not changed
func (m *UpdatePostRequestBody) SetChannelName(value *string)() {
    m.channel_name = value
}
// SetIsFavorite sets the is_favorite property value. true for favorite, false otherwise
func (m *UpdatePostRequestBody) SetIsFavorite(value *bool)() {
    m.is_favorite = value
}
// SetIsPublic sets the is_public property value. Defines if the channel have public access. If set to true, the channel will have a public access link, and anyone with access to this link will be able to see the channel content and join the channel. If not provided, previous status is kept. This property is deprecated. Use allow_join.
// Deprecated: 
func (m *UpdatePostRequestBody) SetIsPublic(value *bool)() {
    m.is_public = value
}
// SetRequireJoinApproval sets the require_join_approval property value. Defines if the user must wait for a manager to approve his join to the channel. Only used when `allow_join` is set. If not provided, previous status is kept.
func (m *UpdatePostRequestBody) SetRequireJoinApproval(value *bool)() {
    m.require_join_approval = value
}
// UpdatePostRequestBodyable 
type UpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowExternalRead()(*bool)
    GetAllowJoin()(*bool)
    GetBlurb()(*string)
    GetChannelId()(*int32)
    GetChannelName()(*string)
    GetIsFavorite()(*bool)
    GetIsPublic()(*bool)
    GetRequireJoinApproval()(*bool)
    SetAllowExternalRead(value *bool)()
    SetAllowJoin(value *bool)()
    SetBlurb(value *string)()
    SetChannelId(value *int32)()
    SetChannelName(value *string)()
    SetIsFavorite(value *bool)()
    SetIsPublic(value *bool)()
    SetRequireJoinApproval(value *bool)()
}
