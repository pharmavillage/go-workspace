package lockacquire

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AcquireResponse 
// Deprecated: This class is obsolete. Use acquirePostResponse instead.
type AcquireResponse struct {
    AcquirePostResponse
}
// NewAcquireResponse instantiates a new acquireResponse and sets the default values.
func NewAcquireResponse()(*AcquireResponse) {
    m := &AcquireResponse{
        AcquirePostResponse: *NewAcquirePostResponse(),
    }
    return m
}
// CreateAcquireResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcquireResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcquireResponse(), nil
}
// AcquireResponseable 
// Deprecated: This class is obsolete. Use acquirePostResponse instead.
type AcquireResponseable interface {
    AcquirePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
