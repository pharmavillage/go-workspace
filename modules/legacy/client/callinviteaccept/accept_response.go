package callinviteaccept

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AcceptResponse 
// Deprecated: This class is obsolete. Use acceptGetResponse instead.
type AcceptResponse struct {
    AcceptGetResponse
}
// NewAcceptResponse instantiates a new acceptResponse and sets the default values.
func NewAcceptResponse()(*AcceptResponse) {
    m := &AcceptResponse{
        AcceptGetResponse: *NewAcceptGetResponse(),
    }
    return m
}
// CreateAcceptResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcceptResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcceptResponse(), nil
}
// AcceptResponseable 
// Deprecated: This class is obsolete. Use acceptGetResponse instead.
type AcceptResponseable interface {
    AcceptGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
