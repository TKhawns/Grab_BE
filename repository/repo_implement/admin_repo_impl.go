package repo_implement

import (
	"context"
	"database/sql"
	"go_grab/banana"
	"go_grab/db"
	"go_grab/model"
	"go_grab/model/request"
	"go_grab/repository"
)

type AdminRepoImplement struct {
	sql *db.Sql
}

func NewAdminRepo(sql *db.Sql) repository.AdminRepo {
	return &AdminRepoImplement{sql: sql}

}

func (u *AdminRepoImplement) GetUserModel(context context.Context) ([]model.Model, error) {

	rows, err := u.sql.Db.QueryContext(context, "SELECT u.full_name, m.model_id, m.name, m.status, m.content, m.format FROM models m JOIN users u ON m.user_id = u.user_id")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var list_model []model.Model
	var model model.Model
	for rows.Next() {
		rows.Scan(&model.Date, &model.Model_id, &model.Name, &model.Status, &model.Content,
			&model.Format)
		list_model = append(list_model, model)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return list_model, banana.UserNotFound
		}
		return list_model, err
	}

	return list_model, nil
}

func (u *AdminRepoImplement) UpdateStatus(context context.Context, req request.RequestUpdateStatus) error {
	statement := `
	UPDATE models SET status= :status WHERE model_id= :model_id;
	`
	_, err := u.sql.Db.NamedExecContext(context, statement, req)
	if err != nil {
		return err
	}

	return nil
}
