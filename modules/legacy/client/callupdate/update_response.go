package callupdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateResponse 
// Deprecated: This class is obsolete. Use updatePostResponse instead.
type UpdateResponse struct {
    UpdatePostResponse
}
// NewUpdateResponse instantiates a new updateResponse and sets the default values.
func NewUpdateResponse()(*UpdateResponse) {
    m := &UpdateResponse{
        UpdatePostResponse: *NewUpdatePostResponse(),
    }
    return m
}
// CreateUpdateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateResponse(), nil
}
// UpdateResponseable 
// Deprecated: This class is obsolete. Use updatePostResponse instead.
type UpdateResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UpdatePostResponseable
}
