package util

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
)

func ToJson(data serialization.Parsable) (res string, err error) {
	if data == nil {
		return "{}", nil
	}
	writer := jsonserialization.NewJsonSerializationWriter()
	//err = data.Serialize(writer)
	err = writer.WriteObjectValue("", data)
	if err != nil {
		return res, err
	}
	content, err := writer.GetSerializedContent()
	if err != nil {
		return res, err
	}
	return string(content), nil
}
