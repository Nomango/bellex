package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:InstitutionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/new`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "Close",
            Router: `/:id([0-9]+)/close`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "Start",
            Router: `/:id([0-9]+)/start`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "StartCurrent",
            Router: `/:id([0-9]+)/start/current`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:MechineController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/new`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:ScheduleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/new`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Method: "Post",
            Router: `/new`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "ChangePassword",
            Router: `/password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserController"],
        beego.ControllerComments{
            Method: "ChangeProfile",
            Router: `/profile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserLoginController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserLoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserLoginController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserLoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserLoginController"] = append(beego.GlobalControllerRouter["github.com/nomango/bellex/server/api/v1:UserLoginController"],
        beego.ControllerComments{
            Method: "Status",
            Router: `/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
