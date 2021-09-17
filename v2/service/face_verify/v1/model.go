// Package face_verify code generated by lark suite oapi sdk gen
package face_verify

import (
	"io"

	"github.com/larksuite/oapi-sdk-go/v2"
)

type FaceVerifyCropFaceImageReqBody struct {
	RawImage io.Reader `json:"raw_image,omitempty"`
}

type FaceVerifyCropFaceImageReq struct {
	Body *FaceVerifyCropFaceImageReqBody `body:""`
}

type FaceVerifyCropFaceImageRespData struct {
	FaceImage *string `json:"face_image,omitempty"`
}

type FaceVerifyCropFaceImageResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *FaceVerifyCropFaceImageRespData `json:"data"`
}

type FaceVerifyQueryAuthResultReq struct {
	ReqOrderNo *string `query:"req_order_no"`
	OpenId     *string `query:"open_id"`
	EmployeeId *string `query:"employee_id"`
}

type FaceVerifyQueryAuthResultRespData struct {
	AuthState     *int `json:"auth_state,omitempty"`
	AuthTimpstamp *int `json:"auth_timpstamp,omitempty"`
}

type FaceVerifyQueryAuthResultResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *FaceVerifyQueryAuthResultRespData `json:"data"`
}

type FaceVerifyUploadFaceImageReqBody struct {
	Image io.Reader `json:"image,omitempty"`
}

type FaceVerifyUploadFaceImageReq struct {
	OpenId     *string                           `query:"open_id"`
	EmployeeId *string                           `query:"employee_id"`
	Body       *FaceVerifyUploadFaceImageReqBody `body:""`
}

type FaceVerifyUploadFaceImageRespData struct {
	FaceUid *string `json:"face_uid,omitempty"`
}

type FaceVerifyUploadFaceImageResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *FaceVerifyUploadFaceImageRespData `json:"data"`
}
