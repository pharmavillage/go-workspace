package actionupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePostRequestBody 
type UpdatePostRequestBody struct {
    // Action description
    action_desc *string
    // Action Due Date in date string format or timestamp integer format. To remove the due date, pass an empty string. If the field is not provided, the due date will be kept as it is.
    action_due_date UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable
    // Action Id
    action_id *string
    // Action name
    action_name *string
    // Action Status
    action_status *int32
    // Action Type
    action_type *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Comma separated user ids of assignees. 91010000,91010001
    user_ids *string
}
// UpdatePostRequestBody_UpdatePostRequestBody_action_due_date composed type wrapper for classes integer, string
type UpdatePostRequestBody_UpdatePostRequestBody_action_due_date struct {
    // Composed type representation for type integer
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewUpdatePostRequestBody_UpdatePostRequestBody_action_due_date instantiates a new updatePostRequestBody_action_due_date and sets the default values.
func NewUpdatePostRequestBody_UpdatePostRequestBody_action_due_date()(*UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) {
    m := &UpdatePostRequestBody_UpdatePostRequestBody_action_due_date{
    }
    return m
}
// CreateUpdatePostRequestBody_UpdatePostRequestBody_action_due_dateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePostRequestBody_UpdatePostRequestBody_action_due_dateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUpdatePostRequestBody_UpdatePostRequestBody_action_due_date()
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
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UpdatePostRequestBody_UpdatePostRequestBody_action_due_date) SetString(value *string)() {
    m.string = value
}
// UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable 
type UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
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
// GetActionDesc gets the action_desc property value. Action description
func (m *UpdatePostRequestBody) GetActionDesc()(*string) {
    return m.action_desc
}
// GetActionDueDate gets the action_due_date property value. Action Due Date in date string format or timestamp integer format. To remove the due date, pass an empty string. If the field is not provided, the due date will be kept as it is.
func (m *UpdatePostRequestBody) GetActionDueDate()(UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable) {
    return m.action_due_date
}
// GetActionId gets the action_id property value. Action Id
func (m *UpdatePostRequestBody) GetActionId()(*string) {
    return m.action_id
}
// GetActionName gets the action_name property value. Action name
func (m *UpdatePostRequestBody) GetActionName()(*string) {
    return m.action_name
}
// GetActionStatus gets the action_status property value. Action Status
func (m *UpdatePostRequestBody) GetActionStatus()(*int32) {
    return m.action_status
}
// GetActionType gets the action_type property value. Action Type
func (m *UpdatePostRequestBody) GetActionType()(*int32) {
    return m.action_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateUpdatePostRequestBody_UpdatePostRequestBody_action_due_dateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionDueDate(val.(UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable))
        }
        return nil
    }
    res["action_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionId(val)
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
// GetUserIds gets the user_ids property value. Comma separated user ids of assignees. 91010000,91010001
func (m *UpdatePostRequestBody) GetUserIds()(*string) {
    return m.user_ids
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("action_id", m.GetActionId())
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
func (m *UpdatePostRequestBody) SetActionDesc(value *string)() {
    m.action_desc = value
}
// SetActionDueDate sets the action_due_date property value. Action Due Date in date string format or timestamp integer format. To remove the due date, pass an empty string. If the field is not provided, the due date will be kept as it is.
func (m *UpdatePostRequestBody) SetActionDueDate(value UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable)() {
    m.action_due_date = value
}
// SetActionId sets the action_id property value. Action Id
func (m *UpdatePostRequestBody) SetActionId(value *string)() {
    m.action_id = value
}
// SetActionName sets the action_name property value. Action name
func (m *UpdatePostRequestBody) SetActionName(value *string)() {
    m.action_name = value
}
// SetActionStatus sets the action_status property value. Action Status
func (m *UpdatePostRequestBody) SetActionStatus(value *int32)() {
    m.action_status = value
}
// SetActionType sets the action_type property value. Action Type
func (m *UpdatePostRequestBody) SetActionType(value *int32)() {
    m.action_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUserIds sets the user_ids property value. Comma separated user ids of assignees. 91010000,91010001
func (m *UpdatePostRequestBody) SetUserIds(value *string)() {
    m.user_ids = value
}
// UpdatePostRequestBodyable 
type UpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionDesc()(*string)
    GetActionDueDate()(UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable)
    GetActionId()(*string)
    GetActionName()(*string)
    GetActionStatus()(*int32)
    GetActionType()(*int32)
    GetUserIds()(*string)
    SetActionDesc(value *string)()
    SetActionDueDate(value UpdatePostRequestBody_UpdatePostRequestBody_action_due_dateable)()
    SetActionId(value *string)()
    SetActionName(value *string)()
    SetActionStatus(value *int32)()
    SetActionType(value *int32)()
    SetUserIds(value *string)()
}
