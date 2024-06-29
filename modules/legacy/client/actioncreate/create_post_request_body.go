package actioncreate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreatePostRequestBody 
type CreatePostRequestBody struct {
    // Action description
    action_desc *string
    // Action Due Date in date string format or timestamp integer format
    action_due_date CreatePostRequestBody_CreatePostRequestBody_action_due_dateable
    // New action name
    action_name *string
    // Action Status
    action_status *int32
    // Action Type
    action_type *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Channel Id
    channel_id *int32
    // Indicates that the new action must be a child of the action which the id is passed on this field
    parent_id *int32
    // Comma separated user ids of assignees. 91010000,91010001
    user_ids *string
}
// CreatePostRequestBody_CreatePostRequestBody_action_due_date composed type wrapper for classes integer, string
type CreatePostRequestBody_CreatePostRequestBody_action_due_date struct {
    // Composed type representation for type integer
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewCreatePostRequestBody_CreatePostRequestBody_action_due_date instantiates a new createPostRequestBody_action_due_date and sets the default values.
func NewCreatePostRequestBody_CreatePostRequestBody_action_due_date()(*CreatePostRequestBody_CreatePostRequestBody_action_due_date) {
    m := &CreatePostRequestBody_CreatePostRequestBody_action_due_date{
    }
    return m
}
// CreateCreatePostRequestBody_CreatePostRequestBody_action_due_dateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreatePostRequestBody_CreatePostRequestBody_action_due_dateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCreatePostRequestBody_CreatePostRequestBody_action_due_date()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type integer
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *CreatePostRequestBody_CreatePostRequestBody_action_due_date) SetString(value *string)() {
    m.string = value
}
// CreatePostRequestBody_CreatePostRequestBody_action_due_dateable 
type CreatePostRequestBody_CreatePostRequestBody_action_due_dateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewCreatePostRequestBody instantiates a new createPostRequestBody and sets the default values.
func NewCreatePostRequestBody()(*CreatePostRequestBody) {
    m := &CreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreatePostRequestBody(), nil
}
// GetActionDesc gets the action_desc property value. Action description
func (m *CreatePostRequestBody) GetActionDesc()(*string) {
    return m.action_desc
}
// GetActionDueDate gets the action_due_date property value. Action Due Date in date string format or timestamp integer format
func (m *CreatePostRequestBody) GetActionDueDate()(CreatePostRequestBody_CreatePostRequestBody_action_due_dateable) {
    return m.action_due_date
}
// GetActionName gets the action_name property value. New action name
func (m *CreatePostRequestBody) GetActionName()(*string) {
    return m.action_name
}
// GetActionStatus gets the action_status property value. Action Status
func (m *CreatePostRequestBody) GetActionStatus()(*int32) {
    return m.action_status
}
// GetActionType gets the action_type property value. Action Type
func (m *CreatePostRequestBody) GetActionType()(*int32) {
    return m.action_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. Channel Id
func (m *CreatePostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action_desc"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionDesc(val)
        }
        return nil
    }
    res["action_due_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCreatePostRequestBody_CreatePostRequestBody_action_due_dateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionDueDate(val.(CreatePostRequestBody_CreatePostRequestBody_action_due_dateable))
        }
        return nil
    }
    res["action_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
    res["action_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionStatus(val)
        }
        return nil
    }
    res["action_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val)
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
    res["parent_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentId(val)
        }
        return nil
    }
    res["user_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserIds(val)
        }
        return nil
    }
    return res
}
// GetParentId gets the parent_id property value. Indicates that the new action must be a child of the action which the id is passed on this field
func (m *CreatePostRequestBody) GetParentId()(*int32) {
    return m.parent_id
}
// GetUserIds gets the user_ids property value. Comma separated user ids of assignees. 91010000,91010001
func (m *CreatePostRequestBody) GetUserIds()(*string) {
    return m.user_ids
}
// Serialize serializes information the current object
func (m *CreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("action_desc", m.GetActionDesc())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("action_due_date", m.GetActionDueDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("action_name", m.GetActionName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("action_status", m.GetActionStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("action_type", m.GetActionType())
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
        err := writer.WriteInt32Value("parent_id", m.GetParentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_ids", m.GetUserIds())
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
// SetActionDesc sets the action_desc property value. Action description
func (m *CreatePostRequestBody) SetActionDesc(value *string)() {
    m.action_desc = value
}
// SetActionDueDate sets the action_due_date property value. Action Due Date in date string format or timestamp integer format
func (m *CreatePostRequestBody) SetActionDueDate(value CreatePostRequestBody_CreatePostRequestBody_action_due_dateable)() {
    m.action_due_date = value
}
// SetActionName sets the action_name property value. New action name
func (m *CreatePostRequestBody) SetActionName(value *string)() {
    m.action_name = value
}
// SetActionStatus sets the action_status property value. Action Status
func (m *CreatePostRequestBody) SetActionStatus(value *int32)() {
    m.action_status = value
}
// SetActionType sets the action_type property value. Action Type
func (m *CreatePostRequestBody) SetActionType(value *int32)() {
    m.action_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. Channel Id
func (m *CreatePostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetParentId sets the parent_id property value. Indicates that the new action must be a child of the action which the id is passed on this field
func (m *CreatePostRequestBody) SetParentId(value *int32)() {
    m.parent_id = value
}
// SetUserIds sets the user_ids property value. Comma separated user ids of assignees. 91010000,91010001
func (m *CreatePostRequestBody) SetUserIds(value *string)() {
    m.user_ids = value
}
// CreatePostRequestBodyable 
type CreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionDesc()(*string)
    GetActionDueDate()(CreatePostRequestBody_CreatePostRequestBody_action_due_dateable)
    GetActionName()(*string)
    GetActionStatus()(*int32)
    GetActionType()(*int32)
    GetChannelId()(*int32)
    GetParentId()(*int32)
    GetUserIds()(*string)
    SetActionDesc(value *string)()
    SetActionDueDate(value CreatePostRequestBody_CreatePostRequestBody_action_due_dateable)()
    SetActionName(value *string)()
    SetActionStatus(value *int32)()
    SetActionType(value *int32)()
    SetChannelId(value *int32)()
    SetParentId(value *int32)()
    SetUserIds(value *string)()
}
