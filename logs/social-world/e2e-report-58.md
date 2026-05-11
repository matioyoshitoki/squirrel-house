## 🤖 E2E 测试报告

| 项目 | 内容 |
|------|------|
| Issue | #58 |
| 分支 | `feat/issue-5` |
| 结果 | ❌ 失败 |
| 时间 | 2026-04-25 20:34:10 |
| 日志 | `e2e-issue-58-20260425-203258.log` |

### 详细日志

[2026-04-25 20:32:58] E2E 任务启动 (由 rework-58)
分支: feat/issue-5
项目: social-world
日志: /Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-issue-58-20260425-203258.log

=== E2E 环境检查 ===
任务类型: full-stack
✅ adb: Android Debug Bridge version 1.0.41
✅ maestro: 1.39.13

=== 检查后端正端 ===
⚠️ 后端未运行，尝试启动 E2E 环境...
🔌 分配端口: MySQL=55348, Redis=55349
time="2026-04-25T20:32:59+08:00" level=warning msg="/private/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-58-489453622/docker-compose.yml: the attribute `version` is obsolete, it will be ignored, please remove it to avoid potential confusion"
 Network dev-issue-58-489453622_sw-network  Creating
 Network dev-issue-58-489453622_sw-network  Created
 Volume "dev-issue-58-489453622_redis_data"  Creating
 Volume "dev-issue-58-489453622_redis_data"  Created
 Volume "dev-issue-58-489453622_mysql_data"  Creating
 Volume "dev-issue-58-489453622_mysql_data"  Created
 Container dev-issue-58-489453622-mysql-1  Creating
 Container dev-issue-58-489453622-redis-1  Creating
 Container dev-issue-58-489453622-mysql-1  Created
 Container dev-issue-58-489453622-redis-1  Created
 Container dev-issue-58-489453622-redis-1  Starting
 Container dev-issue-58-489453622-mysql-1  Starting
 Container dev-issue-58-489453622-mysql-1  Started
 Container dev-issue-58-489453622-redis-1  Started
⏳ 等待数据库就绪...
   等待中... (1/30)
   等待中... (2/30)
   等待中... (3/30)
✅ 数据库已就绪 (MySQL: localhost:55348, Redis: localhost:55349)
🚀 启动 API (pm2)...
[32m[PM2] [39mApplying action restartProcessId on app [sw-api](ids: [ 11 ])
[32m[PM2] [39m[sw-api](11) ✓
┌────┬────────────────────┬─────────────┬─────────┬─────────┬──────────┬────────┬──────┬───────────┬──────────┬──────────┬──────────┬──────────┐
│ id │ name               │ namespace   │ version │ mode    │ pid      │ uptime │ ↺    │ status    │ cpu      │ mem      │ user     │ watching │
├────┼────────────────────┼─────────────┼─────────┼─────────┼──────────┼────────┼──────┼───────────┼──────────┼──────────┼──────────┼──────────┤
│ [1m[36m5[39m[22m  │ agflux-server      │ default     │ 0.1.0   │ [7m[1mfork[22m[27m    │ 85414    │ 2D     │ 2    │ [32m[1monline[22m[39m    │ 0%       │ 22.5mb   │ [1minsulate[22m │ [90mdisabled[39m │
│ [1m[36m1[39m[22m  │ dashboard          │ default     │ 1.0.0   │ [7m[1mfork[22m[27m    │ 99759    │ 6h     │ 52   │ [32m[1monline[22m[39m    │ 0%       │ 20.3mb   │ [1minsulate[22m │ [90mdisabled[39m │
│ [1m[36m6[39m[22m  │ immortal-server    │ default     │ 0.1.0   │ [7m[1mfork[22m[27m    │ 93847    │ 12h    │ 62   │ [32m[1monline[22m[39m    │ 0%       │ 31.8mb   │ [1minsulate[22m │ [90mdisabled[39m │
│ [1m[36m11[39m[22m │ sw-api             │ default     │ N/A     │ [34m[1mcluster[22m[39m │ 19550    │ 0s     │ 8    │ [32m[1monline[22m[39m    │ 0%       │ 37.1mb   │ [1minsulate[22m │ [90mdisabled[39m │
└────┴────────────────────┴─────────────┴─────────┴─────────┴──────────┴────────┴──────┴───────────┴──────────┴──────────┴──────────┴──────────┘
⏳ 等待 API 就绪...
   等待中... (1/30)
   等待中... (2/30)
   等待中... (3/30)
   等待中... (4/30)
   等待中... (5/30)
   等待中... (6/30)
   等待中... (7/30)
   等待中... (8/30)
   等待中... (9/30)
   等待中... (10/30)
   等待中... (11/30)
   等待中... (12/30)
   等待中... (13/30)
   等待中... (14/30)
   等待中... (15/30)
   等待中... (16/30)
   等待中... (17/30)
   等待中... (18/30)
   等待中... (19/30)
   等待中... (20/30)
   等待中... (21/30)
   等待中... (22/30)
   等待中... (23/30)
   等待中... (24/30)
   等待中... (25/30)
   等待中... (26/30)
   等待中... (27/30)
   等待中... (28/30)
   等待中... (29/30)
   等待中... (30/30)
❌ API 未在预期时间内启动，请检查 pm2 logs
make: *** [e2e-env-up] Error 1
❌ 后端 API 启动失败

❌ e2e失败: 步骤 checkEnv 返回 exit code 1


---
*由 Agent Control Center 自动生成*