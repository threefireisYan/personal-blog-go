package service

import (
	"errors"
	"personal-blog-go/dao"
	"personal-blog-go/models"
	"personal-blog-go/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	//密码Md5加密
	passwd = utils.Md5Crypt(passwd, "sanhuo")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成token jwt技术 生成令牌
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
