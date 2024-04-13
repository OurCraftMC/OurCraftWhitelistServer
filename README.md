# OurCraft 白名单申请站

## 功能说明

- 支持中文与英语
- 直接输出whitelist.json文件
- 支持自定义白名单格式

## 如何运行

### 1. 使用Docker

首先Build镜像

```bash
docker build -t ourcraft-whitelist .
```

然后运行容器

```bash
docker run -d --name applyserver -p 8080:8080 -v /path/to/whitelist.json:/whitelist.json ourcraft-whitelist
```

### 2. 不用Docker

使用go进行建构
    
```bash
go build -o ourcraft-whitelist
```

然后运行

```bash
./ourcraft-whitelist /config.yaml
```

## 配置文件

```yaml
whitelist_file_path: "/path/to/whitelist.json" # 白名单文件路径
html_path: "/path/to/html" # 静态文件路径 即库中html文件夹路径
listen_addr: ":8080" # 监听地址 例如"127.0.0.1:8080"
```

## 输出白名单格式
```json
[
  {
    "name" : "player1",
    "uuid" : "uuid1",
    "contact_method" : "qq",
    "contact_id" : "114514"
  },
  {
    "name": "player2",
    "uuid": "uuid2",
    "contact_method": "email",
    "contact_id": "1919810"
  }
]
```