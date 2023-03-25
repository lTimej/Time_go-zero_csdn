package logic

import (
	"context"
	"github.com/pkg/errors"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoEditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoEditLogic {
	return &UserInfoEditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoEditLogic) UserInfoEdit(in *user.UserInfoEditRequest) (*user.UserInfoEditResponse, error) {
	// todo: add your logic here and delete this line
	HeadPhoto := in.HeadPhoto
	OldPwd := in.OldPwd
	NewPwd := in.NewPwd
	UserName := in.UserName
	Gender := in.Gender
	Introduce := in.Introduce
	Tag := in.Tag
	AuthName := in.AuthName
	Birthday := in.Birthday
	Areas := in.Areas
	user_id := in.UserId
	userstr := ""
	userprofilestr := ""
	userinter := []interface{}{}
	profileinter := []interface{}{}
	if HeadPhoto != "" {
		userstr += "profile_photo = ?,"
		userinter = append(userinter, HeadPhoto)
	}
	if NewPwd != "" {
		if OldPwd != "" {
			user_basic, err := l.svcCtx.UserModel.FindOne(l.ctx, user_id)
			if err != nil {
				return nil, err
			}
			if utils.Md5ByString(OldPwd) != user_basic.Password {
				return nil, errors.New("旧密码输入错误")
			} else {
				//userobj.Password = utils.Md5ByString(NewPwd)
				userstr += "password = ?,"
				userinter = append(userinter, NewPwd)
			}
		} else {
			//userobj.Password = utils.Md5ByString(NewPwd)
			userstr += "password = ?,"
			userinter = append(userinter, NewPwd)
		}
	}
	if UserName != "" {
		//userobj.UserName = UserName
		userstr += "user_name = ?,"
		userinter = append(userinter, UserName)
	}
	if Gender != "" {
		//userprofile.Gender = utils.StringToInt64(Gender)
		userprofilestr += "gender = ?,"
		profileinter = append(profileinter, Gender)
	}
	if Introduce != "" {
		//userobj.Introduction = Introduce
		userstr += "introduction = ?,"
		userinter = append(userinter, Introduce)
	}
	if AuthName != "" {
		//userprofile.RealName = AuthName
		userprofilestr += "real_name = ?,"
		profileinter = append(profileinter, AuthName)
	}
	if Tag != "" {
		//userprofile.Tag = Tag
		userprofilestr += "tag = ?,"
		profileinter = append(profileinter, Tag)
	}
	if Birthday != "" {
		//userprofile.Birthday = utils.StringToTime(Birthday)
		userprofilestr += "birthday = ?,"
		profileinter = append(profileinter, Birthday)
	}
	if Areas != "" {
		//userprofile.Area = Areas
		userprofilestr += "area = ?,"
		profileinter = append(profileinter, Areas)
	}

	if userstr != "" {
		userstr = userstr[:len(userstr)-1]
		userinter = append(userinter, user_id)
		l.svcCtx.UserModel.UpdateUserInfo(l.ctx, user_id, userstr, userinter)
	}
	if userprofilestr != "" {
		userprofilestr = userprofilestr[:len(userprofilestr)-1]
		profileinter = append(profileinter, user_id)
		l.svcCtx.UserProfileModel.UpdateProfileInfo(l.ctx, user_id, userprofilestr, profileinter)
	}
	return &user.UserInfoEditResponse{
		HeadPhoto: "http://172.20.16.20:9000/" + HeadPhoto,
		Pwd:       "1",
		UserName:  UserName,
		Gender:    utils.StringToInt64(Gender),
		Introduce: Introduce,
		Tag:       Tag,
		AuthName:  AuthName,
		Birthday:  Birthday,
		Areas:     Areas,
	}, nil
}
