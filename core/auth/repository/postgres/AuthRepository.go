package authRepository

import (
	"GOSecretProject/core/model/base"
	"database/sql"
	"encoding/base64"
	"github.com/kataras/golog"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(user base.User) (err error) {
	var personID uint64

	err = r.db.QueryRow("INSERT INTO users (login, password, phone) VALUES($1, $2, $3) RETURNING id",
		user.Login, user.Password, user.Phone).
		Scan(&personID)

	return err
}

func (r *AuthRepository) Login(user base.User) (userID int, session string, statusCode int, err error) {
	checkUser := "SELECT id, password FROM users WHERE login = $1"
	var hashedPwd string
	rows := r.db.QueryRow(checkUser, user.Login)
	err = rows.Scan(&userID, &hashedPwd)
	if err != nil || user.Password != hashedPwd {
		golog.Error(user.Password, hashedPwd, user.Password != hashedPwd)
		return 0, "", 401, err
	}

	//TODO Сделать генерацию токена
	session = base64.StdEncoding.EncodeToString([]byte(user.Login + user.Password))

	insertSession := `INSERT INTO session (user_id, session_id) 
					VALUES($1, $2)`
	_, err = r.db.Exec(insertSession, userID, session)
	if err != nil {
		return 0, "", 500, err
	}

	return userID, session, 201, nil
}

func (r *AuthRepository) Logout(session string) (err error) {
	deleteRow := "DELETE FROM session WHERE session_id = $1;"
	_, err = r.db.Exec(deleteRow, session)

	return err
}
