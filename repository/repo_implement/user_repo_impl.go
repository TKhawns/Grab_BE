package repo_implement

import (
	"context"
	"database/sql"
	"go_grab/banana"
	"go_grab/db"
	"go_grab/model"
	"go_grab/model/request"
	"go_grab/repository"

	"github.com/lib/pq"
)

type UserRepoImplement struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImplement{sql: sql}

}

func (u *UserRepoImplement) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, phone, password, role, full_name) 
		VALUES (:user_id, :phone, :password, :role, :full_name);
	`
	// user.CreateAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}

	return user, nil
}

func (u *UserRepoImplement) CheckSignIn(context context.Context, signInReq request.RequestSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE phone=$1", signInReq.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		return user, err
	}

	return user, nil
}
