package dbops

import (
	"testing"
)

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t * testing.T)  {
	err := AddUserCredential("zed", "123")
	if err != nil {
		t.Errorf("Error AddUser: %v", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("zed")
	if pwd !="123" || err != nil {
		t.Errorf("Error GetUser")
	}
}

func testDeleteUser(t *testing.T)  {
	err := DeleteUser("zed", "123")
	if err != nil {
		t.Errorf("Error DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("zed")
	if err != nil {
		t.Errorf("Error ReGetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Delte User Test Failed")
	}
}