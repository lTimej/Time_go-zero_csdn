package sms

import (
	"errors"
	"fmt"

	"github.com/cloopen/go-sms-sdk/cloopen"
)

var (
	api_account = "8a216da874af5fff0174bda90e2405e2"
	api_token   = "6b964382a6994b40b9057c006a551bc2"
	app_id      = "8aaf0708776728df01776a45e6f20044"
	to          = "19971251762"
	template_id = "1"
)

func SendSmsCode(code, phone string) error {
	cfg := cloopen.DefaultConfig().
		WithAPIAccount(api_account).
		WithAPIToken(api_token)
	sms := cloopen.NewJsonClient(cfg).SMS()
	fmt.Println("======code:", code, "+++++++++++")
	input := &cloopen.SendRequest{
		AppId:      app_id,
		To:         to,
		TemplateId: template_id,
		Datas:      []string{code},
	}
	resp, err := sms.Send(input)
	if err != nil {
		return err
	}
	if resp.StatusCode != "000000" {
		return errors.New(resp.StatusMsg)
	}
	return nil
}
