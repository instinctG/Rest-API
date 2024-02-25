package coment

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Service - our comment service
type Service struct {
	DB *pgxpool.Pool
}

type Comment struct {
	ID     uint
	Slug   string
	Body   string
	Author string
}

type CommentService interface {
	GetComment(ctx context.Context, ID uint) (Comment, error)
	GetCommentBySlug(ctx context.Context, slug string) ([]Comment, error)
	PostComment(ctx context.Context, comment Comment) (Comment, error)
	UpdateComment(ctx context.Context, ID uint, newComment Comment) (Comment, error)
	DeleteComment(ctx context.Context, ID uint) error
	GetAllComments(ctx context.Context) ([]Comment, error)
}

// NewService - returns a new comments service
func NewService(db *pgxpool.Pool) *Service {
	return &Service{DB: db}
}

func (s *Service) GetComment(ctx context.Context, ID uint) (Comment, error) {
	var comment Comment
	row := s.DB.QueryRow(ctx, `SELECT * FROM comments WHERE id = $1;`, ID)
	err := row.Scan(
		&comment.ID,
		&comment.Slug,
		&comment.Body,
		&comment.Author,
	)

	return comment, err
}

func (s *Service) GetCommentBySlug(ctx context.Context, slug string) ([]Comment, error) {
	var comments []Comment
	rows, err := s.DB.Query(ctx, `SELECT * FROM comments WHERE slug = $1`, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(
			&comment.Slug,
			&comment.Body,
			&comment.Author,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {
	row := s.DB.QueryRow(ctx, `INSERT INTO comments (slug,body,author) VALUES ($1,$2,$3) RETURNING id,slug,body,author`, comment.Slug, comment.Body, comment.Author)

	err := row.Scan(
		&comment.ID,
		&comment.Slug,
		&comment.Body,
		&comment.Author,
	)
	return comment, err
}

func (s *Service) UpdateComment(ctx context.Context, ID uint, newComment Comment) (Comment, error) {
	row := s.DB.QueryRow(ctx, `UPDATE comments SET slug = $2,body = $3,author = $4 WHERE id = $1 RETURNING id,slug,body,author`, ID, newComment.Slug, newComment.Body, newComment.Author)
	var comment Comment
	if err := row.Scan(
		&comment.ID,
		&comment.Slug,
		&comment.Body,
		&comment.Author,
	); err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func (s *Service) DeleteComment(ctx context.Context, ID uint) error {
	_, err := s.DB.Exec(ctx, `DELETE FROM comments WHERE id = $1`, ID)
	return err
}

func (s *Service) GetAllComments(ctx context.Context) ([]Comment, error) {
	rows, err := s.DB.Query(ctx, `SELECT id, slug, body, author FROM comments`)
	var comments []Comment
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.Slug,
			&comment.Body,
			&comment.Author,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
