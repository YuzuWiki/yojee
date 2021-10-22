package pixiv

// pixiv apis
const (
	domain = "https://www.pixiv.net"

	Following      = "/ajax/user/104409/following?offset=0&limit=24&rest=show&tag"
	Bookmarks      = "ajax/user/478993/illusts/bookmarks?tag&offset=0&limit=48&rest=show"
	BookmarksHide  = "/ajax/user/10124151/following?offset=0&limit=24&rest=hide&tag&lang=zh"
	Recommender    = "/rpc/recommender.php?type=illust&sample_illusts=auto&num_recommendations=100&page=discovery&mode=all" // 推荐作品
	RecommenderR18 = "/rpc/recommender.php?type=illust&sample_illusts=auto&num_recommendations=100&page=discovery&mode=r18"
	RecommendUsers = "/rpc/index.php?mode=get_recommend_users_and_works_by_user_ids&user_ids=59401081%2C2868852%2C8349252%2C200059%2C19920821%2C258913%2C13770035%2C10140347%2C29999491%2C3059749&user_num=30&work_num=5"
	Artworks       = "/artworks/91205364"
	Illust         = "ajax/illust/46260941/pages"
	Profile        = "/ajax/user/478993/profile/all"
)

// fanbox apis
const (
	FanBox                = "https://api.fanbox.cc"
	FanBoxFlowingCreators = "/creator.listFollowing"
)
