package handler

import (
	"net/http"

	"github.com/bjhuang103/go_pay/union_gateway/internal/logic"
	"github.com/bjhuang103/go_pay/union_gateway/internal/svc"
	"github.com/bjhuang103/go_pay/union_gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), ctx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
