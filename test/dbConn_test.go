package test

import (
	"testing"
	"userRepository/internal/database"
)

func TestConection(t *testing.T){
	t.Run("test for connect",func(t *testing.T){
	_,err:= database.DbConnect()
	if err!=nil{
		t.Error(err.Error())
	}
	})
}
