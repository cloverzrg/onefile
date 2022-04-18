package util

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

func ToJson(data serialization.Parsable) (res string, err error) {
	if data == nil {
		return "{}", nil
	}
	writer := NewJsonSerializationWriter()
	err = data.Serialize(writer)
	if err != nil {
		return res, err
	}
	content, err := writer.GetSerializedContent()
	if err != nil {
		return res, err
	}
	return string(content), nil
}
