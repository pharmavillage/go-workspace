package channelgroupcreate

import (
    i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1 "kiota_airsend/client/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreatePostResponse 
type CreatePostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The channel_group property
    channel_group i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.ChannelGroupable
    // The meta property
    meta i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable
}
// NewCreatePostResponse instantiates a new createPostResponse and sets the default values.
func NewCreatePostResponse()(*CreatePostResponse) {
    m := &CreatePostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreatePostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreatePostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreatePostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelGroup gets the channel_group property value. The channel_group property
func (m *CreatePostResponse) GetChannelGroup()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.ChannelGroupable) {
    return m.channel_group
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channel_group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateChannelGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelGroup(val.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.ChannelGroupable))
        }
        return nil
    }
    res["meta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.CreateMetaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable))
        }
        return nil
    }
    return res
}
// GetMeta gets the meta property value. The meta property
func (m *CreatePostResponse) GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable) {
    return m.meta
}
// Serialize serializes information the current object
func (m *CreatePostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("channel_group", m.GetChannelGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("meta", m.GetMeta())
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
func (m *CreatePostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelGroup sets the channel_group property value. The channel_group property
func (m *CreatePostResponse) SetChannelGroup(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.ChannelGroupable)() {
    m.channel_group = value
}
// SetMeta sets the meta property value. The meta property
func (m *CreatePostResponse) SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)() {
    m.meta = value
}
// CreatePostResponseable 
type CreatePostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelGroup()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.ChannelGroupable)
    GetMeta()(i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)
    SetChannelGroup(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.ChannelGroupable)()
    SetMeta(value i05dfa36b0c4b2e6d792d7cf603d66eff5d7d5eb1eda7f14fe9003dd30b87f3e1.Metaable)()
}
