package channelactivate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivateResponse 
// Deprecated: This class is obsolete. Use activatePostResponse instead.
type ActivateResponse struct {
    ActivatePostResponse
}
// NewActivateResponse instantiates a new activateResponse and sets the default values.
func NewActivateResponse()(*ActivateResponse) {
    m := &ActivateResponse{
        ActivatePostResponse: *NewActivatePostResponse(),
    }
    return m
}
// CreateActivateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivateResponse(), nil
}
// ActivateResponseable 
// Deprecated: This class is obsolete. Use activatePostResponse instead.
type ActivateResponseable interface {
    ActivatePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
