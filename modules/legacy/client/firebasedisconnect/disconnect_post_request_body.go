package firebasedisconnect

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DisconnectPostRequestBody 
type DisconnectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The token to be removed.
    token *string
}
// NewDisconnectPostRequestBody instantiates a new disconnectPostRequestBody and sets the default values.
func NewDisconnectPostRequestBody()(*DisconnectPostRequestBody) {
    m := &DisconnectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDisconnectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDisconnectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDisconnectPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DisconnectPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DisconnectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetToken gets the token property value. The token to be removed.
func (m *DisconnectPostRequestBody) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *DisconnectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *DisconnectPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetToken sets the token property value. The token to be removed.
func (m *DisconnectPostRequestBody) SetToken(value *string)() {
    m.token = value
}
// DisconnectPostRequestBodyable 
type DisconnectPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetToken()(*string)
    SetToken(value *string)()
}
