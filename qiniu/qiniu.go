package qiniu

import (
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	qiniuqbox "github.com/qiniu/go-sdk/v7/auth/qbox"
)

// Qiniu 七牛sdk
type Qiniu interface {
	// 组合相关接口
	upload   // 上传
	download // 下载
	manager  // 管理
}

type qiniu struct {
	credentials *auth.Credentials
}

// NewQiniu 创建七牛sdk服务
func NewQiniu(accessKey, secretKey string) (Qiniu, error) {
	if accessKey == "" || secretKey == "" {
		return nil, errors.New("accessKey or secretKey empty")
	}
	qiniuMac := qiniuqbox.NewMac(accessKey, secretKey)
	return &qiniu{credentials: qiniuMac}, nil
}
