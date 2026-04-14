.PHONY: dev build build-fe test clean run

# 开发模式：启动后端 + 前端热重载
dev:
	@echo "Starting backend..."
	cd server && go run ./cmd/app/ & \
	echo "Starting frontend dev server..." && \
	cd web && npm run dev

# 构建前端并复制到 embed 目录
build-fe:
	cd web && npm run build
	rm -rf server/pkg/embedfs/dist/assets
	cp -r web/dist/* server/pkg/embedfs/dist/

# 构建单二进制
build: build-fe
	cd server && go build -o ../pawprints ./cmd/app/

# 运行测试
test:
	cd server && go test ./tests/ -v

# 运行构建好的二进制
run: build
	./pawprints

# 清理构建产物
clean:
	rm -f pawprints
	rm -rf server/pkg/embedfs/dist/assets
	rm -rf web/dist
	rm -rf data/pawprints.db

# Docker 构建
docker:
	docker build -t pawprints .

# Docker 运行
docker-run:
	docker run -p 8080:8080 -v pawprints-data:/app/data -v pawprints-uploads:/app/uploads pawprints
