package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Action 
type Action struct {
    // The action_desc property
    action_desc *string
    // The action_name property
    action_name *string
    // The action_status property
    action_status *int32
    // The action_type property
    action_type *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The channel_id property
    channel_id *int32
    // The channel_name property
    channel_name *string
    // The created_by property
    created_by *int32
    // The created_on property
    created_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The due_on property
    due_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *int32
    // The updated_by property
    updated_by *int32
    // The updated_on property
    updated_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAction instantiates a new Action and sets the default values.
func NewAction()(*Action) {
    m := &Action{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAction(), nil
}
// GetActionDesc gets the action_desc property value. The action_desc property
func (m *Action) GetActionDesc()(*string) {
    return m.action_desc
}
// GetActionName gets the action_name property value. The action_name property
func (m *Action) GetActionName()(*string) {
    return m.action_name
}
// GetActionStatus gets the action_status property value. The action_status property
func (m *Action) GetActionStatus()(*int32) {
    return m.action_status
}
// GetActionType gets the action_type property value. The action_type property
func (m *Action) GetActionType()(*int32) {
    return m.action_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Action) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. The channel_id property
func (m *Action) GetChannelId()(*int32) {
    return m.channel_id
}
// GetChannelName gets the channel_name property value. The channel_name property
func (m *Action) GetChannelName()(*string) {
    return m.channel_name
}
// GetCreatedBy gets the created_by property value. The created_by property
func (m *Action) GetCreatedBy()(*int32) {
    return m.created_by
}
// GetCreatedOn gets the created_on property value. The created_on property
func (m *Action) GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_on
}
// GetDueOn gets the due_on property value. The due_on property
func (m *Action) GetDueOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.due_on
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Action) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["created_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["created_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOn(val)
        }
        return nil
    }
    res["due_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueOn(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["updated_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedBy(val)
        }
        return nil
    }
    res["updated_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedOn(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *Action) GetId()(*int32) {
    return m.id
}
// GetUpdatedBy gets the updated_by property value. The updated_by property
func (m *Action) GetUpdatedBy()(*int32) {
    return m.updated_by
}
// GetUpdatedOn gets the updated_on property value. The updated_on property
func (m *Action) GetUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_on
}
// Serialize serializes information the current object
func (m *Action) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("action_desc", m.GetActionDesc())
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
        err := writer.WriteStringValue("channel_name", m.GetChannelName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("created_by", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_on", m.GetCreatedOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("due_on", m.GetDueOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("updated_by", m.GetUpdatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_on", m.GetUpdatedOn())
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
// SetActionDesc sets the action_desc property value. The action_desc property
func (m *Action) SetActionDesc(value *string)() {
    m.action_desc = value
}
// SetActionName sets the action_name property value. The action_name property
func (m *Action) SetActionName(value *string)() {
    m.action_name = value
}
// SetActionStatus sets the action_status property value. The action_status property
func (m *Action) SetActionStatus(value *int32)() {
    m.action_status = value
}
// SetActionType sets the action_type property value. The action_type property
func (m *Action) SetActionType(value *int32)() {
    m.action_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Action) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. The channel_id property
func (m *Action) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetChannelName sets the channel_name property value. The channel_name property
func (m *Action) SetChannelName(value *string)() {
    m.channel_name = value
}
// SetCreatedBy sets the created_by property value. The created_by property
func (m *Action) SetCreatedBy(value *int32)() {
    m.created_by = value
}
// SetCreatedOn sets the created_on property value. The created_on property
func (m *Action) SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_on = value
}
// SetDueOn sets the due_on property value. The due_on property
func (m *Action) SetDueOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.due_on = value
}
// SetId sets the id property value. The id property
func (m *Action) SetId(value *int32)() {
    m.id = value
}
// SetUpdatedBy sets the updated_by property value. The updated_by property
func (m *Action) SetUpdatedBy(value *int32)() {
    m.updated_by = value
}
// SetUpdatedOn sets the updated_on property value. The updated_on property
func (m *Action) SetUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_on = value
}
// Actionable 
type Actionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionDesc()(*string)
    GetActionName()(*string)
    GetActionStatus()(*int32)
    GetActionType()(*int32)
    GetChannelId()(*int32)
    GetChannelName()(*string)
    GetCreatedBy()(*int32)
    GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDueOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int32)
    GetUpdatedBy()(*int32)
    GetUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActionDesc(value *string)()
    SetActionName(value *string)()
    SetActionStatus(value *int32)()
    SetActionType(value *int32)()
    SetChannelId(value *int32)()
    SetChannelName(value *string)()
    SetCreatedBy(value *int32)()
    SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDueOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int32)()
    SetUpdatedBy(value *int32)()
    SetUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
