package models

// Alert Struct
// ToDo ---> add photo field
type Alert struct {
	Id       uint `gorrm:"unique" json:"id"`
	CameraId uint `json:"cameraid"`
}

func (b *Alert) TableName() string {
	return "alert"
}
