// Auto-generated to Go types and interfaces using avdl-compiler v1.4.2 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/login_ui.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ResetPromptType int

const (
	ResetPromptType_COMPLETE         ResetPromptType = 0
	ResetPromptType_ENTER_NO_DEVICES ResetPromptType = 1
	ResetPromptType_ENTER_FORGOT_PW  ResetPromptType = 2
)

func (o ResetPromptType) DeepCopy() ResetPromptType { return o }

var ResetPromptTypeMap = map[string]ResetPromptType{
	"COMPLETE":         0,
	"ENTER_NO_DEVICES": 1,
	"ENTER_FORGOT_PW":  2,
}

var ResetPromptTypeRevMap = map[ResetPromptType]string{
	0: "COMPLETE",
	1: "ENTER_NO_DEVICES",
	2: "ENTER_FORGOT_PW",
}

func (e ResetPromptType) String() string {
	if v, ok := ResetPromptTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type PassphraseRecoveryPromptType int

const (
	PassphraseRecoveryPromptType_ENCRYPTED_PGP_KEYS PassphraseRecoveryPromptType = 0
)

func (o PassphraseRecoveryPromptType) DeepCopy() PassphraseRecoveryPromptType { return o }

var PassphraseRecoveryPromptTypeMap = map[string]PassphraseRecoveryPromptType{
	"ENCRYPTED_PGP_KEYS": 0,
}

var PassphraseRecoveryPromptTypeRevMap = map[PassphraseRecoveryPromptType]string{
	0: "ENCRYPTED_PGP_KEYS",
}

func (e PassphraseRecoveryPromptType) String() string {
	if v, ok := PassphraseRecoveryPromptTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type GetEmailOrUsernameArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type PromptRevokePaperKeysArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Device    Device `codec:"device" json:"device"`
	Index     int    `codec:"index" json:"index"`
}

type DisplayPaperKeyPhraseArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Phrase    string `codec:"phrase" json:"phrase"`
}

type DisplayPrimaryPaperKeyArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Phrase    string `codec:"phrase" json:"phrase"`
}

type PromptResetAccountArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Kind      ResetPromptType `codec:"kind" json:"kind"`
}

type DisplayResetProgressArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Text      string `codec:"text" json:"text"`
}

type ExplainDeviceRecoveryArg struct {
	SessionID int        `codec:"sessionID" json:"sessionID"`
	Kind      DeviceType `codec:"kind" json:"kind"`
	Name      string     `codec:"name" json:"name"`
}

type PromptPassphraseRecoveryArg struct {
	SessionID int                          `codec:"sessionID" json:"sessionID"`
	Kind      PassphraseRecoveryPromptType `codec:"kind" json:"kind"`
}

type LoginUiInterface interface {
	GetEmailOrUsername(context.Context, int) (string, error)
	PromptRevokePaperKeys(context.Context, PromptRevokePaperKeysArg) (bool, error)
	DisplayPaperKeyPhrase(context.Context, DisplayPaperKeyPhraseArg) error
	DisplayPrimaryPaperKey(context.Context, DisplayPrimaryPaperKeyArg) error
	// Called during login / provisioning flows to ask the user whether they
	// would like to either enter the autoreset pipeline or perform the reset
	// of the account.
	PromptResetAccount(context.Context, PromptResetAccountArg) (bool, error)
	// In some flows the user will get notified of the reset progress
	DisplayResetProgress(context.Context, DisplayResetProgressArg) error
	// During recovery the service might want to explain to the user how they can change
	// their password by using the "change password" functionality on other devices.
	ExplainDeviceRecovery(context.Context, ExplainDeviceRecoveryArg) error
	PromptPassphraseRecovery(context.Context, PromptPassphraseRecoveryArg) (bool, error)
}

func LoginUiProtocol(i LoginUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.loginUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getEmailOrUsername": {
				MakeArg: func() interface{} {
					var ret [1]GetEmailOrUsernameArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]GetEmailOrUsernameArg)
					if !ok {
						err = rpc.NewTypeError((*[1]GetEmailOrUsernameArg)(nil), args)
						return
					}
					ret, err = i.GetEmailOrUsername(ctx, typedArgs[0].SessionID)
					return
				},
			},
			"promptRevokePaperKeys": {
				MakeArg: func() interface{} {
					var ret [1]PromptRevokePaperKeysArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]PromptRevokePaperKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[1]PromptRevokePaperKeysArg)(nil), args)
						return
					}
					ret, err = i.PromptRevokePaperKeys(ctx, typedArgs[0])
					return
				},
			},
			"displayPaperKeyPhrase": {
				MakeArg: func() interface{} {
					var ret [1]DisplayPaperKeyPhraseArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]DisplayPaperKeyPhraseArg)
					if !ok {
						err = rpc.NewTypeError((*[1]DisplayPaperKeyPhraseArg)(nil), args)
						return
					}
					err = i.DisplayPaperKeyPhrase(ctx, typedArgs[0])
					return
				},
			},
			"displayPrimaryPaperKey": {
				MakeArg: func() interface{} {
					var ret [1]DisplayPrimaryPaperKeyArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]DisplayPrimaryPaperKeyArg)
					if !ok {
						err = rpc.NewTypeError((*[1]DisplayPrimaryPaperKeyArg)(nil), args)
						return
					}
					err = i.DisplayPrimaryPaperKey(ctx, typedArgs[0])
					return
				},
			},
			"promptResetAccount": {
				MakeArg: func() interface{} {
					var ret [1]PromptResetAccountArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]PromptResetAccountArg)
					if !ok {
						err = rpc.NewTypeError((*[1]PromptResetAccountArg)(nil), args)
						return
					}
					ret, err = i.PromptResetAccount(ctx, typedArgs[0])
					return
				},
			},
			"displayResetProgress": {
				MakeArg: func() interface{} {
					var ret [1]DisplayResetProgressArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]DisplayResetProgressArg)
					if !ok {
						err = rpc.NewTypeError((*[1]DisplayResetProgressArg)(nil), args)
						return
					}
					err = i.DisplayResetProgress(ctx, typedArgs[0])
					return
				},
			},
			"explainDeviceRecovery": {
				MakeArg: func() interface{} {
					var ret [1]ExplainDeviceRecoveryArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]ExplainDeviceRecoveryArg)
					if !ok {
						err = rpc.NewTypeError((*[1]ExplainDeviceRecoveryArg)(nil), args)
						return
					}
					err = i.ExplainDeviceRecovery(ctx, typedArgs[0])
					return
				},
			},
			"promptPassphraseRecovery": {
				MakeArg: func() interface{} {
					var ret [1]PromptPassphraseRecoveryArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]PromptPassphraseRecoveryArg)
					if !ok {
						err = rpc.NewTypeError((*[1]PromptPassphraseRecoveryArg)(nil), args)
						return
					}
					ret, err = i.PromptPassphraseRecovery(ctx, typedArgs[0])
					return
				},
			},
		},
	}
}

type LoginUiClient struct {
	Cli rpc.GenericClient
}

func (c LoginUiClient) GetEmailOrUsername(ctx context.Context, sessionID int) (res string, err error) {
	__arg := GetEmailOrUsernameArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.loginUi.getEmailOrUsername", []interface{}{__arg}, &res)
	return
}

func (c LoginUiClient) PromptRevokePaperKeys(ctx context.Context, __arg PromptRevokePaperKeysArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.promptRevokePaperKeys", []interface{}{__arg}, &res)
	return
}

func (c LoginUiClient) DisplayPaperKeyPhrase(ctx context.Context, __arg DisplayPaperKeyPhraseArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.displayPaperKeyPhrase", []interface{}{__arg}, nil)
	return
}

func (c LoginUiClient) DisplayPrimaryPaperKey(ctx context.Context, __arg DisplayPrimaryPaperKeyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.displayPrimaryPaperKey", []interface{}{__arg}, nil)
	return
}

// Called during login / provisioning flows to ask the user whether they
// would like to either enter the autoreset pipeline or perform the reset
// of the account.
func (c LoginUiClient) PromptResetAccount(ctx context.Context, __arg PromptResetAccountArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.promptResetAccount", []interface{}{__arg}, &res)
	return
}

// In some flows the user will get notified of the reset progress
func (c LoginUiClient) DisplayResetProgress(ctx context.Context, __arg DisplayResetProgressArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.displayResetProgress", []interface{}{__arg}, nil)
	return
}

// During recovery the service might want to explain to the user how they can change
// their password by using the "change password" functionality on other devices.
func (c LoginUiClient) ExplainDeviceRecovery(ctx context.Context, __arg ExplainDeviceRecoveryArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.explainDeviceRecovery", []interface{}{__arg}, nil)
	return
}

func (c LoginUiClient) PromptPassphraseRecovery(ctx context.Context, __arg PromptPassphraseRecoveryArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.loginUi.promptPassphraseRecovery", []interface{}{__arg}, &res)
	return
}
