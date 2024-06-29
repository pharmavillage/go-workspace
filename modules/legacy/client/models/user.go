package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// User 
type User struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_on property
    created_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The created_on_ts property
    created_on_ts *int32
    // The display_name property
    display_name *string
    // The email property
    email *string
    // The has_avatar property
    has_avatar *bool
    // The id property
    id *int32
    // The invited_by property
    invited_by *int32
    // The is_terms_agreed property
    is_terms_agreed *bool
    // The is_tour_complete property
    is_tour_complete *bool
    // The lang_code property
    lang_code *string
    // The last_active_on property
    last_active_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last_active_on_ts property
    last_active_on_ts *int32
    // Notification option (can be any of the values below):  * `0` - No notifications (completely unsubscribe)  * `1` - Only notify mentions  * `2` - Notify all unread messages
    notifications_config *int32
    // The online_status property
    online_status *int32
    // The phone property
    phone *string
    // ISO text representation of the user timezone
    timezone *string
    // The updated_by property
    updated_by *int32
    // The user_fs_roots property
    user_fs_roots []User_fs_rootsable
    // The user_fs_stats property
    user_fs_stats User_fs_statsable
    // The user_role property
    user_role *int32
}
// NewUser instantiates a new User and sets the default values.
func NewUser()(*User) {
    m := &User{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *User) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedOn gets the created_on property value. The created_on property
func (m *User) GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_on
}
// GetCreatedOnTs gets the created_on_ts property value. The created_on_ts property
func (m *User) GetCreatedOnTs()(*int32) {
    return m.created_on_ts
}
// GetDisplayName gets the display_name property value. The display_name property
func (m *User) GetDisplayName()(*string) {
    return m.display_name
}
// GetEmail gets the email property value. The email property
func (m *User) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *User) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["created_on_ts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnTs(val)
        }
        return nil
    }
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
    res["invited_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedBy(val)
        }
        return nil
    }
    res["is_terms_agreed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTermsAgreed(val)
        }
        return nil
    }
    res["is_tour_complete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTourComplete(val)
        }
        return nil
    }
    res["lang_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLangCode(val)
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
    res["notifications_config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationsConfig(val)
        }
        return nil
    }
    res["online_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineStatus(val)
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["timezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimezone(val)
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
    res["user_fs_roots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUser_fs_rootsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]User_fs_rootsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(User_fs_rootsable)
                }
            }
            m.SetUserFsRoots(res)
        }
        return nil
    }
    res["user_fs_stats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUser_fs_statsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFsStats(val.(User_fs_statsable))
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
// GetHasAvatar gets the has_avatar property value. The has_avatar property
func (m *User) GetHasAvatar()(*bool) {
    return m.has_avatar
}
// GetId gets the id property value. The id property
func (m *User) GetId()(*int32) {
    return m.id
}
// GetInvitedBy gets the invited_by property value. The invited_by property
func (m *User) GetInvitedBy()(*int32) {
    return m.invited_by
}
// GetIsTermsAgreed gets the is_terms_agreed property value. The is_terms_agreed property
func (m *User) GetIsTermsAgreed()(*bool) {
    return m.is_terms_agreed
}
// GetIsTourComplete gets the is_tour_complete property value. The is_tour_complete property
func (m *User) GetIsTourComplete()(*bool) {
    return m.is_tour_complete
}
// GetLangCode gets the lang_code property value. The lang_code property
func (m *User) GetLangCode()(*string) {
    return m.lang_code
}
// GetLastActiveOn gets the last_active_on property value. The last_active_on property
func (m *User) GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_active_on
}
// GetLastActiveOnTs gets the last_active_on_ts property value. The last_active_on_ts property
func (m *User) GetLastActiveOnTs()(*int32) {
    return m.last_active_on_ts
}
// GetNotificationsConfig gets the notifications_config property value. Notification option (can be any of the values below):  * `0` - No notifications (completely unsubscribe)  * `1` - Only notify mentions  * `2` - Notify all unread messages
func (m *User) GetNotificationsConfig()(*int32) {
    return m.notifications_config
}
// GetOnlineStatus gets the online_status property value. The online_status property
func (m *User) GetOnlineStatus()(*int32) {
    return m.online_status
}
// GetPhone gets the phone property value. The phone property
func (m *User) GetPhone()(*string) {
    return m.phone
}
// GetTimezone gets the timezone property value. ISO text representation of the user timezone
func (m *User) GetTimezone()(*string) {
    return m.timezone
}
// GetUpdatedBy gets the updated_by property value. The updated_by property
func (m *User) GetUpdatedBy()(*int32) {
    return m.updated_by
}
// GetUserFsRoots gets the user_fs_roots property value. The user_fs_roots property
func (m *User) GetUserFsRoots()([]User_fs_rootsable) {
    return m.user_fs_roots
}
// GetUserFsStats gets the user_fs_stats property value. The user_fs_stats property
func (m *User) GetUserFsStats()(User_fs_statsable) {
    return m.user_fs_stats
}
// GetUserRole gets the user_role property value. The user_role property
func (m *User) GetUserRole()(*int32) {
    return m.user_role
}
// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_on", m.GetCreatedOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("created_on_ts", m.GetCreatedOnTs())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteInt32Value("invited_by", m.GetInvitedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_terms_agreed", m.GetIsTermsAgreed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_tour_complete", m.GetIsTourComplete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lang_code", m.GetLangCode())
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
        err := writer.WriteInt32Value("notifications_config", m.GetNotificationsConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("online_status", m.GetOnlineStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timezone", m.GetTimezone())
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
    if m.GetUserFsRoots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserFsRoots()))
        for i, v := range m.GetUserFsRoots() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("user_fs_roots", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user_fs_stats", m.GetUserFsStats())
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
func (m *User) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedOn sets the created_on property value. The created_on property
func (m *User) SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_on = value
}
// SetCreatedOnTs sets the created_on_ts property value. The created_on_ts property
func (m *User) SetCreatedOnTs(value *int32)() {
    m.created_on_ts = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *User) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetEmail sets the email property value. The email property
func (m *User) SetEmail(value *string)() {
    m.email = value
}
// SetHasAvatar sets the has_avatar property value. The has_avatar property
func (m *User) SetHasAvatar(value *bool)() {
    m.has_avatar = value
}
// SetId sets the id property value. The id property
func (m *User) SetId(value *int32)() {
    m.id = value
}
// SetInvitedBy sets the invited_by property value. The invited_by property
func (m *User) SetInvitedBy(value *int32)() {
    m.invited_by = value
}
// SetIsTermsAgreed sets the is_terms_agreed property value. The is_terms_agreed property
func (m *User) SetIsTermsAgreed(value *bool)() {
    m.is_terms_agreed = value
}
// SetIsTourComplete sets the is_tour_complete property value. The is_tour_complete property
func (m *User) SetIsTourComplete(value *bool)() {
    m.is_tour_complete = value
}
// SetLangCode sets the lang_code property value. The lang_code property
func (m *User) SetLangCode(value *string)() {
    m.lang_code = value
}
// SetLastActiveOn sets the last_active_on property value. The last_active_on property
func (m *User) SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_active_on = value
}
// SetLastActiveOnTs sets the last_active_on_ts property value. The last_active_on_ts property
func (m *User) SetLastActiveOnTs(value *int32)() {
    m.last_active_on_ts = value
}
// SetNotificationsConfig sets the notifications_config property value. Notification option (can be any of the values below):  * `0` - No notifications (completely unsubscribe)  * `1` - Only notify mentions  * `2` - Notify all unread messages
func (m *User) SetNotificationsConfig(value *int32)() {
    m.notifications_config = value
}
// SetOnlineStatus sets the online_status property value. The online_status property
func (m *User) SetOnlineStatus(value *int32)() {
    m.online_status = value
}
// SetPhone sets the phone property value. The phone property
func (m *User) SetPhone(value *string)() {
    m.phone = value
}
// SetTimezone sets the timezone property value. ISO text representation of the user timezone
func (m *User) SetTimezone(value *string)() {
    m.timezone = value
}
// SetUpdatedBy sets the updated_by property value. The updated_by property
func (m *User) SetUpdatedBy(value *int32)() {
    m.updated_by = value
}
// SetUserFsRoots sets the user_fs_roots property value. The user_fs_roots property
func (m *User) SetUserFsRoots(value []User_fs_rootsable)() {
    m.user_fs_roots = value
}
// SetUserFsStats sets the user_fs_stats property value. The user_fs_stats property
func (m *User) SetUserFsStats(value User_fs_statsable)() {
    m.user_fs_stats = value
}
// SetUserRole sets the user_role property value. The user_role property
func (m *User) SetUserRole(value *int32)() {
    m.user_role = value
}
// Userable 
type Userable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedOnTs()(*int32)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetHasAvatar()(*bool)
    GetId()(*int32)
    GetInvitedBy()(*int32)
    GetIsTermsAgreed()(*bool)
    GetIsTourComplete()(*bool)
    GetLangCode()(*string)
    GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastActiveOnTs()(*int32)
    GetNotificationsConfig()(*int32)
    GetOnlineStatus()(*int32)
    GetPhone()(*string)
    GetTimezone()(*string)
    GetUpdatedBy()(*int32)
    GetUserFsRoots()([]User_fs_rootsable)
    GetUserFsStats()(User_fs_statsable)
    GetUserRole()(*int32)
    SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedOnTs(value *int32)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetHasAvatar(value *bool)()
    SetId(value *int32)()
    SetInvitedBy(value *int32)()
    SetIsTermsAgreed(value *bool)()
    SetIsTourComplete(value *bool)()
    SetLangCode(value *string)()
    SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastActiveOnTs(value *int32)()
    SetNotificationsConfig(value *int32)()
    SetOnlineStatus(value *int32)()
    SetPhone(value *string)()
    SetTimezone(value *string)()
    SetUpdatedBy(value *int32)()
    SetUserFsRoots(value []User_fs_rootsable)()
    SetUserFsStats(value User_fs_statsable)()
    SetUserRole(value *int32)()
}
