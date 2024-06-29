package lockrefresh

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RefreshPostRequestBody 
type RefreshPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // seconds to automatically release the lock after acquring.
    expires_after_sec *int32
    // File path to lock
    lock_id *string
}
// NewRefreshPostRequestBody instantiates a new refreshPostRequestBody and sets the default values.
func NewRefreshPostRequestBody()(*RefreshPostRequestBody) {
    m := &RefreshPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRefreshPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRefreshPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRefreshPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RefreshPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresAfterSec gets the expires_after_sec property value. seconds to automatically release the lock after acquring.
func (m *RefreshPostRequestBody) GetExpiresAfterSec()(*int32) {
    return m.expires_after_sec
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RefreshPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expires_after_sec"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAfterSec(val)
        }
        return nil
    }
    res["lock_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockId(val)
        }
        return nil
    }
    return res
}
// GetLockId gets the lock_id property value. File path to lock
func (m *RefreshPostRequestBody) GetLockId()(*string) {
    return m.lock_id
}
// Serialize serializes information the current object
func (m *RefreshPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("expires_after_sec", m.GetExpiresAfterSec())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lock_id", m.GetLockId())
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
func (m *RefreshPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresAfterSec sets the expires_after_sec property value. seconds to automatically release the lock after acquring.
func (m *RefreshPostRequestBody) SetExpiresAfterSec(value *int32)() {
    m.expires_after_sec = value
}
// SetLockId sets the lock_id property value. File path to lock
func (m *RefreshPostRequestBody) SetLockId(value *string)() {
    m.lock_id = value
}
// RefreshPostRequestBodyable 
type RefreshPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAfterSec()(*int32)
    GetLockId()(*string)
    SetExpiresAfterSec(value *int32)()
    SetLockId(value *string)()
}
