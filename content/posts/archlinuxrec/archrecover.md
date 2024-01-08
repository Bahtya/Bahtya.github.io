---
title: "ArchLinux用户绕不开的系统回滚"
date: 2024-01-08T22:47:16+08:00
tags: [btrfs, archlinux]
categories: ["技术大杂烩"]
draft: true
---

原本以为archlinux的回滚是个玩笑，怎么会轻易更新挂掉，但是自动用了holoiso系统，也是恢复过多次文件系统了，很可能是holoiso本身的问题，更新完总是会挂掉。还好是btrfs文件系统，恢复起来很容易，这里使用的是manjaro的CDLive

```
sudo -o subvol=/ /dev/nvmen1p2 /mnt
sudo mv @ @bad
# 找到最后一个可用快照，进行创建快照
sudo btrfs subvolume snap @snapshots/921/snapshot @
sudo btrfs property set -ts @ ro false
sudo btrfs subvolume list
# 记下@的volid
sudo btrfs subvolume set-default 922 /mnt
# 修改fstab的挂载子卷subvol
sudo vi /@/etc/fstab
```