package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelRoot 
type ChannelRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The location property
    location *string
    // the type of channel root, "files" or "wiki"
    typeEscaped *string
}
// NewChannelRoot instantiates a new ChannelRoot and sets the default values.
func NewChannelRoot()(*ChannelRoot) {
    m := &ChannelRoot{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChannelRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelRoot(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelRoot) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. The location property
func (m *ChannelRoot) GetLocation()(*string) {
    return m.location
}
// GetTypeEscaped gets the type property value. the type of channel root, "files" or "wiki"
func (m *ChannelRoot) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ChannelRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *ChannelRoot) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocation sets the location property value. The location property
func (m *ChannelRoot) SetLocation(value *string)() {
    m.location = value
}
// SetTypeEscaped sets the type property value. the type of channel root, "files" or "wiki"
func (m *ChannelRoot) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// ChannelRootable 
type ChannelRootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(*string)
    GetTypeEscaped()(*string)
    SetLocation(value *string)()
    SetTypeEscaped(value *string)()
}
