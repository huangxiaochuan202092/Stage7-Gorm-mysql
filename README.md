

# 项目介绍
本项目实现了登录即注册的功能，并在此基础上依次实现了用户功能模块、任务功能模块,博客功能模块,docker,docker-compose部署项目
gorm 快速上手


## 项目部署流程

### 1. 克隆项目
使用 `git clone` 命令克隆项目到本地：
```bash
git clone <项目地址>
```
请将 `<项目地址>` 替换为实际的项目仓库地址。

### 2. 修改配置
需要修改以下配置：
- **MySQL 连接配置**：在项目配置文件中MySQL 的连接信息，如地址、端口、用户名、密码等。

### 3. 初始化项目
在项目根目录下执行以下命令初始化项目依赖：
```bash
go mod tidy
```

### 4. 运行项目
在 `resul` 目录下执行以下命令启动项目：
```bash
go run database.go model.go

```

### 5. 访问数据库
<img width="1160" alt="image" src="https://github.com/user-attachments/assets/0b9aa970-590c-4fa4-978a-3e5716b59c3f" />


## 视频教程
- **哔哩哔哩**：[https://space.bilibili.com/3546867629558058?spm_id_from=333.337.0.0](https://space.bilibili.com/3546867629558058?spm_id_from=333.337.0.0)
- **抖音**：[https://www.douyin.com/user/MS4wLjABAAAAutuiF-v06OCpXGOjaUDTGT6u4WG4kadCuRbZEvLRY1s?from_tab_name=main](https://www.douyin.com/user/MS4wLjABAAAAutuiF-v06OCpXGOjaUDTGT6u4WG4kadCuRbZEvLRY1s?from_tab_name=main)

## 视频内容
- 【第1期】go快速上手web应用项目登陆即注册，Gin Web开发的涡轮增压引擎，技术点 Gin + Gorm + Reids
- 【第2期】go快速上手web应用项目用户管理模块
- 【第3期】go快速上手web应用项目任务管理模块
- 【第4期】go快速上手web应用项目博客管理模块
- 【第5期】go快速上手web应用项目问答管理模块
- 【第6期】go快速上手web应用部署项目
- 【第7期】详细步骤快速上手 Gorm


