package channelgroupmove

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MovePostRequestBody 
type MovePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The channel group id where the group should be moved after. It can be an integer (for user defined groups) or a string (for virtual/system defined groups)
    after MovePostRequestBody_MovePostRequestBody_afterable
    // The channel group id to be moved. It can be an integer (for user defined groups) or a string (for virtual/system defined groups)
    group_id MovePostRequestBody_MovePostRequestBody_group_idable
}
// MovePostRequestBody_MovePostRequestBody_after composed type wrapper for classes integer, string
type MovePostRequestBody_MovePostRequestBody_after struct {
    // Composed type representation for type integer
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewMovePostRequestBody_MovePostRequestBody_after instantiates a new movePostRequestBody_after and sets the default values.
func NewMovePostRequestBody_MovePostRequestBody_after()(*MovePostRequestBody_MovePostRequestBody_after) {
    m := &MovePostRequestBody_MovePostRequestBody_after{
    }
    return m
}
// CreateMovePostRequestBody_MovePostRequestBody_afterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMovePostRequestBody_MovePostRequestBody_afterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMovePostRequestBody_MovePostRequestBody_after()
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
func (m *MovePostRequestBody_MovePostRequestBody_after) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *MovePostRequestBody_MovePostRequestBody_after) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *MovePostRequestBody_MovePostRequestBody_after) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
func (m *MovePostRequestBody_MovePostRequestBody_after) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *MovePostRequestBody_MovePostRequestBody_after) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MovePostRequestBody_MovePostRequestBody_after) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *MovePostRequestBody_MovePostRequestBody_after) SetString(value *string)() {
    m.string = value
}
// MovePostRequestBody_MovePostRequestBody_group_id composed type wrapper for classes integer, string
type MovePostRequestBody_MovePostRequestBody_group_id struct {
    // Composed type representation for type integer
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewMovePostRequestBody_MovePostRequestBody_group_id instantiates a new movePostRequestBody_group_id and sets the default values.
func NewMovePostRequestBody_MovePostRequestBody_group_id()(*MovePostRequestBody_MovePostRequestBody_group_id) {
    m := &MovePostRequestBody_MovePostRequestBody_group_id{
    }
    return m
}
// CreateMovePostRequestBody_MovePostRequestBody_group_idFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMovePostRequestBody_MovePostRequestBody_group_idFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMovePostRequestBody_MovePostRequestBody_group_id()
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
func (m *MovePostRequestBody_MovePostRequestBody_group_id) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *MovePostRequestBody_MovePostRequestBody_group_id) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *MovePostRequestBody_MovePostRequestBody_group_id) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
func (m *MovePostRequestBody_MovePostRequestBody_group_id) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *MovePostRequestBody_MovePostRequestBody_group_id) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MovePostRequestBody_MovePostRequestBody_group_id) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *MovePostRequestBody_MovePostRequestBody_group_id) SetString(value *string)() {
    m.string = value
}
// MovePostRequestBody_MovePostRequestBody_afterable 
type MovePostRequestBody_MovePostRequestBody_afterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// MovePostRequestBody_MovePostRequestBody_group_idable 
type MovePostRequestBody_MovePostRequestBody_group_idable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewMovePostRequestBody instantiates a new movePostRequestBody and sets the default values.
func NewMovePostRequestBody()(*MovePostRequestBody) {
    m := &MovePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMovePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMovePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMovePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MovePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfter gets the after property value. The channel group id where the group should be moved after. It can be an integer (for user defined groups) or a string (for virtual/system defined groups)
func (m *MovePostRequestBody) GetAfter()(MovePostRequestBody_MovePostRequestBody_afterable) {
    return m.after
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MovePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["after"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMovePostRequestBody_MovePostRequestBody_afterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfter(val.(MovePostRequestBody_MovePostRequestBody_afterable))
        }
        return nil
    }
    res["group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMovePostRequestBody_MovePostRequestBody_group_idFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val.(MovePostRequestBody_MovePostRequestBody_group_idable))
        }
        return nil
    }
    return res
}
// GetGroupId gets the group_id property value. The channel group id to be moved. It can be an integer (for user defined groups) or a string (for virtual/system defined groups)
func (m *MovePostRequestBody) GetGroupId()(MovePostRequestBody_MovePostRequestBody_group_idable) {
    return m.group_id
}
// Serialize serializes information the current object
func (m *MovePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("after", m.GetAfter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("group_id", m.GetGroupId())
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
func (m *MovePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfter sets the after property value. The channel group id where the group should be moved after. It can be an integer (for user defined groups) or a string (for virtual/system defined groups)
func (m *MovePostRequestBody) SetAfter(value MovePostRequestBody_MovePostRequestBody_afterable)() {
    m.after = value
}
// SetGroupId sets the group_id property value. The channel group id to be moved. It can be an integer (for user defined groups) or a string (for virtual/system defined groups)
func (m *MovePostRequestBody) SetGroupId(value MovePostRequestBody_MovePostRequestBody_group_idable)() {
    m.group_id = value
}
// MovePostRequestBodyable 
type MovePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAfter()(MovePostRequestBody_MovePostRequestBody_afterable)
    GetGroupId()(MovePostRequestBody_MovePostRequestBody_group_idable)
    SetAfter(value MovePostRequestBody_MovePostRequestBody_afterable)()
    SetGroupId(value MovePostRequestBody_MovePostRequestBody_group_idable)()
}
