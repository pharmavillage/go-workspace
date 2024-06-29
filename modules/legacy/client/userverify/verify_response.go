package userverify

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifyResponse 
// Deprecated: This class is obsolete. Use verifyPostResponse instead.
type VerifyResponse struct {
    VerifyPostResponse
}
// NewVerifyResponse instantiates a new verifyResponse and sets the default values.
func NewVerifyResponse()(*VerifyResponse) {
    m := &VerifyResponse{
        VerifyPostResponse: *NewVerifyPostResponse(),
    }
    return m
}
// CreateVerifyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifyResponse(), nil
}
// VerifyResponseable 
// Deprecated: This class is obsolete. Use verifyPostResponse instead.
type VerifyResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VerifyPostResponseable
}
