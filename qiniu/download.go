package qiniu

import (
	"time"

	"github.com/qiniu/go-sdk/v7/storage"
)

// 下载相关接口
type download interface {
	// 返回公开存储的指定文件URL
	MakePublicURL(domain string, key string) (url string)
	// 获取私有存储的指定文件的URL
	// 仅在指定expire时间段内有效
	MakePrivateURL(domain string, key string, expire time.Duration) (url string)
}

// 返回公开存储的指定文件URL
func (q *qiniu) MakePublicURL(domain string, key string) (url string) {
	return storage.MakePublicURLv2(domain, key)
}

// 获取私有存储的指定文件的URL
// 仅在指定expire时间段内有效
func (q *qiniu) MakePrivateURL(domain string, key string, expire time.Duration) (url string) {
	return storage.MakePrivateURLv2(q.credentials, domain, key, time.Now().Add(expire).Unix())
}
