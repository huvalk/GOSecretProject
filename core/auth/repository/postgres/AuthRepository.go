package authRepository

import (
	"GOSecretProject/core/model/base"
	"database/sql"
	"encoding/base64"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(user baseModels.User) (err error) {
	var personID uint64

	err = r.db.QueryRow("INSERT INTO users (login, password, phone) VALUES($1, $2, $3) RETURNING id",
		user.Login, user.Password, user.Phone).
		Scan(&personID)

	return err
}

func (r *AuthRepository) Login(user baseModels.User) (userID int, session string, statusCode int, err error) {
	checkUser := "SELECT id, password FROM users WHERE login = $1"
	var hashedPwd string
	rows := r.db.QueryRow(checkUser, user.Login)
	err = rows.Scan(&userID, &hashedPwd)
	if err != nil || user.Password != hashedPwd {
		return 0, "", 401, err
	}

	//TODO Сделать генерацию токена, разрешить повторную авторизацию
	session = base64.StdEncoding.EncodeToString([]byte(user.Login + user.Password + time.Now().String()))

	insertSession := `INSERT INTO session (user_id, session_id) 
					VALUES($1, $2)`
	_, err = r.db.Exec(insertSession, userID, session)
	if err != nil {
		return 0, "", 500, err
	}

	return userID, session, 201, nil
}

func (r *AuthRepository) RestorePassword(userLogin string) (user baseModels.User, err error) {
	checkUser := "SELECT phone, password FROM users WHERE login = $1"
	rows := r.db.QueryRow(checkUser, userLogin)
	err = rows.Scan(&user.Phone, &user.Password)

	return user, err
}

func (r *AuthRepository) Logout(session string) (err error) {
	deleteRow := "DELETE FROM session WHERE session_id = $1;"
	_, err = r.db.Exec(deleteRow, session)

	return err
}

func (r *AuthRepository) CheckSession(session string) (user baseModels.User, err error) {
	check := "SELECT user_id FROM session WHERE session_id = $1"
	err = r.db.QueryRow(check, session).Scan(&user.ID)
	if err != nil {
		return user, err
	}

	getUser := "SELECT phone, login FROM users WHERE id = $1"
	err = r.db.QueryRow(getUser, user.ID).Scan(&user.Phone, &user.Login)
	user.Session = session

	return user, err
}

func (r *AuthRepository) CheckPhone(phone string) (res bool, err error) {
	check := "SELECT exists(select 1 from users where phone = $1)"
	err = r.db.QueryRow(check, phone).Scan(res)

	return res, err
}