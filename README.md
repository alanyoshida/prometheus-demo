# Prometheus Demonstration

## Get grafana admin password

`kubectl get secret local-grafana -ojsonpath="{.data.admin-password}" | base64 -d`

## Add scrape to prometheus server

`kubectl edit cm local-prometheus-server`

```yaml
    - job_name: golangapi
      static_configs:
      - targets: ["go-api:9191"]
      metrics_path: "/metrics"
```

## Prometheus UI

`kubectl port-forward svc/local-prometheus-server 9191:80`
