package handler

import (
	"augment-gateway/internal/config"

	"github.com/gin-gonic/gin"
)

// Version 系统版本号
const Version = "v0.0.1"

// SystemHandler 系统信息处理器
type SystemHandler struct {
	cfg *config.Config
}

// NewSystemHandler 创建系统信息处理器
func NewSystemHandler(cfg *config.Config) *SystemHandler {
	return &SystemHandler{cfg: cfg}
}

// GetVersion 获取系统版本号
func (h *SystemHandler) GetVersion(c *gin.Context) {
	ResponseSuccess(c, gin.H{
		"version": Version,
	})
}

// GetFrontendConfig 获取前端公开配置
func (h *SystemHandler) GetFrontendConfig(c *gin.Context) {
	ResponseSuccess(c, gin.H{
		"turnstile_enabled": h.cfg.Turnstile.Enabled,
	})
}
