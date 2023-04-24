package httpResp

import (
	"liujun/Time_go-zero_csdn/common/xerr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func HttpResp(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusCreated, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}
		logx.WithContext(r.Context()).Errorf("【===========API-ERR===========】 : %+v ", err)
		errs := Error(errcode, errmsg)
		httpx.WriteJson(w, int(errs.Code), errs)
	}
}
