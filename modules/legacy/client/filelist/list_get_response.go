package filelist

import (
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListGetResponse 
type ListGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The files property
    files []i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable
    // The meta property
    meta i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable
    // The total property
    total *int32
}
// NewListGetResponse instantiates a new listGetResponse and sets the default values.
func NewListGetResponse()(*ListGetResponse) {
    m := &ListGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateListGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ListGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable)
                }
            }
            m.SetFiles(res)
        }
        return nil
    }
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
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The files property
func (m *ListGetResponse) GetFiles()([]i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable) {
    return m.files
}
// GetMeta gets the meta property value. The meta property
func (m *ListGetResponse) GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable) {
    return m.meta
}
// GetTotal gets the total property value. The total property
func (m *ListGetResponse) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *ListGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("meta", m.GetMeta())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *ListGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFiles sets the files property value. The files property
func (m *ListGetResponse) SetFiles(value []i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable)() {
    m.files = value
}
// SetMeta sets the meta property value. The meta property
func (m *ListGetResponse) SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)() {
    m.meta = value
}
// SetTotal sets the total property value. The total property
func (m *ListGetResponse) SetTotal(value *int32)() {
    m.total = value
}
// ListGetResponseable 
type ListGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFiles()([]i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable)
    GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)
    GetTotal()(*int32)
    SetFiles(value []i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Fileable)()
    SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)()
    SetTotal(value *int32)()
}
