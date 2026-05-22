package post

import (
	"database/sql"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(post *domain.Post) error {
	query := `
        INSERT INTO posts (id, "createdAt", "updatedAt", "authorId", "content", "entityId", "entityType", "mediaUrls")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `
	_, err := r.db.Exec(query, *post.Id, post.CreatedAt, post.UpdatedAt, post.AuthorId, post.Content, post.EntityId, post.EntityType, pq.Array(post.MediaUrls))
	return err
}

func (r *Repository) FindAll(currentUserId string, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error) {
	query := `
		SELECT 
			p.id, p."createdAt", p."updatedAt", p."authorId", p."content", p."entityId", p."entityType", p."mediaUrls",
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

	posts := make([]domain.Post, 0)
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(
			&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType, pq.Array(&p.MediaUrls),
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

func (r *Repository) FindById(id string) (*domain.Post, error) {
	query := `
		SELECT id, "createdAt", "updatedAt", "authorId", "content", "entityId", "entityType", "mediaUrls"
		FROM posts WHERE id = $1 AND "deletedAt" IS NULL
	`
	var p domain.Post
	err := r.db.QueryRow(query, id).Scan(&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType, pq.Array(&p.MediaUrls))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(post *domain.Post) error {
	query := `UPDATE posts SET content = $1, "updatedAt" = $2 WHERE id = $3`
	_, err := r.db.Exec(query, post.Content, post.UpdatedAt, *post.Id)
	return err
}

func (r *Repository) Delete(id string) error {
	query := `UPDATE posts SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
