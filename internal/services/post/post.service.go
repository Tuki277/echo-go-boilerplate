package post

import (
	"github.com/tuki277/golang-boilerplate/internal/models"
	"github.com/tuki277/golang-boilerplate/internal/requests"
)

type postRepository interface {
	Create(post *models.Post)
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, id int)
	Update(post *models.Post)
	Delete(post *models.Post)
}

type Service struct {
	postRepository postRepository
}

func NewPostService(postRepository postRepository) Service {
	return Service{postRepository: postRepository}
}

func (s Service) Create(post *models.Post) {
	s.postRepository.Create(post)
}

func (s Service) GetPosts(posts *[]models.Post) {
	s.postRepository.GetPosts(posts)
}

func (s Service) GetPost(post *models.Post, id int) {
	s.postRepository.GetPost(post, id)
}

func (s Service) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title

	s.postRepository.Update(post)
}

func (s Service) Delete(post *models.Post) {
	s.postRepository.Delete(post)
}
