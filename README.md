# 魔兽世界绿玩助手
使用Python爬取网易处罚公告内容或处理网易发布的处罚PDF文件（两种方式），将爬取到的结果插入到数据库中。项目采用前后端分离技术，使用Golang+Gin处理数据提供查询服务，Vue3+Vite7构建前端项目。

## 构建前端项目（打包到服务器）

pnpm install

pnpm run build

## 二次开发项目

pnpm install

pnpm run dev

## 构建后端项目(Go部分)

go build main.go