package channelleave

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LeaveResponse 
// Deprecated: This class is obsolete. Use leavePostResponse instead.
type LeaveResponse struct {
    LeavePostResponse
}
// NewLeaveResponse instantiates a new leaveResponse and sets the default values.
func NewLeaveResponse()(*LeaveResponse) {
    m := &LeaveResponse{
        LeavePostResponse: *NewLeavePostResponse(),
    }
    return m
}
// CreateLeaveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLeaveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLeaveResponse(), nil
}
// LeaveResponseable 
// Deprecated: This class is obsolete. Use leavePostResponse instead.
type LeaveResponseable interface {
    LeavePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
