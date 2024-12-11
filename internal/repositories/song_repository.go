package repositories

import (
	"Songs/Song-library/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SongRepository struct {
	db *pgxpool.Pool
}

func NewSongRepository(db *pgxpool.Pool) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(ctx context.Context, song *models.Song) error {
	query := `
		INSERT INTO songs (group_name, song_name, release_date, text, link, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(ctx, query,
		song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, time.Now(), time.Now())
	return err
}

func (r *SongRepository) GetAll(ctx context.Context) ([]models.Song, error) {
	query := `SELECT id, group_name, song_name, release_date, text, link, created_at, updated_at FROM songs`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongRepository) GetByGroupAndSong(ctx context.Context, group, song string) (*models.Song, error) {
	query := `
		SELECT id, group_name, song_name, release_date, text, link, created_at, updated_at
		FROM songs
		WHERE group_name = $1 AND song_name = $2
	`
	row := r.db.QueryRow(ctx, query, group, song)

	var foundSong models.Song
	err := row.Scan(&foundSong.ID, &foundSong.Group, &foundSong.Song, &foundSong.ReleaseDate, &foundSong.Text, &foundSong.Link, &foundSong.CreatedAt, &foundSong.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &foundSong, nil
}

func (r *SongRepository) Update(ctx context.Context, song *models.Song) error {
	query := `
		UPDATE songs
		SET release_date = $1, text = $2, link = $3, updated_at = $4
		WHERE group_name = $5 AND song_name = $6
	`
	_, err := r.db.Exec(ctx, query, song.ReleaseDate, song.Text, song.Link, time.Now(), song.Group, song.Song)
	return err
}

func (r *SongRepository) DeleteByGroupAndSong(ctx context.Context, group, song string) error {
	query := `
		DELETE FROM songs
		WHERE group_name = $1 AND song_name = $2
	`
	_, err := r.db.Exec(ctx, query, group, song)
	return err
}
