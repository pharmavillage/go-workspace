package passwordrecover

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecoverResponse 
// Deprecated: This class is obsolete. Use recoverPostResponse instead.
type RecoverResponse struct {
    RecoverPostResponse
}
// NewRecoverResponse instantiates a new recoverResponse and sets the default values.
func NewRecoverResponse()(*RecoverResponse) {
    m := &RecoverResponse{
        RecoverPostResponse: *NewRecoverPostResponse(),
    }
    return m
}
// CreateRecoverResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecoverResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecoverResponse(), nil
}
// RecoverResponseable 
// Deprecated: This class is obsolete. Use recoverPostResponse instead.
type RecoverResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RecoverPostResponseable
}
