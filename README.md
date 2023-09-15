# Prometheus Demonstration

Demonstration of an API with `/metrics`, an prometheus scrape, and grafana dashboard.

## Requirements
- Go
- Tilt
- microk8s, kind, or another with local registry

## Steps

### 1 - Create cluster with registry

`./kind-with-registry.sh`

### 2 - Tilt

`tilt up`

### 3 - Open browser in `http://localhost:10350/`, this is Tilt dashboard

### 4 - Login in Grafana: `http://localhost:8080/`

user: admin

password: you get from secrets.

### 5 - Configure a new prometheus Data source

Fill the following field:

Prometheus server URL: http://local-prometheus-server

Save & Test

### 6 - Create a new Dashboard

Create a new visualization

Use the following query, this will show total requests grouped by status code

`sum by (status_code)(rate(http_requests_total{method="GET",service="my-service-name"}[1m]))`

Save and Apply

## Get grafana admin password

`kubectl get secret local-grafana -ojsonpath="{.data.admin-password}" | base64 -d`

## Add scrape to prometheus server

`kubectl edit cm local-prometheus-server`

Change scrape interval

```yaml
global:
  evaluation_interval: 15s
  scrape_interval: 15s
  scrape_timeout: 10s
```

You can make scrape in the same config, or just use the annotation
```yaml
    - job_name: golangapi
      static_configs:
      - targets: ["go-api:9191"]
      metrics_path: "/metrics"
```

The annotation is already configured in pod deployment, so you dont need to configure the scraper

```yaml
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/path: /metrics
    prometheus.io/port: "3000"
```

## Prometheus UI

`kubectl port-forward svc/local-prometheus-server 9191:80`
