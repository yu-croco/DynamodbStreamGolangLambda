package converter

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"yu-croco.com/DynamodbStreamGolangLambda/app/model"
)

// map[string]events.DynamoDBAttributeValueを
// map[string]*dynamodb.AttributeValueに変換して
// structに変換してる
func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue) (model.RecordImage, error) {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)
	var image model.RecordImage

	for idx, value := range attribute {
		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := value.MarshalJSON(); if marshalErr != nil {
			return image, marshalErr
		}

		if unmarshalErr := json.Unmarshal(bytes, &dbAttr); unmarshalErr != nil {
			return image, unmarshalErr
		}

		dbAttrMap[idx] = &dbAttr
	}

	if unmarshalErr := dynamodbattribute.UnmarshalMap(dbAttrMap, &image); unmarshalErr != nil {
		return image, unmarshalErr
	}

	return image, nil
}

func ToModel(event events.DynamoDBEvent) (*model.DynamodbStream, error) {
	var records model.Records
	for _, e := range event.Records {
		newImage, newImageErr := unmarshalStreamImage(e.Change.NewImage)
		if newImageErr != nil {
			return nil, newImageErr
		}

		oldImage, oldImageErr := unmarshalStreamImage(e.Change.OldImage)
		if oldImageErr != nil {
			return nil, oldImageErr
		}

		record := model.Record{
			Dynamodb: model.DynamoDBImages{
				NewImage: newImage,
				OldImage: oldImage,
			},
			EventName: e.EventName,
		}
		records = append(records, record)
	}

	events := model.DynamodbStream{Records: records}

	return &events, nil
}
