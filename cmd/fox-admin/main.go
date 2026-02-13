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
