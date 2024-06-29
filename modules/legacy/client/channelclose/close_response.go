package channelclose

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloseResponse 
// Deprecated: This class is obsolete. Use closePostResponse instead.
type CloseResponse struct {
    ClosePostResponse
}
// NewCloseResponse instantiates a new closeResponse and sets the default values.
func NewCloseResponse()(*CloseResponse) {
    m := &CloseResponse{
        ClosePostResponse: *NewClosePostResponse(),
    }
    return m
}
// CreateCloseResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloseResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloseResponse(), nil
}
// CloseResponseable 
// Deprecated: This class is obsolete. Use closePostResponse instead.
type CloseResponseable interface {
    ClosePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
