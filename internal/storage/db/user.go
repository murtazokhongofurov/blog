package db

import "github.com/blog/internal/controller/http/models"

func (r storagePg) UserPost(req *models.UserReq) (*models.UserRes, error) {
	res := models.UserRes{}
	query := `
	INSERT INTO users(
		full_name, login, password
	)
	VALUES($1, $2, $3) RETURNING id, full_name, login, password, created_at
	`
	err := r.db.QueryRow(query, req.FullName, req.Login, req.Password).Scan(
		&res.Id, &res.FullName, &res.Login, &res.Password, &res.CreatedAt,
	)
	if err != nil {
		return &models.UserRes{}, err
	}
	return &res, nil
}

func (r storagePg) UserGet(id int) (*models.UserRes, error) {
	res := models.UserRes{}
	query := `
	SELECT
	 	id,
		full_name, 
		login, 
		password, 
		photo,
		created_at
	FROM 
		users
	where id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&res.Id, &res.FullName, &res.Login, &res.Password, &res.Photo, &res.CreatedAt,
	)
	if err != nil {
		return &models.UserRes{}, err
	}
	return &res, nil
}

func (r storagePg) UserPut(req *models.UserRes) error {
	query := `
	UPDATE users SET
		full_name=$1, login=$2, password=$3, photo=$4
	
	WHERE id=$5
	`
	_, err := r.db.Exec(query, req.FullName, req.Login, req.Password, req.Photo, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r storagePg) UserDel(id int) error {
	query := `
	DELETE FROM users WHERE id=$1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
