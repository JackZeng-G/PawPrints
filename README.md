# PawPrints (萌宠档案馆)

个人私密宠物日记应用，记录宠物从出生到离世的完整生命历程。单用户、无认证、单二进制部署。

## 技术栈

| 层 | 技术 |
|---|------|
| 后端 | Go 1.22+ / Gin / GORM |
| 数据库 | SQLite (modernc.org/sqlite，纯 Go，无 CGO) |
| 前端 | Vue 3 + TypeScript + Vite + Pinia |
| UI | Tailwind CSS v4 |
| 图片 | disintegration/imaging (缩略图) |
| 部署 | go:embed 单二进制 / Docker |

## 功能

- **宠物档案** - 管理多只宠物，支持 7 大分类 29+ 品种
- **日记系统** - 图文日记，关联多只宠物，关键词搜索
- **健康记录** - 疫苗/驱虫/体重/医疗四种类型
- **提醒管理** - 定时检查，完成标记，延后提醒
- **数据管理** - JSON 导入导出，SQLite 备份
- **图片上传** - 自动缩略图生成，支持 JPG/PNG/WebP

## 快速开始

### 直接运行

```bash
# 构建（需 Go 1.22+ 和 Node 18+）
make build

# 启动
./pawprints
```

访问 http://localhost:8080

### Docker

```bash
docker compose up -d
```

### 开发模式

```bash
# 后端
cd server && go run ./cmd/app/

# 前端（另一个终端）
cd web && npm install && npm run dev
```

前端开发服务器 http://localhost:5173，自动代理 `/api` 到后端 8080。

## 项目结构

```
PawPrints/
├── server/
│   ├── cmd/app/main.go           # 入口
│   ├── internal/
│   │   ├── config/               # Viper 配置
│   │   ├── bootstrap/            # 初始化、迁移、种子
│   │   ├── model/                # GORM 数据模型 (8 表)
│   │   ├── repository/           # 数据访问层
│   │   ├── service/              # 业务逻辑层
│   │   ├── handler/              # Gin HTTP handler
│   │   ├── middleware/           # CORS
│   │   ├── router/               # 路由配置
│   │   └── task/                 # 定时任务
│   ├── pkg/embedfs/              # go:embed 前端
│   └── tests/                    # 测试
├── web/
│   ├── src/
│   │   ├── views/                # 页面组件
│   │   ├── stores/               # Pinia stores
│   │   ├── api/                  # Axios API 层
│   │   ├── router/               # Vue Router
│   │   └── types/                # TypeScript 类型
│   └── vite.config.ts
├── config.yaml                   # 默认配置
├── Makefile
├── Dockerfile
└── docker-compose.yml
```

## REST API

| 模块 | 端点 |
|------|------|
| 宠物 | `GET/POST /api/pets`, `GET/PUT/DELETE /api/pets/:id`, `PUT /api/pets/:id/memorial` |
| 日记 | `GET/POST /api/diaries`, `GET/PUT/DELETE /api/diaries/:id` |
| 健康 | `GET/POST /api/health`, `GET/PUT/DELETE /api/health/:id`, `GET /api/health/:id/weight-chart` |
| 提醒 | `GET /api/reminders/active`, `POST/PUT/DELETE /api/reminders/:id`, `PUT complete/snooze` |
| 上传 | `POST /api/upload`, `GET /api/upload/:filename` |
| 统计 | `GET /api/stats/overview`, `GET /api/stats/timeline` |
| 数据 | `GET /api/data/export`, `POST /api/data/import`, `GET /api/data/backup` |
| 分类 | `GET /api/categories`, `GET /api/categories/:id/breeds` |

## 配置

`config.yaml`：

```yaml
server:
  port: 8080
  upload_dir: ./uploads
  max_upload_size: 10MB

database:
  path: ./data/pawprints.db

thumbnail:
  max_width: 400
  max_height: 400
  quality: 80
```

环境变量覆盖：`PAWPRINTS_SERVER_PORT`、`PAWPRINTS_DATABASE_PATH`、`PAWPRINTS_SERVER_UPLOAD_DIR`

## 测试

```bash
cd server && go test ./tests/ -v
```

## 许可证

本项目基于 [GNU General Public License v3.0](LICENSE) 开源。

简而言之：你可以自由使用、修改和分发本软件，但修改后的版本必须同样以 GPL-3.0 发布，并保留原作者的版权声明。详见 [LICENSE](LICENSE) 文件。
