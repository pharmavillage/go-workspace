package userfinalize

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FinalizeResponse 
// Deprecated: This class is obsolete. Use finalizePostResponse instead.
type FinalizeResponse struct {
    FinalizePostResponse
}
// NewFinalizeResponse instantiates a new finalizeResponse and sets the default values.
func NewFinalizeResponse()(*FinalizeResponse) {
    m := &FinalizeResponse{
        FinalizePostResponse: *NewFinalizePostResponse(),
    }
    return m
}
// CreateFinalizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFinalizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFinalizeResponse(), nil
}
// FinalizeResponseable 
// Deprecated: This class is obsolete. Use finalizePostResponse instead.
type FinalizeResponseable interface {
    FinalizePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
