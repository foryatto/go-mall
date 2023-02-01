package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall/app/user/api/internal/logic/user"
	"mall/app/user/api/internal/svc"
	"mall/app/user/api/internal/types"
)

func DelAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddressDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDelAddressLogic(r.Context(), svcCtx)
		resp, err := l.DelAddress(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
