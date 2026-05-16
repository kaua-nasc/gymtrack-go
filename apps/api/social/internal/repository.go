package internal

import (
	"database/sql"

	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *Post) error {
	query := `
		INSERT INTO public.posts (id, "createdAt", "updatedAt", "authorId", "content", "entityId", "entityType")
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query, post.Id, post.CreatedAt, post.UpdatedAt, post.AuthorId, post.Content, post.EntityId, post.EntityType)
	return err
}

func (r *PostRepository) FindAll(currentUserId string) ([]Post, error) {
	query := `
		SELECT 
			p.id, p."createdAt", p."updatedAt", p."authorId", p."content", p."entityId", p."entityType",
			(SELECT COUNT(*) FROM public.post_likes WHERE "postId" = p.id) as likes_count,
			(SELECT COUNT(*) FROM public.post_comments WHERE "postId" = p.id) as comments_count,
			EXISTS(SELECT 1 FROM public.post_likes WHERE "postId" = p.id AND "userId" = $1) as liked_by_me
		FROM public.posts p
		ORDER BY p."createdAt" DESC
	`
	rows, err := r.db.Query(query, currentUserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]Post, 0)
	for rows.Next() {
		var p Post
		err := rows.Scan(
			&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType,
			&p.LikesCount, &p.CommentsCount, &p.LikedByCurrentUser,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *PostRepository) ToggleLike(postId, userId string) error {
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

func (r *PostRepository) AddComment(comment *Comment) error {
	query := `
		INSERT INTO public.post_comments (id, "createdAt", "updatedAt", "content", "authorId", "postId")
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query, comment.Id, comment.CreatedAt, comment.UpdatedAt, comment.Content, comment.AuthorId, comment.PostId)
	return err
}

func (r *PostRepository) GetComments(postId string) ([]Comment, error) {
	query := `
		SELECT id, "createdAt", "updatedAt", content, "authorId", "postId"
		FROM public.post_comments
		WHERE "postId" = $1
		ORDER BY "createdAt" ASC
	`
	rows, err := r.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]Comment, 0)
	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.Content, &c.AuthorId, &c.PostId)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
