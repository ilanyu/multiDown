package multiDown

import "testing"

func TestDownload(t *testing.T) {
	_, err := Download("http://music.163.com/song/media/outer/url?id=25703039&amp;userid=57263927", 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDownloadToFile(t *testing.T) {
	err := DownloadToFile("25703039.mp3", "http://music.163.com/song/media/outer/url?id=25703039&amp;userid=57263927", 10)
	if err != nil {
		t.Fatal(err)
	}
}
