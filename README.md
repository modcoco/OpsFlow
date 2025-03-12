```bash
docker buildx build --platform linux/amd64 \
  -f container/Dockerfile.gnu \
  -t modco/opsflow:2025.0312.1013 \
  --push \
  .

# Before run
deepseek_r1_pvc_model.yaml
deepseek_r1_cm_runcode.yaml

# After run
deepseek_r1_svc.yaml

```
