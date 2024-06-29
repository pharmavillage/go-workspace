package callstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusResponse 
// Deprecated: This class is obsolete. Use statusGetResponse instead.
type StatusResponse struct {
    StatusGetResponse
}
// NewStatusResponse instantiates a new statusResponse and sets the default values.
func NewStatusResponse()(*StatusResponse) {
    m := &StatusResponse{
        StatusGetResponse: *NewStatusGetResponse(),
    }
    return m
}
// CreateStatusResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStatusResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatusResponse(), nil
}
// StatusResponseable 
// Deprecated: This class is obsolete. Use statusGetResponse instead.
type StatusResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StatusGetResponseable
}
