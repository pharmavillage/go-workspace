package actionmove

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MoveResponse 
// Deprecated: This class is obsolete. Use movePostResponse instead.
type MoveResponse struct {
    MovePostResponse
}
// NewMoveResponse instantiates a new moveResponse and sets the default values.
func NewMoveResponse()(*MoveResponse) {
    m := &MoveResponse{
        MovePostResponse: *NewMovePostResponse(),
    }
    return m
}
// CreateMoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMoveResponse(), nil
}
// MoveResponseable 
// Deprecated: This class is obsolete. Use movePostResponse instead.
type MoveResponseable interface {
    MovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
