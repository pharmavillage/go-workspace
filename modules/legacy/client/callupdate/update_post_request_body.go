package callupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePostRequestBody 
type UpdatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Comma seperated ids of users allowed to join the call
    allowed_users *string
    // hash of the previously created call
    call_hash *string
    // Convert to a public call to allow anyone to join
    is_public *bool
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
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedUsers gets the allowed_users property value. Comma seperated ids of users allowed to join the call
func (m *UpdatePostRequestBody) GetAllowedUsers()(*string) {
    return m.allowed_users
}
// GetCallHash gets the call_hash property value. hash of the previously created call
func (m *UpdatePostRequestBody) GetCallHash()(*string) {
    return m.call_hash
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedUsers(val)
        }
        return nil
    }
    res["call_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallHash(val)
        }
        return nil
    }
    res["is_public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublic(val)
        }
        return nil
    }
    return res
}
// GetIsPublic gets the is_public property value. Convert to a public call to allow anyone to join
func (m *UpdatePostRequestBody) GetIsPublic()(*bool) {
    return m.is_public
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("allowed_users", m.GetAllowedUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("call_hash", m.GetCallHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_public", m.GetIsPublic())
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
func (m *UpdatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedUsers sets the allowed_users property value. Comma seperated ids of users allowed to join the call
func (m *UpdatePostRequestBody) SetAllowedUsers(value *string)() {
    m.allowed_users = value
}
// SetCallHash sets the call_hash property value. hash of the previously created call
func (m *UpdatePostRequestBody) SetCallHash(value *string)() {
    m.call_hash = value
}
// SetIsPublic sets the is_public property value. Convert to a public call to allow anyone to join
func (m *UpdatePostRequestBody) SetIsPublic(value *bool)() {
    m.is_public = value
}
// UpdatePostRequestBodyable 
type UpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedUsers()(*string)
    GetCallHash()(*string)
    GetIsPublic()(*bool)
    SetAllowedUsers(value *string)()
    SetCallHash(value *string)()
    SetIsPublic(value *bool)()
}
