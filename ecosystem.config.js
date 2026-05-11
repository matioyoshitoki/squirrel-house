module.exports = {
  apps: [{
    name: 'dashboard',
    script: './build/dashboard',
    cwd: __dirname,
    instances: 1,
    exec_mode: 'fork',

    // 优雅关闭：给 dashboard 30 秒时间完成 running 任务
    kill_timeout: 30000,
    listen_timeout: 10000,

    // 资源限制
    max_memory_restart: '500M',

    // 重启策略
    min_uptime: '10s',
    max_restarts: 5,
    autorestart: true,

    // 不自动监视文件变化
    watch: false,

    // 日志
    log_file: './logs/pm2/dashboard.log',
    error_file: './logs/pm2/dashboard-error.log',
    out_file: './logs/pm2/dashboard-out.log',
    merge_logs: true,
    time: true,

    // 环境变量
    env: {
      NODE_ENV: 'production',
      DASHBOARD_PORT: '9090',
      PATH: '/Users/insulate/Library/Android/sdk/platform-tools:/Users/insulate/.maestro/maestro/bin:/Users/insulate/bin:/Users/insulate/.local/bin:/Library/Frameworks/Python.framework/Versions/3.11/bin:/Library/Frameworks/Python.framework/Versions/3.14/bin:/opt/homebrew/bin:/opt/homebrew/sbin:/usr/local/bin:/System/Cryptexes/App/usr/bin:/usr/bin:/bin:/usr/sbin:/sbin:/Users/insulate/flutter/bin',
      ANDROID_HOME: process.env.ANDROID_HOME || '/Users/insulate/Library/Android/sdk',
      ANDROID_SDK_ROOT: process.env.ANDROID_SDK_ROOT || '/Users/insulate/Library/Android/sdk'
    },

    // 不强制使用特定用户运行
    // pm2 会以当前用户身份启动进程
  }]
};
