package main

// pixiv API
const (
	FollowingAPI      = "/ajax/user/{id}/following"         // ?offset=0&limit=24&rest=show&tag
	BookmarksAPI      = "/ajax/user/{id}/illusts/bookmarks" // ?tag&offset=0&limit=48&rest=show"
	BookmarksHideAPI  = "/ajax/user/{id}/following"         // ?offset=0&limit=24&rest=hide&tag&lang=zh
	RecommenderAPI    = "/rpc/recommender.php"              // ?type=illust&sample_illusts=auto&num_recommendations=100&page=discovery&mode=all
	RecommendUsersAPI = "/rpc/index.php"                    // mode=get_recommend_users_and_works_by_user_ids&user_ids=[xxx, xxx]&user_num=30&work_num=5
	ArtworksAPI       = "/artworks/{id}"
	IllustAPI         = "ajax/illust/{id}/pages"
	ProfileAPI        = "/ajax/user/{id}/profile/all"
)
