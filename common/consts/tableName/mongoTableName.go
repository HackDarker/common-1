package tableName

const (
	//数据库相关
	DB_ENSURECOUNTER_KEY string = "id" //自增键

	//大厅
	DBT_T_USER                      string = "t_user"                      //用户表
	DBT_T_USER_DIAMOND_DETAILS      string = "t_user_diamond_details"      //钻石交易记录
	DBT_T_TH_NOTICE                 string = "t_th_notice"                 //公告
	DBT_T_RECHARGE_DETAILS          string = "t_recharge_details"          //充值记录
	DBT_T_ACTIVE_LIST               string = "t_config_active_list"        //活动列表
	DBT_T_TASK_INFO                 string = "t_config_task_info"          //公共任务信息表
	DBT_T_USER_TASK                 string = "t_user_task"                 //用户任务状态表
	DBT_T_USER_BAG                  string = "t_user_bag"                  //用户背包道具表
	DBT_T_MAIL_CONTENT              string = "t_mail_content"              //公共邮件表
	DBT_T_USER_MAIL_list            string = "t_user_mail_list"            //用户邮件表
	DBT_T_USER_FRIENDS_LIST         string = "t_user_friends_list"         //用户好友关系表
	DBT_T_PAYBASEDETAILS            string = "t_PayBaseDetails"            //重新订单信息
	DBT_T_USER_STRONGBOX            string = "t_user_strongbox"            //保险箱
	DBT_T_GOODS_INFO                string = "t_goods_info"                //商品信息表
	DBT_T_CONFIG_SYS                string = "t_config_sys"                //系统的配置表
	DBT_T_USER_ATTACH               string = "t_user_attach"               //用户签到/抽奖/补助 领取记录表
	DBT_T_DRAW_LOTTERY              string = "t_config_draw_lottery"       //转盘物品信息表
	DBT_T_SIGN_REWARD               string = "t_config_sign_reward"        //签到奖励表
	DBT_T_GAME_LOG                  string = "t_game_log"                  //游戏记录表
	DBT_T_GAME_COUNT                string = "t_game_count"                //游戏记录表
	DBT_T_GAME_DAY_COUNT            string = "t_game_day_count"            //游戏记录表
	DBT_T_BONUS_TASK_COMPUTE_CONFIG string = "t_bonus_task_compute_config" //红包任务计算配置表
	DBT_T_TRADE_LOG                 string = "t_game_trade_log"            //游戏货币流水记录表
	DBT_T_GRAY_RELEASE_USER         string = "t_gray_release_user"         //灰度发布的数据库表

	DBT_STATISTIC_BTN_CLICK         string = "T_statistic_btn_click"       //统计相关的数据库表名字
	DBT_ADMIN_EXCHANGE_RECORD       string = "t_admin_exchange_record"     //红包、实物兑换记录表
	DBT_IP_ADDRESS                  string = "dbt_ip_address"              //存储Ip对应的地址

	//代理充值系统
	DBT_AGENT_GOODS        = "t_agent_goods"        //商品信息表
	DBT_AGENT_RECHARGE_LOG = "t_agent_recharge_log" //代理商充值记录表
	DBT_AGENT_SALES_LOG    = "t_agent_sales_log"    //代理商销售记录表
	DBT_AGENT_APPLY_LOG    = "t_agent_apply_log"    //代理商申请记录表
	DBT_AGENT_REBATE_LOG   = "t_agent_rebate_log"   //代理商领取返利记录表
	DBT_AGENT_INFO         = "t_agent_info"         //代理关系表

	//数据分析
	ADMIN_USER_ATHOME     = "t_user_athome"
	ADMIN_USER_ONLINEHOUR = "t_data_onlinehour"
	ADMIN_USER_ONLINEDAY  = "t_data_onlineday"

	//接入服务器
	DBT_GAME_CONFIG_LOGIN      = "t_game_config_login"      //登陆服配置	//游戏配置表
	DBT_GAME_CONFIG_LOGIN_LIST = "t_game_config_login_list" //登陆服配置	//登陆服游戏列表

	//麻将
	DBT_MJ_DESK             = "t_mj_desk"           //桌子的信息
	DBT_MJ_DESK_ROUND       = "t_mj_desk_round"     //一把麻将结束
	DBT_T_TH_GAMENUMBER_SEQ = "t_th_gamenumber_seq" //麻将 编号

	//斗地主
	DBT_DDZ_DESK_ROUND = "t_ddz_desk_round" //一把斗地主结束

	//跑得快
	DBT_PDK_DESK_ROUND          = "t_pdk_desk_round"          //一把跑得快结束战绩
	DBT_PDK_DESK_ROUND_PLAYBACK = "t_pdk_desk_round_playback" //一把跑得快回放

	//牛牛
	DBT_NIU_DESK_ROUND_ONE          = "t_niuniu_desk_round_one"       //牛牛1局结束战绩
	DBT_NIU_DESK_ROUND_ALL          = "t_niuniu_desk_round_all"       //牛牛10局结束战绩

	//拼二章
	DBT_PEZ_DESK_ROUND  = "dbt_pez_desk_round"	//
)
