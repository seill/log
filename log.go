package log

import (
	"encoding/json"
	"fmt"
)

func StartJson(logId string, title string, data interface{}) {
	Start(StructToString(buildJson(logId, title, "START", data)))
}

func Start(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Println(message)
}

func InfoJson(logId string, title string, data interface{}) {
	Info(StructToString(buildJson(logId, title, "INFO", data)))
}

func Info(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Println(message)
}

func ErrorJson(logId string, title string, data interface{}) {
	Error(StructToString(buildJson(logId, title, "ERROR", data)))
}

func Error(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Println(message)
}

func DebugJson(logId string, title string, data interface{}) {
	Debug(StructToString(buildJson(logId, title, "DEBUG", data)))
}

func Debug(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Println(message)
}

func buildJson(logId string, title string, level string, data interface{}) (dataMap map[string]interface{}) {
	var _err error

	dataMap = map[string]interface{}{
		"logId": logId,
		"title": title,
		"level": level,
	}

	switch dataType := data.(type) {
	case nil:
	case map[string]interface{}:
		dataMap["message"] = map[string]interface{}{}

		for k, v := range dataType {
			dataMap["message"].(map[string]interface{})[k] = v
		}
	case string:
		dataMap["message"] = dataType
	case bool:
		dataMap["message"] = dataType
	default:
		var dataBytes []byte
		mapFromData := map[string]interface{}{}

		dataBytes, _err = json.Marshal(dataType)
		if nil != _err {
			dataMap["message"] = fmt.Sprintf("error while marshaling (%v):%s", dataType, _err.Error())
			return
		}

		_err = json.Unmarshal(dataBytes, &mapFromData)
		if nil != _err {
			dataMap["message"] = fmt.Sprintf("error while unmarshaling (%s):%s", string(dataBytes), _err.Error())
			return
		}

		dataMap["message"] = map[string]interface{}{}

		for k, v := range mapFromData {
			dataMap["message"].(map[string]interface{})[k] = v
		}
	}

	return
}

func StructToString(s interface{}) string {
	bytes, err := json.Marshal(s)
	if nil != err {
		return ""
	}

	return string(bytes)
}
