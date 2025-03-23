package queue

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/redis/go-redis/v9"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type TaskProcessorConfig struct {
	Clientset   *kubernetes.Clientset
	CRDClient   dynamic.NamespaceableResourceInterface
	RedisClient *redis.ClusterClient
	WorkerCount int
	QueueName   string
}

func StartTaskQueueProcessor(config TaskProcessorConfig) {
	ctx := context.Background()
	if err := config.RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	taskChannel := make(chan Task)

	go monitorTaskQueue(ctx, config.RedisClient, config.QueueName, taskChannel)

	var wg sync.WaitGroup
	processor := NewTaskProcessor(config.Clientset, &config.CRDClient)

	for i := range config.WorkerCount {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			processTasks(ctx, workerID, taskChannel, processor)
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	close(taskChannel)
	wg.Wait()
	log.Println("All workers have exited.")
}
