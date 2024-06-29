package searchquery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// QueryResponse 
// Deprecated: This class is obsolete. Use queryGetResponse instead.
type QueryResponse struct {
    QueryGetResponse
}
// NewQueryResponse instantiates a new queryResponse and sets the default values.
func NewQueryResponse()(*QueryResponse) {
    m := &QueryResponse{
        QueryGetResponse: *NewQueryGetResponse(),
    }
    return m
}
// CreateQueryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQueryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQueryResponse(), nil
}
// QueryResponseable 
// Deprecated: This class is obsolete. Use queryGetResponse instead.
type QueryResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    QueryGetResponseable
}
