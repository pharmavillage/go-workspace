package actionmove

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MovePostRequestBody 
type MovePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The moved action will be placed after the action defined by this id. If not provided, the moved action will be at the first position.
    after *int32
    // The moved action will be placed under the action defined by this id (as a child) Shouldn't be provided if the action is been moved to the root (no parent).
    under *int32
}
// NewMovePostRequestBody instantiates a new movePostRequestBody and sets the default values.
func NewMovePostRequestBody()(*MovePostRequestBody) {
    m := &MovePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMovePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMovePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMovePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MovePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfter gets the after property value. The moved action will be placed after the action defined by this id. If not provided, the moved action will be at the first position.
func (m *MovePostRequestBody) GetAfter()(*int32) {
    return m.after
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MovePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["after"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfter(val)
        }
        return nil
    }
    res["under"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnder(val)
        }
        return nil
    }
    return res
}
// GetUnder gets the under property value. The moved action will be placed under the action defined by this id (as a child) Shouldn't be provided if the action is been moved to the root (no parent).
func (m *MovePostRequestBody) GetUnder()(*int32) {
    return m.under
}
// Serialize serializes information the current object
func (m *MovePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("after", m.GetAfter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("under", m.GetUnder())
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
func (m *MovePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfter sets the after property value. The moved action will be placed after the action defined by this id. If not provided, the moved action will be at the first position.
func (m *MovePostRequestBody) SetAfter(value *int32)() {
    m.after = value
}
// SetUnder sets the under property value. The moved action will be placed under the action defined by this id (as a child) Shouldn't be provided if the action is been moved to the root (no parent).
func (m *MovePostRequestBody) SetUnder(value *int32)() {
    m.under = value
}
// MovePostRequestBodyable 
type MovePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAfter()(*int32)
    GetUnder()(*int32)
    SetAfter(value *int32)()
    SetUnder(value *int32)()
}
