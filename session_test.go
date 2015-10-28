package session_test

import (
    "testing"
    "encoding/base64"
    "crypto/rand"
    "io"
    "log")

func TestSessionId(t *testing.T){
    b := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        log.Printf("io.ReadFull err :%+v", err)
        return
    }
    log.Printf("session id :%s", base64.URLEncoding.EncodeToString(b))
}
