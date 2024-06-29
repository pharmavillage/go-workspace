package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// User_fs_stats 
type User_fs_stats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The quota property
    quota *int32
    // The total_file_count property
    total_file_count *int32
    // The total_folder_count property
    total_folder_count *int32
    // The total_fs_count property
    total_fs_count *int32
    // The total_fs_size property
    total_fs_size *int32
}
// NewUser_fs_stats instantiates a new user_fs_stats and sets the default values.
func NewUser_fs_stats()(*User_fs_stats) {
    m := &User_fs_stats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUser_fs_statsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUser_fs_statsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser_fs_stats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *User_fs_stats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *User_fs_stats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quota"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuota(val)
        }
        return nil
    }
    res["total_file_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFileCount(val)
        }
        return nil
    }
    res["total_folder_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFolderCount(val)
        }
        return nil
    }
    res["total_fs_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFsCount(val)
        }
        return nil
    }
    res["total_fs_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFsSize(val)
        }
        return nil
    }
    return res
}
// GetQuota gets the quota property value. The quota property
func (m *User_fs_stats) GetQuota()(*int32) {
    return m.quota
}
// GetTotalFileCount gets the total_file_count property value. The total_file_count property
func (m *User_fs_stats) GetTotalFileCount()(*int32) {
    return m.total_file_count
}
// GetTotalFolderCount gets the total_folder_count property value. The total_folder_count property
func (m *User_fs_stats) GetTotalFolderCount()(*int32) {
    return m.total_folder_count
}
// GetTotalFsCount gets the total_fs_count property value. The total_fs_count property
func (m *User_fs_stats) GetTotalFsCount()(*int32) {
    return m.total_fs_count
}
// GetTotalFsSize gets the total_fs_size property value. The total_fs_size property
func (m *User_fs_stats) GetTotalFsSize()(*int32) {
    return m.total_fs_size
}
// Serialize serializes information the current object
func (m *User_fs_stats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("quota", m.GetQuota())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_file_count", m.GetTotalFileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_folder_count", m.GetTotalFolderCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_fs_count", m.GetTotalFsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_fs_size", m.GetTotalFsSize())
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
func (m *User_fs_stats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetQuota sets the quota property value. The quota property
func (m *User_fs_stats) SetQuota(value *int32)() {
    m.quota = value
}
// SetTotalFileCount sets the total_file_count property value. The total_file_count property
func (m *User_fs_stats) SetTotalFileCount(value *int32)() {
    m.total_file_count = value
}
// SetTotalFolderCount sets the total_folder_count property value. The total_folder_count property
func (m *User_fs_stats) SetTotalFolderCount(value *int32)() {
    m.total_folder_count = value
}
// SetTotalFsCount sets the total_fs_count property value. The total_fs_count property
func (m *User_fs_stats) SetTotalFsCount(value *int32)() {
    m.total_fs_count = value
}
// SetTotalFsSize sets the total_fs_size property value. The total_fs_size property
func (m *User_fs_stats) SetTotalFsSize(value *int32)() {
    m.total_fs_size = value
}
// User_fs_statsable 
type User_fs_statsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetQuota()(*int32)
    GetTotalFileCount()(*int32)
    GetTotalFolderCount()(*int32)
    GetTotalFsCount()(*int32)
    GetTotalFsSize()(*int32)
    SetQuota(value *int32)()
    SetTotalFileCount(value *int32)()
    SetTotalFolderCount(value *int32)()
    SetTotalFsCount(value *int32)()
    SetTotalFsSize(value *int32)()
}
