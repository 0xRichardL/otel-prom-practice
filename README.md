# otel-prom-practice

A hands-on practice project for learning OpenTelemetry with Prometheus integration.

## 📋 Practice Checklist

### Level 1: Fundamentals

#### Understanding the Stack

- [x] Read OpenTelemetry documentation overview
- [x] Understand the three pillars of observability (metrics, traces, logs)
- [x] Learn Prometheus basics (scraping, time-series, PromQL)
- [x] Understand the difference between push vs pull metrics

#### Basic Setup

- [x] Install Prometheus locally
- [x] Run Prometheus and access the UI (localhost:9090)
- [x] Install OpenTelemetry Collector
- [x] Configure a basic OTel Collector pipeline (receiver → processor → exporter)
- [x] Create a simple application in your preferred language (Python/Go/Node.js)

#### First Metrics

- [x] Add OpenTelemetry SDK to your application
- [x] Create a counter metric
- [x] Create a histogram metric
- [x] Create a gauge metric
- [x] Export metrics to console/stdout

---

### Level 2: Integration

#### OTel → Prometheus Integration

- [x] Configure OTel Collector to export metrics in Prometheus format
- [x] Set up Prometheus to scrape OTel Collector endpoint
- [x] Verify metrics appear in Prometheus UI
- [x] Write basic PromQL queries for your metrics
- [x] Create labels/attributes for your metrics

#### Application Instrumentation

- [x] Add automatic instrumentation to your application
- [x] Instrument HTTP endpoints with custom metrics
- [x] Track request duration with histograms
- [x] Track error rates with counters
- [x] Add business-specific metrics (e.g., items processed, queue depth)

#### Visualization

- [x] Install Grafana
- [x] Connect Grafana to Prometheus
- [x] Create your first dashboard
- [x] Build panels for key metrics (QPS, latency, error rate)
- [x] Set up dashboard variables and filters

---

### Level 3: Advanced Practices

#### Distributed Tracing

- [ ] Add trace instrumentation to your application
- [ ] Configure trace sampling strategies
- [ ] Export traces to Jaeger or Tempo
- [ ] Correlate traces with metrics (exemplars)
- [ ] Implement context propagation across services

#### Production-Ready Setup

- [ ] Implement metric cardinality best practices
- [ ] Configure appropriate bucket boundaries for histograms
- [ ] Set up delta temporality vs cumulative temporality
- [ ] Implement resource detection and attributes
- [ ] Configure batch processors for performance
- [ ] Set up memory limiters in OTel Collector

#### Multi-Service Architecture

- [ ] Create a second microservice
- [ ] Implement service-to-service communication
- [ ] Track distributed metrics across services
- [ ] Set up service mesh observability (optional)
- [ ] Implement RED metrics (Rate, Errors, Duration) for all services

#### Advanced Prometheus Features

- [ ] Create recording rules in Prometheus
- [ ] Set up alerting rules
- [ ] Configure Alertmanager
- [ ] Implement metric relabeling
- [ ] Set up Prometheus federation or remote write

---

### Level 4: Expert

#### Custom Components

- [ ] Build a custom OTel Collector processor
- [ ] Create custom metric views and aggregations
- [ ] Implement custom samplers for traces
- [ ] Build a custom exporter

#### Performance & Scale

- [ ] Benchmark OTel Collector throughput
- [ ] Implement sharding strategies
- [ ] Set up high-availability Prometheus
- [ ] Configure Prometheus remote storage
- [ ] Optimize memory usage and cardinality

#### Advanced Scenarios

- [ ] Implement OpenMetrics format
- [ ] Set up cross-cluster metric collection
- [ ] Configure service level objectives (SLOs) using metrics
- [ ] Implement cost attribution with metrics
- [ ] Set up multi-tenancy in OTel Collector

#### Integration & Ecosystem

- [ ] Integrate with cloud provider metrics (AWS/GCP/Azure)
- [ ] Set up logs correlation with traces and metrics
- [ ] Implement OpenTelemetry Protocol (OTLP) end-to-end
- [ ] Configure tail-based sampling
- [ ] Integrate with APM platforms

---

## 🎯 Suggested Learning Path

1. **Week 1-2**: Complete Level 1 - Get comfortable with basic concepts
2. **Week 3-4**: Complete Level 2 - Build integration skills
3. **Week 5-6**: Complete Level 3 - Implement production patterns
4. **Week 7+**: Explore Level 4 - Advanced topics based on interest

## 📚 Resources

- [OpenTelemetry Documentation](https://opentelemetry.io/docs/)
- [Prometheus Documentation](https://prometheus.io/docs/)
- [OTel Collector Documentation](https://opentelemetry.io/docs/collector/)
- [Grafana Getting Started](https://grafana.com/docs/grafana/latest/getting-started/)

## 🚀 Quick Start

```bash
# Start all services (Prometheus, OpenTelemetry Collector, Grafana)
docker-compose up -d

# Check if services are running
docker-compose ps

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

### Access the Services

- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)
- **OpenTelemetry Collector**: Receivers on ports 4317 (gRPC) and 4318 (HTTP)

### First Time Setup in Grafana

1. Open http://localhost:3000
2. Login with username: `admin`, password: `admin`
3. Prometheus datasource is automatically configured
4. Create your first dashboard by clicking "+" → "Dashboard"

---

**Track your progress by checking off items as you complete them!**
