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

func path(prefix string, paths ...interface{}) string {
	sep := "/"
	if !strings.HasPrefix(prefix, sep) {
		prefix = sep + prefix
	}

	elems := []string{prefix}
	for i := range paths {

		switch reflect.TypeOf(paths[i]).Kind() {
		case reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int:
			elems = append(elems, fmt.Sprintf("%d", paths[i]))
		case reflect.String:
			elems = append(elems, paths[i].(string))
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
	for setK := range values {
		var setV string

		switch reflect.TypeOf(values[setK]).Kind() {
		case reflect.String:
			setV = values[setK].(string)
		case reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int:
			setV = fmt.Sprintf("%d", values[setK])
		case reflect.Slice:
			_setV, err := json.Marshal(values[setK])
			if err != nil {
				return nil, err
			}
			setV = string(_setV)
		default:
			return nil, errors.New(fmt.Sprintf("Query error: unsupported type = %s", reflect.TypeOf(values[setK]).String()))
		}

		query.Set(setK, setV)
	}
	return &query, nil
}
