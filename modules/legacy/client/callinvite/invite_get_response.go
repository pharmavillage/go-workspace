package callinvite

import (
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InviteGetResponse 
type InviteGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The meta property
    meta i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable
}
// NewInviteGetResponse instantiates a new inviteGetResponse and sets the default values.
func NewInviteGetResponse()(*InviteGetResponse) {
    m := &InviteGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInviteGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInviteGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInviteGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InviteGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InviteGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["meta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateMetaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable))
        }
        return nil
    }
    return res
}
// GetMeta gets the meta property value. The meta property
func (m *InviteGetResponse) GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable) {
    return m.meta
}
// Serialize serializes information the current object
func (m *InviteGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("meta", m.GetMeta())
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
func (m *InviteGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMeta sets the meta property value. The meta property
func (m *InviteGetResponse) SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)() {
    m.meta = value
}
// InviteGetResponseable 
type InviteGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)
    SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)()
}