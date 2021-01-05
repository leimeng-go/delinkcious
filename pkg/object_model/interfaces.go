package object_model

//LinkManager 链接管理接口
type LinkManager interface {
	GetLinks(request GetLinksRequest) (GetLinksResult, error)
	AddLink(request AddLinkRequest) error
	UpdateLink(request UpdateLinkRequest) error
	DeleteLink(username string, url string) error
}
//UserManager 用户管理接口
type UserManager interface {
	Register(user User) error
	Login(username string, authToken string) (session string, err error)
	Logout(username string, session string) error
}
//SocialGraphManager 社交图谱管理接口
type SocialGraphManager interface {
	Follow(followed string, follower string) error
	Unfollow(followed string, follower string) error

	GetFollowing(username string) (map[string]bool, error)
	GetFollowers(username string) (map[string]bool, error)

	//AcceptFollowRequest(followed string, follower string) error
	//RejectFollowRequest(followed string, follower string) error
	//KickFollower(followed string, follower string) error
}
//NewsManager 新闻管理接口
type NewsManager interface {
	GetNews(request GetNewsRequest) (GetNewsResult, error)
}
//LinkManagerEvents 链接管理回调
type LinkManagerEvents interface {
	OnLinkAdded(username string, link *Link)
	OnLinkUpdated(username string, link *Link)
	OnLinkDeleted(username string, url string)
}
//LinkCheckerEvents 链接检查回调
type LinkCheckerEvents interface {
	OnLinkChecked(username string, url string, status LinkStatus)
}
