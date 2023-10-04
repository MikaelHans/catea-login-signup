package login
import (
	"database/sql"

	"github.com/MikaelHans/catea/login-signup/util"
)
func GetMemberWithLoginInfo(logininfo util.LoginInfo) (*sql.Rows, error) {
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `SELECT *
	FROM catea_member
	WHERE email = $1 AND password = $2`

	rows, err := db.Query(query, logininfo.Email, logininfo.Pass)
	db.Close()
	return rows, err
}