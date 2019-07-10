package dbop

import "testing"

func dbClear(){
	dbConn.Exec("truncate user_information")
	dbConn.Exec(("truncate user_identity_information"))
}

func TestMain(m *testing.M) {
	//dbClear()
	m.Run()
	//dbClear()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("register", testUserRegister)
	t.Run("login", testUserLogin)
}

func testUserRegister(t *testing.T) {
	err := userRegister("zheng","1258@gmail.com","000000")
	if err != nil {
		t.Errorf("Error of register: %v", err)
	}
}

func testUserLogin(t *testing.T) {
	password,err := userLogin("zheng")
	if password != "000000" {
		t.Errorf("Error of user login for wrong password:%s",password)
	}
	if err != nil {
		t.Errorf("Error of login: %v", err)
	}
}