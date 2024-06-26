package channelremove

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveResponse 
// Deprecated: This class is obsolete. Use removePostResponse instead.
type RemoveResponse struct {
    RemovePostResponse
}
// NewRemoveResponse instantiates a new removeResponse and sets the default values.
func NewRemoveResponse()(*RemoveResponse) {
    m := &RemoveResponse{
        RemovePostResponse: *NewRemovePostResponse(),
    }
    return m
}
// CreateRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveResponse(), nil
}
// RemoveResponseable 
// Deprecated: This class is obsolete. Use removePostResponse instead.
type RemoveResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RemovePostResponseable
}
