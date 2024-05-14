package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *customResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func RequestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &customResponseWriter{ResponseWriter: w}
		t1 := time.Now()

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("[error-io.ReadAll()] \n%v\n", err)
		}

		defer func() {
			log.Printf("%s from %s - %d in %s \n", fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto), r.RemoteAddr, ww.statusCode, time.Since(t1).Abs().String())

			token, _ := GetTokenFromHeader(r)
			claims := ParseWithoutVerified(token)
			if token != "" && claims != nil {
				log.Printf(`{"@auth":{"user_id"%s}}`, claims.UserID)
			}

			err = r.Body.Close()
			if err != nil {
				log.Printf("[error-body.Close()] \n%v\n", err)
			}

			if !strings.Contains(r.RequestURI, "image") && len(b) > 0 {
				log.Printf(`{"@request":%s}`, string(b))
			}
		}()

		r.Body = io.NopCloser(bytes.NewBuffer(b))

		h.ServeHTTP(ww, r)
	})
}
