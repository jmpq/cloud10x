package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	v1 := app.Party("/v1")

	v1.Get("/VCode/{phone}", apiVCodeRequest)

	v1.Get("/Tenants", apiTenantList)
	v1.Get("/Tenants/{id}", apiTenantGet)
	v1.Post("/Tenants", apiTenantPost)
	v1.Delete("/Tenants/{id}", apiUserDelete)

	v1.Get("/Users", apiUserList)
	v1.Get("/Users/{id}", apiUserGet)
	v1.Post("/Users", apiUserPost)
	v1.Delete("/Users/{id}", apiUserDelete)

	v1.Get("/Clusters", apiClusterList)
	v1.Get("/Clusters/{id}", apiClusterGet)

	v1.Get("/Projects", apiProjectList)
	v1.Get("/Projects/{id}", apiProjectGet)

	app.Run(iris.Addr(":8080"))
}
