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

func (u *UserRepoImplement) CheckIfExist(context context.Context, signUpReq request.RequestSignUp) error {
	statement := `
		SELECT COUNT(*) FROM users where phone=$1;
	`
	var count int
	u.sql.Db.QueryRowContext(context, statement, signUpReq.Phone).Scan(&count)

	if count > 0 {
		return banana.UserConflict
	}
	return nil
}

func (u *UserRepoImplement) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, phone, password, email, gender, full_name) 
		VALUES (:user_id, :phone, :password, :email, :gender, :full_name);
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

func (u *UserRepoImplement) SaveModel(context context.Context, model model.Model) (model.Model, error) {
	statement := `
		INSERT INTO models(model_id, name, description, status, model_date, content) 
		VALUES (:model_id, :name, :description, :status, :model_date, :content);
	`
	// user.CreateAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, model)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return model, banana.UserConflict
			}
		}
		return model, banana.SignUpFail
	}

	return model, nil
}

func (u *UserRepoImplement) CheckSignIn(context context.Context, signInReq request.RequestSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", signInReq.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		return user, err
	}

	return user, nil
}
