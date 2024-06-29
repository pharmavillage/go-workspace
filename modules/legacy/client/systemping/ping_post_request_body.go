package systemping

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PingPostRequestBody 
type PingPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Finger print id of the client from RTM connect.
    finger_print *string
    // Unique token for correlation when received in WS.
    ping_token *string
}
// NewPingPostRequestBody instantiates a new pingPostRequestBody and sets the default values.
func NewPingPostRequestBody()(*PingPostRequestBody) {
    m := &PingPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePingPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePingPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPingPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PingPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PingPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["finger_print"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerPrint(val)
        }
        return nil
    }
    res["ping_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPingToken(val)
        }
        return nil
    }
    return res
}
// GetFingerPrint gets the finger_print property value. Finger print id of the client from RTM connect.
func (m *PingPostRequestBody) GetFingerPrint()(*string) {
    return m.finger_print
}
// GetPingToken gets the ping_token property value. Unique token for correlation when received in WS.
func (m *PingPostRequestBody) GetPingToken()(*string) {
    return m.ping_token
}
// Serialize serializes information the current object
func (m *PingPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("finger_print", m.GetFingerPrint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ping_token", m.GetPingToken())
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
func (m *PingPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFingerPrint sets the finger_print property value. Finger print id of the client from RTM connect.
func (m *PingPostRequestBody) SetFingerPrint(value *string)() {
    m.finger_print = value
}
// SetPingToken sets the ping_token property value. Unique token for correlation when received in WS.
func (m *PingPostRequestBody) SetPingToken(value *string)() {
    m.ping_token = value
}
// PingPostRequestBodyable 
type PingPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFingerPrint()(*string)
    GetPingToken()(*string)
    SetFingerPrint(value *string)()
    SetPingToken(value *string)()
}
