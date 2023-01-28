package getInfo

import "testing"

func TestGetOutOfFHD(t *testing.T) {
	dir := "/Users/zen/Downloads/Telegram Desktop/Nier/2B x 9S"
	pattern := "mp4"
	GetOutOfFHD(dir, pattern)

}
