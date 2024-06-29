package channelinvite

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InviteResponse 
// Deprecated: This class is obsolete. Use invitePostResponse instead.
type InviteResponse struct {
    InvitePostResponse
}
// NewInviteResponse instantiates a new inviteResponse and sets the default values.
func NewInviteResponse()(*InviteResponse) {
    m := &InviteResponse{
        InvitePostResponse: *NewInvitePostResponse(),
    }
    return m
}
// CreateInviteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInviteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInviteResponse(), nil
}
// InviteResponseable 
// Deprecated: This class is obsolete. Use invitePostResponse instead.
type InviteResponseable interface {
    InvitePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
