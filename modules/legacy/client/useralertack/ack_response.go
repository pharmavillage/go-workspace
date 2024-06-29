package useralertack

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AckResponse 
// Deprecated: This class is obsolete. Use ackPostResponse instead.
type AckResponse struct {
    AckPostResponse
}
// NewAckResponse instantiates a new ackResponse and sets the default values.
func NewAckResponse()(*AckResponse) {
    m := &AckResponse{
        AckPostResponse: *NewAckPostResponse(),
    }
    return m
}
// CreateAckResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAckResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAckResponse(), nil
}
// AckResponseable 
// Deprecated: This class is obsolete. Use ackPostResponse instead.
type AckResponseable interface {
    AckPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
