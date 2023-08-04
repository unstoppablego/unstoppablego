package controls

import "github.com/unstoppablego/framework/httpapi"

type ReqGet struct {
	Id    string
	Hello string
	World string `json:"world"`
}

func Api() {
	httpapi.Get[ReqGet]("/fuck", Fuckfunc[ReqGet], true)

	httpapi.Post[ReqGet]("/fuck2", httpapi.CustomXSSMiddleWare[ReqGet](Fuckfunc[ReqGet]), true)

	httpapi.Get[ReqGet]("/fuck3", func(ctx *httpapi.Context, query ReqGet) (interface{}, error) {

		return query, nil
	}, true)

	httpapi.AddFileUpload("/upload")

}

func Fuckfunc[ReqGet any](ctx *httpapi.Context, query ReqGet) (interface{}, error) {

	return query, nil
}
