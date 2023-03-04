package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/minIO"
	"liujun/Time_go-zero_csdn/common/snowflak"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	var user_id int64
	var err error
	switch in.AuthType {
	case model.UserAuthTypeUsername:
		user_id, err = l.UserNameLogin(in.Account, in.Password)
	default:
		user_id, err = l.PhoneLogin(in.Account, in.Password)
	}
	if err != nil {
		return nil, err
	}
	generatoken := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token_reps, err := generatoken.GenerateToken(&user.GenerateTokenRequest{UserId: user_id})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user_id)
	}
	return &user.LoginResponse{
		Token:        token_reps.AccessToken,
		RefreshToken: token_reps.AccessToken,
	}, nil
}

func (l *UserLoginLogic) UserNameLogin(username, password string) (int64, error) {
	user_basic, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, username)
	if err != nil && err != model.ErrNotFound {
		//code = Unknown desc = 根据用户名查询用户信息失败，user_name:19971251761,err:sql: Scan error on column index 8, name \"last_login\": unsupported Scan, storing driver.Value type \u003cnil\u003e into type *time.Time: ErrCode:403，ErrMsg:数据库错误"
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.OTHER_ERROR), "根据用户名查询用户信息失败，user_name:%s,err:%v", username, err)
	}
	if user_basic == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "user_name：%s", username)
	}
	fmt.Println(utils.Md5ByString(password), "=======================", user_basic.Password)
	if utils.Md5ByString(password) != user_basic.Password {
		return 0, errors.Wrapf(ErrUsernamePwdError, "密码匹配出错")
	}
	return user_basic.UserId, nil
}
func (l *UserLoginLogic) PhoneLogin(phone, code string) (int64, error) {
	key := "sms:code:" + phone
	user_basic, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, phone)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.OTHER_ERROR), "根据手机号查询用户信息失败，phone:%s,err:%v", phone, err)
	}
	//手机号不存在则注册
	if user_basic == nil {
		WI, _ := snowflak.NewSnowFlak(1, 2, 0, -1)
		user_id := WI.GetId()
		user_basic := model.UserBasic{
			UserId:       user_id,
			Mobile:       phone,
			Password:     utils.Md5ByString("12345678"),
			ProfilePhoto: minIO.DefaultHeadPhoto,
		}
		l.svcCtx.UserModel.Insert(l.ctx, &user_basic)
	}
	if ok, _ := l.svcCtx.RedisClient.Exists(key); !ok {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.OTHER_ERROR), "验证码已过期")
	}
	val, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.OTHER_ERROR), "验证码获取失败")
	}
	if val != code {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.OTHER_ERROR), "验证码输入错误")
	}
	return user_basic.UserId, nil
}
