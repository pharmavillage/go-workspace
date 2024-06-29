package rtmconnect

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectGetResponse_rtm 
type ConnectGetResponse_rtm struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ws_endpoint property
    ws_endpoint *string
    // The ws_token property
    ws_token *string
}
// NewConnectGetResponse_rtm instantiates a new connectGetResponse_rtm and sets the default values.
func NewConnectGetResponse_rtm()(*ConnectGetResponse_rtm) {
    m := &ConnectGetResponse_rtm{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConnectGetResponse_rtmFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectGetResponse_rtmFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectGetResponse_rtm(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectGetResponse_rtm) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectGetResponse_rtm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ws_endpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWsEndpoint(val)
        }
        return nil
    }
    res["ws_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWsToken(val)
        }
        return nil
    }
    return res
}
// GetWsEndpoint gets the ws_endpoint property value. The ws_endpoint property
func (m *ConnectGetResponse_rtm) GetWsEndpoint()(*string) {
    return m.ws_endpoint
}
// GetWsToken gets the ws_token property value. The ws_token property
func (m *ConnectGetResponse_rtm) GetWsToken()(*string) {
    return m.ws_token
}
// Serialize serializes information the current object
func (m *ConnectGetResponse_rtm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ws_endpoint", m.GetWsEndpoint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ws_token", m.GetWsToken())
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
func (m *ConnectGetResponse_rtm) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWsEndpoint sets the ws_endpoint property value. The ws_endpoint property
func (m *ConnectGetResponse_rtm) SetWsEndpoint(value *string)() {
    m.ws_endpoint = value
}
// SetWsToken sets the ws_token property value. The ws_token property
func (m *ConnectGetResponse_rtm) SetWsToken(value *string)() {
    m.ws_token = value
}
// ConnectGetResponse_rtmable 
type ConnectGetResponse_rtmable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWsEndpoint()(*string)
    GetWsToken()(*string)
    SetWsEndpoint(value *string)()
    SetWsToken(value *string)()
}
