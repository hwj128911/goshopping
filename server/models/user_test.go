package models

import "testing"

func Test_AddUser(t *testing.T) {
	if has, _ := x.IsTableExist(new(User)); has {
		x.DropTables(new(User))
		x.CreateTables(new(User))
	}
	user := User{Phone: "12345678911", Password: "123", Iconpath: "/default.png"}
	err := AddUser(user)
	if err != nil {
		t.Error("AddUser测试失败", err)
		return
	}
	userget, err := GetUser(1)
	if err != nil {
		t.Error("AddUser测试失败", err)
		return
	}
	if userget.Phone != "12345678911" || userget.Password != "123" {
		t.Error("AddUser测试失败")
	}
}
func Test_GetUser(t *testing.T) {
	userget, err := GetUser(1)
	if err != nil {
		t.Error("GetUser测试失败", err)
		return
	}
	if userget.Phone != "12345678911" || userget.Password != "123" {
		t.Error("AddUser测试失败")
	}
}
func Test_UpdateUser(t *testing.T) {
	user := User{Id: 1, Phone: "14725836900", Password: "456"}
	err := UpdateUser(user)
	if err != nil {
		t.Error("UpdateUser测试失败", err)
		return
	}
	userget, err := GetUser(1)
	if err != nil {
		t.Error("UpdateUser测试失败", err)
		return
	}
	if userget.Phone != "14725836900" || userget.Password != "456" || userget.Iconpath != "/default.png" || userget.Bonuspoints != 10 {
		t.Error("UpdateUser测试失败")
	}
}
func Test_DeleteUser(t *testing.T) {
	err := DeleteUser(1)
	if err != nil {
		t.Error("DeleteUser测试失败", err)
		return
	}
	if err = DeleteUser(1); err == nil {
		t.Error("DeleteUser测试失败")
	}
}
