package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	textserialization "github.com/microsoft/kiota-serialization-text-go"
)

func WriteJson(data serialization.Parsable, c *gin.Context) {
	//writer := jsonserialization.NewJsonSerializationWriter()
	writer := textserialization.NewTextSerializationWriter()
	err := data.Serialize(writer)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	content, err := writer.GetSerializedContent()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.String(200, "{"+string(content)+"}")
}
