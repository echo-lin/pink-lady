// Package statics embed 静态文件
package statics

import "embed"

// Files 静态文件资源
//go:embed favicon.ico robots.txt
//go:embed css/* font/* html/* img/* js/*
//go:embed i18n/*
var Files embed.FS
