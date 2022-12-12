package qiniu_test

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/iKuiki/go-sdk/qiniu"
	"github.com/stretchr/testify/assert"
)

var q qiniu.Qiniu

const ( // 如需测试，填入以下信息
	accessKey = ""
	secretKey = ""
	// 七牛的存储桶
	bucket = ""
	// 该桶的外网访问域名
	domain = ""
)

func init() {
	var err error
	q, err = qiniu.NewQiniu(accessKey, secretKey)
	if err != nil {
		panic(err)
	}
}

const (
	testFilename = "testuploadfile.txt"
)

func TestUpload(t *testing.T) {
	// 构造测试用上传文件
	err := ioutil.WriteFile(testFilename, []byte("this is a test file"), fs.ModePerm)
	assert.NoError(t, err)
	defer os.Remove(testFilename) // 随后移除该文件
	// 上传该文件到云
	ret, err := q.UploadFile(bucket, testFilename)
	assert.NoError(t, err)
	t.Log(ret)
	// 获取文件管理器
	manager, err := q.GetBucketManager()
	assert.NoError(t, err)
	// 查看该文件状态
	stat, err := manager.Stat(bucket, ret.Key)
	assert.NoError(t, err)
	t.Log(stat)
	// 获取该文件外网URL
	remoteURL := q.MakePrivateURL(domain, ret.Key, time.Minute)
	t.Log(remoteURL)
	// 最后删除该文件
	err = manager.Delete(bucket, ret.Key)
	assert.NoError(t, err)
}
