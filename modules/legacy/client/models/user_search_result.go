package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSearchResult the search result for an user
type UserSearchResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display_name property
    display_name *string
    // The email property
    email *string
    // The has_avatar property
    has_avatar *bool
    // The user_id property
    user_id *int32
}
// NewUserSearchResult instantiates a new UserSearchResult and sets the default values.
func NewUserSearchResult()(*UserSearchResult) {
    m := &UserSearchResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserSearchResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSearchResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSearchResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSearchResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the display_name property value. The display_name property
func (m *UserSearchResult) GetDisplayName()(*string) {
    return m.display_name
}
// GetEmail gets the email property value. The email property
func (m *UserSearchResult) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSearchResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetHasAvatar gets the has_avatar property value. The has_avatar property
func (m *UserSearchResult) GetHasAvatar()(*bool) {
    return m.has_avatar
}
// GetUserId gets the user_id property value. The user_id property
func (m *UserSearchResult) GetUserId()(*int32) {
    return m.user_id
}
// Serialize serializes information the current object
func (m *UserSearchResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserSearchResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *UserSearchResult) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetEmail sets the email property value. The email property
func (m *UserSearchResult) SetEmail(value *string)() {
    m.email = value
}
// SetHasAvatar sets the has_avatar property value. The has_avatar property
func (m *UserSearchResult) SetHasAvatar(value *bool)() {
    m.has_avatar = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *UserSearchResult) SetUserId(value *int32)() {
    m.user_id = value
}
// UserSearchResultable 
type UserSearchResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetHasAvatar()(*bool)
    GetUserId()(*int32)
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetHasAvatar(value *bool)()
    SetUserId(value *int32)()
}
