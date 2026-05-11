## 🤖 E2E 测试报告

| 项目 | 内容 |
|------|------|
| Issue | #68 |
| 分支 | `feat/issue-67` |
| 结果 | 🛑 已停止 |
| 时间 | 2026-05-06 15:22:52 |
| 日志 | `e2e-issue-68-20260506-152237.log` |

### 详细日志

[2026-05-06 15:22:37] E2E 任务启动 (由 rework-68)
分支: feat/issue-67
项目: social-world
日志: /Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-issue-68-20260506-152237.log

=== E2E 环境检查 ===
任务类型: full-stack
✅ adb: Android Debug Bridge version 1.0.41
✅ maestro: 1.39.13

=== 检查后端正端 ===
⚠️ 后端未运行，尝试启动 E2E 环境...
sw-mysql
sw-redis
docker compose up --build -d mysql redis
time="2026-05-06T15:22:39+08:00" level=warning msg="/private/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/rework-pr-68-3002107877/docker-compose.yml: the attribute `version` is obsolete, it will be ignored, please remove it to avoid potential confusion"
 Network rework-pr-68-3002107877_sw-network  Creating
 Network rework-pr-68-3002107877_sw-network  Created
 Volume "rework-pr-68-3002107877_mysql_data"  Creating
 Volume "rework-pr-68-3002107877_mysql_data"  Created
 Volume "rework-pr-68-3002107877_redis_data"  Creating
 Volume "rework-pr-68-3002107877_redis_data"  Created
 Container sw-mysql  Creating
 Container sw-redis  Creating
 Container sw-mysql  Created
 Container sw-redis  Created
 Container sw-mysql  Starting
 Container sw-redis  Starting
 Container sw-redis  Started
 Container sw-mysql  Started
⏳ 等待数据库就绪...
   等待中... (1/30)
   等待中... (2/30)
   等待中... (3/30)
✅ 数据库已就绪
🚀 启动 API (pm2)...
[32m[PM2] [39mApplying action restartProcessId on app [sw-api](ids: [ 4 ])
[32m[PM2] [39m[sw-api](4) ✓
┌────┬────────────────────┬─────────────┬─────────┬─────────┬──────────┬────────┬──────┬───────────┬──────────┬──────────┬──────────┬──────────┐
│ id │ name               │ namespace   │ version │ mode    │ pid      │ uptime │ ↺    │ status    │ cpu      │ mem      │ user     │ watching │
├────┼────────────────────┼─────────────┼─────────┼─────────┼──────────┼────────┼──────┼───────────┼──────────┼──────────┼──────────┼──────────┤
│ [1m[36m3[39m[22m  │ dashboard          │ default     │ 1.0.0   │ [7m[1mfork[22m[27m    │ 91375    │ 75m    │ 2    │ [32m[1monline[22m[39m    │ 0%       │ 16.3mb   │ [1minsulate[22m │ [90mdisabled[39m │
│ [1m[36m1[39m[22m  │ immortal-client    │ default     │ 0.1.0   │ [7m[1mfork[22m[27m    │ 18437    │ 21m    │ 109  │ [32m[1monline[22m[39m    │ 0%       │ 54.3mb   │ [1minsulate[22m │ [90mdisabled[39m │
│ [1m[36m0[39m[22m  │ immortal-server    │ default     │ N/A     │ [7m[1mfork[22m[27m    │ 81317    │ 98m    │ 70   │ [32m[1monline[22m[39m    │ 0%       │ 35.0mb   │ [1minsulate[22m │ [90mdisabled[39m │
│ [1m[36m4[39m[22m  │ sw-api             │ default     │ N/A     │ [34m[1mcluster[22m[39m │ 31701    │ 0s     │ 8    │ [32m[1monline[22m[39m    │ 0%       │ 37.2mb   │ [1minsulate[22m │ [90mdisabled[39m │
└────┴────────────────────┴─────────────┴─────────┴─────────┴──────────┴────────┴──────┴───────────┴──────────┴──────────┴──────────┴──────────┘
⏳ 等待 API 就绪...
   等待中... (1/30)
   等待中... (2/30)
   等待中... (3/30)

🛑 任务已停止

⏹️ 任务已手动停止


---
*由 Agent Control Center 自动生成*