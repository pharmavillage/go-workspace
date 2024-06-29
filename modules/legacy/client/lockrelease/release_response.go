package lockrelease

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReleaseResponse 
// Deprecated: This class is obsolete. Use releasePostResponse instead.
type ReleaseResponse struct {
    ReleasePostResponse
}
// NewReleaseResponse instantiates a new releaseResponse and sets the default values.
func NewReleaseResponse()(*ReleaseResponse) {
    m := &ReleaseResponse{
        ReleasePostResponse: *NewReleasePostResponse(),
    }
    return m
}
// CreateReleaseResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReleaseResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReleaseResponse(), nil
}
// ReleaseResponseable 
// Deprecated: This class is obsolete. Use releasePostResponse instead.
type ReleaseResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReleasePostResponseable
}
