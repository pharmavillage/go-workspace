package lockrelease

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReleasePostRequestBody 
type ReleasePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // File path to lock
    fspath *string
}
// NewReleasePostRequestBody instantiates a new releasePostRequestBody and sets the default values.
func NewReleasePostRequestBody()(*ReleasePostRequestBody) {
    m := &ReleasePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReleasePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReleasePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReleasePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReleasePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReleasePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ReleasePostRequestBody) GetFspath()(*string) {
    return m.fspath
}
// Serialize serializes information the current object
func (m *ReleasePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReleasePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFspath sets the fspath property value. File path to lock
func (m *ReleasePostRequestBody) SetFspath(value *string)() {
    m.fspath = value
}
// ReleasePostRequestBodyable 
type ReleasePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFspath()(*string)
    SetFspath(value *string)()
}
