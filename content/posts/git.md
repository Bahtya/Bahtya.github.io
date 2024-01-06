---
title: "Github git SSH 代理设置"
date: 2024-01-06T17:43:47+08:00
draft: flase
---


为了加速 Git 的访问速度，可以按照下面的方式对 SSH 进行配置来加速。
## 准备
准备http代理服务端
    
    http://127.0.0.1:7890

准备仓库和验证token

```
 git config --global https.https://github.com.proxy http://127.0.0.1:7890
 # 使用github Token： Personal access tokens (classic)
 git remote remove origin
 git remote add origin https://token@github.com/Bahtya/Bahtya.github.io.git
```