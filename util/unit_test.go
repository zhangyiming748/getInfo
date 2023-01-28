package util

import "testing"

func TestGetMediaInfo(t *testing.T) {

	absPath := "/Users/zen/Downloads/Telegram Desktop/Nier/2B x 9S/2B #4 [4K].mp4"
	md, err := GetMediaInfo(absPath)
	if err != nil {
		return
	}
	t.Logf("%+v\n", md)
}
