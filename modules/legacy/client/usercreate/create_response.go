package usercreate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateResponse 
// Deprecated: This class is obsolete. Use createPostResponse instead.
type CreateResponse struct {
    CreatePostResponse
}
// NewCreateResponse instantiates a new createResponse and sets the default values.
func NewCreateResponse()(*CreateResponse) {
    m := &CreateResponse{
        CreatePostResponse: *NewCreatePostResponse(),
    }
    return m
}
// CreateCreateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateResponse(), nil
}
// CreateResponseable 
// Deprecated: This class is obsolete. Use createPostResponse instead.
type CreateResponseable interface {
    CreatePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
