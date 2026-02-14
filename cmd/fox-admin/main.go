// Copyright by AcmaTvirus
package main

import (
	"io/fs"
	"net/http"
	"log"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/acmavirus/foxdocker-panel"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/disk"
)

type SystemStats struct {
	CPU    float64 `json:"cpu"`
	RAM    float64 `json:"ram"`
	Disk   float64 `json:"disk"`
	Uptime string  `json:"uptime"`
}

type FirewallRule struct {
	ID      int    `json:"id"`
	IP      string `json:"ip"`
	Action  string `json:"action"`
	Reason  string `json:"reason"`
	Target  string `json:"target"`
	Time    string `json:"time"`
}

type AttackTrend struct {
	Hour  string `json:"hour"`
	Count int    `json:"count"`
}

type AuditLog struct {
	ID      int    `json:"id"`
	Time    string `json:"time"`
	User    string `json:"user"`
	Action  string `json:"action"`
	Target  string `json:"target"`
}

type SecurityStats struct {
	AttackBlocked24h int           `json:"attackBlocked24h"`
	Score            int           `json:"score"`
	ActiveWafRules   int           `json:"activeWafRules"`
	FirewallRules    int           `json:"firewallRules"`
	BlockedIps       int           `json:"blockedIps"`
	ScannedImages    int           `json:"scannedImages"`
	AttackTrend      []AttackTrend `json:"attackTrend"`
}

type FirewallConfig struct {
	Enabled bool     `json:"enabled"`
	Ports   []string `json:"ports"`
}

var startTime = time.Now()

func getSystemStats() SystemStats {
	c, _ := cpu.Percent(time.Second, false)
	vm, _ := mem.VirtualMemory()
	d, _ := disk.Usage("/")

	cpuVal := 0.0
	if len(c) > 0 {
		cpuVal = c[0]
	}

	uptime := time.Since(startTime)
	uptimeStr := fmt.Sprintf("%dd %dh %dm", int(uptime.Hours())/24, int(uptime.Hours())%24, int(uptime.Minutes())%60)

	return SystemStats{
		CPU:    cpuVal,
		RAM:    vm.UsedPercent,
		Disk:   d.UsedPercent,
		Uptime: uptimeStr,
	}
}

func main() {
	// Set Gin to release mode if needed
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// API Routes
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
				"status":  "FoxDocker Panel is running",
			})
		})

		api.GET("/system/stats", func(c *gin.Context) {
			c.JSON(http.StatusOK, getSystemStats())
		})

		// Security Endpoints
		api.GET("/security/stats", func(c *gin.Context) {
			trend := []AttackTrend{
				{Hour: "14:00", Count: 45},
				{Hour: "15:00", Count: 68},
				{Hour: "16:00", Count: 124},
				{Hour: "17:00", Count: 89},
				{Hour: "18:00", Count: 210},
				{Hour: "19:00", Count: 156},
			}
			c.JSON(http.StatusOK, SecurityStats{
				AttackBlocked24h: 1243,
				Score:            85,
				ActiveWafRules:   42,
				FirewallRules:    12,
				BlockedIps:       154,
				ScannedImages:    48,
				AttackTrend:      trend,
			})
		})

		api.GET("/security/firewall", func(c *gin.Context) {
			rules := []FirewallRule{
				{ID: 1, IP: "192.168.1.105", Action: "Blocked", Reason: "Fail2Ban: SSH Brute Force", Target: "SSH", Time: "2 mins ago"},
				{ID: 2, IP: "45.12.33.1", Action: "Allowed", Reason: "WAF: Valid Traffic", Target: "1998.best", Time: "15 mins ago"},
				{ID: 3, IP: "103.4.15.22", Action: "Blocked", Reason: "SQL Injection", Target: "weatheraxis.com", Time: "1 hour ago"},
			}
			c.JSON(http.StatusOK, rules)
		})

		api.GET("/security/firewall/config", func(c *gin.Context) {
			c.JSON(http.StatusOK, FirewallConfig{
				Enabled: true,
				Ports:   []string{"80", "443", "22", "8888"},
			})
		})

		api.POST("/security/firewall/toggle", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success", "enabled": true})
		})

		api.GET("/security/audit", func(c *gin.Context) {
			logs := []AuditLog{
				{ID: 1, Time: "2026-02-14 18:30:00", User: "AcmaTvirus", Action: "Update Project", Target: "1998.best"},
				{ID: 2, Time: "2026-02-14 18:45:00", User: "AcmaTvirus", Action: "Stop Container", Target: "db-mysql"},
			}
			c.JSON(http.StatusOK, logs)
		})

		api.GET("/security/logs", func(c *gin.Context) {
			logs := []gin.H{
				{"time": time.Now().Format(time.RFC3339), "level": "CRITICAL", "message": "SQL Injection attempt blocked on 1998.best from 45.155.205.233"},
				{"time": time.Now().Add(-5 * time.Minute).Format(time.RFC3339), "level": "WARNING", "message": "Failed SSH login attempt from 185.224.128.11"},
			}
			c.JSON(http.StatusOK, logs)
		})

		api.POST("/security/scan", func(c *gin.Context) {
			// Simulate scan logic
			time.Sleep(2 * time.Second)
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "Security scan completed. 0 new vulnerabilities found.",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		})
	}

	// Serve Static Web Files (UI)
	dist, err := fs.Sub(foxdocker.WebDist, "web/dist")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/dashboard", http.FS(dist))

	// Redirect root to dashboard
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/dashboard/")
	})

	log.Println("FoxDocker Panel starting on :8888...")
	r.Run(":8888")
}
