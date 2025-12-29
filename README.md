<div align="center">
  <img src="assets/icon.svg" alt="webdav-115drive" width="200" style="display:block;margin:0">
  <h1 style="margin:0;">webdav-115drive</h1>

**一个 115 网盘 WebDAV 只读服务**

![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)
[![Docker Pulls](https://img.shields.io/docker/pulls/heartleo/webdav-115drive.svg)](https://hub.docker.com/r/heartleo/webdav-115drive)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

</div>

## 🐳 Docker 运行

```bash
docker run --rm -d \
  --name webdav-115drive \
  -p 8090:8090 \
  -e DRIVE_UID=115-uid \
  -e DRIVE_CID=115-cid \
  -e DRIVE_SEID=115-seid \
  -e DRIVE_KID=115-kid \
  heartleo/webdav-115drive
```

## 🐋 Docker Compose 运行

```bash
cat > docker-compose.yml <<EOF
services:
  webdav:
    container_name: webdav-115drive
    image: "heartleo/webdav-115drive:latest"
    ports:
      - "8090:8090"
    env_file:
      - .env
    restart: unless-stopped
EOF
```

```bash
cat > .env <<EOF
DRIVE_UID=115-uid
DRIVE_CID=115-cid
DRIVE_SEID=115-seid
DRIVE_KID=115-kid
EOF
```

```bash
docker-compose up -d
```

## 🚀 编译运行

### 1. ⚒️ 编译

```bash
git clone https://github.com/heartleo/webdav-115drive.git
cd webdav-115drive
go build -o webdav-115drive .
```

### 2. ⚙️ 配置

**使用 `.env` 文件**

```bash
cp .env.example .env
```

**使用 `config.yaml`**

```bash
cp config.yaml.example config.yaml
```

### 3. ✈️ 运行

```bash
./webdav-115drive
```

### 🔑 环境变量

| 变量名                  | 说明              | 默认值       | 必填 |
|----------------------|-----------------|-----------|----|
| `SERVER_HOST`        | 监听主机            | `0.0.0.0` | ❌  |
| `SERVER_PORT`        | 监听端口            | `8090`    | ❌  |
| `SERVER_PATH`        | WebDAV 路径       | `/dav`    | ❌  |
| `SERVER_USER`        | 用户名             | user      | ❌  |
| `SERVER_PWD`         | 密码              | password  | ❌  |
| `DRIVE_UID`          | 115 Cookie UID  | -         | ✅  |
| `DRIVE_CID`          | 115 Cookie CID  | -         | ✅  |
| `DRIVE_SEID`         | 115 Cookie SEID | -         | ✅  |
| `DRIVE_KID`          | 115 Cookie KID  | -         | ✅  |
| `DRIVE_RATE`         | API 请求速率（次/秒）   | `3`       | ❌  |
| `DRIVE_CACHE_EXPIRE` | 缓存过期时间（分钟）      | `1`       | ❌  |

## 🍪 获取 115 Cookies

1. 登录 [115.com](https://115.com/)
2. 打开浏览器开发者工具（F12）
3. 切换到 `Application`
4. 找到 `Cookies` → `https://115.com`
5. 复制以下字段的值：
    - `UID` → `DRIVE_UID`
    - `CID` → `DRIVE_CID`
    - `SEID` → `DRIVE_SEID`
    - `KID` → `DRIVE_KID`

---

<div align="center">

Made with ❤️ by [heartleo](https://github.com/heartleo)

</div>