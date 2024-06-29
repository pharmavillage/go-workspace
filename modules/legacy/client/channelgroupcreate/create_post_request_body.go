package channelgroupcreate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreatePostRequestBody 
type CreatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Comma separated list of channels to be included on the group
    channels *string
    // New group name
    name *string
}
// NewCreatePostRequestBody instantiates a new createPostRequestBody and sets the default values.
func NewCreatePostRequestBody()(*CreatePostRequestBody) {
    m := &CreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannels gets the channels property value. Comma separated list of channels to be included on the group
func (m *CreatePostRequestBody) GetChannels()(*string) {
    return m.channels
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannels(val)
        }
        return nil
    }
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
    return res
}
// GetName gets the name property value. New group name
func (m *CreatePostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *CreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("channels", m.GetChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *CreatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannels sets the channels property value. Comma separated list of channels to be included on the group
func (m *CreatePostRequestBody) SetChannels(value *string)() {
    m.channels = value
}
// SetName sets the name property value. New group name
func (m *CreatePostRequestBody) SetName(value *string)() {
    m.name = value
}
// CreatePostRequestBodyable 
type CreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannels()(*string)
    GetName()(*string)
    SetChannels(value *string)()
    SetName(value *string)()
}
