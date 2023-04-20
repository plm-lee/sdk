package utils

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func UUID(n int) string {
	return RandStringRunes(n)
}

// 1041267613438251008
func MsgId(n int) int64 {
	uid := RandStringRunes(n)
	strings.TrimLeft(uid, "0")
	msgid, err := strconv.Atoi(uid)
	if err != nil {
		log.Println("get msgId err:", err)
	}
	return int64(msgid)
}
