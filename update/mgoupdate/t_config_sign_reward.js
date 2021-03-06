conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_config_sign_reward').drop();

db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8d5de36aeb81eb71a723"),
    "id" : 1.0,
    "day" : 1.0,
    "type" : 1.0,
    "name" : "50金币",
    "amount" : 50.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8d5de36aeb81eb71a724"),
    "id" : 1.0,
    "day" : 2.0,
    "type" : 1.0,
    "name" : "50金币",
    "amount" : 50.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8d79e36aeb81eb71a725"),
    "id" : 2.0,
    "day" : 5.0,
    "type" : 1.0,
    "name" : "100金币",
    "amount" : 100.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8da0e36aeb81eb71a726"),
    "id" : 3.0,
    "day" : 7.0,
    "type" : 1.0,
    "name" : "200金币",
    "amount" : 200.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8dbbe36aeb81eb71a727"),
    "id" : 4.0,
    "day" : 10.0,
    "type" : 1.0,
    "name" : "500金币",
    "amount" : 500.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8dd4e36aeb81eb71a728"),
    "id" : 5.0,
    "day" : 15.0,
    "type" : 1.0,
    "name" : "1000金币",
    "amount" : 1000.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8eb5e36aeb81eb71a729"),
    "id" : 6.0,
    "day" : 20.0,
    "type" : 1.0,
    "name" : "2000金币",
    "amount" : 2000.0
}
);
db.t_config_sign_reward.insert(
{
    "_id" : ObjectId("587c8f16e36aeb81eb71a72a"),
    "id" : 7.0,
    "day" : 50.0,
    "type" : 1.0,
    "name" : "5000金币",
    "amount" : 5000.0
}
);
