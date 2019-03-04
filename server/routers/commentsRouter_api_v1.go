package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}