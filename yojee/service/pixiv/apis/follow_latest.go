package apis

import (
	"context"
	"fmt"
)

/*
 已关注用户的作品:
		mod all && r18
	漫画: https://www.pixiv.net/ajax/follow_latest/illust?p=1&mode=all&lang=zh
	小说: https://www.pixiv.net/ajax/follow_latest/novel?p=1&mode=all&lang=zh
 追更列表中的作品:
	漫画: https://www.pixiv.net/ajax/watch_list/manga?p=1&lang=zh
	小说: https://www.pixiv.net/ajax/watch_list/novel?p=1&new=1&lang=zh

 好P朋作品:
	漫画: https://www.pixiv.net/ajax/mypixiv_latest/illust?p=1&lang=zh
	小说: https://www.pixiv.net/ajax/mypixiv_latest/novel?p=1&lang=zh

*/

type FollowLatestAPI struct{}

func (api FollowLatestAPI) path(method string) string {
	return fmt.Sprintf("/ajax/follow_latest/%s", method)
}

func (api FollowLatestAPI) get(ctx context.Context, page int, mode string) {
	//	mod: all, r18,
}

func (api FollowLatestAPI) Illust(ctx context.Context, page int, mode string) {

}

func (api FollowLatestAPI) Novel(ctx context.Context, page int, mode string) {

}
