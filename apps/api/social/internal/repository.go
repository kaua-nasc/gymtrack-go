package internal

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *Post) error {
	query := `
        INSERT INTO posts (id, "createdAt", "updatedAt", "authorId", "content", "entityId", "entityType")
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := r.db.Exec(query, *post.Id, post.CreatedAt, post.UpdatedAt, post.AuthorId, post.Content, post.EntityId, post.EntityType)
	return err
}

func (r *PostRepository) FindAll(currentUserId string, cursor *utils.CursorData, limit int) ([]Post, *utils.CursorData, error) {
	query := `
		SELECT 
			p.id, p."createdAt", p."updatedAt", p."authorId", p."content", p."entityId", p."entityType",
			(SELECT COUNT(*) FROM public.post_likes WHERE "postId" = p.id) as likes_count,
			(SELECT COUNT(*) FROM public.post_comments WHERE "postId" = p.id) as comments_count,
			EXISTS(SELECT 1 FROM public.post_likes WHERE "postId" = p.id AND "userId" = $1) as liked_by_me
		FROM posts p
		WHERE ("deletedAt" IS NULL) AND ($2::timestamp IS NULL OR p."createdAt" < $2)
		ORDER BY p."createdAt" DESC
		LIMIT $3
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.Query(query, currentUserId, cursorTime, limit)
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}
		posts = append(posts, p)
	}

	var nextCursor *utils.CursorData
	if len(posts) > 0 && len(posts) == limit {
		lastPost := posts[len(posts)-1]
		nextCursor = &utils.CursorData{
			ID:        *lastPost.Id,
			CreatedAt: lastPost.CreatedAt,
		}
	}

	return posts, nextCursor, nil
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

func (r *PostRepository) GetComments(postId string, cursor *utils.CursorData, limit int) ([]Comment, *utils.CursorData, error) {
	query := `
		SELECT id, "createdAt", "updatedAt", content, "authorId", "postId"
		FROM public.post_comments
		WHERE "postId" = $1 AND "deletedAt" IS NULL
		AND ($2::timestamp IS NULL OR "createdAt" > $2)
		ORDER BY "createdAt" ASC
		LIMIT $3
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.Query(query, postId, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	comments := make([]Comment, 0)
	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.Content, &c.AuthorId, &c.PostId)
		if err != nil {
			return nil, nil, err
		}
		comments = append(comments, c)
	}

	var nextCursor *utils.CursorData
	if len(comments) > 0 && len(comments) == limit {
		lastComment := comments[len(comments)-1]
		nextCursor = &utils.CursorData{
			ID:        *lastComment.Id,
			CreatedAt: lastComment.CreatedAt,
		}
	}

	return comments, nextCursor, nil
}
