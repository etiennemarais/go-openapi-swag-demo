package handler

import "net/http"

// ListDemo godoc
// @Summary ListDemo will return the demos created by the user making the request.
// @Description ListDemo will return the demos created by the user making the request.
// @ID list-demos
// @Tags Demo
// @Accept json
// @Produce json
// @Param offset query int false "Position in template feed" default(0)
// @Param limit query int false "Number of templates to return" default(20)
// @Success 200 "Success"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorised"
// @Security ApiKeyAuth
// @Router /demo [get]
func (h DemoHandler) ListDemo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("List Demo"))
}

// CreateDemo godoc
// @Summary CreateDemo will create a new demo.
// @Description CreateDemo will create a new demo.
// @ID create-demo
// @Tags Demo
// @Accept json
// @Produce json
// @Success 201 "Success"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorised"
// @Failure 404 "Not Found"
// @Failure 409 "Conflict"
// @Failure 500 "Internal error"
// @Security ApiKeyAuth
// @Router /demo/{id} [post]
func (h DemoHandler) CreateDemo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Create Demo"))
}
