```bash
docker buildx build --platform linux/amd64 \
  -f container/Dockerfile.gnu \
  -t modco/opsflow:2025.0302.1921 \
  --push \
  .

```
