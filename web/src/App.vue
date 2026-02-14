<script setup lang="ts">
// Copyright by AcmaTvirus
import { ref, onMounted, onUnmounted } from 'vue'
import axios from 'axios'
import appsData from './apps.json' // Import templates
import { 
  LayoutDashboard, 
  Layers, 
  Box, 
  Globe, 
  Settings, 
  Search, 
  Plus, 
  Cpu, 
  MemoryStick as Memory, 
  HardDrive,
  ShoppingBag,
  Database,
  FolderOpen,
  Terminal,
  ScrollText,
  LogOut,
  User,
  ShieldCheck,
  Bell,
  Activity,
  History,
  Timer,
  Network,
  Download
} from 'lucide-vue-next'

const currentTab = ref('dashboard')

interface SidebarItem {
  id: string
  label: string
  icon: any
  count?: number
  badge?: string
  alert?: boolean
}

interface SidebarGroup {
  title: string
  items: SidebarItem[]
}

// Real-time System Stats
const sysStats = ref({
  cpu: 0,
  ram: 0,
  disk: 0,
  uptime: '0d 0h 0m'
})

const fetchStats = async () => {
  try {
    const response = await axios.get('/api/system/stats')
    sysStats.value = response.data
  } catch (error) {
    console.error('Failed to fetch system stats:', error)
  }
}

let statsInterval: any = null

onMounted(() => {
  fetchStats()
  statsInterval = setInterval(fetchStats, 3000)
})

onUnmounted(() => {
  if (statsInterval) clearInterval(statsInterval)
})

// Sidebar menu groups
const sidebarGroups: SidebarGroup[] = [
  {
    title: 'Resources',
    items: [
      { id: 'dashboard', label: 'Dashboard', icon: LayoutDashboard },
      { id: 'projects', label: 'Projects', icon: Layers, count: 8 },
      { id: 'appstore', label: 'App Store', icon: ShoppingBag },
      { id: 'containers', label: 'Containers', icon: Box, count: 24 },
    ]
  },
  {
    title: 'Management',
    items: [
      { id: 'databases', label: 'Databases', icon: Database },
      { id: 'files', label: 'Files', icon: FolderOpen },
      { id: 'domains', label: 'Domains', icon: Globe },
      { id: 'security', label: 'Security', icon: ShieldCheck, alert: true },
      { id: 'backup', label: 'Backups', icon: History },
    ]
  },
  {
    title: 'System',
    items: [
      { id: 'cron', label: 'Cron Jobs', icon: Timer },
      { id: 'network', label: 'Networks', icon: Network },
      { id: 'terminal', label: 'Terminal', icon: Terminal },
      { id: 'logs', label: 'System Logs', icon: ScrollText, badge: '5' },
      { id: 'settings', label: 'Settings', icon: Settings },
    ]
  }
]

const recentProjects = [
  { name: 'mysite.com', status: 'online', type: 'WordPress' },
  { name: 'api-service', status: 'online', type: 'Node.js' },
  { name: 'blog-dev', status: 'offline', type: 'PHP 8.2' },
]

const vulnerabilities = [
  { id: 1, target: 'nginx:latest', severity: 'High', desc: 'CVE-2023-XXXX', status: 'pending' },
  { id: 2, target: 'wordpress:latest', severity: 'Medium', desc: 'Outdated Plugin', status: 'scanned' },
]

// App Store Logic
const installApp = (app: any) => {
  alert(`Đang cài đặt ${app.name}...\n(Tính năng backend đang được phát triển)`)
}
</script>

<template>
  <div class="flex h-screen bg-slate-50 dark:bg-dark-bg text-slate-900 dark:text-slate-100 font-sans">
    <!-- Sidebar -->
    <aside class="w-72 bg-white dark:bg-dark-card border-r border-slate-200 dark:border-dark-border p-6 flex flex-col overflow-y-auto custom-scrollbar">
      <div class="flex items-center justify-between mb-8 px-2">
        <div class="flex items-center space-x-3">
          <div class="w-10 h-10 bg-fox-500 rounded-xl flex items-center justify-center shadow-lg shadow-fox-500/30 transition-transform hover:rotate-12">
            <Box class="text-white w-6 h-6" />
          </div>
          <h1 class="text-2xl font-black tracking-tight">FoxDocker</h1>
        </div>
        <button class="p-1.5 bg-fox-500/10 text-fox-500 rounded-lg hover:bg-fox-500 hover:text-white transition-all shadow-sm group" title="Quick Create">
          <Plus class="w-5 h-5 group-hover:rotate-90 transition-transform" />
        </button>
      </div>

      <!-- Resource Widget (LIVE) -->
      <div class="mb-8 p-4 bg-slate-50 dark:bg-slate-800/40 rounded-2xl border border-slate-100 dark:border-dark-border space-y-4 shadow-sm">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-2">
            <Activity class="w-3 h-3 text-fox-500 animate-pulse" />
            <span class="text-[10px] font-black uppercase tracking-widest text-slate-400">Live Health</span>
          </div>
          <span class="text-[10px] font-bold text-green-500 bg-green-500/10 px-1.5 py-0.5 rounded-md tracking-tighter shrink-0">REALTIME</span>
        </div>
        <div class="space-y-2">
          <div class="flex justify-between text-[10px] font-black">
            <span class="text-slate-400">CPU</span>
            <span class="text-fox-500">{{ sysStats.cpu.toFixed(1) }}%</span>
          </div>
          <div class="h-1 w-full bg-slate-200 dark:bg-slate-700 rounded-full overflow-hidden">
            <div class="h-full bg-fox-500 transition-all duration-1000 shadow-[0_0_8px_rgba(6,182,212,0.5)]" :style="{ width: sysStats.cpu + '%' }"></div>
          </div>
        </div>
        <div class="space-y-2">
          <div class="flex justify-between text-[10px] font-black">
            <span class="text-slate-400">RAM</span>
            <span class="text-purple-500">{{ sysStats.ram.toFixed(1) }}%</span>
          </div>
          <div class="h-1 w-full bg-slate-200 dark:bg-slate-700 rounded-full overflow-hidden">
            <div class="h-full bg-purple-500 transition-all duration-1000 shadow-[0_0_8px_rgba(168,85,247,0.5)]" :style="{ width: sysStats.ram + '%' }"></div>
          </div>
        </div>
      </div>

      <nav class="flex-1 space-y-6">
        <div v-for="group in sidebarGroups" :key="group.title" class="space-y-2">
          <p class="px-4 text-[9px] font-black uppercase tracking-[0.2em] text-slate-400 dark:text-slate-500">{{ group.title }}</p>
          <div class="space-y-0.5">
            <div 
              v-for="item in group.items" 
              :key="item.id"
              @click="currentTab = item.id"
              class="sidebar-item group flex items-center"
              :class="{ 'active': currentTab === item.id }"
            >
              <component :is="item.icon" class="w-4.5 h-4.5 transition-colors shrink-0" :class="currentTab === item.id ? 'text-white' : 'text-slate-400 group-hover:text-fox-500'" />
              <span class="font-bold text-sm ml-3 transition-colors">{{ item.label }}</span>
              
              <!-- Counts & Badges -->
              <span v-if="item.count" class="ml-auto text-[10px] font-black bg-slate-100 dark:bg-slate-800 text-slate-500 px-1.5 py-0.5 rounded-md group-hover:bg-fox-500 group-hover:text-white transition-colors">
                {{ item.count }}
              </span>
              <span v-if="item.badge" class="ml-auto text-[10px] font-black bg-red-500 text-white px-1.5 py-0.5 rounded-full shadow-lg shadow-red-500/20">
                {{ item.badge }}
              </span>
              <div v-if="item.id === 'security'" class="ml-auto w-2 h-2 rounded-full bg-red-500 animate-ping shadow-[0_0_8px_rgba(239,68,68,0.8)]"></div>
            </div>
          </div>
        </div>
      </nav>

      <!-- User Profile -->
      <div class="mt-8 p-4 bg-slate-100 dark:bg-slate-800/50 rounded-2xl border border-transparent dark:border-dark-border transition-all hover:border-fox-500/30 group">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <div class="w-9 h-9 rounded-xl bg-gradient-to-br from-fox-400 to-fox-600 flex items-center justify-center text-white text-sm font-bold shadow-md shadow-fox-500/20 group-hover:scale-110 transition-transform">AV</div>
            <div>
              <p class="text-sm font-bold">AcmaTvirus</p>
              <p class="text-[10px] text-slate-500 leading-none font-bold uppercase tracking-tighter">Super Admin</p>
            </div>
          </div>
          <div class="flex items-center space-x-1 shrink-0">
            <button class="p-1.5 text-slate-400 hover:bg-fox-500 hover:text-white rounded-lg transition-all" title="Settings">
              <User class="w-4 h-4" />
            </button>
            <button class="p-1.5 text-slate-400 hover:bg-red-500 hover:text-white rounded-lg transition-all" title="Logout">
              <LogOut class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 flex flex-col overflow-hidden">
      <!-- Top Bar -->
      <header class="h-20 bg-white/50 dark:bg-dark-card/50 backdrop-blur-md border-b border-slate-200 dark:border-dark-border px-8 flex items-center justify-between z-10">
        <div class="relative w-96">
          <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input 
            type="text" 
            placeholder="Search resources..." 
            class="w-full bg-slate-100 dark:bg-slate-800 border-none rounded-xl py-2.5 pl-11 pr-4 focus:ring-2 focus:ring-fox-500 outline-none transition-all text-sm font-medium"
          >
        </div>

        <div class="flex items-center space-x-4">
          <button class="relative p-2.5 text-slate-400 hover:bg-slate-200 dark:hover:bg-slate-800 rounded-xl transition-all">
            <Bell class="w-5 h-5" />
            <span class="absolute top-2 right-2 w-2 h-2 bg-red-500 rounded-full border-2 border-white dark:border-dark-card animate-pulse"></span>
          </button>
          <button class="button-primary flex items-center space-x-2 group">
            <Plus class="w-4 h-4 group-hover:rotate-90 transition-transform" />
            <span class="text-sm">Quick Deploy</span>
          </button>
        </div>
      </header>

      <!-- Views -->
      <div class="flex-1 overflow-y-auto p-8 custom-scrollbar">
        <!-- Dashboard View -->
        <div v-if="currentTab === 'dashboard'" class="max-w-6xl mx-auto space-y-8 animate-in fade-in duration-500">
          <div class="flex items-center justify-between">
            <div>
              <h2 class="text-3xl font-black tracking-tight flex items-center space-x-3">
                <span>Infrastructure</span>
                <span class="text-xs bg-fox-500/10 text-fox-500 px-2 py-1 rounded-md font-black tracking-widest uppercase">STABLE</span>
              </h2>
              <p class="text-slate-500 mt-1 font-medium italic">Everything is running smoothly.</p>
            </div>
            <div class="text-right hidden sm:block font-mono">
              <p class="text-[10px] text-slate-400 font-bold uppercase tracking-widest mb-1">Server Metadata</p>
              <p class="text-xs text-slate-500">IP: 160.191.214.103</p>
              <p class="text-fox-500 text-xs font-bold mt-0.5 uppercase tracking-tighter">Uptime: {{ sysStats.uptime }}</p>
            </div>
          </div>

          <!-- Stats Grid (REAL LIVE) -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="glass-card p-6 flex items-center space-x-4 hover:border-fox-500/30 transition-all group cursor-default shadow-sm text-blue-500">
              <div class="p-4 rounded-2xl bg-slate-50 dark:bg-slate-800/80 shadow-inner group-hover:scale-110 transition-transform">
                <Cpu class="w-8 h-8" />
              </div>
              <div>
                <p class="text-[10px] text-slate-500 font-black uppercase tracking-wider">CPU Usage</p>
                <div class="flex items-baseline space-x-2">
                  <p class="text-2xl font-black text-slate-900 dark:text-white">{{ sysStats.cpu.toFixed(1) }}%</p>
                  <div class="flex items-center space-x-1">
                    <div class="w-1.5 h-1.5 rounded-full" :class="sysStats.cpu < 80 ? 'bg-green-500' : 'bg-red-500'"></div>
                    <span class="text-[10px] font-black tracking-tighter uppercase" :class="sysStats.cpu < 80 ? 'text-green-500' : 'text-red-500'">{{ sysStats.cpu < 80 ? 'Healthy' : 'High Load' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <div class="glass-card p-6 flex items-center space-x-4 hover:border-purple-500/30 transition-all group cursor-default shadow-sm text-purple-500">
              <div class="p-4 rounded-2xl bg-slate-50 dark:bg-slate-800/80 shadow-inner group-hover:scale-110 transition-transform">
                <Memory class="w-8 h-8" />
              </div>
              <div>
                <p class="text-[10px] text-slate-500 font-black uppercase tracking-wider">RAM Usage</p>
                <div class="flex items-baseline space-x-2">
                  <p class="text-2xl font-black text-slate-900 dark:text-white">{{ sysStats.ram.toFixed(1) }}%</p>
                  <div class="flex items-center space-x-1">
                    <div class="w-1.5 h-1.5 rounded-full" :class="sysStats.ram < 90 ? 'bg-green-500' : 'bg-red-500'"></div>
                    <span class="text-[10px] font-black tracking-tighter uppercase" :class="sysStats.ram < 90 ? 'text-green-500' : 'text-red-500'">{{ sysStats.ram < 90 ? 'Optimal' : 'Exhausted' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <div class="glass-card p-6 flex items-center space-x-4 hover:border-green-500/30 transition-all group cursor-default shadow-sm text-green-500">
              <div class="p-4 rounded-2xl bg-slate-50 dark:bg-slate-800/80 shadow-inner group-hover:scale-110 transition-transform">
                <HardDrive class="w-8 h-8" />
              </div>
              <div>
                <p class="text-[10px] text-slate-500 font-black uppercase tracking-wider">Disk Usage</p>
                <div class="flex items-baseline space-x-2">
                  <p class="text-2xl font-black text-slate-900 dark:text-white">{{ sysStats.disk.toFixed(1) }}%</p>
                  <div class="flex items-center space-x-1">
                    <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
                    <span class="text-[10px] text-green-500 font-black tracking-tighter uppercase">Sufficient</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Quick Actions -->
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <button class="glass-card p-6 flex flex-col items-center justify-center space-y-3 hover:bg-fox-500/5 transition-colors border-dashed border-2 group shadow-sm">
              <Plus class="w-8 h-8 text-fox-500 group-hover:rotate-90 transition-transform" />
              <span class="font-black text-xs uppercase tracking-widest text-slate-600 dark:text-slate-400">Add Site</span>
            </button>
            <button class="glass-card p-6 flex flex-col items-center justify-center space-y-3 hover:bg-fox-500/5 transition-colors group shadow-sm">
              <History class="w-8 h-8 text-blue-500 group-hover:rotate-[-45deg] transition-transform" />
              <span class="font-black text-xs uppercase tracking-widest text-slate-600 dark:text-slate-400">Backup</span>
            </button>
            <button class="glass-card p-6 flex flex-col items-center justify-center space-y-3 hover:bg-fox-500/5 transition-colors group shadow-sm">
              <Timer class="w-8 h-8 text-purple-500 group-hover:scale-110 transition-transform" />
              <span class="font-black text-xs uppercase tracking-widest text-slate-600 dark:text-slate-400">Cron Jobs</span>
            </button>
            <button class="glass-card p-6 flex flex-col items-center justify-center space-y-3 hover:bg-fox-500/5 transition-colors group shadow-sm">
              <ShieldCheck class="w-8 h-8 text-red-500 group-hover:animate-pulse" />
              <span class="font-black text-xs uppercase tracking-widest text-slate-600 dark:text-slate-400" @click="currentTab = 'security'">Security</span>
            </button>
          </div>

          <!-- Recent Stacks -->
          <div class="glass-card overflow-hidden shadow-md">
            <div class="p-6 border-b border-slate-200 dark:border-dark-border flex items-center justify-between bg-slate-50/50 dark:bg-slate-800/30">
              <h3 class="text-xl font-black tracking-tight">Active Stacks</h3>
              <button class="text-fox-500 text-[10px] font-black uppercase tracking-widest hover:underline">View Infrastructure</button>
            </div>
            <div class="divide-y divide-slate-200 dark:divide-dark-border">
              <div v-for="project in recentProjects" :key="project.name" class="p-6 flex items-center justify-between hover:bg-slate-50 dark:hover:bg-slate-800/20 transition-colors cursor-pointer group">
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 bg-slate-100 dark:bg-slate-800 rounded-xl flex items-center justify-center transition-colors group-hover:bg-fox-500/10">
                    <Globe class="w-6 h-6 text-slate-400 group-hover:text-fox-500 transition-colors" />
                  </div>
                  <div>
                    <p class="font-black group-hover:text-fox-500 transition-colors font-mono tracking-tight">{{ project.name }}</p>
                    <div class="flex items-center space-x-2 mt-0.5">
                      <span class="text-[9px] bg-slate-100 dark:bg-slate-800 text-slate-500 px-1.5 py-0.5 rounded font-black tracking-widest uppercase">{{ project.type }}</span>
                    </div>
                  </div>
                </div>
                <div class="flex items-center space-x-6">
                  <div class="flex items-center space-x-2 px-3 py-1 bg-slate-50 dark:bg-slate-800/50 rounded-full border border-slate-100 dark:border-dark-border">
                    <div :class="`w-2 h-2 rounded-full ${project.status === 'online' ? 'bg-green-500 animate-pulse shadow-[0_0_8px_rgba(34,197,94,0.6)]' : 'bg-red-500'}`"></div>
                    <span class="text-[10px] font-black uppercase tracking-widest text-slate-500">{{ project.status }}</span>
                  </div>
                  <button class="p-2 hover:bg-fox-500/10 hover:text-fox-500 rounded-lg transition-all" title="Detailed Control">
                    <Layers class="w-5 h-5" />
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Security View -->
        <div v-else-if="currentTab === 'security'" class="max-w-6xl mx-auto space-y-8 animate-in slide-in-from-right-10 duration-500">
          <div class="flex items-center justify-between">
            <div>
              <h2 class="text-3xl font-black tracking-tight flex items-center space-x-3">
                <ShieldCheck class="w-10 h-10 text-red-500" />
                <span>Security Audit</span>
              </h2>
              <p class="text-slate-500 mt-1 font-medium">Quét và bảo vệ hạ tầng Docker bằng FoxAudit Engine.</p>
            </div>
            <button class="button-primary bg-red-500 hover:bg-red-600 shadow-red-500/30 group">
               <Activity class="w-4 h-4 group-hover:animate-spin" />
               <span>Run Security Scan</span>
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="glass-card p-6 border-l-4 border-red-500 shadow-md">
              <h3 class="font-black text-xs uppercase tracking-widest mb-4 flex items-center space-x-2 text-red-500">
                <Bell class="w-4 h-4" />
                <span>Critical Alerts</span>
              </h3>
              <div class="space-y-4">
                <div v-for="v in vulnerabilities" :key="v.id" class="p-4 bg-red-500/5 rounded-xl border border-red-500/10 flex justify-between items-center group cursor-pointer hover:bg-red-500/10 transition-colors">
                  <div>
                    <p class="font-black text-sm">{{ v.target }}</p>
                    <p class="text-xs text-red-400 font-medium">{{ v.desc }}</p>
                  </div>
                  <span class="px-2 py-1 bg-red-100 dark:bg-red-500/20 text-red-600 text-[9px] font-black rounded uppercase tracking-widest">{{ v.severity }}</span>
                </div>
              </div>
            </div>
            <div class="glass-card p-6 shadow-md border border-fox-500/20">
              <h3 class="font-black text-xs uppercase tracking-widest mb-4 flex items-center space-x-2 text-fox-500">
                <Activity class="w-4 h-4" />
                <span>Integration Webhooks</span>
              </h3>
              <div class="space-y-4">
                <div class="flex justify-between items-center p-3.5 bg-slate-50 dark:bg-slate-800/40 rounded-xl border border-slate-100 dark:border-dark-border">
                  <div class="flex items-center space-x-3">
                    <div class="p-2 bg-blue-500/10 rounded-lg text-blue-500">
                       <Bell class="w-4 h-4" />
                    </div>
                    <span class="text-sm font-bold">Telegram Bot</span>
                  </div>
                  <div class="w-10 h-5 bg-fox-500 rounded-full flex items-center px-1 shadow-inner cursor-pointer">
                    <div class="w-3.5 h-3.5 bg-white rounded-full ml-auto shadow-sm"></div>
                  </div>
                </div>
                <div class="flex justify-between items-center p-3.5 bg-slate-50 dark:bg-slate-800/40 rounded-xl border border-slate-100 dark:border-dark-border">
                  <div class="flex items-center space-x-3">
                    <div class="p-2 bg-purple-500/10 rounded-lg text-purple-500">
                       <Activity class="w-4 h-4" />
                    </div>
                    <span class="text-sm font-bold">Discord Webhook</span>
                  </div>
                  <div class="w-10 h-5 bg-slate-200 dark:bg-slate-700 rounded-full flex items-center px-1 shadow-inner cursor-pointer">
                    <div class="w-3.5 h-3.5 bg-white rounded-full shadow-sm"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- App Store View -->
        <div v-else-if="currentTab === 'appstore'" class="max-w-6xl mx-auto space-y-8 animate-in slide-in-from-right-10 duration-500">
          <div class="flex items-center justify-between">
            <div>
              <h2 class="text-3xl font-black tracking-tight flex items-center space-x-3">
                <ShoppingBag class="w-10 h-10 text-fox-500" />
                <span>App Store</span>
              </h2>
              <p class="text-slate-500 mt-1 font-medium">Kho ứng dụng Docker 1-click cài đặt.</p>
            </div>
            <div class="flex space-x-2">
              <button class="px-4 py-2 bg-slate-100 dark:bg-slate-800 rounded-xl text-xs font-bold uppercase tracking-widest text-slate-500 hover:text-fox-500">CMS</button>
              <button class="px-4 py-2 bg-slate-100 dark:bg-slate-800 rounded-xl text-xs font-bold uppercase tracking-widest text-slate-500 hover:text-fox-500">Database</button>
              <button class="px-4 py-2 bg-slate-100 dark:bg-slate-800 rounded-xl text-xs font-bold uppercase tracking-widest text-slate-500 hover:text-fox-500">Tools</button>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
            <div v-for="app in appsData" :key="app.id" class="glass-card p-6 flex flex-col justify-between hover:border-fox-500/50 transition-all group">
              <div class="flex items-start justify-between mb-4">
                <img :src="app.icon" class="w-16 h-16 object-contain group-hover:scale-110 transition-transform drop-shadow-md" alt="icon">
                <span class="px-2 py-1 bg-slate-100 dark:bg-slate-800 text-[9px] font-black uppercase tracking-widest text-slate-500 rounded">{{ app.category }}</span>
              </div>
              <div class="mb-6">
                <h3 class="text-lg font-black mb-1">{{ app.name }}</h3>
                <p class="text-xs text-slate-500 line-clamp-2 h-8">{{ app.description }}</p>
              </div>
              <button @click="installApp(app)" class="w-full py-2.5 bg-slate-900 dark:bg-white text-white dark:text-slate-900 rounded-xl font-bold text-xs uppercase tracking-widest hover:bg-fox-500 hover:text-white dark:hover:bg-fox-500 dark:hover:text-white transition-all shadow-md flex items-center justify-center space-x-2">
                <Download class="w-4 h-4" />
                <span>Install</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Placeholder for other tabs -->
        <div v-else class="flex flex-col items-center justify-center h-full space-y-6 animate-in zoom-in duration-300">
          <div class="w-24 h-24 bg-slate-100 dark:bg-slate-800 rounded-full flex items-center justify-center animate-pulse border border-slate-200 dark:border-dark-border shadow-inner">
            <component :is="sidebarGroups.flatMap(g => g.items).find(i => i.id === currentTab)?.icon || LayoutDashboard" class="w-12 h-12 text-slate-300" />
          </div>
          <div class="text-center space-y-2">
            <h2 class="text-2xl font-black text-slate-400 capitalize tracking-tight">{{ currentTab }}</h2>
            <p class="text-slate-500 text-sm font-medium">Phòng Lab đang hoàn thiện tính năng này.</p>
          </div>
          <button @click="currentTab = 'dashboard'" class="px-6 py-2 bg-slate-200 dark:bg-slate-800 text-slate-500 font-black rounded-xl text-xs uppercase tracking-[0.2em] hover:bg-fox-500 hover:text-white transition-all shadow-md">Quay lại Dashboard</button>
        </div>
      </div>
    </main>
  </div>
</template>

<style>
/* Custom animations */
.animate-in {
  animation: animate-in-kf 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}
@keyframes animate-in-kf {
  from { opacity: 0; transform: translateY(10px) scale(0.95); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}

.slide-in-from-right-10 {
  animation: slide-in-kf 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}
@keyframes slide-in-kf {
  from { opacity: 0; transform: translateX(20px); }
  to { opacity: 1; transform: translateX(0); }
}

/* Custom scrollbar */
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.05);
}

.sidebar-item {
  @apply flex items-center space-x-3 px-4 py-2.5 rounded-xl transition-all duration-200 cursor-pointer hover:bg-fox-500/10 hover:text-fox-500 mx-1;
}

.sidebar-item.active {
  @apply bg-fox-500 text-white shadow-lg shadow-fox-500/40 mx-1;
}
</style>
