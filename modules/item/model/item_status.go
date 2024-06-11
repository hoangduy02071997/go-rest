package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusTodo ItemStatus = iota
	ItemStatusDoing
	ItemStatusDone
	ItemStatuDeleted
)

var allItemStatuses = [4]string{"todo", "doing", "done", "deleted"}

func (i *ItemStatus) String() string { return allItemStatuses[*i] }

func ParseItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), fmt.Errorf("invalid item status: %q", s)
}

// DataDB -> struct // Problem with Postgre
//func (i *ItemStatus) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//
//	if !ok {
//		return errors.New(fmt.Sprintf("fail to scan data from sql %s", value))
//	}
//
//	v, err := ParseItemStatus(string(bytes))
//	if err != nil {
//		return errors.New(fmt.Sprintf("fail to scan data from sql %s", value))
//	}
//	*i = v
//	return nil
//}

// DataDB -> struct // work on 2 db
func (i *ItemStatus) Scan(value interface{}) error {
	if value == nil {
		*i = ItemStatusTodo // or some default value
		return nil
	}

	var strValue string
	switch v := value.(type) {
	case string:
		strValue = v
	case []byte:
		strValue = string(v)
	default:
		return fmt.Errorf("fail to scan data from sql: %v", value)
	}

	status, err := ParseItemStatus(strValue)
	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %v", value)
	}
	*i = status
	return nil
}

// Data struct -> json (json encoding)
func (i *ItemStatus) MarshalJSON() ([]byte, error) {
	if i == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", i.String())), nil
}

// Data json -> struct (json decoding)
func (i *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemStr, err := ParseItemStatus(str)
	if err != nil {
		return err
	}

	*i = itemStr
	return nil
}

// Data struct -> db
func (i *ItemStatus) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.String(), nil
}
