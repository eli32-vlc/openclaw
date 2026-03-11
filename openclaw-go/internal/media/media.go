package media

// MediaType identifies the type of media.
type MediaType string

const (
	MediaTypeImage    MediaType = "image"
	MediaTypeVideo    MediaType = "video"
	MediaTypeAudio    MediaType = "audio"
	MediaTypeDocument MediaType = "document"
	MediaTypeSticker  MediaType = "sticker"
)

// Media represents a media attachment.
type Media struct {
	Type     MediaType `json:"type"`
	URL      string    `json:"url,omitempty"`
	Data     []byte    `json:"data,omitempty"`
	MimeType string    `json:"mimeType,omitempty"`
	Filename string    `json:"filename,omitempty"`
	Size     int64     `json:"size,omitempty"`
}

// ProcessMedia processes a media item (resize, transcode, etc.).
func ProcessMedia(m *Media) (*Media, error) {
	// In a full implementation, this would resize images, transcode videos, etc.
	return m, nil
}

// IsImage returns true if the media is an image.
func (m *Media) IsImage() bool {
	return m.Type == MediaTypeImage
}

// IsVideo returns true if the media is a video.
func (m *Media) IsVideo() bool {
	return m.Type == MediaTypeVideo
}

// IsAudio returns true if the media is audio.
func (m *Media) IsAudio() bool {
	return m.Type == MediaTypeAudio
}
