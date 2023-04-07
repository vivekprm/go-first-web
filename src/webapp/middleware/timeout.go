package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()
	ctx, _ = context.WithTimeout(ctx, 3 * time.Second)
	// replace request with new context
	r.WithContext(ctx)

	<- ctx.Done()

	ch := make(chan struct{})

	go func ()  {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <- ch: 
	return
	case <- ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}