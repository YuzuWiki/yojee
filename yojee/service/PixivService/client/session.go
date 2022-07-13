package client

// Session session: unit -> PHPSESSID
var (
	Session = make(map[string]*Client)
)

func NewSession(phpSessid string) *Client {
	_, isOk := Session[phpSessid]
	if !isOk {
		Session[phpSessid] = NewClient(phpSessid)
	}

	return Session[phpSessid]
}
