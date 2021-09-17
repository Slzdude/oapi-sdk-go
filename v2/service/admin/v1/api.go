// Package admin code generated by lark suite oapi sdk gen
package admin

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v2"
)

type AdminService struct {
	AdminDeptStats *adminDeptStats
	AdminUserStats *adminUserStats
}

func New(app *lark.App) *AdminService {
	a := &AdminService{}
	a.AdminDeptStats = &adminDeptStats{app: app}
	a.AdminUserStats = &adminUserStats{app: app}
	return a
}

type adminDeptStats struct {
	app *lark.App
}
type adminUserStats struct {
	app *lark.App
}

func (a *adminDeptStats) List(ctx context.Context, req *AdminDeptStatListReq, options ...lark.RequestOptionFunc) (*AdminDeptStatListResp, error) {
	rawResp, err := a.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/admin/v1/admin_dept_stats", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &AdminDeptStatListResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *adminUserStats) List(ctx context.Context, req *AdminUserStatListReq, options ...lark.RequestOptionFunc) (*AdminUserStatListResp, error) {
	rawResp, err := a.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/admin/v1/admin_user_stats", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &AdminUserStatListResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
