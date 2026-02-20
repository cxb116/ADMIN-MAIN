package dsp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

// EtcdSyncService etcd 同步服务
type EtcdSyncService struct {
	prefix string
}

// NewEtcdSyncService 创建 etcd 同步服务
func NewEtcdSyncService() *EtcdSyncService {
	prefix := "/dsp/config"
	if global.GVA_CONFIG.Etcd.Prefix != "" {
		prefix = global.GVA_CONFIG.Etcd.Prefix
	}
	return &EtcdSyncService{prefix: prefix}
}

// IsEnabled 检查 etcd 是否可用
func (s *EtcdSyncService) IsEnabled() bool {
	return global.GVA_ETCD != nil && global.GVA_CONFIG.Etcd.Enabled
}

// Put 写入数据到 etcd
func (s *EtcdSyncService) Put(ctx context.Context, key string, value interface{}) error {
	if !s.IsEnabled() {
		return fmt.Errorf("etcd 未启用或未初始化")
	}

	// 序列化为 JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("序列化失败: %w", err)
	}

	// 生成完整的 key
	fullKey := s.prefix + "/" + key

	// 写入 etcd（带超时）
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err = global.GVA_ETCD.Put(timeoutCtx, fullKey, string(jsonData))
	if err != nil {
		return fmt.Errorf("写入 etcd 失败: %w", err)
	}

	global.GVA_LOG.Info("写入 etcd 成功",
		zap.String("key", fullKey),
		zap.String("value", string(jsonData)))

	return nil
}

// Delete 从 etcd 删除数据
func (s *EtcdSyncService) Delete(ctx context.Context, key string) error {
	if !s.IsEnabled() {
		return fmt.Errorf("etcd 未启用或未初始化")
	}

	// 生成完整的 key
	fullKey := s.prefix + "/" + key

	// 删除（带超时）
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := global.GVA_ETCD.Delete(timeoutCtx, fullKey)
	if err != nil {
		return fmt.Errorf("删除 etcd 失败: %w", err)
	}

	global.GVA_LOG.Info("删除 etcd 成功", zap.String("key", fullKey))
	return nil
}

// generateKey 生成 etcd key
func (s *EtcdSyncService) generateKey(entityType, action string, id int64) string {
	return fmt.Sprintf("%s/%s/%d", entityType, action, id)
}
