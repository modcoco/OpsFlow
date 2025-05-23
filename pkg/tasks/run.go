package tasks

import (
	"context"
	"time"

	"github.com/modcoco/OpsFlow/pkg/core"
	"github.com/modcoco/OpsFlow/pkg/crd"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type TaskConfig struct {
	Duration          time.Duration // 任务调度周期
	TaskFunc          TaskFunc      // 任务函数
	WaitForCompletion bool          // 是否等待上一个任务完成
}

func InitializeTasks(clent core.Client, redisClient redis.Cmdable, grpc *grpc.ClientConn) map[string]TaskConfig {
	updateNodeInfoConfig := &QueueConfig{
		Clientset:   clent.Core(),
		RedisClient: redisClient,
		QueueName:   "task_queue",
		PageSize:    50,
	}

	nodeInfoConfig := crd.NodeResourceInfoOptions{
		CRDClient:   clent.DynamicNRI(),
		KubeClient:  clent.Core(),
		GRPCClient:  grpc,
		Parallelism: 3,
	}

	return map[string]TaskConfig{
		"task1": {10 * time.Second, func(ctx context.Context) error {
			return task1Func(ctx)
		}, true},
		"task2": {20 * time.Second, func(ctx context.Context) error {
			return task2Func(ctx)
		}, false},
		"node_heartbeat": {30 * time.Second, func(ctx context.Context) error {
			return NodeHeartbeat(ctx, nodeInfoConfig)
		}, true},
		"add_update_node_info": {
			Duration: 30 * time.Second,
			TaskFunc: func(ctx context.Context) error {
				return UpdateNodeInfo(ctx, updateNodeInfoConfig)
			},
			WaitForCompletion: true,
		},
		"del_node_info": {
			Duration: 40 * time.Second,
			TaskFunc: func(ctx context.Context) error {
				return DeleteNonExistingNodeResourceInfoTask(ctx, nodeInfoConfig)
			},
			WaitForCompletion: true,
		},
	}
}

func StartTaskScheduler(redisClient redis.Cmdable, tasks map[string]TaskConfig) {
	for taskName, taskConfig := range tasks {
		go scheduleTask(context.Background(), redisClient, taskName, taskConfig.Duration, taskConfig.TaskFunc, taskConfig.WaitForCompletion)
	}
}
