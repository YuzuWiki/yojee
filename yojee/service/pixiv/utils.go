package pixiv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	netUrl "net/url"
	"reflect"
	"strings"

	"github.com/like9th/yojee/yojee/service/pixiv/client"
)

func Path(prefix string, paths ...interface{}) string {
	return path(prefix, paths...)
}

func trimString(s string, sep string) string {
	if len(s) == 0 || len(sep) == 0 {
		return s
	}

	s = strings.TrimPrefix(s, sep)
	s = strings.TrimSuffix(s, sep)
	return s
}

func path(prefix string, paths ...interface{}) string {
	sep := "/"
	prefix = trimString(prefix, sep)

	elems := []string{prefix}
	for i := range paths {
		switch reflect.TypeOf(paths[i]).Kind() {
		case reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int:
			elems = append(elems, fmt.Sprintf("%d", paths[i]))
		case reflect.String:
			elems = append(elems, trimString(paths[i].(string), sep))
		default:
			log.Printf("Warning: Unsupported path = %v", paths[i])
		}

	}
	return strings.Join(elems, sep)
}

func Get(ctx context.Context, url string, query *netUrl.Values) ([]byte, error) {
	c := client.For(ctx)

	url = c.EndpointULR(url, query).String()
	resp, err := c.GetWithContext(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := struct {
		Error bool        `json:"error"`
		Body  interface{} `json:"body"`
	}{}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if data.Error {
		return nil, errors.New(fmt.Sprintf("Response Error: api = %s; body = %s", url, body))
	}

	ret, err := json.Marshal(data.Body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

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
