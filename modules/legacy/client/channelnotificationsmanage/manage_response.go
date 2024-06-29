package channelnotificationsmanage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManageResponse 
// Deprecated: This class is obsolete. Use managePostResponse instead.
type ManageResponse struct {
    ManagePostResponse
}
// NewManageResponse instantiates a new manageResponse and sets the default values.
func NewManageResponse()(*ManageResponse) {
    m := &ManageResponse{
        ManagePostResponse: *NewManagePostResponse(),
    }
    return m
}
// CreateManageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManageResponse(), nil
}
// ManageResponseable 
// Deprecated: This class is obsolete. Use managePostResponse instead.
type ManageResponseable interface {
    ManagePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
