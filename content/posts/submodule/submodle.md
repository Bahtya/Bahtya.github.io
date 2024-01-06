---
title: "小试Submodule"
author: "Bahtyar"
date: 2024-01-06T21:25:24+08:00
tags: [git, submodule]
categories: ["技术大杂烩"]
resources:
- name: "featured-image"
  src: "featured-image.webp"
draft: false
---



创建完hugo之后，换了一台电脑后将博客仓库拉下来后发现主题不见了，才反应过来是不是因为当时是子模块方式创建的仓库，默认pull不包括。故又去查了一下用法。

git介绍[7.11 Git 工具 - 子模块](https://git-scm.com/book/zh/v2/Git-%E5%B7%A5%E5%85%B7-%E5%AD%90%E6%A8%A1%E5%9D%97)

```
git submodule init
git pull --recurse-submodules
```