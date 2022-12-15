package service

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type commentService struct {
	commentRepo  repository.CommentRepo
	reactionRepo repository.CommentReactionRepo // make set reaction for comments
}

func NewCommentService(cRepo repository.CommentRepo, reactionRepo repository.CommentReactionRepo) *commentService {
	log.Println("| | comment service is done!")
	return &commentService{
		commentRepo: cRepo,
	}
}

func (c *commentService) CreateComment(ctx context.Context, comment entity.Comment) (int64, error) {
	return c.commentRepo.CreateComment(ctx, comment)
}

func (c *commentService) GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error) {
	return c.commentRepo.GetCommentByID(ctx, commentID)
}

func (c *commentService) GetCommentsByPostID(ctx context.Context, postID uint64) ([]entity.Comment, error) {
	return c.commentRepo.GetCommentsByPostID(ctx, postID)
}
