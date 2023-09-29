package tracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"os"
)

func ConnectJaegerTracing() (opentracing.Tracer, io.Closer, error) {
	jaegerHost := os.Getenv("JAEGER_HOST")

	cfg := config.Configuration{
		ServiceName: "tiket-service",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 100,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%v:6831", jaegerHost),
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	return tracer, closer, err
}
