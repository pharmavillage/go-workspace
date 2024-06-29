package channelrename

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RenameResponse 
// Deprecated: This class is obsolete. Use renamePostResponse instead.
type RenameResponse struct {
    RenamePostResponse
}
// NewRenameResponse instantiates a new renameResponse and sets the default values.
func NewRenameResponse()(*RenameResponse) {
    m := &RenameResponse{
        RenamePostResponse: *NewRenamePostResponse(),
    }
    return m
}
// CreateRenameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRenameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRenameResponse(), nil
}
// RenameResponseable 
// Deprecated: This class is obsolete. Use renamePostResponse instead.
type RenameResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RenamePostResponseable
}
