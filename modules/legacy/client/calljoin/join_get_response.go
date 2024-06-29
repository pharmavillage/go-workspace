package calljoin

import (
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JoinGetResponse 
type JoinGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The User property
    user i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.User_Abbrable
}
// NewJoinGetResponse instantiates a new joinGetResponse and sets the default values.
func NewJoinGetResponse()(*JoinGetResponse) {
    m := &JoinGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJoinGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["User"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateUser_AbbrFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.User_Abbrable))
        }
        return nil
    }
    return res
}
// GetUser gets the User property value. The User property
func (m *JoinGetResponse) GetUser()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.User_Abbrable) {
    return m.user
}
// Serialize serializes information the current object
func (m *JoinGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("User", m.GetUser())
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
func (m *JoinGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUser sets the User property value. The User property
func (m *JoinGetResponse) SetUser(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.User_Abbrable)() {
    m.user = value
}
// JoinGetResponseable 
type JoinGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUser()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.User_Abbrable)
    SetUser(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.User_Abbrable)()
}
