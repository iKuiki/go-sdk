package qiniu

import (
	"context"
	"path"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 上传相关接口
type upload interface {
	// 获取七牛上传凭据
	// bucket 要上传的储存区
	// expire 凭据有效期
	GenUploadToken(bucket string, expire time.Duration, remoteFilename ...string) (upToken string, err error)
	// 上传本地文件
	// 可以指定远程文件名
	UploadFile(bucket string, localFilename string, remoteFilename ...string) (ret storage.PutRet, err error)
}

// 获取七牛上传凭据
// bucket 要上传的储存区
// expire 凭据有效期
func (q *qiniu) GenUploadToken(bucket string, expire time.Duration, remoteKey ...string) (upToken string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	if len(remoteKey) > 0 {
		putPolicy.Scope += ":" + remoteKey[0] // 如果有指定文件名，则增加文件名
	}
	putPolicy.Expires = uint64(expire.Seconds()) //示例2小时有效期
	return putPolicy.UploadToken(q.credentials), nil
}

// 上传本地文件
// 可以指定远程文件名
func (q *qiniu) UploadFile(bucket string, localFilename string, remoteKey ...string) (ret storage.PutRet, err error) {
	formUploader := storage.NewFormUploader(&storage.Config{})
	key := path.Base(localFilename)

	if len(remoteKey) > 0 {
		key = remoteKey[0] // 如果有指定文件名，则使用指定文件名
	}
	// 创建上传凭据
	uploadToken, err := q.GenUploadToken(bucket, time.Minute)
	if err != nil {
		return
	}
	err = formUploader.PutFile(context.Background(), &ret, uploadToken, key, localFilename, nil)
	if err != nil {
		err = errors.Wrapf(err, "PutFile(remoteKey:%s,localFile:%s)", key, localFilename)
		return
	}
	return
}
