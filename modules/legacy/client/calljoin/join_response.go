package calljoin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JoinResponse 
// Deprecated: This class is obsolete. Use joinGetResponse instead.
type JoinResponse struct {
    JoinGetResponse
}
// NewJoinResponse instantiates a new joinResponse and sets the default values.
func NewJoinResponse()(*JoinResponse) {
    m := &JoinResponse{
        JoinGetResponse: *NewJoinGetResponse(),
    }
    return m
}
// CreateJoinResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinResponse(), nil
}
// JoinResponseable 
// Deprecated: This class is obsolete. Use joinGetResponse instead.
type JoinResponseable interface {
    JoinGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
