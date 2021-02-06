package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func dashboardHTML(c *gin.Context) {
	c.String(http.StatusOK, "Halo, Selamat Datang "+c.GetString("name"))
	// Halo, Selamat Datang Rubi
}
