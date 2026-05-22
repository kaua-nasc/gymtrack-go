package like

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ToggleLike(postId, userId string) error {
	var likeId string
	queryCheck := `SELECT id FROM public.post_likes WHERE "postId" = $1 AND "userId" = $2`
	err := r.db.QueryRow(queryCheck, postId, userId).Scan(&likeId)

	if err == sql.ErrNoRows {
		// Like
		id := uuid.New().String()
		queryInsert := `INSERT INTO public.post_likes (id, "userId", "postId") VALUES ($1, $2, $3)`
		_, err = r.db.Exec(queryInsert, id, userId, postId)
		return err
	} else if err != nil {
		return err
	}

	// Unlike
	queryDelete := `DELETE FROM public.post_likes WHERE id = $1`
	_, err = r.db.Exec(queryDelete, likeId)
	return err
}
