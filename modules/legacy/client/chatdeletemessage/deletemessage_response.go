package chatdeletemessage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletemessageResponse 
// Deprecated: This class is obsolete. Use deletemessagePostResponse instead.
type DeletemessageResponse struct {
    DeletemessagePostResponse
}
// NewDeletemessageResponse instantiates a new deletemessageResponse and sets the default values.
func NewDeletemessageResponse()(*DeletemessageResponse) {
    m := &DeletemessageResponse{
        DeletemessagePostResponse: *NewDeletemessagePostResponse(),
    }
    return m
}
// CreateDeletemessageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletemessageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletemessageResponse(), nil
}
// DeletemessageResponseable 
// Deprecated: This class is obsolete. Use deletemessagePostResponse instead.
type DeletemessageResponseable interface {
    DeletemessagePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
