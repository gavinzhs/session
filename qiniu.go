package main

import (
	"github.com/qiniu/api/auth/digest"
	qiniuio "github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"log"
)

const (
	QINIU_DOMAIN = "7xk2ua.com1.z0.glb.clouddn.com"
	ACCESS_KEY   = "-rMJ-C0aIB1XSYmi2qCjOv8B2lwgoJH_di1wDQTQ"
	SECRET_KEY   = "19euIosNnN-KIY-mk8rm7nCD_NrfLAq2vFXPL8Cq"

	BUCKET       = "gavin"
	CALLBACK_URL = ""
)

var (
	mac = &digest.Mac{
		AccessKey: ACCESS_KEY,
		SecretKey: []byte(SECRET_KEY),
	}

	c rs.Client
)

func init() {
	c = rs.New(mac)
}

func upFile(bucket string, key string, localPath string) error {
	var ret qiniuio.PutRet

	policy := rs.PutPolicy{
		Scope: bucket + ":" + key,
	}
	err := qiniuio.PutFile(nil, &ret, policy.Token(mac), key, localPath, nil)
	log.Printf("ret : %+v", ret)
	return err
}

func downloadUrl(domain, key string) string {
	baseUrl := rs.MakeBaseUrl(domain, key)
	policy := rs.GetPolicy{}
	return policy.MakeRequest(baseUrl, mac)
}
