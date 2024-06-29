package useralertack

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AckPostRequestBody 
type AckPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id of the alert record to mark as read
    alert_id *int32
}
// NewAckPostRequestBody instantiates a new ackPostRequestBody and sets the default values.
func NewAckPostRequestBody()(*AckPostRequestBody) {
    m := &AckPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAckPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAckPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAckPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AckPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlertId gets the alert_id property value. The id of the alert record to mark as read
func (m *AckPostRequestBody) GetAlertId()(*int32) {
    return m.alert_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AckPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alert_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AckPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("alert_id", m.GetAlertId())
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
func (m *AckPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlertId sets the alert_id property value. The id of the alert record to mark as read
func (m *AckPostRequestBody) SetAlertId(value *int32)() {
    m.alert_id = value
}
// AckPostRequestBodyable 
type AckPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertId()(*int32)
    SetAlertId(value *int32)()
}
