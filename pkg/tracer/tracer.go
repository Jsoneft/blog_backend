package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJeagerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Disabled:    false,
		RPCMetrics:  false,
		Tags:        nil,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			QueueSize:                  0,
			BufferFlushInterval:        1 * time.Second,
			LogSpans:                   true,
			LocalAgentHostPort:         agentHostPort,
			DisableAttemptReconnecting: false,
			AttemptReconnectInterval:   0,
			CollectorEndpoint:          "",
			User:                       "",
			Password:                   "",
			HTTPHeaders:                nil,
		},
		Headers:             nil,
		BaggageRestrictions: nil,
		Throttler:           nil,
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
