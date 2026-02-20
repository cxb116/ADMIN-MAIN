package initialize

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	file, err := os.Open("go.mod")
	if err == nil && global.GVA_CONFIG.AutoCode.Module == "" {
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		global.GVA_CONFIG.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
	}

	// 初始化 etcd 客户端
	etcdClient()
}

// etcdClient 初始化 etcd 客户端
func etcdClient() {
	etcdCfg := global.GVA_CONFIG.Etcd

	// 如果未启用，跳过初始化
	if !etcdCfg.Enabled {
		global.GVA_LOG.Info("etcd 未启用，跳过初始化")
		return
	}

	// 创建 etcd 客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdCfg.Endpoints,
		DialTimeout: time.Duration(etcdCfg.Timeout) * time.Second,
		Username:    etcdCfg.Username,
		Password:    etcdCfg.Password,
	})

	if err != nil {
		global.GVA_LOG.Error("etcd 初始化失败",
			zap.String("endpoints", fmt.Sprintf("%v", etcdCfg.Endpoints)),
			zap.Error(err))
		// 不影响服务启动，只记录日志
		return
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = client.Get(ctx, "test")
	if err != nil && err != context.DeadlineExceeded {
		global.GVA_LOG.Warn("etcd 连接测试失败，服务将继续运行", zap.Error(err))
	}

	global.GVA_ETCD = client
	global.GVA_LOG.Info("etcd 初始化成功",
		zap.String("endpoints", fmt.Sprintf("%v", etcdCfg.Endpoints)),
		zap.String("prefix", etcdCfg.Prefix))
}
