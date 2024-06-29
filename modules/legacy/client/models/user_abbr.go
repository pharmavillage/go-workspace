package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// User_Abbr 
type User_Abbr struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display_name property
    display_name *string
    // The email property
    email *string
    // The has_avatar property
    has_avatar *bool
    // The id property
    id *int32
    // The last_active_on property
    last_active_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last_active_on_ts property
    last_active_on_ts *int32
    // The online_status property
    online_status *bool
}
// NewUser_Abbr instantiates a new User_Abbr and sets the default values.
func NewUser_Abbr()(*User_Abbr) {
    m := &User_Abbr{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUser_AbbrFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUser_AbbrFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser_Abbr(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *User_Abbr) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the display_name property value. The display_name property
func (m *User_Abbr) GetDisplayName()(*string) {
    return m.display_name
}
// GetEmail gets the email property value. The email property
func (m *User_Abbr) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *User_Abbr) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["has_avatar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAvatar(val)
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
    res["last_active_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActiveOn(val)
        }
        return nil
    }
    res["last_active_on_ts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActiveOnTs(val)
        }
        return nil
    }
    res["online_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineStatus(val)
        }
        return nil
    }
    return res
}
// GetHasAvatar gets the has_avatar property value. The has_avatar property
func (m *User_Abbr) GetHasAvatar()(*bool) {
    return m.has_avatar
}
// GetId gets the id property value. The id property
func (m *User_Abbr) GetId()(*int32) {
    return m.id
}
// GetLastActiveOn gets the last_active_on property value. The last_active_on property
func (m *User_Abbr) GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_active_on
}
// GetLastActiveOnTs gets the last_active_on_ts property value. The last_active_on_ts property
func (m *User_Abbr) GetLastActiveOnTs()(*int32) {
    return m.last_active_on_ts
}
// GetOnlineStatus gets the online_status property value. The online_status property
func (m *User_Abbr) GetOnlineStatus()(*bool) {
    return m.online_status
}
// Serialize serializes information the current object
func (m *User_Abbr) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_avatar", m.GetHasAvatar())
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
        err := writer.WriteTimeValue("last_active_on", m.GetLastActiveOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("last_active_on_ts", m.GetLastActiveOnTs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("online_status", m.GetOnlineStatus())
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
func (m *User_Abbr) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *User_Abbr) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetEmail sets the email property value. The email property
func (m *User_Abbr) SetEmail(value *string)() {
    m.email = value
}
// SetHasAvatar sets the has_avatar property value. The has_avatar property
func (m *User_Abbr) SetHasAvatar(value *bool)() {
    m.has_avatar = value
}
// SetId sets the id property value. The id property
func (m *User_Abbr) SetId(value *int32)() {
    m.id = value
}
// SetLastActiveOn sets the last_active_on property value. The last_active_on property
func (m *User_Abbr) SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_active_on = value
}
// SetLastActiveOnTs sets the last_active_on_ts property value. The last_active_on_ts property
func (m *User_Abbr) SetLastActiveOnTs(value *int32)() {
    m.last_active_on_ts = value
}
// SetOnlineStatus sets the online_status property value. The online_status property
func (m *User_Abbr) SetOnlineStatus(value *bool)() {
    m.online_status = value
}
// User_Abbrable 
type User_Abbrable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetHasAvatar()(*bool)
    GetId()(*int32)
    GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastActiveOnTs()(*int32)
    GetOnlineStatus()(*bool)
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetHasAvatar(value *bool)()
    SetId(value *int32)()
    SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastActiveOnTs(value *int32)()
    SetOnlineStatus(value *bool)()
}
