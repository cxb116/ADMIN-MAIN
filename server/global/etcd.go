package global

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	GVA_ETCD *clientv3.Client
)
