package controllers

import (
	"github.com/jailcomfranssa/Go_GORM_Postgres_Mysql_Testing/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
