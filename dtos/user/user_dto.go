package user

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type FindByEmailDTO struct {
	// 这个错误是因为验证规则要求 Email 字段必填，但是在查询参数中没有找到。需要修改 DTO 的标签，将 json 改为 form，因为是从 URL 查询参数中获取数据：
	Email string `form:"email" binding:"required,email"`
}
