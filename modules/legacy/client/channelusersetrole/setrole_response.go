package channelusersetrole

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetroleResponse 
// Deprecated: This class is obsolete. Use setrolePostResponse instead.
type SetroleResponse struct {
    SetrolePostResponse
}
// NewSetroleResponse instantiates a new setroleResponse and sets the default values.
func NewSetroleResponse()(*SetroleResponse) {
    m := &SetroleResponse{
        SetrolePostResponse: *NewSetrolePostResponse(),
    }
    return m
}
// CreateSetroleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetroleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetroleResponse(), nil
}
// SetroleResponseable 
// Deprecated: This class is obsolete. Use setrolePostResponse instead.
type SetroleResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SetrolePostResponseable
}
