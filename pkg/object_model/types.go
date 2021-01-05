package object_model

import "time"

//LinkManagerEventTypeEnum 链接管理回调枚举类型
//定义新类型
type LinkManagerEventTypeEnum int

const (
	LinkAdded LinkManagerEventTypeEnum = iota
	LinkUpdated
	LinkDeleted
)

const (
	LinkStatusPending = "pending"
	LinkStatusValid   = "valid"
	LinkStatusInvalid = "invalid"
)
//LinkStatus 链接别名
type LinkStatus = string
//Link 链接结构体
type Link struct {
	Url         string
	Title       string
	Description string
	Status      LinkStatus
	Tags        map[string]bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
//GetLinksRequest 获取链接请求结构体
type GetLinksRequest struct {
	UrlRegex         string
	TitleRegex       string
	DescriptionRegex string
	Username         string
	Tag              string
	StartToken       string
}
//GetLinksResult 获取链接响应结构体
type GetLinksResult struct {
	Links         []Link
	NextPageToken string
}
//AddLinkRequest 添加链接请求结构体
type AddLinkRequest struct {
	Url         string
	Title       string
	Description string
	Username    string
	Tags        map[string]bool
}
//UpdateLinkRequest 修改链接请求结构体
type UpdateLinkRequest struct {
	Url         string
	Title       string
	Description string
	Username    string
	AddTags     map[string]bool
	RemoveTags  map[string]bool
}
//User 用户结构体
type User struct {
	Email string
	Name  string
}
//LinkManagerEvent 链接管理回调结构体
type LinkManagerEvent struct {
	EventType LinkManagerEventTypeEnum
	Username  string
	Url       string
	Timestamp time.Time
}
//GetNewsRequest 获取新闻请求结构体
type GetNewsRequest struct {
	Username   string
	StartToken string
}
//GetNewsResult 获取新闻响应结构体
type GetNewsResult struct {
	Events    []*LinkManagerEvent
	NextToken string
}
//CheckLinkRequest 检查链接请求结构体
type CheckLinkRequest struct {
	Username string
	Url      string
}
