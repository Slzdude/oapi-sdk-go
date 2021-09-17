// Package optical_char_recognition code generated by lark suite oapi sdk gen
package optical_char_recognition

import (
	"github.com/larksuite/oapi-sdk-go/v2"
)

type ImageBasicRecognizeReqBody struct {
	Image *string `json:"image,omitempty"`
}

type ImageBasicRecognizeReq struct {
	Body *ImageBasicRecognizeReqBody `body:""`
}

type ImageBasicRecognizeRespData struct {
	TextList []string `json:"text_list,omitempty"`
}

type ImageBasicRecognizeResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *ImageBasicRecognizeRespData `json:"data"`
}
