package passwordreset

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResetPostRequestBody 
type ResetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The password property
    password *string
    // The reset_code property
    reset_code *string
    // The user_id property
    user_id *int32
}
// NewResetPostRequestBody instantiates a new resetPostRequestBody and sets the default values.
func NewResetPostRequestBody()(*ResetPostRequestBody) {
    m := &ResetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResetPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["reset_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetCode(val)
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
// GetPassword gets the password property value. The password property
func (m *ResetPostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetResetCode gets the reset_code property value. The reset_code property
func (m *ResetPostRequestBody) GetResetCode()(*string) {
    return m.reset_code
}
// GetUserId gets the user_id property value. The user_id property
func (m *ResetPostRequestBody) GetUserId()(*int32) {
    return m.user_id
}
// Serialize serializes information the current object
func (m *ResetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reset_code", m.GetResetCode())
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
func (m *ResetPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPassword sets the password property value. The password property
func (m *ResetPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetResetCode sets the reset_code property value. The reset_code property
func (m *ResetPostRequestBody) SetResetCode(value *string)() {
    m.reset_code = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *ResetPostRequestBody) SetUserId(value *int32)() {
    m.user_id = value
}
// ResetPostRequestBodyable 
type ResetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    GetResetCode()(*string)
    GetUserId()(*int32)
    SetPassword(value *string)()
    SetResetCode(value *string)()
    SetUserId(value *int32)()
}
