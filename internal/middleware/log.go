package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogPayloadAndResponse(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	span := trace.SpanFromContext(c.Request.Context())
	span.AddEvent("payload", trace.WithAttributes(
		attribute.String("payload", string(body)),
	))

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()

	// record errors
	if len(c.Errors.Errors()) != 0 {
		for _, err := range c.Errors {
			span.RecordError(err, trace.WithAttributes(attribute.String("error string", err.Error())))
		}
		span.SetStatus(codes.Error, "error occurred")
	}

	span.AddEvent("response", trace.WithAttributes(attribute.String("response", blw.body.String())))
}
