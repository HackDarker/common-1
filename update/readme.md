## 更新及部署说明：
### 2017-03-18:
1. 更新t_config_active_list表
2. 在各个游戏初始化时初始化common/service下的pushService，
例：pushService.PoolInit("192.168.199.155:2801"),建议将HallTcpAddr字段配置到json文件中。
3. 先部署大厅
4. 再部署casino_super、casino_mj、casino_ddz、casino_zjh
### 2017-03-20:
1. 更新t_th_notice表。