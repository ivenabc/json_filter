package json_filter

import (
	"github.com/Jeffail/gabs"
)

const (
	BaseKeys   = "baseKeys"
	ObjectKeys = "objectKeys"
	ArrayKeys  = "arrayKeys"
)

func GenerateFilterJson(jsonBytes, parseBytes []byte) interface{} {
	oldObj, err := gabs.ParseJSON(jsonBytes)
	if err != nil {
		return nil
	}

	parse, err := gabs.ParseJSON(parseBytes)
	if err != nil {
		return nil
	}

	jsonObj := gabs.New()

	transfer := transfer{
		OldObj:   oldObj,
		ParseObj: parse,
		NewObj:   jsonObj,
	}

	transfer.setKeys()
	return transfer.NewObj.Data()
}

type transfer struct {
	OldObj   *gabs.Container
	ParseObj *gabs.Container
	NewObj   *gabs.Container
}

func (t transfer) setKeys() {
	setObjectKeys(t.OldObj, t.NewObj, t.ParseObj)
}

func setObjectKeys(oldObj *gabs.Container, newObj *gabs.Container, parseObj *gabs.Container) {
	baseChildren, err := parseObj.S(BaseKeys).Children()
	if err != nil {
		goto ObjectParse
	}

	for _, child := range baseChildren {
		// set base keys
		key := child.Data().(string)
		value := getData(oldObj, key)
		if value != nil {
			newObj.Set(value, key)
		}
	}

ObjectParse:
	objChildren, err := parseObj.S(ObjectKeys).Children()
	if err != nil {
		goto ArrayParse
	}

	for _, child := range objChildren {
		key := child.Data().(string)
		newContainer, err := newObj.Set(make(map[string]interface{}, 0), key)
		if err != nil {
			continue
		}
		oldContainer := oldObj.S(key)

		parseContainer := parseObj.S(key)

		setObjectKeys(oldContainer, newContainer, parseContainer)
	}

ArrayParse:
	arrayChildren, err := parseObj.S(ArrayKeys).Children()
	for _, child := range arrayChildren {
		key := child.Data().(string)

		_, err := newObj.Array(key)
		if err != nil {
			continue
		}

		children, _ := oldObj.S(key).Children()
		for _, c := range children {
			jsonObj := gabs.New()
			parseContainer := parseObj.S(key)
			setObjectKeys(c, jsonObj, parseContainer)

			err := newObj.ArrayAppend(jsonObj.Data(), key)
			if err != nil {
				continue
			}
		}

	}
}

func getData(obj *gabs.Container, key string) interface{} {

	return obj.Path(key).Data()
}

//func getObjectDat
