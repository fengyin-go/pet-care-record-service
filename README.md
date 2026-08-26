# 宠物管理系统 (Pet Care)

纯 Go 标准库后端项目，零第三方依赖。

## 运行

```bash
cd origin
go run ./cmd/server
```

HTTP API 默认监听 `http://localhost:8080/`

## 项目结构

```
origin/
├── cmd/server/main.go          # 入口
├── internal/
│   ├── app/app.go              # 依赖装配
│   ├── config/config.go        # 配置
│   ├── handler/                # HTTP 处理器（7 实体 + stats + middleware）
│   ├── middleware/middleware.go # 中间件
│   ├── model/                  # 领域模型（7 实体 + errors）
│   ├── report/report.go        # 报表生成
│   ├── notification/notification.go # 通知服务
│   ├── health/health.go        # 健康检查
│   ├── service/                # 业务逻辑（7 实体 + stats）
│   ├── store/                  # 内存存储（7 实体）
│   └── utils/                  # 工具函数
└── pkg/
    ├── converter/converter.go  # 单位转换
    ├── export/export.go        # 数据导出
    ├── httpx/httpx.go          # HTTP 工具
    ├── idgen/idgen.go          # ID 生成
    ├── logger/logger.go        # 日志
    ├── security/security.go    # 安全工具
    ├── sorter/sorter.go        # 排序工具
    └── validator/validator.go  # 验证器
```

## 模块名

`pet-care`

## 技术栈

- Go 1.22
- 标准库：`net/http`、`encoding/json`、`sync`、`sort`、`time`、`crypto/rand` 等
- 零第三方依赖

## API 概览

### 主人 (Owner)
- `POST /api/owners` - 创建主人
- `GET /api/owners` - 列表（分页 + 筛选）
- `GET /api/owners/{id}` - 详情
- `PUT /api/owners/{id}` - 更新
- `DELETE /api/owners/{id}` - 删除

### 品种 (Species)
- `POST /api/species` - 创建品种
- `GET /api/species` - 列表（分页 + 筛选）
- `GET /api/species/{id}` - 详情
- `PUT /api/species/{id}` - 更新
- `DELETE /api/species/{id}` - 删除

### 宠物 (Pet)
- `POST /api/pets` - 创建宠物（外键校验）
- `GET /api/pets` - 列表（分页 + 多条件筛选）
- `GET /api/pets/{id}` - 详情
- `PUT /api/pets/{id}` - 更新
- `DELETE /api/pets/{id}` - 删除
- `POST /api/pets/{id}/transition` - 状态流转（active ↔ archived）
- `DELETE /api/pets/{id}/with-records` - 批量删除宠物及全部记录

### 就诊记录 (MedicalRecord)
- `POST /api/medical-records` - 创建就诊记录
- `GET /api/medical-records` - 列表（分页 + 筛选）
- `GET /api/medical-records/{id}` - 详情
- `PUT /api/medical-records/{id}` - 更新
- `DELETE /api/medical-records/{id}` - 删除

### 疫苗记录 (VaccineRecord)
- `POST /api/vaccine-records` - 创建疫苗记录
- `GET /api/vaccine-records` - 列表（分页 + 筛选，含状态筛选）
- `GET /api/vaccine-records/{id}` - 详情
- `PUT /api/vaccine-records/{id}` - 更新
- `DELETE /api/vaccine-records/{id}` - 删除
- `GET /api/vaccine-records/expiring?days=30` - 即将到期疫苗列表
- `GET /api/vaccine-records/expired` - 已过期疫苗列表

### 体重记录 (WeightRecord)
- `POST /api/weight-records` - 创建体重记录
- `GET /api/weight-records` - 列表（分页 + 筛选）
- `GET /api/weight-records/{id}` - 详情
- `PUT /api/weight-records/{id}` - 更新
- `DELETE /api/weight-records/{id}` - 删除
- `GET /api/pets/{id}/weight-trend` - 体重趋势（按时间排序）

### 喂养记录 (FeedingRecord)
- `POST /api/feeding-records` - 创建喂养记录
- `GET /api/feeding-records` - 列表（分页 + 筛选）
- `GET /api/feeding-records/{id}` - 详情
- `PUT /api/feeding-records/{id}` - 更新
- `DELETE /api/feeding-records/{id}` - 删除

### 统计接口
- `GET /api/stats/global` - 全局统计（owner/pet/各类记录数）
- `GET /api/stats/pets/{id}/records` - 单宠物各类型记录数
- `GET /api/stats/pets-by-species` - 按品种分类宠物数
- `GET /api/pets/{id}/health-profile` - 导出宠物完整健康档案（JSON）

## 安全特性

- Bearer Token 鉴权中间件
- 请求日志中间件
- Recovery（panic 恢复）中间件
- 限流中间件（每 IP 每分钟 60 次）
- CORS 中间件
- 请求 ID 注入

## API 客户端

服务只提供 JSON HTTP API，可由任意客户端调用宠物档案、记录和提醒接口。

## 环境变量

- `PORT` - 服务端口（默认 8080）
- `ADDR` - 服务地址（覆盖 PORT）
- `MAX_PAGE_SIZE` - 最大分页大小（默认 100）
- `API_TOKEN` - API 鉴权令牌（默认 `petcare-demo-token`）
- `LOG_LEVEL` - 日志级别（默认 info）

## 测试

```bash
go test ./...
```

## 编译

```bash
go build ./cmd/server
```
