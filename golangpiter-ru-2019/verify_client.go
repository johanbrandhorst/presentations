package main

import "net/http"

func VerifyClient(next http.Handler, allowedCNs ...string) http.Handler {
	m := map[string]struct{}{}
	for _, cn := range allowedCNs {
		m[cn] = struct{}{}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, c := range r.TLS.VerifiedChains {
			if _, ok := m[c[0].Subject.CommonName]; ok {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "client is unauthenticated", http.StatusUnauthorized)
	})
}
