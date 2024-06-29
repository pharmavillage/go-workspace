package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Channel 
type Channel struct {
    // Total number of actions in this channel
    action_count *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The blurb property
    blurb *string
    // Defines the group of this channel for the logged user
    channel_group_id *int32
    // The channel_name property
    channel_name *string
    // The channel_roots property
    channel_roots []ChannelRootable
    // 1 to indicate channel is OPEN and 2 to indicate it is CLOSED
    channel_status *int32
    // Automatically close the channel after these many days.
    close_after_days *int32
    // User id of the channel creator
    created_by *int32
    // The created_on property
    created_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates if a custom background is setup for this channel
    has_background *bool
    // Indicates if a custom logo is setup for this channel
    has_logo *bool
    // The id property
    id *int32
    // Enable if this channel setup for auto closing
    is_auto_closed *bool
    // Indicates if this user has favorited this channel
    is_favorite *bool
    // The last_active_on property
    last_active_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // This is a message object.message_type is an enum.message_type=1 - Regular Message,message_type=3 - Quoted message. The parent_message will have an object as describedmessage_type=5 - Bot Message.content can be a string or JSON encoded string depending on the message_type If message_type is 5, then it is a JSON encoded string containing a type field and bot_message field when parsed. The type field can be used to render this message differently from a user generated message.Attachment is an array of JSON encoded string. Once it is parsed, This contains two fields. A ctp field and a content JSON object. The ctp can be ATTACHMENT_TYPE_FILE or ATTACHMENT_TYPE_UNFURL. Depending on the ctp value, the content JSON object will contain different data.
    latest Messageable
    // Email generated for each member of the channel, that allows the member to send messages through email
    member_email *string
    // The members property
    members []User_Abbrable
    // The oldest unread message for requesting user. This can be 0 if there are no message in the channel
    oldest_unread_message_id *int32
    // Defines if the channel is a 1x1 channel (don't accept user invites)
    one_one *bool
    // The url that gives public access to the channel. Not available when the channel is private.
    public_url *string
    // The team_id property
    team_id *int32
    // total number of files in a channel
    total_file_count *int32
    // total number of folders in a channel
    total_folder_count *int32
    // total files and folders in a channel
    total_fs_count *int32
    // total size of files in a channel in bytes
    total_fs_size *int32
    // Total number of messages not read by this user in this channel
    unread_count *int32
    // The updated_by property
    updated_by *int32
    // The updated_on property
    updated_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
    user_role *int32
}
// NewChannel instantiates a new Channel and sets the default values.
func NewChannel()(*Channel) {
    m := &Channel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChannelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannel(), nil
}
// GetActionCount gets the action_count property value. Total number of actions in this channel
func (m *Channel) GetActionCount()(*int32) {
    return m.action_count
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Channel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlurb gets the blurb property value. The blurb property
func (m *Channel) GetBlurb()(*string) {
    return m.blurb
}
// GetChannelGroupId gets the channel_group_id property value. Defines the group of this channel for the logged user
func (m *Channel) GetChannelGroupId()(*int32) {
    return m.channel_group_id
}
// GetChannelName gets the channel_name property value. The channel_name property
func (m *Channel) GetChannelName()(*string) {
    return m.channel_name
}
// GetChannelRoots gets the channel_roots property value. The channel_roots property
func (m *Channel) GetChannelRoots()([]ChannelRootable) {
    return m.channel_roots
}
// GetChannelStatus gets the channel_status property value. 1 to indicate channel is OPEN and 2 to indicate it is CLOSED
func (m *Channel) GetChannelStatus()(*int32) {
    return m.channel_status
}
// GetCloseAfterDays gets the close_after_days property value. Automatically close the channel after these many days.
func (m *Channel) GetCloseAfterDays()(*int32) {
    return m.close_after_days
}
// GetCreatedBy gets the created_by property value. User id of the channel creator
func (m *Channel) GetCreatedBy()(*int32) {
    return m.created_by
}
// GetCreatedOn gets the created_on property value. The created_on property
func (m *Channel) GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_on
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Channel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionCount(val)
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
    res["channel_group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelGroupId(val)
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
    res["channel_roots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChannelRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChannelRootable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ChannelRootable)
                }
            }
            m.SetChannelRoots(res)
        }
        return nil
    }
    res["channel_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelStatus(val)
        }
        return nil
    }
    res["close_after_days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloseAfterDays(val)
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
    res["has_background"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasBackground(val)
        }
        return nil
    }
    res["has_logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasLogo(val)
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
    res["is_auto_closed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutoClosed(val)
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
    res["latest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatest(val.(Messageable))
        }
        return nil
    }
    res["member_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEmail(val)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUser_AbbrFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]User_Abbrable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(User_Abbrable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["oldest_unread_message_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldestUnreadMessageId(val)
        }
        return nil
    }
    res["one_one"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneOne(val)
        }
        return nil
    }
    res["public_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicUrl(val)
        }
        return nil
    }
    res["team_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    res["total_file_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFileCount(val)
        }
        return nil
    }
    res["total_folder_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFolderCount(val)
        }
        return nil
    }
    res["total_fs_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFsCount(val)
        }
        return nil
    }
    res["total_fs_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFsSize(val)
        }
        return nil
    }
    res["unread_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnreadCount(val)
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
// GetHasBackground gets the has_background property value. Indicates if a custom background is setup for this channel
func (m *Channel) GetHasBackground()(*bool) {
    return m.has_background
}
// GetHasLogo gets the has_logo property value. Indicates if a custom logo is setup for this channel
func (m *Channel) GetHasLogo()(*bool) {
    return m.has_logo
}
// GetId gets the id property value. The id property
func (m *Channel) GetId()(*int32) {
    return m.id
}
// GetIsAutoClosed gets the is_auto_closed property value. Enable if this channel setup for auto closing
func (m *Channel) GetIsAutoClosed()(*bool) {
    return m.is_auto_closed
}
// GetIsFavorite gets the is_favorite property value. Indicates if this user has favorited this channel
func (m *Channel) GetIsFavorite()(*bool) {
    return m.is_favorite
}
// GetLastActiveOn gets the last_active_on property value. The last_active_on property
func (m *Channel) GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_active_on
}
// GetLatest gets the latest property value. This is a message object.message_type is an enum.message_type=1 - Regular Message,message_type=3 - Quoted message. The parent_message will have an object as describedmessage_type=5 - Bot Message.content can be a string or JSON encoded string depending on the message_type If message_type is 5, then it is a JSON encoded string containing a type field and bot_message field when parsed. The type field can be used to render this message differently from a user generated message.Attachment is an array of JSON encoded string. Once it is parsed, This contains two fields. A ctp field and a content JSON object. The ctp can be ATTACHMENT_TYPE_FILE or ATTACHMENT_TYPE_UNFURL. Depending on the ctp value, the content JSON object will contain different data.
func (m *Channel) GetLatest()(Messageable) {
    return m.latest
}
// GetMemberEmail gets the member_email property value. Email generated for each member of the channel, that allows the member to send messages through email
func (m *Channel) GetMemberEmail()(*string) {
    return m.member_email
}
// GetMembers gets the members property value. The members property
func (m *Channel) GetMembers()([]User_Abbrable) {
    return m.members
}
// GetOldestUnreadMessageId gets the oldest_unread_message_id property value. The oldest unread message for requesting user. This can be 0 if there are no message in the channel
func (m *Channel) GetOldestUnreadMessageId()(*int32) {
    return m.oldest_unread_message_id
}
// GetOneOne gets the one_one property value. Defines if the channel is a 1x1 channel (don't accept user invites)
func (m *Channel) GetOneOne()(*bool) {
    return m.one_one
}
// GetPublicUrl gets the public_url property value. The url that gives public access to the channel. Not available when the channel is private.
func (m *Channel) GetPublicUrl()(*string) {
    return m.public_url
}
// GetTeamId gets the team_id property value. The team_id property
func (m *Channel) GetTeamId()(*int32) {
    return m.team_id
}
// GetTotalFileCount gets the total_file_count property value. total number of files in a channel
func (m *Channel) GetTotalFileCount()(*int32) {
    return m.total_file_count
}
// GetTotalFolderCount gets the total_folder_count property value. total number of folders in a channel
func (m *Channel) GetTotalFolderCount()(*int32) {
    return m.total_folder_count
}
// GetTotalFsCount gets the total_fs_count property value. total files and folders in a channel
func (m *Channel) GetTotalFsCount()(*int32) {
    return m.total_fs_count
}
// GetTotalFsSize gets the total_fs_size property value. total size of files in a channel in bytes
func (m *Channel) GetTotalFsSize()(*int32) {
    return m.total_fs_size
}
// GetUnreadCount gets the unread_count property value. Total number of messages not read by this user in this channel
func (m *Channel) GetUnreadCount()(*int32) {
    return m.unread_count
}
// GetUpdatedBy gets the updated_by property value. The updated_by property
func (m *Channel) GetUpdatedBy()(*int32) {
    return m.updated_by
}
// GetUpdatedOn gets the updated_on property value. The updated_on property
func (m *Channel) GetUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_on
}
// GetUserRole gets the user_role property value. 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *Channel) GetUserRole()(*int32) {
    return m.user_role
}
// Serialize serializes information the current object
func (m *Channel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("action_count", m.GetActionCount())
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
        err := writer.WriteInt32Value("channel_group_id", m.GetChannelGroupId())
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
    if m.GetChannelRoots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChannelRoots()))
        for i, v := range m.GetChannelRoots() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("channel_roots", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("channel_status", m.GetChannelStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("close_after_days", m.GetCloseAfterDays())
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
        err := writer.WriteBoolValue("has_background", m.GetHasBackground())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_logo", m.GetHasLogo())
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
        err := writer.WriteBoolValue("is_auto_closed", m.GetIsAutoClosed())
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
        err := writer.WriteTimeValue("last_active_on", m.GetLastActiveOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("latest", m.GetLatest())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("member_email", m.GetMemberEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("oldest_unread_message_id", m.GetOldestUnreadMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("one_one", m.GetOneOne())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_url", m.GetPublicUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("team_id", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_file_count", m.GetTotalFileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_folder_count", m.GetTotalFolderCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_fs_count", m.GetTotalFsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_fs_size", m.GetTotalFsSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unread_count", m.GetUnreadCount())
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
// SetActionCount sets the action_count property value. Total number of actions in this channel
func (m *Channel) SetActionCount(value *int32)() {
    m.action_count = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Channel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlurb sets the blurb property value. The blurb property
func (m *Channel) SetBlurb(value *string)() {
    m.blurb = value
}
// SetChannelGroupId sets the channel_group_id property value. Defines the group of this channel for the logged user
func (m *Channel) SetChannelGroupId(value *int32)() {
    m.channel_group_id = value
}
// SetChannelName sets the channel_name property value. The channel_name property
func (m *Channel) SetChannelName(value *string)() {
    m.channel_name = value
}
// SetChannelRoots sets the channel_roots property value. The channel_roots property
func (m *Channel) SetChannelRoots(value []ChannelRootable)() {
    m.channel_roots = value
}
// SetChannelStatus sets the channel_status property value. 1 to indicate channel is OPEN and 2 to indicate it is CLOSED
func (m *Channel) SetChannelStatus(value *int32)() {
    m.channel_status = value
}
// SetCloseAfterDays sets the close_after_days property value. Automatically close the channel after these many days.
func (m *Channel) SetCloseAfterDays(value *int32)() {
    m.close_after_days = value
}
// SetCreatedBy sets the created_by property value. User id of the channel creator
func (m *Channel) SetCreatedBy(value *int32)() {
    m.created_by = value
}
// SetCreatedOn sets the created_on property value. The created_on property
func (m *Channel) SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_on = value
}
// SetHasBackground sets the has_background property value. Indicates if a custom background is setup for this channel
func (m *Channel) SetHasBackground(value *bool)() {
    m.has_background = value
}
// SetHasLogo sets the has_logo property value. Indicates if a custom logo is setup for this channel
func (m *Channel) SetHasLogo(value *bool)() {
    m.has_logo = value
}
// SetId sets the id property value. The id property
func (m *Channel) SetId(value *int32)() {
    m.id = value
}
// SetIsAutoClosed sets the is_auto_closed property value. Enable if this channel setup for auto closing
func (m *Channel) SetIsAutoClosed(value *bool)() {
    m.is_auto_closed = value
}
// SetIsFavorite sets the is_favorite property value. Indicates if this user has favorited this channel
func (m *Channel) SetIsFavorite(value *bool)() {
    m.is_favorite = value
}
// SetLastActiveOn sets the last_active_on property value. The last_active_on property
func (m *Channel) SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_active_on = value
}
// SetLatest sets the latest property value. This is a message object.message_type is an enum.message_type=1 - Regular Message,message_type=3 - Quoted message. The parent_message will have an object as describedmessage_type=5 - Bot Message.content can be a string or JSON encoded string depending on the message_type If message_type is 5, then it is a JSON encoded string containing a type field and bot_message field when parsed. The type field can be used to render this message differently from a user generated message.Attachment is an array of JSON encoded string. Once it is parsed, This contains two fields. A ctp field and a content JSON object. The ctp can be ATTACHMENT_TYPE_FILE or ATTACHMENT_TYPE_UNFURL. Depending on the ctp value, the content JSON object will contain different data.
func (m *Channel) SetLatest(value Messageable)() {
    m.latest = value
}
// SetMemberEmail sets the member_email property value. Email generated for each member of the channel, that allows the member to send messages through email
func (m *Channel) SetMemberEmail(value *string)() {
    m.member_email = value
}
// SetMembers sets the members property value. The members property
func (m *Channel) SetMembers(value []User_Abbrable)() {
    m.members = value
}
// SetOldestUnreadMessageId sets the oldest_unread_message_id property value. The oldest unread message for requesting user. This can be 0 if there are no message in the channel
func (m *Channel) SetOldestUnreadMessageId(value *int32)() {
    m.oldest_unread_message_id = value
}
// SetOneOne sets the one_one property value. Defines if the channel is a 1x1 channel (don't accept user invites)
func (m *Channel) SetOneOne(value *bool)() {
    m.one_one = value
}
// SetPublicUrl sets the public_url property value. The url that gives public access to the channel. Not available when the channel is private.
func (m *Channel) SetPublicUrl(value *string)() {
    m.public_url = value
}
// SetTeamId sets the team_id property value. The team_id property
func (m *Channel) SetTeamId(value *int32)() {
    m.team_id = value
}
// SetTotalFileCount sets the total_file_count property value. total number of files in a channel
func (m *Channel) SetTotalFileCount(value *int32)() {
    m.total_file_count = value
}
// SetTotalFolderCount sets the total_folder_count property value. total number of folders in a channel
func (m *Channel) SetTotalFolderCount(value *int32)() {
    m.total_folder_count = value
}
// SetTotalFsCount sets the total_fs_count property value. total files and folders in a channel
func (m *Channel) SetTotalFsCount(value *int32)() {
    m.total_fs_count = value
}
// SetTotalFsSize sets the total_fs_size property value. total size of files in a channel in bytes
func (m *Channel) SetTotalFsSize(value *int32)() {
    m.total_fs_size = value
}
// SetUnreadCount sets the unread_count property value. Total number of messages not read by this user in this channel
func (m *Channel) SetUnreadCount(value *int32)() {
    m.unread_count = value
}
// SetUpdatedBy sets the updated_by property value. The updated_by property
func (m *Channel) SetUpdatedBy(value *int32)() {
    m.updated_by = value
}
// SetUpdatedOn sets the updated_on property value. The updated_on property
func (m *Channel) SetUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_on = value
}
// SetUserRole sets the user_role property value. 100 - Owner, 50 - Manager, 20 - Collaborator, 10 - Viewer
func (m *Channel) SetUserRole(value *int32)() {
    m.user_role = value
}
// Channelable 
type Channelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionCount()(*int32)
    GetBlurb()(*string)
    GetChannelGroupId()(*int32)
    GetChannelName()(*string)
    GetChannelRoots()([]ChannelRootable)
    GetChannelStatus()(*int32)
    GetCloseAfterDays()(*int32)
    GetCreatedBy()(*int32)
    GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHasBackground()(*bool)
    GetHasLogo()(*bool)
    GetId()(*int32)
    GetIsAutoClosed()(*bool)
    GetIsFavorite()(*bool)
    GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLatest()(Messageable)
    GetMemberEmail()(*string)
    GetMembers()([]User_Abbrable)
    GetOldestUnreadMessageId()(*int32)
    GetOneOne()(*bool)
    GetPublicUrl()(*string)
    GetTeamId()(*int32)
    GetTotalFileCount()(*int32)
    GetTotalFolderCount()(*int32)
    GetTotalFsCount()(*int32)
    GetTotalFsSize()(*int32)
    GetUnreadCount()(*int32)
    GetUpdatedBy()(*int32)
    GetUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserRole()(*int32)
    SetActionCount(value *int32)()
    SetBlurb(value *string)()
    SetChannelGroupId(value *int32)()
    SetChannelName(value *string)()
    SetChannelRoots(value []ChannelRootable)()
    SetChannelStatus(value *int32)()
    SetCloseAfterDays(value *int32)()
    SetCreatedBy(value *int32)()
    SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHasBackground(value *bool)()
    SetHasLogo(value *bool)()
    SetId(value *int32)()
    SetIsAutoClosed(value *bool)()
    SetIsFavorite(value *bool)()
    SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLatest(value Messageable)()
    SetMemberEmail(value *string)()
    SetMembers(value []User_Abbrable)()
    SetOldestUnreadMessageId(value *int32)()
    SetOneOne(value *bool)()
    SetPublicUrl(value *string)()
    SetTeamId(value *int32)()
    SetTotalFileCount(value *int32)()
    SetTotalFolderCount(value *int32)()
    SetTotalFsCount(value *int32)()
    SetTotalFsSize(value *int32)()
    SetUnreadCount(value *int32)()
    SetUpdatedBy(value *int32)()
    SetUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserRole(value *int32)()
}
