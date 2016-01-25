package RS

const (
	RS_failed  = -1 // 操作失败
	RS_success = 1  // 操作成功

	RS_user_exist       = 100 // 账号已存在
	RS_user_inexistence = 101 // 账号不存在
	RS_activate_failed  = 102 // 激活失败
	RS_password_error   = 103 // 密码错误

	RS_query_failed  = 200 // 查询失败
	RS_update_failed = 201 // 更新失败
	RS_create_failed = 202 // 创建失败
	RS_delete_failed = 203 // 删除失败

	RS_user_not_activate = 300 // 用户暂未激活
	RS_user_not_login    = 301 // 用户没有登录

	RS_params_error = 400 // 参数错误
)
