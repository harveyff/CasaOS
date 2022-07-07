/*
 * @Author: LinkLeong link@icewhale.com
 * @Date: 2022-07-06 09:46:06
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-07-06 10:29:17
 * @FilePath: /CasaOS/graph/resolver/user.resolvers.go
 * @Description:
 * @Website: https://www.casaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package graph_resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	graph_generated_model "github.com/IceWhaleTech/CasaOS/graph/generated/model"
	"github.com/IceWhaleTech/CasaOS/pkg/utils/common_err"
	"github.com/IceWhaleTech/CasaOS/pkg/utils/encryption"
	"github.com/IceWhaleTech/CasaOS/pkg/utils/jwt"
	"github.com/IceWhaleTech/CasaOS/service"
	model2 "github.com/IceWhaleTech/CasaOS/service/model"
)

// SetNamePwd is the resolver for the setNamePwd field.
func (r *mutationResolver) SetNamePwd(ctx context.Context, username string, pwd string) (*graph_generated_model.SetNamePwdOutput, error) {
	m := graph_generated_model.SetNamePwdOutput{}
	message := ""
	success := 200
	if len(username) == 0 || len(pwd) == 0 {
		success = common_err.INVALID_PARAMS
		message = common_err.GetMsg(common_err.INVALID_PARAMS)
		m.Message = &message
		m.Success = &success
		return &m, nil
	}
	user := model2.UserDBModel{}
	user.UserName = username
	user.Password = encryption.GetMD5ByStr(pwd)
	user.Role = "admin"

	user = service.MyService.User().CreateUser(user)
	if user.Id == 0 {
		success = common_err.ERROR
		message = common_err.GetMsg(common_err.ERROR)
		m.Message = &message
		m.Success = &success
		return &m, nil
	}

	m.Message = &message
	m.Success = &success
	return &m, nil
}

// UserLogin is the resolver for the userLogin field.
func (r *mutationResolver) UserLogin(ctx context.Context, username string, pwd string) (*graph_generated_model.UserLoginOutput, error) {
	m := graph_generated_model.UserLoginOutput{}
	message := ""
	success := 200
	if len(username) == 0 || len(pwd) == 0 {
		success = common_err.INVALID_PARAMS
		message = common_err.GetMsg(common_err.INVALID_PARAMS)
		m.Message = &message
		m.Success = &success
		return &m, nil
	}
	user := service.MyService.User().GetUserAllInfoByName(username)
	if user.Id == 0 {
		success = common_err.USER_NOT_EXIST
		message = common_err.GetMsg(common_err.USER_NOT_EXIST)
		m.Message = &message
		m.Success = &success
		return &m, nil
	}
	if user.Password != encryption.GetMD5ByStr(pwd) {
		success = common_err.PWD_INVALID
		message = common_err.GetMsg(common_err.PWD_INVALID)
		m.Message = &message
		m.Success = &success
		return &m, nil
	}

	success = common_err.SUCCESS
	message = common_err.GetMsg(common_err.SUCCESS)
	m.Message = &message
	m.Success = &success
	var v string = "v0.3.3"
	token := jwt.GetToken(username, pwd)
	user.Password = ""
	id := strconv.Itoa(user.Id)
	u := graph_generated_model.User{}
	u.Avatar = &user.Avatar
	u.Description = &user.Description
	u.Email = &user.Email
	u.ID = &id
	u.Username = &user.UserName
	u.Nickname = &user.NickName
	u.Role = &user.Role
	lo := graph_generated_model.UserLoginData{}
	lo.Token = &token
	lo.User = &u
	lo.Version = &v
	m.Data = &lo

	return &m, nil
}

// GetUserInfo is the resolver for the getUserInfo field.
func (r *queryResolver) GetUserInfo(ctx context.Context, id string) (*graph_generated_model.GetUserInfoOutput, error) {

	m := graph_generated_model.GetUserInfoOutput{}
	message := ""
	success := 200
	user := service.MyService.User().GetUserInfoById(id)

	//*****
	// var u = make(map[string]string, 5)
	// u["user_name"] = user.UserName
	// u["head"] = user.Avatar
	// u["email"] = user.Email
	// u["description"] = user.NickName
	// u["nick_name"] = user.NickName
	// u["id"] = strconv.Itoa(user.Id)

	m.Message = &message
	m.Success = &success

	uid := strconv.Itoa(user.Id)
	u := graph_generated_model.User{}
	u.Avatar = &user.Avatar
	u.Description = &user.Description
	u.Email = &user.Email
	u.ID = &uid
	u.Username = &user.UserName
	u.Nickname = &user.NickName
	u.Role = &user.Role
	m.Data = &u
	return &m, nil
}
