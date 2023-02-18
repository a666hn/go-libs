package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// =====================================
// Create Enum Sort Direction ==========
// =====================================

type OrderDirection uint8

const (
	Ascending OrderDirection = iota + 1
	Descending
)

var mapStringDirection = map[string]OrderDirection{
	"ASC":  Ascending,
	"DESC": Descending,
}

var mapDirectionString = map[OrderDirection]string{
	Ascending:  "ASC",
	Descending: "DESC",
}

func (w *OrderDirection) Index() OrderDirection {
	return *w
}

func (w *OrderDirection) Pointer() *OrderDirection {
	return w
}

func (w *OrderDirection) String() string {
	return mapDirectionString[*w]
}

// BuildOrderDirectionFromString create order direction from one of value [asc,desc]
func BuildOrderDirectionFromString(
	drc string,
	dir interface{},
) error {
	dc := strings.ToUpper(drc)
	if d, ok := mapStringDirection[dc]; ok {
		dir = d.Pointer()
		return nil
	}

	return errors.New("invalid value of order direction, should one of \"asc\" or \"desc\"")
}

// =======================================
// Create Pagination Middleware ==========
// =======================================

type MiddlewarePagination struct {
	minPageLoad int
	maxPageLoad int
}

func (m *MiddlewarePagination) MinPageLoad() int {
	return m.minPageLoad
}

func (m *MiddlewarePagination) MaxPageLoad() int {
	return m.maxPageLoad
}

type MiddlewarePaginationOptions func(p *MiddlewarePagination)

func SetMinimumPageLoaded(val int) MiddlewarePaginationOptions {
	return func(p *MiddlewarePagination) {
		p.minPageLoad = val
	}
}

func SetMaximumPageLoaded(val int) MiddlewarePaginationOptions {
	return func(p *MiddlewarePagination) {
		p.maxPageLoad = val
	}
}

func InitPagination(opts ...MiddlewarePaginationOptions) *MiddlewarePagination {
	opt := &MiddlewarePagination{}

	// Default setting for min and max page loaded from pagination
	opt.minPageLoad = 5
	opt.maxPageLoad = 50

	for _, option := range opts {
		option(opt)
	}

	return opt
}

type ctxPagination uint8

const (
	CtxPagination ctxPagination = iota + 1
)

type ContextPagination struct {
	Page           int
	PerPage        int
	OrderDirection *OrderDirection
	OrderKey       string
	Keyword        string
}

func WithPagination(pagination *MiddlewarePagination) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			var initPage = 1
			var initPerPage = 5
			var direction OrderDirection

			min := pagination.MinPageLoad()
			max := pagination.MaxPageLoad()

			// Get Pagination parameter from url query
			p := getQueryParam(r, "page")
			pp := getQueryParam(r, "perPage")
			sd := getQueryParam(r, "sortDirection")
			st := getQueryParam(r, "sortType")
			k := getQueryParam(r, "keyword")

			// Try to transform to int type for Page
			if page, err := strconv.Atoi(p); err == nil {
				if page >= initPage {
					initPage = page
				}
			}

			// Try to transform to int type for PerPage
			if perPage, err := strconv.Atoi(pp); err == nil {
				if perPage >= min && perPage <= max {
					initPerPage = perPage
				}
			}

			if sd != "" {
				if err := BuildOrderDirectionFromString(sd, &direction); err != nil {
					mesErr := map[string]interface{}{
						"message": "Invalid value",
						"error":   err.Error(),
					}
					m, _ := json.Marshal(mesErr)

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(500)
					_, _ = w.Write(m)
					return
				}
			}

			contextVal := ContextPagination{
				Page:           initPage,
				PerPage:        initPerPage,
				OrderDirection: &direction,
				OrderKey:       st,
				Keyword:        k,
			}

			ctx := r.Context()

			ctx = context.WithValue(ctx, CtxPagination, contextVal)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func getQueryParam(
	r *http.Request,
	key string,
) string {
	return r.URL.Query().Get(key)
}

// =================================
// Get Pagination Context ==========
// =================================

func GetPaginationContext(ctx context.Context) ContextPagination {
	val, ok := ctx.Value(CtxPagination).(ContextPagination)
	if !ok {
		panic("Unimplemented pagination middleware, setup your route using middleware \"WithPagination\"")
	}
	return val
}
