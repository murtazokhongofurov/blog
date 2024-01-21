package db

import "github.com/blog/internal/controller/http/models"

func (r storagePg) PostCreate(req *models.PostReq) (*models.PostRes, error) {
	res := models.PostRes{}
	query := `
	INSERT INTO post(
		title, text, post_image
	)
	VALUES($1, $2, $3) RETURNING id, title, text, post_image, created_at
	`
	err := r.db.QueryRow(query, req.Title, req.Text, req.PostImage).Scan(
		&res.Id, &res.Title, &res.Text, &res.PostImage, &res.CreatedAt,
	)
	if err != nil {
		return &models.PostRes{}, err
	}
	return &res, nil
}

func (r storagePg) PostGet(id int) (*models.PostRes, error) {
	res := models.PostRes{}
	query := `
	SELECT
		id, 
		title, 
		text, 
		post_image, 
		created_at
	FROM 
		post
	where id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&res.Id, &res.Title, &res.Text, &res.PostImage, &res.CreatedAt,
	)
	if err != nil {
		return &models.PostRes{}, err
	}
	return &res, nil
}

func (r storagePg) PostPut(req *models.PostRes) error {
	query := `
	UPDATE post SET
		title=$1, text=$2, post_image=$3
	WHERE id=$4
	`
	_, err := r.db.Exec(query, req.Title, req.Text, req.PostImage, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r storagePg) PostDel(id int) error {
	query := `
	DELETE FROM post WHERE id=$1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
