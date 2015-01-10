package gocontrol

import (
	"reflect"
	"strconv"
	"strings"
	"errors"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Render(template string, args ...string) string {
	dat, err := ioutil.ReadFile("gocontrol/views/" + template + ".go.html")
    check(err)
    templateString := string(dat)
    /* TODO: Optimize this */
    for counter, arg := range args {
		counterVariable := "{{ " + strconv.Itoa(counter) + " }}"
		templateString = strings.Replace(templateString, counterVariable, arg, -1)
	}
    return string(templateString)
}

func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
    f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        err = errors.New("The number of params is not adapted.")
        return
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return
}

func CallController(controller string, headers map[string]string) string {
	controllers := map[string]interface{}{"index": index, "test": test, "ip": ip, "browser": browser}
	x, _ := Call(controllers, controller, headers)
	return x[0].String()
}