package contact_formfill

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FillResponse 
// Deprecated: This class is obsolete. Use fillPostResponse instead.
type FillResponse struct {
    FillPostResponse
}
// NewFillResponse instantiates a new fillResponse and sets the default values.
func NewFillResponse()(*FillResponse) {
    m := &FillResponse{
        FillPostResponse: *NewFillPostResponse(),
    }
    return m
}
// CreateFillResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFillResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFillResponse(), nil
}
// FillResponseable 
// Deprecated: This class is obsolete. Use fillPostResponse instead.
type FillResponseable interface {
    FillPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
