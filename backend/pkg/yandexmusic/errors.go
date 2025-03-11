package yandexmusic

import "errors"

var (
	ErrInvalidToken   = errors.New("invalid OAuth token")
	ErrTrackNotFound  = errors.New("track not found")
	ErrPlaybackFailed = errors.New("playback failed")
	ErrSearchFailed   = errors.New("search failed")
)
