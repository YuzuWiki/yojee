package pixivService

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	netUrl "net/url"
	"reflect"
	"strings"
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
