package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"log"
)

func ConnectJaegerTracing() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 100,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	tracer, _, err := cfg.NewTracer()

	if err != nil {
		log.Println("error cant initialize jarger tracer :", err.Error())
	}

	opentracing.SetGlobalTracer(tracer)
}
