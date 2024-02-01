package v1

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	pgLimitCtx  = "pgLimit"
	pgOffsetCtx = "pgOffset"
)

func getPagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlRaw := r.URL.RawQuery
		ctx := r.Context()

		regex := regexp.MustCompile(`\[([^{}]*)]$`)
		regexUrlString := regex.FindStringSubmatch(urlRaw)
		if regexUrlString != nil {
			splitPagination := strings.Split(regexUrlString[1], ":")
			for index, element := range splitPagination {
				if element != "" {
					switch index {
					case 0:
						ctx = context.WithValue(ctx, pgOffsetCtx, element)
					case 1:
						ctx = context.WithValue(ctx, pgLimitCtx, element)
					default:
						fmt.Println("Error getPagination foreach splitPagination")
					}
				}
			}
			//почистим следы пагинации в query запросе
			r.URL.RawQuery = regex.ReplaceAllString(r.URL.RawQuery, "")
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
