package passwordupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePostRequestBody 
type UpdatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The current_password property
    current_password *string
    // The new_password property
    new_password *string
    // The user_id property
    user_id *int32
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
// GetCurrentPassword gets the current_password property value. The current_password property
func (m *UpdatePostRequestBody) GetCurrentPassword()(*string) {
    return m.current_password
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["current_password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPassword(val)
        }
        return nil
    }
    res["new_password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
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
// GetNewPassword gets the new_password property value. The new_password property
func (m *UpdatePostRequestBody) GetNewPassword()(*string) {
    return m.new_password
}
// GetUserId gets the user_id property value. The user_id property
func (m *UpdatePostRequestBody) GetUserId()(*int32) {
    return m.user_id
}
// Serialize serializes information the current object
func (m *UpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("current_password", m.GetCurrentPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("new_password", m.GetNewPassword())
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
func (m *UpdatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentPassword sets the current_password property value. The current_password property
func (m *UpdatePostRequestBody) SetCurrentPassword(value *string)() {
    m.current_password = value
}
// SetNewPassword sets the new_password property value. The new_password property
func (m *UpdatePostRequestBody) SetNewPassword(value *string)() {
    m.new_password = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *UpdatePostRequestBody) SetUserId(value *int32)() {
    m.user_id = value
}
// UpdatePostRequestBodyable 
type UpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPassword()(*string)
    GetNewPassword()(*string)
    GetUserId()(*int32)
    SetCurrentPassword(value *string)()
    SetNewPassword(value *string)()
    SetUserId(value *int32)()
}
