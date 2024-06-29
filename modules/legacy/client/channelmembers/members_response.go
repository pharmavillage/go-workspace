package channelmembers

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MembersResponse 
// Deprecated: This class is obsolete. Use membersGetResponse instead.
type MembersResponse struct {
    MembersGetResponse
}
// NewMembersResponse instantiates a new membersResponse and sets the default values.
func NewMembersResponse()(*MembersResponse) {
    m := &MembersResponse{
        MembersGetResponse: *NewMembersGetResponse(),
    }
    return m
}
// CreateMembersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMembersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMembersResponse(), nil
}
// MembersResponseable 
// Deprecated: This class is obsolete. Use membersGetResponse instead.
type MembersResponseable interface {
    MembersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
