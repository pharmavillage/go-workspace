package chatpostmessage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostmessageResponse 
// Deprecated: This class is obsolete. Use postmessagePostResponse instead.
type PostmessageResponse struct {
    PostmessagePostResponse
}
// NewPostmessageResponse instantiates a new postmessageResponse and sets the default values.
func NewPostmessageResponse()(*PostmessageResponse) {
    m := &PostmessageResponse{
        PostmessagePostResponse: *NewPostmessagePostResponse(),
    }
    return m
}
// CreatePostmessageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePostmessageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostmessageResponse(), nil
}
// PostmessageResponseable 
// Deprecated: This class is obsolete. Use postmessagePostResponse instead.
type PostmessageResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PostmessagePostResponseable
}
