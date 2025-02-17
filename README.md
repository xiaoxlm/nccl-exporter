# 构建二进制
```shell
make build-amd-linux
```

# 构建镜像
```shell
 make build-image
```

# docker 运行
```shell
docker run -d -p 9134:9134 \
 --name  nccl-exporter \
 -e NCCL_METRICS_LABEL=nccl_error \
 -e LOKI_URL="your loki url" \
 nccl-exporter:v1.0.0
```
