package core

import (
	"testing"

	"github.com/orzzzli/orz_cms/src/source"
)

var testScheme = "root:root@tcp(127.0.0.1:3306)/cms?charset=utf8"

func TestCheckTableComplete(t *testing.T) {
	m := source.NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	complete := CheckTableComplete(m)
	if complete {
		t.Log("table complete.")
	} else {
		t.Log("table not complete.")
	}
}

func TestInstallTable(t *testing.T) {
	m := source.NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	complete := InstallTable(m)
	if complete {
		t.Log("table install success.")
	} else {
		t.Log("table install fail.")
	}
}

func TestConfirmInstall(t *testing.T) {
	m := source.NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	complete := ConfirmInstall(m)
	if complete {
		t.Log("confirm install success.")
	} else {
		t.Log("confirm install fail.")
	}
}

func TestIsInstalled(t *testing.T) {
	m := source.NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	complete := IsInstalled(m)
	if complete {
		t.Log("is installed.")
	} else {
		t.Log("not installed.")
	}
}

func TestInstallDefaultData(t *testing.T) {
	m := source.NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	complete := InstallDefaultData(m)
	if complete {
		t.Log("data installed.")
	} else {
		t.Log("data not installed.")
	}
}
