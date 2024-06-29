package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Message this is a message object.message_type is an enum.message_type=1 - Regular Message,message_type=3 - Quoted message. The parent_message will have an object as describedmessage_type=5 - Bot Message.content can be a string or JSON encoded string depending on the message_type If message_type is 5, then it is a JSON encoded string containing a type field and bot_message field when parsed. The type field can be used to render this message differently from a user generated message.Attachment is an array of JSON encoded string. Once it is parsed, This contains two fields. A ctp field and a content JSON object. The ctp can be ATTACHMENT_TYPE_FILE or ATTACHMENT_TYPE_UNFURL. Depending on the ctp value, the content JSON object will contain different data.
type Message struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attachments property
    attachments []Message_Message_attachmentsable
    // The channel_id property
    channel_id *int32
    // The channel_name property
    channel_name *string
    // The content_text property
    content_text *string
    // The created_on property
    created_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display_name property
    display_name *string
    // The emoticons property
    emoticons []Message_emoticonsable
    // The is_edited property
    is_edited *bool
    // The message_type property
    message_type *int32
    // The parent_message property
    parent_message Message_parent_messageable
    // The read_user_count property
    read_user_count *int32
    // The read_users property
    read_users []Message_read_usersable
    // Defines the source of the message
    source *Message_source
    // The user_id property
    user_id *int32
}
// Message_Message_attachments composed type wrapper for classes FileAttachment, UnfurlAttachment
type Message_Message_attachments struct {
    // Composed type representation for type FileAttachment
    fileAttachment FileAttachmentable
    // Composed type representation for type UnfurlAttachment
    unfurlAttachment UnfurlAttachmentable
}
// NewMessage_Message_attachments instantiates a new Message_attachments and sets the default values.
func NewMessage_Message_attachments()(*Message_Message_attachments) {
    m := &Message_Message_attachments{
    }
    return m
}
// CreateMessage_Message_attachmentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessage_Message_attachmentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMessage_Message_attachments()
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
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message_Message_attachments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFileAttachment gets the FileAttachment property value. Composed type representation for type FileAttachment
func (m *Message_Message_attachments) GetFileAttachment()(FileAttachmentable) {
    return m.fileAttachment
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *Message_Message_attachments) GetIsComposedType()(bool) {
    return true
}
// GetUnfurlAttachment gets the UnfurlAttachment property value. Composed type representation for type UnfurlAttachment
func (m *Message_Message_attachments) GetUnfurlAttachment()(UnfurlAttachmentable) {
    return m.unfurlAttachment
}
// Serialize serializes information the current object
func (m *Message_Message_attachments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFileAttachment() != nil {
        err := writer.WriteObjectValue("", m.GetFileAttachment())
        if err != nil {
            return err
        }
    } else if m.GetUnfurlAttachment() != nil {
        err := writer.WriteObjectValue("", m.GetUnfurlAttachment())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileAttachment sets the FileAttachment property value. Composed type representation for type FileAttachment
func (m *Message_Message_attachments) SetFileAttachment(value FileAttachmentable)() {
    m.fileAttachment = value
}
// SetUnfurlAttachment sets the UnfurlAttachment property value. Composed type representation for type UnfurlAttachment
func (m *Message_Message_attachments) SetUnfurlAttachment(value UnfurlAttachmentable)() {
    m.unfurlAttachment = value
}
// Message_Message_attachmentsable 
type Message_Message_attachmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileAttachment()(FileAttachmentable)
    GetUnfurlAttachment()(UnfurlAttachmentable)
    SetFileAttachment(value FileAttachmentable)()
    SetUnfurlAttachment(value UnfurlAttachmentable)()
}
// NewMessage instantiates a new Message and sets the default values.
func NewMessage()(*Message) {
    m := &Message{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Message) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachments gets the attachments property value. The attachments property
func (m *Message) GetAttachments()([]Message_Message_attachmentsable) {
    return m.attachments
}
// GetChannelId gets the channel_id property value. The channel_id property
func (m *Message) GetChannelId()(*int32) {
    return m.channel_id
}
// GetChannelName gets the channel_name property value. The channel_name property
func (m *Message) GetChannelName()(*string) {
    return m.channel_name
}
// GetContentText gets the content_text property value. The content_text property
func (m *Message) GetContentText()(*string) {
    return m.content_text
}
// GetCreatedOn gets the created_on property value. The created_on property
func (m *Message) GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_on
}
// GetDisplayName gets the display_name property value. The display_name property
func (m *Message) GetDisplayName()(*string) {
    return m.display_name
}
// GetEmoticons gets the emoticons property value. The emoticons property
func (m *Message) GetEmoticons()([]Message_emoticonsable) {
    return m.emoticons
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessage_Message_attachmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Message_Message_attachmentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Message_Message_attachmentsable)
                }
            }
            m.SetAttachments(res)
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
    res["content_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentText(val)
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
    res["emoticons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessage_emoticonsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Message_emoticonsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Message_emoticonsable)
                }
            }
            m.SetEmoticons(res)
        }
        return nil
    }
    res["is_edited"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEdited(val)
        }
        return nil
    }
    res["message_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageType(val)
        }
        return nil
    }
    res["parent_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessage_parent_messageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentMessage(val.(Message_parent_messageable))
        }
        return nil
    }
    res["read_user_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadUserCount(val)
        }
        return nil
    }
    res["read_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessage_read_usersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Message_read_usersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Message_read_usersable)
                }
            }
            m.SetReadUsers(res)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessage_source)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*Message_source))
        }
        return nil
    }
    res["user_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetIsEdited gets the is_edited property value. The is_edited property
func (m *Message) GetIsEdited()(*bool) {
    return m.is_edited
}
// GetMessageType gets the message_type property value. The message_type property
func (m *Message) GetMessageType()(*int32) {
    return m.message_type
}
// GetParentMessage gets the parent_message property value. The parent_message property
func (m *Message) GetParentMessage()(Message_parent_messageable) {
    return m.parent_message
}
// GetReadUserCount gets the read_user_count property value. The read_user_count property
func (m *Message) GetReadUserCount()(*int32) {
    return m.read_user_count
}
// GetReadUsers gets the read_users property value. The read_users property
func (m *Message) GetReadUsers()([]Message_read_usersable) {
    return m.read_users
}
// GetSource gets the source property value. Defines the source of the message
func (m *Message) GetSource()(*Message_source) {
    return m.source
}
// GetUserId gets the user_id property value. The user_id property
func (m *Message) GetUserId()(*int32) {
    return m.user_id
}
// Serialize serializes information the current object
func (m *Message) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttachments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attachments", cast)
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
        err := writer.WriteStringValue("content_text", m.GetContentText())
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
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEmoticons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmoticons()))
        for i, v := range m.GetEmoticons() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("emoticons", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_edited", m.GetIsEdited())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("message_type", m.GetMessageType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent_message", m.GetParentMessage())
        if err != nil {
            return err
        }
    }
    if m.GetReadUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReadUsers()))
        for i, v := range m.GetReadUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("read_users", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("read_user_count", m.GetReadUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err := writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_id", m.GetUserId())
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
func (m *Message) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachments sets the attachments property value. The attachments property
func (m *Message) SetAttachments(value []Message_Message_attachmentsable)() {
    m.attachments = value
}
// SetChannelId sets the channel_id property value. The channel_id property
func (m *Message) SetChannelId(value *int32)() {
    m.channel_id = value
}
// SetChannelName sets the channel_name property value. The channel_name property
func (m *Message) SetChannelName(value *string)() {
    m.channel_name = value
}
// SetContentText sets the content_text property value. The content_text property
func (m *Message) SetContentText(value *string)() {
    m.content_text = value
}
// SetCreatedOn sets the created_on property value. The created_on property
func (m *Message) SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_on = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *Message) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetEmoticons sets the emoticons property value. The emoticons property
func (m *Message) SetEmoticons(value []Message_emoticonsable)() {
    m.emoticons = value
}
// SetIsEdited sets the is_edited property value. The is_edited property
func (m *Message) SetIsEdited(value *bool)() {
    m.is_edited = value
}
// SetMessageType sets the message_type property value. The message_type property
func (m *Message) SetMessageType(value *int32)() {
    m.message_type = value
}
// SetParentMessage sets the parent_message property value. The parent_message property
func (m *Message) SetParentMessage(value Message_parent_messageable)() {
    m.parent_message = value
}
// SetReadUserCount sets the read_user_count property value. The read_user_count property
func (m *Message) SetReadUserCount(value *int32)() {
    m.read_user_count = value
}
// SetReadUsers sets the read_users property value. The read_users property
func (m *Message) SetReadUsers(value []Message_read_usersable)() {
    m.read_users = value
}
// SetSource sets the source property value. Defines the source of the message
func (m *Message) SetSource(value *Message_source)() {
    m.source = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *Message) SetUserId(value *int32)() {
    m.user_id = value
}
// Messageable 
type Messageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()([]Message_Message_attachmentsable)
    GetChannelId()(*int32)
    GetChannelName()(*string)
    GetContentText()(*string)
    GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetEmoticons()([]Message_emoticonsable)
    GetIsEdited()(*bool)
    GetMessageType()(*int32)
    GetParentMessage()(Message_parent_messageable)
    GetReadUserCount()(*int32)
    GetReadUsers()([]Message_read_usersable)
    GetSource()(*Message_source)
    GetUserId()(*int32)
    SetAttachments(value []Message_Message_attachmentsable)()
    SetChannelId(value *int32)()
    SetChannelName(value *string)()
    SetContentText(value *string)()
    SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetEmoticons(value []Message_emoticonsable)()
    SetIsEdited(value *bool)()
    SetMessageType(value *int32)()
    SetParentMessage(value Message_parent_messageable)()
    SetReadUserCount(value *int32)()
    SetReadUsers(value []Message_read_usersable)()
    SetSource(value *Message_source)()
    SetUserId(value *int32)()
}
