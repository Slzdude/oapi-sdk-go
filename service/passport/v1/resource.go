// Code generated by Lark OpenAPI.

package larkpassport

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type V1 struct {
	Session *session // 登录态
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		Session: &session{config: config},
	}
}

type session struct {
	config *larkcore.Config
}

// Query 批量获取用户登录信息（脱敏）
//
// - 该接口用于查询用户的登录信息
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/passport-v1/session/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/passportv1/query_session.go
func (s *session) Query(ctx context.Context, req *QuerySessionReq, options ...larkcore.RequestOptionFunc) (*QuerySessionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/passport/v1/sessions/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySessionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
