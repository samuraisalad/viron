package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// AdminUserMediaType of media type.
var AdminUserMediaType = MediaType("application/vnd.admin_user+json", func() {
	Description("A Admin User")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", Integer, "unique id")
		Attribute("email", String, "login id")
		Attribute("role_id", String, "role id")

		Required("id", "email")
	})

	largeView := func() {
		Attribute("id")
		Attribute("email")
		Attribute("role_id")
	}

	View("default", largeView)
	View("large", largeView)
	View("medium", func() {
		Attribute("id")
		Attribute("email")
		Attribute("role_id")
	})
	View("small", func() {
		Attribute("id")
		Attribute("email")
	})
})

var _ = Resource("admin_user", func() {
	Origin(OriginURL, OriginAllowAll)
	BasePath("/adminuser")
	DefaultMedia(AdminUserMediaType)

	// TODO: ログイン画面できるまでは外しておく
	//Security(JWT, func() {
	//	Scope("api:access")
	//})

	Action("list", func() {
		Description("get admin users")
		Routing(GET(""))
		Response(OK, func() {
			Media(CollectionOf(AdminUserMediaType, func() {
				ContentType("application/json")
				View("default")
				View("large")
				View("medium")
				View("small")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError)
	})

	Action("show", func() {
		Description("get the admin user")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "id")
		})
		Response(OK, func() { Media(AdminUserMediaType) })
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError)
	})

	Action("create", func() {
		Description("create a admin user")
		Routing(POST(""))
		Payload(func() {
			Member("email", String)
			Member("password", String)
		})
		Response(OK, func() { Media(AdminUserMediaType) })
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("update the admin user")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id", Integer, "id")
		})
		Payload(func() {
			Member("password", String)
			Member("role_id", String)
		})
		Response(OK, func() { Media(AdminUserMediaType) })
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError)
	})

	Action("delete", func() {
		Description("delete the user")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError)
	})
})
