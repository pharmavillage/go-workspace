package callend

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EndPostRequestBody 
type EndPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // hash of the previously created call
    call_hash *string
    // Rtm token of the  call creator
    rtm_token *string
}
// NewEndPostRequestBody instantiates a new endPostRequestBody and sets the default values.
func NewEndPostRequestBody()(*EndPostRequestBody) {
    m := &EndPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEndPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EndPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCallHash gets the call_hash property value. hash of the previously created call
func (m *EndPostRequestBody) GetCallHash()(*string) {
    return m.call_hash
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EndPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["call_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallHash(val)
        }
        return nil
    }
    res["rtm_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRtmToken(val)
        }
        return nil
    }
    return res
}
// GetRtmToken gets the rtm_token property value. Rtm token of the  call creator
func (m *EndPostRequestBody) GetRtmToken()(*string) {
    return m.rtm_token
}
// Serialize serializes information the current object
func (m *EndPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("call_hash", m.GetCallHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rtm_token", m.GetRtmToken())
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
func (m *EndPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCallHash sets the call_hash property value. hash of the previously created call
func (m *EndPostRequestBody) SetCallHash(value *string)() {
    m.call_hash = value
}
// SetRtmToken sets the rtm_token property value. Rtm token of the  call creator
func (m *EndPostRequestBody) SetRtmToken(value *string)() {
    m.rtm_token = value
}
// EndPostRequestBodyable 
type EndPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallHash()(*string)
    GetRtmToken()(*string)
    SetCallHash(value *string)()
    SetRtmToken(value *string)()
}
