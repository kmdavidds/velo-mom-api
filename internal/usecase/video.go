package usecase

import (
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"github.com/valentinusdelvin/velo-mom-api/pkg/video"
)

type InterVideoUsecase interface {
	CreateVideo(param models.CreateVideo) error
	DeleteVideo(id string) error
	GetVideos(page, size int) ([]models.CreateVideo, error)
	GetVideoByID(id string) (entity.Video, error)
	GetVideosBySearch(param models.CreateVideo, page, size int) ([]entity.Video, error)
	GetVideoByFilter(param models.CreateVideo, page, size int) ([]entity.Video, error)
}

type VideoUsecase struct {
	vrsc repository.InterVideoRepository
}

func NewVideoUsecase(videoRepo repository.InterVideoRepository) InterVideoUsecase {
	return &VideoUsecase{
		vrsc: videoRepo,
	}
}

func (v *VideoUsecase) CreateVideo(param models.CreateVideo) error {
	vidID := video.ExtractVideoID(param.YoutubeURL)
	thumbnailIMG := video.GenerateThumbnail(vidID)

	videoPost := entity.Video{
		Title:       param.Title,
		YoutubeURL:  param.YoutubeURL,
		YoutubeID:   vidID,
		Thumbnail:   thumbnailIMG,
		Description: param.Description,
		Filter:      entity.Filter(param.Filter),
	}
	_, err := v.vrsc.CreateVideo(videoPost)
	if err != nil {
		return err
	}
	return nil
}

func (v *VideoUsecase) DeleteVideo(id string) error {
	return v.vrsc.DeleteVideo(id)
}

func (v *VideoUsecase) GetVideos(page, size int) ([]models.CreateVideo, error) {
	videos, err := v.vrsc.GetVideos(page, size)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoUsecase) GetVideoByID(id string) (entity.Video, error) {
	video, err := v.vrsc.GetVideoByID(id)
	if err != nil {
		return entity.Video{}, err
	}
	return video, nil
}

func (v *VideoUsecase) GetVideosBySearch(param models.CreateVideo, page, size int) ([]entity.Video, error) {
	videos, err := v.vrsc.GetVideosBySearch(param, page, size)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoUsecase) GetVideoByFilter(param models.CreateVideo, page, size int) ([]entity.Video, error) {
	videos, err := v.vrsc.GetVideoByFilter(param, page, size)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
