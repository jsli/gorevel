package controllers

const (
	_ = iota
	AdminGroup
	MemberGroup
)

var Permissions = map[string]int{
	// Admin
	"Admin.Index":            AdminGroup,
	"Admin.ListUser":         AdminGroup,
	"Admin.DeleteUser":       AdminGroup,
	"Admin.ListCategory":     AdminGroup,
	"Admin.DeleteCategory":   AdminGroup,
	"Admin.NewCategory":      AdminGroup,
	"Admin.NewCategoryPost":  AdminGroup,
	"Admin.EditCategory":     AdminGroup,
	"Admin.EditCategoryPost": AdminGroup,

	// User
	"User.Edit":     MemberGroup,
	"User.EditPost": MemberGroup,

	// Topic
	"Topic.New":      MemberGroup,
	"Topic.NewPost":  MemberGroup,
	"Topic.Edit":     MemberGroup,
	"Topic.EditPost": MemberGroup,
	"Topic.Reply":    MemberGroup,
}
