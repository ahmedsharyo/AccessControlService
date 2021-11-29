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

//GetSecurityMangerCamerasByID ... Fetch all SecurityMangerCameras by cameraId
func GetSecurityMangerCamerasByCameraId(securityMangerCameras *SecurityMangerCameras, camera_id string) (err error) {
	print("--------------+-++---------------")
	print(securityMangerCameras)
	print("-----------------------------")
	if err = Config.DB.Find(securityMangerCameras).Error; err != nil {

		print("--------------++++---------------")
		print(err)
		print("-----------------------------")
		return err
	}
	print("-----------------***------------")
	print(camera_id)
	print("-----------------------------")
	return nil
}

//UpdateSecurityMangerCameras ... Update SecurityMangerCameras
func UpdateSecurityMangerCameras(SecurityMangerCameras *SecurityMangerCameras, id uint) (err error) {
	fmt.Println(SecurityMangerCameras)
	Config.DB.Save(SecurityMangerCameras)
	return nil
}

//DeleteSecurityMangerCameras ... Delete SecurityMangerCameras
func DeleteSecurityMangerCameras(SecurityMangerCameras *SecurityMangerCameras, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(SecurityMangerCameras)
	return nil
}
