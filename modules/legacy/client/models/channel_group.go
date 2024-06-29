package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelGroup 
type ChannelGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Must be a int when the group is a user defined group (non virtual). Must be a string when the group is a virtual group (system defined group).
    id ChannelGroup_ChannelGroup_idable
    // Group name that should be rendered (it varies depending on the user language).
    name *string
    // When `virtual` is false, it's expected that the id is an integer. When it's true, the id should be a string. Virtual groups have their own logic handled by the UI when grouping the channels. Non virtual groups are user-created groups, the it's id will be found on the channel object, on `channel_group_id` param.
    virtual *bool
}
// ChannelGroup_ChannelGroup_id composed type wrapper for classes integer, string
type ChannelGroup_ChannelGroup_id struct {
    // Composed type representation for type integer
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewChannelGroup_ChannelGroup_id instantiates a new ChannelGroup_id and sets the default values.
func NewChannelGroup_ChannelGroup_id()(*ChannelGroup_ChannelGroup_id) {
    m := &ChannelGroup_ChannelGroup_id{
    }
    return m
}
// CreateChannelGroup_ChannelGroup_idFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelGroup_ChannelGroup_idFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewChannelGroup_ChannelGroup_id()
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
func (m *ChannelGroup_ChannelGroup_id) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *ChannelGroup_ChannelGroup_id) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *ChannelGroup_ChannelGroup_id) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
func (m *ChannelGroup_ChannelGroup_id) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ChannelGroup_ChannelGroup_id) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ChannelGroup_ChannelGroup_id) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ChannelGroup_ChannelGroup_id) SetString(value *string)() {
    m.string = value
}
// ChannelGroup_ChannelGroup_idable 
type ChannelGroup_ChannelGroup_idable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewChannelGroup instantiates a new ChannelGroup and sets the default values.
func NewChannelGroup()(*ChannelGroup) {
    m := &ChannelGroup{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChannelGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelGroup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelGroup) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChannelGroup_ChannelGroup_idFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val.(ChannelGroup_ChannelGroup_idable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["virtual"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtual(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Must be a int when the group is a user defined group (non virtual). Must be a string when the group is a virtual group (system defined group).
func (m *ChannelGroup) GetId()(ChannelGroup_ChannelGroup_idable) {
    return m.id
}
// GetName gets the name property value. Group name that should be rendered (it varies depending on the user language).
func (m *ChannelGroup) GetName()(*string) {
    return m.name
}
// GetVirtual gets the virtual property value. When `virtual` is false, it's expected that the id is an integer. When it's true, the id should be a string. Virtual groups have their own logic handled by the UI when grouping the channels. Non virtual groups are user-created groups, the it's id will be found on the channel object, on `channel_group_id` param.
func (m *ChannelGroup) GetVirtual()(*bool) {
    return m.virtual
}
// Serialize serializes information the current object
func (m *ChannelGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("virtual", m.GetVirtual())
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
func (m *ChannelGroup) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Must be a int when the group is a user defined group (non virtual). Must be a string when the group is a virtual group (system defined group).
func (m *ChannelGroup) SetId(value ChannelGroup_ChannelGroup_idable)() {
    m.id = value
}
// SetName sets the name property value. Group name that should be rendered (it varies depending on the user language).
func (m *ChannelGroup) SetName(value *string)() {
    m.name = value
}
// SetVirtual sets the virtual property value. When `virtual` is false, it's expected that the id is an integer. When it's true, the id should be a string. Virtual groups have their own logic handled by the UI when grouping the channels. Non virtual groups are user-created groups, the it's id will be found on the channel object, on `channel_group_id` param.
func (m *ChannelGroup) SetVirtual(value *bool)() {
    m.virtual = value
}
// ChannelGroupable 
type ChannelGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(ChannelGroup_ChannelGroup_idable)
    GetName()(*string)
    GetVirtual()(*bool)
    SetId(value ChannelGroup_ChannelGroup_idable)()
    SetName(value *string)()
    SetVirtual(value *bool)()
}
