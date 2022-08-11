package pixiv

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YuzuWiki/yojee/global"
	"io/ioutil"
	"log"
	"net/http"
	netUrl "net/url"
	"reflect"
	"strings"

	"github.com/YuzuWiki/yojee/common/requests"
)

// NewQuery return get params
func NewQuery(values map[string]interface{}) (*netUrl.Values, error) {
	query := netUrl.Values{}
	for k := range values {
		var v string

		switch reflect.TypeOf(values[k]).Kind() {
		case reflect.String:
			v = values[k].(string)
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			v = fmt.Sprintf("%d", values[k])
		case reflect.Slice:
			_setV, err := json.Marshal(values[k])
			if err != nil {
				return nil, err
			}
			v = string(_setV)
		default:
			return nil, errors.New(fmt.Sprintf("Query error: unsupported type = %s", reflect.TypeOf(values[k]).String()))
		}

		query.Set(k, v)
	}
	return &query, nil
}

func strip(s string, sep string) string {
	if len(s) == 0 || len(sep) == 0 {
		return s
	}

	s = strings.TrimPrefix(s, sep)
	s = strings.TrimSuffix(s, sep)
	return s
}

// Path returns pixiv api
func Path(paths ...interface{}) string {
	sep := "/"
	elems := []string{"https://" + PixivHost}
	for i := range paths {
		switch reflect.TypeOf(paths[i]).Kind() {
		case reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int:
			elems = append(elems, fmt.Sprintf("%d", paths[i]))
		case reflect.String:
			elems = append(elems, strip(paths[i].(string), sep))
		default:
			log.Printf("Warning: Unsupported path = %v", paths[i])
		}

	}
	return strings.Join(elems, sep)
}

// Request return http.body && error
func Request(ctx Context, method string, u string, query *requests.Query, params *requests.Params) ([]byte, error) {
	method = strings.ToUpper(method)
	var fn func(string, *requests.Query, *requests.Params) (*http.Response, error)

	global.Logger.Info().Msg(fmt.Sprintf("ctx is %+v", ctx))
	client := ctx.Client()
	switch method {
	case http.MethodGet:
		fn = client.Get
	case http.MethodPost:
		fn = client.Post
	case http.MethodPut:
		fn = client.Put
	case http.MethodDelete:
		fn = client.Delete
	default:
		return nil, errors.New("UnSupport method: " + method)
	}

	resp, err := fn(u, query, params)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer resp.Body.Close()

	return ioutil.ReadAll(body)
}
