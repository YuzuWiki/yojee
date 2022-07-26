package pixivService

const (
	PixivDomain = ".pixiv.net"

	PixivHost = "www.pixiv.net"
	PximgHost = "i.pximg.net"

	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0"
	Phpsessid = "PHPSESSID"
)

const (
	/*
		作品类型
	*/
	Illust = "illust"
	Manga  = "manga"
	Novel  = "novel"

	/*
		作品类别, Category
	*/
	FollowLatest  = "follow_latest"
	WatchList     = "watch_list"
	MypixivLatest = "mypixiv_latest"

	/*
		filter
	*/
	Top = "top" // 主页 选项卡中总览信息
	All = "all" // 漫画 + 插画

	Show = "show"
	Hide = "hide"

	R18 = "r18"
)
