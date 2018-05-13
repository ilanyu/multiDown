package multiDown

import "testing"

func TestDownload(t *testing.T) {
	_, err := Download("http://192.168.123.12:61323/mp3?filepath=/storage/emulated/0/netease/cloudmusic/Cache/Music1/3413713-128000-c9c99b69897e438a9fbc1d25c3a26399.mp3.uc!&decode=true", 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDownloadToFile(t *testing.T) {
	err := DownloadToFile("25703039.mp3", "http://192.168.123.12:61323/mp3?filepath=/storage/emulated/0/netease/cloudmusic/Cache/Music1/3413713-128000-c9c99b69897e438a9fbc1d25c3a26399.mp3.uc!&decode=true", 10)
	if err != nil {
		t.Fatal(err)
	}
}
