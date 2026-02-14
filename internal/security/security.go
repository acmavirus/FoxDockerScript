// Copyright by AcmaTvirus
package security

import (
	"encoding/json"
	"os"
	"os/exec"
	"sync"
	"time"
)

type AuditLog struct {
	ID      int    `json:"id"`
	Time    string `json:"time"`
	User    string `json:"user"`
	Action  string `json:"action"`
	Target  string `json:"target"`
}

type FirewallRule struct {
	ID      int    `json:"id"`
	IP      string `json:"ip"`
	Action  string `json:"action"`
	Reason  string `json:"reason"`
	Target  string `json:"target"`
	Time    string `json:"time"`
}

type FirewallConfig struct {
	Enabled bool     `json:"enabled"`
	Ports   []string `json:"ports"`
}

type SecurityData struct {
	AuditLogs   []AuditLog     `json:"audit_logs"`
	Firewall    []FirewallRule `json:"firewall"`
	Config      FirewallConfig `json:"config"`
	BlockedIps  []string       `json:"blocked_ips"`
}

var (
	data     SecurityData
	mu       sync.RWMutex
	dataPath = "data/security.json"
)

func Init() error {
	mu.Lock()
	defer mu.Unlock()

	// Ensure data directory exists
	os.MkdirAll("data", 0755)

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		data = SecurityData{
			AuditLogs: []AuditLog{},
			Firewall:  []FirewallRule{},
			Config: FirewallConfig{
				Enabled: true,
				Ports:   []string{"80", "443", "22", "8888"},
			},
			BlockedIps: []string{},
		}
		return Save()
	}

	file, err := os.ReadFile(dataPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func Save() error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataPath, bytes, 0644)
}

func LogAction(user, action, target string) {
	mu.Lock()
	defer mu.Unlock()

	log := AuditLog{
		ID:     int(time.Now().UnixNano()),
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		User:   user,
		Action: action,
		Target: target,
	}
	data.AuditLogs = append([]AuditLog{log}, data.AuditLogs...)
	if len(data.AuditLogs) > 1000 {
		data.AuditLogs = data.AuditLogs[:1000]
	}
	Save()
}

func AddFirewallRule(rule FirewallRule) {
	mu.Lock()
	defer mu.Unlock()

	rule.ID = int(time.Now().UnixNano())
	rule.Time = "Just now"
	data.Firewall = append([]FirewallRule{rule}, data.Firewall...)
	if rule.Action == "Blocked" {
		data.BlockedIps = append(data.BlockedIps, rule.IP)
	}
	Save()
}

func GetData() SecurityData {
	mu.RLock()
	defer mu.RUnlock()
	return data
}

func ToggleFirewall() bool {
	mu.Lock()
	defer mu.Unlock()
	data.Config.Enabled = !data.Config.Enabled
	Save()
	
	// Real system call
	go func() {
		if data.Config.Enabled {
			exec.Command("ufw", "enable").Run()
		} else {
			exec.Command("ufw", "disable").Run()
		}
	}()
	
	return data.Config.Enabled
}

func GetFail2BanStatus() string {
	// Real logic: check if fail2ban socket exists or run command
	_, err := os.Stat("/var/run/fail2ban/fail2ban.sock")
	if err == nil {
		return "Running"
	}
	return "Stopped"
}
func GetSecurityScore() int {
	mu.RLock()
	defer mu.RUnlock()

	score := 100
	if !data.Config.Enabled {
		score -= 40
	}
	if len(data.Config.Ports) > 10 {
		score -= 10
	}
	// Check if common ports are open
	for _, p := range data.Config.Ports {
		if p == "22" {
			score -= 5 // SSH on default port is a minor risk
		}
	}
	if score < 0 { score = 0 }
	return score
}
