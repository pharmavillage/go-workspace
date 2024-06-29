package chatcommand

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommandResponse 
// Deprecated: This class is obsolete. Use commandPostResponse instead.
type CommandResponse struct {
    CommandPostResponse
}
// NewCommandResponse instantiates a new commandResponse and sets the default values.
func NewCommandResponse()(*CommandResponse) {
    m := &CommandResponse{
        CommandPostResponse: *NewCommandPostResponse(),
    }
    return m
}
// CreateCommandResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommandResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommandResponse(), nil
}
// CommandResponseable 
// Deprecated: This class is obsolete. Use commandPostResponse instead.
type CommandResponseable interface {
    CommandPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
