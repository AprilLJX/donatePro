## 物资捐赠平台

### 数据库

先新建一个donate_google的数据库，再运行DATABASE里的SQL文件

![image-20200605120517868](README_pic/sql.png)
<br>

### 前端配置

1. 安装node
2. 安装cnpm：npm install -g cnpm --registry=http://registry.npm.taobao.org
3. 安装webpack：npm install webpack -g
4. 安装vue-cli：cnpm install vue-cli -g
5. cd /项目名称（donate-sys-google)
6. 安装element：npm i element-ui -S
7. 安装依赖：npm install
8. 安装地区级联选择插件：npm install element-china-area-data -S
9. 启动项目：npm run dev

### 后端配置

1. 配置golang环境
2. 安装gin框架
3. 启动项目：go run Donate_gin
4. 安装tesseract模块，并下载相应中文语言包chi_sim.traineddata 

### 数据库配置

1. 配置mysql环境
2. 新建donate_google数据库
3. 在donate_google数据库中运行donate_google.sql文件