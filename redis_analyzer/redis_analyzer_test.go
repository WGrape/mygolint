package redis_analyzer

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestRedisAnalyzer(t *testing.T) {
	// 在本地运行TestRedisAnalyzer的时候第二个dir参数可以自定义修改为下载依赖的地方, 并把testdata中的测试文件数据也移动过去, 防止出现找不到依赖的问题
	analysistest.Run(t, "/Users/lvsi/go/pkg/mod/github.com/", RedisAnalyzer, "redis_analyzer_testdata")
}
