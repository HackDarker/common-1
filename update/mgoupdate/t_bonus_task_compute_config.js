conn = new Mongo();
db = conn.getDB("test");


db.getCollection('t_bonus_task_compute_config').drop();

//麻将
db.t_bonus_task_compute_config.insert({
    "gameid": 2,
    "coinpercent": 1,    //一万金币与红包的兑率
    "bonuspercent": 0.7,   //比例系数
    "min": 0.1,            //一次领取可得最小红包数
    "max": 5             //一次领取可得最大红包数
});

//斗地主
db.t_bonus_task_compute_config.insert({
    "gameid": 3,
    "coinpercent": 1,
    "bonuspercent": 0.7,
    "min": 0.1,
    "max": 5
});

//炸金花
db.t_bonus_task_compute_config.insert({
    "gameid": 4,
    "coinpercent": 1,
    "bonuspercent": 0.7,
    "min": 0.1,
    "max": 5
});

//计算公式为
// 门票（金币数）x 兑率 x 比例系数，且受最小、最大红包数约束。
