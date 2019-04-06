package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Status",
            Router: `/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
