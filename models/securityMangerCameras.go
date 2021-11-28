package models

import (
	"fmt"

	Config "github.com/ahmedsharyo/AccessControlService/config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllSecurityMangerCamerass Fetch all SecurityMangerCameras data
func GetAllSecurityMangerCamerass(SecurityMangerCameras *[]SecurityMangerCameras) (err error) {
	if err = Config.DB.Find(SecurityMangerCameras).Error; err != nil {
		return err
	}
	return nil
}

//CreateSecurityMangerCameras ... Insert New data
func CreateSecurityMangerCameras(SecurityMangerCameras *SecurityMangerCameras) (err error) {
	if err = Config.DB.Create(SecurityMangerCameras).Error; err != nil {
		return err
	}
	return nil
}

//GetSecurityMangerCamerasByID ... Fetch only one SecurityMangerCameras by Id
func GetSecurityMangerCamerasByEmail(SecurityMangerCameras *SecurityMangerCameras, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(SecurityMangerCameras).Error; err != nil {
		return err
	}
	return nil
}

//UpdateSecurityMangerCameras ... Update SecurityMangerCameras
func UpdateSecurityMangerCameras(SecurityMangerCameras *SecurityMangerCameras, id string) (err error) {
	fmt.Println(SecurityMangerCameras)
	Config.DB.Save(SecurityMangerCameras)
	return nil
}

//DeleteSecurityMangerCameras ... Delete SecurityMangerCameras
func DeleteSecurityMangerCameras(SecurityMangerCameras *SecurityMangerCameras, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(SecurityMangerCameras)
	return nil
}
