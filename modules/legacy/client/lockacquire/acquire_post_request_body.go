package lockacquire

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AcquirePostRequestBody 
type AcquirePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional - seconds to automatically release the lock after acquring.
    expires_after_sec *int32
    // File path to lock
    fspath *string
}
// NewAcquirePostRequestBody instantiates a new acquirePostRequestBody and sets the default values.
func NewAcquirePostRequestBody()(*AcquirePostRequestBody) {
    m := &AcquirePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAcquirePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcquirePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcquirePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AcquirePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresAfterSec gets the expires_after_sec property value. Optional - seconds to automatically release the lock after acquring.
func (m *AcquirePostRequestBody) GetExpiresAfterSec()(*int32) {
    return m.expires_after_sec
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AcquirePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["fspath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFspath(val)
        }
        return nil
    }
    return res
}
// GetFspath gets the fspath property value. File path to lock
func (m *AcquirePostRequestBody) GetFspath()(*string) {
    return m.fspath
}
// Serialize serializes information the current object
func (m *AcquirePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("expires_after_sec", m.GetExpiresAfterSec())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fspath", m.GetFspath())
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
func (m *AcquirePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresAfterSec sets the expires_after_sec property value. Optional - seconds to automatically release the lock after acquring.
func (m *AcquirePostRequestBody) SetExpiresAfterSec(value *int32)() {
    m.expires_after_sec = value
}
// SetFspath sets the fspath property value. File path to lock
func (m *AcquirePostRequestBody) SetFspath(value *string)() {
    m.fspath = value
}
// AcquirePostRequestBodyable 
type AcquirePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAfterSec()(*int32)
    GetFspath()(*string)
    SetExpiresAfterSec(value *int32)()
    SetFspath(value *string)()
}
