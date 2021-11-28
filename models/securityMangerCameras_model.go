package models

// SecurityMangerCameras Struct
type SecurityMangerCameras struct {
	Id               uint   `gorrm:"unique" json:"id"`
	CameraId         string `json:"cameraid"`
	SecurityMangerId int    `json:"securitymangerid"`
}

func (b *SecurityMangerCameras) TableName() string {
	return "securityMangerCameras"
}
