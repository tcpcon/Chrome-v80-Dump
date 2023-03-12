package chromium

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func randHex(len int) string {
	bytes := make([]byte, len)

	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

func chromeTimeToDateTime(microseconds int) string {
	return time.Unix(time.Date(1601, time.January, 1, 0, 0, 0, 0, time.UTC).Unix() + int64(microseconds / 1000000), 0).Format("2006-01-02 15:04:05")
}
