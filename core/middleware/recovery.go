package middleware

import (
	"GOSecretProject/core/utils/custom_http"
	"context"
	"encoding/json"
	"github.com/kataras/golog"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type RecoveryHandler struct{}

func NewMiddleware() *RecoveryHandler {
	return &RecoveryHandler{}
}

var numbers = []rune("1234567890")

func genRequestNumber(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(s)
}

var rIDKey = "rID"

func formatPath(path string) string {
	pathArray := strings.Split(path[1:], "/")
	for i := range pathArray {
		if _, err := strconv.Atoi(pathArray[i]); err == nil {
			pathArray[i] = "*"
		}
	}
	return "/" + strings.Join(pathArray, "/")
}

func (m *RecoveryHandler) ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (m *RecoveryHandler) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := custom_http.NewStatusResponseWriter(w)

		requestNumber := genRequestNumber(6)

		r = r.WithContext(context.WithValue(r.Context(), rIDKey, requestNumber))

		golog.Infof("#%s: %s %s", requestNumber, r.Method, r.URL)

		next.ServeHTTP(sw, r)

		golog.Infof("#%s: code %d", requestNumber, sw.StatusCode)
	})
}

func (m *RecoveryHandler) RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rID, ok := r.Context().Value("rID").(string)

			err := recover()
			if err != nil {
				if ok {
					golog.Errorf("#%s Panic: %s", rID, err.(error).Error())
				} else {
					golog.Errorf("Panic with no id: %s", err.(error).Error())
				}

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal haha error",
				})

				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
