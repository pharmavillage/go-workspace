package filelist

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListResponse 
// Deprecated: This class is obsolete. Use listGetResponse instead.
type ListResponse struct {
    ListGetResponse
}
// NewListResponse instantiates a new listResponse and sets the default values.
func NewListResponse()(*ListResponse) {
    m := &ListResponse{
        ListGetResponse: *NewListGetResponse(),
    }
    return m
}
// CreateListResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListResponse(), nil
}
// ListResponseable 
// Deprecated: This class is obsolete. Use listGetResponse instead.
type ListResponseable interface {
    ListGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
