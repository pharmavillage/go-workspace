package passwordreset

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResetResponse 
// Deprecated: This class is obsolete. Use resetPostResponse instead.
type ResetResponse struct {
    ResetPostResponse
}
// NewResetResponse instantiates a new resetResponse and sets the default values.
func NewResetResponse()(*ResetResponse) {
    m := &ResetResponse{
        ResetPostResponse: *NewResetPostResponse(),
    }
    return m
}
// CreateResetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResetResponse(), nil
}
// ResetResponseable 
// Deprecated: This class is obsolete. Use resetPostResponse instead.
type ResetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResetPostResponseable
}
