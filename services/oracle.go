package services

import (
	"fmt"
	"strings"
	"database/sql"

	_ "github.com/mattn/go-oci8"
)

// create oracle database and database user.
// grant user to database
// return databaseName, userName, userPassword, error.
func createDatabaseAndUser(conn string, tableSpace string, bigFile bool) (string, string, string, error) {
	var bindSucceeded = false

	db, err := sql.Open("oci8", conn)
	if err != nil {
		return "", "", "", err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return "", "", "", err
	}

	databaseName := "ts" + generateGuid()[0:14]
	sqlCreateTS := ""
	if bigFile {
		sqlCreateTS = fmt.Sprintf("create BIGFILE tablespace %s", databaseName)
	} else {
		sqlCreateTS = fmt.Sprintf("create tablespace %s", databaseName)
	}

	println(sqlCreateTS)

	_, err = db.Query(sqlCreateTS)
	if err != nil {
		println("create tablespace err:", err.Error())
		return "", "", "", err
	}

	defer func() {
		if bindSucceeded {
			return
		}

		_, err := db.Query(fmt.Sprintf("drop tablespace %s including contents and datafiles", databaseName))
		if err != nil {
			println("bind failed: drop tablespace", databaseName, "failed:", err)
			return
		}

		println("bind failed: drop tablespace", databaseName, "succeeded.")
	}()

	sqlAlterTS := fmt.Sprintf("alter tablespace %s resize %s", databaseName, tableSpace)

	println(sqlAlterTS)

	_, err = db.Query(sqlAlterTS)
	if err != nil {
		println("alert tablespace err:", err.Error())
		return "", "", "", err
	}

	newUsername := "u" + generateGuid()[0:15]
	newPassword := "p" + generateGuid()[0:15]

	sqlCreateUser := fmt.Sprintf(`CREATE USER %s IDENTIFIED BY %s`, newUsername, newPassword)

	println(sqlCreateUser)

	_, err = db.Query(sqlCreateUser)
	if err != nil {
		println("create user err:", err.Error())
		return "", "", "", err
	}

	defer func() {
		if bindSucceeded {
			return
		}

		_, err := db.Query(fmt.Sprintf("drop user %s cascade", newUsername))
		if err != nil {
			println("bind failed: drop user", newUsername, "failed:", err)
			return
		}

		println("bind failed: drop user", newUsername, "succeeded.")
	}()

	sqlAlterUser := fmt.Sprintf(`ALTER USER %s quota unlimited on %s`, newUsername, databaseName)
	println("alter user: ", sqlAlterUser)
	_, err = db.Query(sqlAlterUser)
	if err != nil {
		return "", "", "", err
	}

	sqlGrantUser := fmt.Sprintf(`GRANT CREATE SESSION, CREATE TABLE, CREATE VIEW, SELECT_CATALOG_ROLE, EXECUTE_CATALOG_ROLE TO %s`, newUsername)
	println("grant user: ", sqlGrantUser)
	_, err = db.Query(sqlGrantUser)
	if err != nil {
		return "", "", "", err
	}

	bindSucceeded = true

	return databaseName, newUsername, newPassword, nil
}

// delete oracle database and database user.
// unGrant user to database
// return error.
func deleteDatabaseAndUser(conn, database, username string) error {
	db, err := sql.Open("oci8", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	_, errDropUser := db.Query(fmt.Sprintf("drop user %s cascade", username))
	if errDropUser == nil {
		println("user", username, "was dropped")
	} else if strings.Index (errDropUser.Error(), "ORA-01918") >= 0 {
		errDropUser = nil
	}

	_, errDropDatabase := db.Query(fmt.Sprintf("drop tablespace %s including contents and datafiles", database))
	if errDropDatabase == nil {
		println("tablespace", database, "was dropped")
	} else if strings.Index (errDropDatabase.Error(), "ORA-00959") >= 0 {
		errDropDatabase = nil
	}

	if errDropUser != nil {
		println("unbind drop user failed:", errDropUser.Error())
		if errDropDatabase == nil {
			return errDropUser
		}
	}

	if errDropDatabase != nil {
		println("unbind drop tablesapce failed:", errDropDatabase.Error())
		return errDropDatabase
	}

	return nil
}