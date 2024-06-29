package filecopy

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopyPostRequestBody 
type CopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // path to location to copy fromExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
    fsfrompath *string
    // path to location to copy toExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
    fstopath *string
}
// NewCopyPostRequestBody instantiates a new copyPostRequestBody and sets the default values.
func NewCopyPostRequestBody()(*CopyPostRequestBody) {
    m := &CopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopyPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fsfrompath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFsfrompath(val)
        }
        return nil
    }
    res["fstopath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFstopath(val)
        }
        return nil
    }
    return res
}
// GetFsfrompath gets the fsfrompath property value. path to location to copy fromExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
func (m *CopyPostRequestBody) GetFsfrompath()(*string) {
    return m.fsfrompath
}
// GetFstopath gets the fstopath property value. path to location to copy toExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
func (m *CopyPostRequestBody) GetFstopath()(*string) {
    return m.fstopath
}
// Serialize serializes information the current object
func (m *CopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fsfrompath", m.GetFsfrompath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fstopath", m.GetFstopath())
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
func (m *CopyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFsfrompath sets the fsfrompath property value. path to location to copy fromExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
func (m *CopyPostRequestBody) SetFsfrompath(value *string)() {
    m.fsfrompath = value
}
// SetFstopath sets the fstopath property value. path to location to copy toExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
func (m *CopyPostRequestBody) SetFstopath(value *string)() {
    m.fstopath = value
}
// CopyPostRequestBodyable 
type CopyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFsfrompath()(*string)
    GetFstopath()(*string)
    SetFsfrompath(value *string)()
    SetFstopath(value *string)()
}
