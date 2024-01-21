package db

import (
	"github.com/blog/internal/controller/http/models"
)

func (r storagePg) CommentPost(req *models.CommentReq) (*models.CommentRes, error) {
	res := models.CommentRes{}
	query := `
	INSERT INTO comment(
		post_id, user_id, text
	)
	VALUES($1, $2, $3) RETURNING id, post_id, user_id, text, created_at
	`
	err := r.db.QueryRow(query, req.PostId, req.UserId, req.Text).Scan(
		&res.Id, &res.PostId, &res.UserId, &res.Text, &res.CreatedAt,
	)
	if err != nil {
		return &models.CommentRes{}, err
	}
	return &res, nil
}

func (r storagePg) CommentGet(id int) (*models.CommentRes, error) {
	res := models.CommentRes{}
	query := `
	SELECT
		id, 
		post_id, 
		user_id, 
		text, 
		created_at
	FROM 
		comment
	where id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&res.Id, &res.PostId, &res.UserId, &res.Text, &res.CreatedAt,
	)
	if err != nil {
		return &models.CommentRes{}, err
	}
	return &res, nil
}

func (r storagePg) CommentPut(req *models.CommentRes) error {
	query := `
	UPDATE comment SET
		text=$1
	WHERE id=$2
	`
	_, err := r.db.Exec(query, req.Text, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r storagePg) CommentDel(id int) error {
	query := `
	DELETE FROM comment WHERE id=$1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
