package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/modcoco/OpsFlow/pkg/core"
	"github.com/modcoco/OpsFlow/pkg/handler"
	"github.com/modcoco/OpsFlow/pkg/queue"
	"github.com/modcoco/OpsFlow/pkg/tasks"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func CreateGinRouter(client core.Client) *gin.Engine {
	r := gin.Default()
	r.Use(core.AppContextMiddleware(client))

	r.GET("/api/v1/pod", handler.GetPodInfo)
	r.POST("/api/v1/raycluster", handler.GetCreateRayClusterInfo)
	r.POST("/api/v1/rayjob", handler.CreateRayJobHandle)
	r.GET("/api/v1/rayjob/:namespace/:name", handler.RayJobInfoHandle)
	r.DELETE("/api/v1/rayjob/:namespace/:name", handler.RemoveRayJobHandle)

	return r
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := core.NewClient()
	if err != nil {
		log.Fatalf("Failed to initialize Kubernetes client: %v", err)
	}

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 不启用安全连接
	)
	if err != nil {
		log.Fatalf("did not connect to rpc: %v", err)
	}
	defer conn.Close()

	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"10.187.6.3:31000",
			"10.187.6.4:31001",
			"10.187.6.5:31002",
			"10.187.6.3:31100",
			"10.187.6.4:31101",
			"10.187.6.5:31102",
		},
		Password: "pass12345",
	})
	config := queue.TaskProcessorConfig{
		Clientset:   client.Core(),
		CRDClient:   client.DynamicNRI(),
		RpcConn:     conn,
		RedisClient: redisClient,
		WorkerCount: 1,
		QueueName:   "task_queue",
	}
	go queue.StartTaskQueueProcessor(ctx, config)

	tasksConfig := tasks.InitializeTasks(client, redisClient, conn)
	tasks.StartTaskScheduler(redisClient, tasksConfig)

	r := CreateGinRouter(client)
	if err := r.Run(":8090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
