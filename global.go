package main

import (
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

var mainDB *gorm.DB
var goCache *cache.Cache
