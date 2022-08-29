package requests

type IContext interface {
	GetPid(phpSessID string) (bool, int64)
}
