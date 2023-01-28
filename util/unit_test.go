package util

import "testing"

func TestGetMediaInfo(t *testing.T) {
	absPath := "/Users/zen/Downloads/iso/2b5.mp4"
	md, err := GetMediaInfo(absPath)
	if err != nil {
		return
	}
	t.Logf("%+v\n", md.Media.Track[1].Width)
	t.Logf("%+v\n", md.Media.Track[1].Height)

}
