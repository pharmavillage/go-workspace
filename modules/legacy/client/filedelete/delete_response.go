package filedelete

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeleteResponse 
// Deprecated: This class is obsolete. Use deletePostResponse instead.
type DeleteResponse struct {
    DeletePostResponse
}
// NewDeleteResponse instantiates a new deleteResponse and sets the default values.
func NewDeleteResponse()(*DeleteResponse) {
    m := &DeleteResponse{
        DeletePostResponse: *NewDeletePostResponse(),
    }
    return m
}
// CreateDeleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteResponse(), nil
}
// DeleteResponseable 
// Deprecated: This class is obsolete. Use deletePostResponse instead.
type DeleteResponseable interface {
    DeletePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
