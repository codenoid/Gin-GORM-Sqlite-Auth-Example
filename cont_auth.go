package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

func authLoginHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "auth_login.html", gin.H{})
}

func authLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user UserModel
	tx := mainDB.First(&user, "username = ?", username)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		c.Redirect(http.StatusFound, "/auth/login?error=true&message=User Not Found")
		return
	}
	if !CheckPasswordHash(password, user.Password) {
		c.Redirect(http.StatusFound, "/auth/login?error=true&message=Wrong Password")
		return
	}

	cookie, _ := c.Cookie("guid")
	goCache.Set("session:"+cookie, username, cache.NoExpiration)

	c.Redirect(http.StatusFound, "/")
}

func authMiddleware(c *gin.Context) {

	// kita whitelist route /auth agar tidak di cek middleware
	path := strings.Split(c.FullPath(), "/")
	// /auth/something
	if path[1] == "auth" {
		c.Next()
		return
	}

	cookie, err := c.Cookie("guid")
	if err != nil {
		guid := uuid.New().String()
		c.SetCookie("guid", guid, 60*60*420, "/", "", false, true)
		c.Redirect(http.StatusFound, "/auth/login")
		c.Abort()
		return
	}

	username, found := goCache.Get("session:" + cookie)
	if !found {
		// session invalid / expired
		c.Redirect(http.StatusFound, "/auth/login")
		c.Abort()
		return
	}

	var user UserModel
	tx := mainDB.First(&user, "username = ?", username)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		c.Redirect(http.StatusFound, "/auth/login?error=true&message=User Not Found")
		c.Abort() // karena ini di dalam middleware, maka kita perlu panggil c.Abort()
		return
	}

	c.Set("user", user)
	c.Set("username", user.Username)
	c.Set("name", user.Name)
}
