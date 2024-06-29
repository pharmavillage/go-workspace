package channeloneonone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OneOnOneResponse 
// Deprecated: This class is obsolete. Use oneOnOnePostResponse instead.
type OneOnOneResponse struct {
    OneOnOnePostResponse
}
// NewOneOnOneResponse instantiates a new oneOnOneResponse and sets the default values.
func NewOneOnOneResponse()(*OneOnOneResponse) {
    m := &OneOnOneResponse{
        OneOnOnePostResponse: *NewOneOnOnePostResponse(),
    }
    return m
}
// CreateOneOnOneResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOneOnOneResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOneOnOneResponse(), nil
}
// OneOnOneResponseable 
// Deprecated: This class is obsolete. Use oneOnOnePostResponse instead.
type OneOnOneResponseable interface {
    OneOnOnePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
