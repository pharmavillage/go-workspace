package filecreate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreatePostRequestBody 
type CreatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // name of the folder to create
    fsname *string
    // path to parent where the new folder should be createdExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
    fsparent *string
}
// NewCreatePostRequestBody instantiates a new createPostRequestBody and sets the default values.
func NewCreatePostRequestBody()(*CreatePostRequestBody) {
    m := &CreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fsname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFsname(val)
        }
        return nil
    }
    res["fsparent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFsparent(val)
        }
        return nil
    }
    return res
}
// GetFsname gets the fsname property value. name of the folder to create
func (m *CreatePostRequestBody) GetFsname()(*string) {
    return m.fsname
}
// GetFsparent gets the fsparent property value. path to parent where the new folder should be createdExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
func (m *CreatePostRequestBody) GetFsparent()(*string) {
    return m.fsparent
}
// Serialize serializes information the current object
func (m *CreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fsname", m.GetFsname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fsparent", m.GetFsparent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFsname sets the fsname property value. name of the folder to create
func (m *CreatePostRequestBody) SetFsname(value *string)() {
    m.fsname = value
}
// SetFsparent sets the fsparent property value. path to parent where the new folder should be createdExamples* /f/342322/path/to/resource* /cf/23232/path/to/resource
func (m *CreatePostRequestBody) SetFsparent(value *string)() {
    m.fsparent = value
}
// CreatePostRequestBodyable 
type CreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFsname()(*string)
    GetFsparent()(*string)
    SetFsname(value *string)()
    SetFsparent(value *string)()
}
