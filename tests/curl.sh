curl -X POST http://opsflow-service.idp-kuberay.svc.cluster.local:8080/api/v1/rayjob \
     -H "Content-Type: application/json" \
     -d '{
  "job": {
    "kind": "vllmOnRaySimpleAutoJob",
    "name": "deepseek-r1-671b",
    "args": [
      {
        "label": {
          "vllmRuncodeCustomParams": "true"
        },
        "params": {
          "--tensor-parallel-size": "8",
          "--pipeline-parallel-size": "6",
          "--swap-space": "8"
        }
      }
    ]
  },
  "namespace": "idp-kuberay",
  "rayVersion":"2.43.0",
  "rayImage":"rayproject/ray:2.43.0-vllm_0.7.3",
  "machines": [
    {
      "machineType": "single",
      "name": "ray-head",
      "cpu": "32",
      "memory": "512Gi",
      "customResources": {
        "nvidia.com/gpu": {
          "quantity": "10"
        }
      },
      "volumes": [
        {
          "name": "volume-deepseek-r1",
          "label": {
            "model": "true",
            "actualModelPathInPod": "/mnt/data/models/DeepSeek-R1"
          },
          "mountPath": "/mnt/data/models/DeepSeek-R1",
          "source": {
            "pvc": {
              "claimName": "nfs-pvc-model-deepseek-r1-671b"
            }
          }
        }
      ],
      "ports": [
        {
          "name": "dashboard",
          "containerPort": 8265
        }
      ],
      "isHeadNode": true
    },
    {
      "machineType": "group",
      "name": "ray-worker",
      "cpu": "32",
      "memory": "512Gi",
      "customResources": {
        "nvidia.com/gpu": {
          "quantity": "10"
        }
      },
      "volumes": [
        {
          "name": "volume-deepseek-r1",
          "label": {
            "model": "true"
          },
          "mountPath": "/mnt/data/models/DeepSeek-R1",
          "source": {
            "pvc": {
              "claimName": "nfs-pvc-model-deepseek-r1-671b"
            }
          }
        }
      ],
      "groupName": "workergroup",
      "replicas": 4,
      "minReplicas": 4,
      "maxReplicas": 4
    }
  ]
}'