package main

// fallbackTemplate 当无法加载外部模板时使用
const fallbackTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, sans-serif; margin: 0; padding: 2rem; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 2rem; border-radius: 8px; }
        h1 { color: #333; }
        .error { color: #d32f2f; padding: 1rem; background: #ffebee; border-radius: 4px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>{{.Title}}</h1>
        <div class="error">
            ⚠️ 静态文件加载失败，这是备用模板。<br>
            请确保 static/index.html 存在。
        </div>
    </div>
</body>
</html>`
