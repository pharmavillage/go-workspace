package systemping

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PingResponse 
// Deprecated: This class is obsolete. Use pingPostResponse instead.
type PingResponse struct {
    PingPostResponse
}
// NewPingResponse instantiates a new pingResponse and sets the default values.
func NewPingResponse()(*PingResponse) {
    m := &PingResponse{
        PingPostResponse: *NewPingPostResponse(),
    }
    return m
}
// CreatePingResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePingResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPingResponse(), nil
}
// PingResponseable 
// Deprecated: This class is obsolete. Use pingPostResponse instead.
type PingResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PingPostResponseable
}
