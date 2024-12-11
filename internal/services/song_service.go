package services

import (
	"Songs/Song-library/internal/models"
	"Songs/Song-library/internal/repositories"
	"context"
)

type SongService struct {
	repo *repositories.SongRepository
}

func NewSongService(repo *repositories.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(ctx context.Context, song *models.Song) error {
	return s.repo.Create(ctx, song)
}

func (s *SongService) GetAllSongs(ctx context.Context) ([]models.Song, error) {
	return s.repo.GetAll(ctx)
}

func (s *SongService) GetSong(ctx context.Context, group, song string) (*models.Song, error) {
	return s.repo.GetByGroupAndSong(ctx, group, song)
}

func (s *SongService) UpdateSong(ctx context.Context, song *models.Song) error {
	return s.repo.Update(ctx, song)
}

func (s *SongService) DeleteSong(ctx context.Context, group, song string) error {
	return s.repo.DeleteByGroupAndSong(ctx, group, song)
}
