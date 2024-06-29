package usernotificationsmanage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagePostRequestBody 
type ManagePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Notification option (can be any of the values below):  * `0` - No notifications (completely unsubscribe)  * `1` - Only notify mentions  * `2` - Notify all unread messages
    notification_option *int32
}
// NewManagePostRequestBody instantiates a new managePostRequestBody and sets the default values.
func NewManagePostRequestBody()(*ManagePostRequestBody) {
    m := &ManagePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notification_option"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationOption(val)
        }
        return nil
    }
    return res
}
// GetNotificationOption gets the notification_option property value. Notification option (can be any of the values below):  * `0` - No notifications (completely unsubscribe)  * `1` - Only notify mentions  * `2` - Notify all unread messages
func (m *ManagePostRequestBody) GetNotificationOption()(*int32) {
    return m.notification_option
}
// Serialize serializes information the current object
func (m *ManagePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("notification_option", m.GetNotificationOption())
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
func (m *ManagePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNotificationOption sets the notification_option property value. Notification option (can be any of the values below):  * `0` - No notifications (completely unsubscribe)  * `1` - Only notify mentions  * `2` - Notify all unread messages
func (m *ManagePostRequestBody) SetNotificationOption(value *int32)() {
    m.notification_option = value
}
// ManagePostRequestBodyable 
type ManagePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNotificationOption()(*int32)
    SetNotificationOption(value *int32)()
}
