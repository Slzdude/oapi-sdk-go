// Package speech_to_text code generated by lark suite oapi sdk gen
package speech_to_text

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v2"
)

type SpeechToTextService struct {
	Speech *speech
}

func New(app *lark.App) *SpeechToTextService {
	s := &SpeechToTextService{}
	s.Speech = &speech{app: app}
	return s
}

type speech struct {
	app *lark.App
}

func (s *speech) FileRecognize(ctx context.Context, req *SpeechFileRecognizeReq, options ...lark.RequestOptionFunc) (*SpeechFileRecognizeResp, error) {
	rawResp, err := s.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/speech_to_text/v1/speech/file_recognize", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &SpeechFileRecognizeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *speech) StreamRecognize(ctx context.Context, req *SpeechStreamRecognizeReq, options ...lark.RequestOptionFunc) (*SpeechStreamRecognizeResp, error) {
	rawResp, err := s.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/speech_to_text/v1/speech/stream_recognize", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &SpeechStreamRecognizeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
