package channeloneonone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OneOnOnePostRequestBody 
type OneOnOnePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the other user on the channel
    with_user_id *int32
}
// NewOneOnOnePostRequestBody instantiates a new oneOnOnePostRequestBody and sets the default values.
func NewOneOnOnePostRequestBody()(*OneOnOnePostRequestBody) {
    m := &OneOnOnePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOneOnOnePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOneOnOnePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOneOnOnePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OneOnOnePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OneOnOnePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["with_user_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWithUserId(val)
        }
        return nil
    }
    return res
}
// GetWithUserId gets the with_user_id property value. Id of the other user on the channel
func (m *OneOnOnePostRequestBody) GetWithUserId()(*int32) {
    return m.with_user_id
}
// Serialize serializes information the current object
func (m *OneOnOnePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("with_user_id", m.GetWithUserId())
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
func (m *OneOnOnePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWithUserId sets the with_user_id property value. Id of the other user on the channel
func (m *OneOnOnePostRequestBody) SetWithUserId(value *int32)() {
    m.with_user_id = value
}
// OneOnOnePostRequestBodyable 
type OneOnOnePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWithUserId()(*int32)
    SetWithUserId(value *int32)()
}
