package searchquery

import (
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// QueryGetResponse 
type QueryGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The has_more property
    has_more *bool
    // The meta property
    meta i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable
    // The next_cursor property
    next_cursor *string
    // The results property
    results []i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable
}
// NewQueryGetResponse instantiates a new queryGetResponse and sets the default values.
func NewQueryGetResponse()(*QueryGetResponse) {
    m := &QueryGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateQueryGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQueryGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQueryGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *QueryGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *QueryGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["has_more"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMore(val)
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
    res["next_cursor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextCursor(val)
        }
        return nil
    }
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateSearchResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable)
                }
            }
            m.SetResults(res)
        }
        return nil
    }
    return res
}
// GetHasMore gets the has_more property value. The has_more property
func (m *QueryGetResponse) GetHasMore()(*bool) {
    return m.has_more
}
// GetMeta gets the meta property value. The meta property
func (m *QueryGetResponse) GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable) {
    return m.meta
}
// GetNextCursor gets the next_cursor property value. The next_cursor property
func (m *QueryGetResponse) GetNextCursor()(*string) {
    return m.next_cursor
}
// GetResults gets the results property value. The results property
func (m *QueryGetResponse) GetResults()([]i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable) {
    return m.results
}
// Serialize serializes information the current object
func (m *QueryGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("has_more", m.GetHasMore())
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
        err := writer.WriteStringValue("next_cursor", m.GetNextCursor())
        if err != nil {
            return err
        }
    }
    if m.GetResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("results", cast)
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
func (m *QueryGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHasMore sets the has_more property value. The has_more property
func (m *QueryGetResponse) SetHasMore(value *bool)() {
    m.has_more = value
}
// SetMeta sets the meta property value. The meta property
func (m *QueryGetResponse) SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)() {
    m.meta = value
}
// SetNextCursor sets the next_cursor property value. The next_cursor property
func (m *QueryGetResponse) SetNextCursor(value *string)() {
    m.next_cursor = value
}
// SetResults sets the results property value. The results property
func (m *QueryGetResponse) SetResults(value []i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable)() {
    m.results = value
}
// QueryGetResponseable 
type QueryGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHasMore()(*bool)
    GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)
    GetNextCursor()(*string)
    GetResults()([]i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable)
    SetHasMore(value *bool)()
    SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)()
    SetNextCursor(value *string)()
    SetResults(value []i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.SearchResultable)()
}
