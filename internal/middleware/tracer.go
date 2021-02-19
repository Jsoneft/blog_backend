package middleware

import (
	"context"
	"ginblog_backend/global"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(spanCtx), opentracing.Tag{
				Key:   string(ext.Component),
				Value: "HTTP",
			})
		}
		defer span.Finish()

		var (
			traceID     string
			spanID      string
			spanContext = span.Context()
		)

		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerCtx := spanContext.(jaeger.SpanContext)
			traceID = jaegerCtx.TraceID().String()
			spanID = jaegerCtx.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
