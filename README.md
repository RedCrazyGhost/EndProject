# 此项目为毕设项目
## 介绍
>此系统由现实需求出发，主要针对单一银行账户多笔资金流向链条的数据处理展示提供参考，降低人工筛查
成本。系统使用前后端分离的架构，降低系统耦合度并且提高后续功能和性能的扩展性。

## 后端框架：
- [X] gin：Web框架
- [X] gorm：持久化层
- [X] viper：配置文件
- [X] logrus：日志框架
- [X] Casbin：角色权限控制
- [X] excelize：表格解析



## 技术细节

通过Golang反射创建结构体，完成对用户的数据注入数据库并且做到数据库表级别数据隔离。

通过 SQL语句上下错行比对，检测交易数据存在不一致的情况，从而达到针对异常数据的标识。