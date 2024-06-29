package channelfavorite

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FavoritePostRequestBody 
type FavoritePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id of the channel to be set as favorite.
    channel_id *int32
}
// NewFavoritePostRequestBody instantiates a new favoritePostRequestBody and sets the default values.
func NewFavoritePostRequestBody()(*FavoritePostRequestBody) {
    m := &FavoritePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFavoritePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFavoritePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFavoritePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FavoritePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannelId gets the channel_id property value. The id of the channel to be set as favorite.
func (m *FavoritePostRequestBody) GetChannelId()(*int32) {
    return m.channel_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FavoritePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channel_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FavoritePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("channel_id", m.GetChannelId())
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
func (m *FavoritePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannelId sets the channel_id property value. The id of the channel to be set as favorite.
func (m *FavoritePostRequestBody) SetChannelId(value *int32)() {
    m.channel_id = value
}
// FavoritePostRequestBodyable 
type FavoritePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*int32)
    SetChannelId(value *int32)()
}
