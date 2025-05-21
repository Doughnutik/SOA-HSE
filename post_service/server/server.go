package server

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gen "post_service/gen"
)

type PostServiceServer struct {
	gen.UnimplementedPostServiceServer
	mutex sync.RWMutex
	posts map[string]*gen.Post
}

func NewPostServiceServer() *PostServiceServer {
	return &PostServiceServer{
		posts: make(map[string]*gen.Post),
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *gen.PostCreateData) (*gen.PostCreateResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := time.Now().Format("20060102150405")
	s.posts[id] = &gen.Post{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		CreatorId:   req.CreatorId,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsPrivate:   req.IsPrivate,
		Tags:        req.Tags,
	}

	return &gen.PostCreateResponse{Id: id}, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *gen.PostGetData) (*gen.PostGetResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	post, exists := s.posts[req.Id]
	if !exists {
		return nil, status.Errorf(status.Code(status.Error(codes.NotFound, "post not found")), "post not found")
	}
	if req.CreatorId != post.GetCreatorId() {
		return nil, status.Errorf(status.Code(status.Error(codes.PermissionDenied, "неверный creatorId")), "неверный creatorId")
	}

	return &gen.PostGetResponse{Post: post}, nil
}

func (s *PostServiceServer) UpdatePost(ctx context.Context, req *gen.PostUpdateData) (*gen.PostUpdateResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	post, exists := s.posts[req.Id]
	if !exists {
		return nil, status.Errorf(status.Code(status.Error(codes.NotFound, "post not found")), "post not found")
	}
	if req.CreatorId != post.GetCreatorId() {
		return nil, status.Errorf(status.Code(status.Error(codes.PermissionDenied, "неверный creatorId")), "неверный creatorId")
	}

	post.Title = req.Title
	post.Description = req.Description
	post.IsPrivate = req.IsPrivate
	post.Tags = req.Tags
	post.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return &gen.PostUpdateResponse{Success: true}, nil
}

func (s *PostServiceServer) DeletePost(ctx context.Context, req *gen.PostDeleteData) (*gen.PostDeleteResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	post, exists := s.posts[req.Id]
	if !exists {
		return nil, status.Errorf(status.Code(status.Error(codes.NotFound, "post not found")), "post not found")
	}
	if req.CreatorId != post.GetCreatorId() {
		return nil, status.Errorf(status.Code(status.Error(codes.PermissionDenied, "неверный creatorId")), "неверный creatorId")
	}

	delete(s.posts, req.Id)
	return &gen.PostDeleteResponse{Success: true}, nil
}

func (s *PostServiceServer) ListPosts(ctx context.Context, req *gen.ListPostsData) (*gen.ListPostsResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var posts []*gen.Post
	for _, post := range s.posts {
		if post.CreatorId == req.CreatorId {
			posts = append(posts, post)
		}
	}

	return &gen.ListPostsResponse{Posts: posts}, nil
}
