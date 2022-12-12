package qiniu

import "github.com/qiniu/go-sdk/v7/storage"

// 七牛的存储管理接口
type manager interface {
	// 获取存储桶管理器
	GetBucketManager() (manager *storage.BucketManager, err error)
}

// 获取存储桶管理器
func (q *qiniu) GetBucketManager() (manager *storage.BucketManager, err error) {
	bucketManager := storage.NewBucketManager(q.credentials, nil)
	return bucketManager, nil
}
