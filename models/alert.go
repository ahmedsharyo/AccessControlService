package models

import (
	"fmt"

	Config "github.com/ahmedsharyo/AlertService/config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllAlerts Fetch all Alert data
func GetAllAlerts(alert *[]Alert) (err error) {
	if err = Config.DB.Find(alert).Error; err != nil {
		return err
	}
	return nil
}

//CreateAlert ... Insert New data
func CreateAlert(alert *Alert) (err error) {
	if err = Config.DB.Create(alert).Error; err != nil {
		return err
	}
	return nil
}

//GetAlertByID ... Fetch only one Alert by Id
func GetAlertByEmail(alert *Alert, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(alert).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAlert ... Update Alert
func UpdateAlert(alert *Alert, id string) (err error) {
	fmt.Println(alert)
	Config.DB.Save(alert)
	return nil
}

//DeleteAlert ... Delete Alert
func DeleteAlert(alert *Alert, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(alert)
	return nil
}
