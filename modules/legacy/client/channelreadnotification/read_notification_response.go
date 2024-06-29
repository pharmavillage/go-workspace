package channelreadnotification

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReadNotificationResponse 
// Deprecated: This class is obsolete. Use readNotificationPostResponse instead.
type ReadNotificationResponse struct {
    ReadNotificationPostResponse
}
// NewReadNotificationResponse instantiates a new readNotificationResponse and sets the default values.
func NewReadNotificationResponse()(*ReadNotificationResponse) {
    m := &ReadNotificationResponse{
        ReadNotificationPostResponse: *NewReadNotificationPostResponse(),
    }
    return m
}
// CreateReadNotificationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReadNotificationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReadNotificationResponse(), nil
}
// ReadNotificationResponseable 
// Deprecated: This class is obsolete. Use readNotificationPostResponse instead.
type ReadNotificationResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReadNotificationPostResponseable
}
