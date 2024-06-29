package systeminfo

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InfoGetResponse_info 
type InfoGetResponse_info struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The version property
    version *string
}
// NewInfoGetResponse_info instantiates a new infoGetResponse_info and sets the default values.
func NewInfoGetResponse_info()(*InfoGetResponse_info) {
    m := &InfoGetResponse_info{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInfoGetResponse_infoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInfoGetResponse_infoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInfoGetResponse_info(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InfoGetResponse_info) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InfoGetResponse_info) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *InfoGetResponse_info) GetName()(*string) {
    return m.name
}
// GetVersion gets the version property value. The version property
func (m *InfoGetResponse_info) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *InfoGetResponse_info) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *InfoGetResponse_info) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *InfoGetResponse_info) SetName(value *string)() {
    m.name = value
}
// SetVersion sets the version property value. The version property
func (m *InfoGetResponse_info) SetVersion(value *string)() {
    m.version = value
}
// InfoGetResponse_infoable 
type InfoGetResponse_infoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetVersion()(*string)
    SetName(value *string)()
    SetVersion(value *string)()
}
