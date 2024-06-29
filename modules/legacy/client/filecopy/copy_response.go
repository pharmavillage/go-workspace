package filecopy

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopyResponse 
// Deprecated: This class is obsolete. Use copyPostResponse instead.
type CopyResponse struct {
    CopyPostResponse
}
// NewCopyResponse instantiates a new copyResponse and sets the default values.
func NewCopyResponse()(*CopyResponse) {
    m := &CopyResponse{
        CopyPostResponse: *NewCopyPostResponse(),
    }
    return m
}
// CreateCopyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopyResponse(), nil
}
// CopyResponseable 
// Deprecated: This class is obsolete. Use copyPostResponse instead.
type CopyResponseable interface {
    CopyPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
