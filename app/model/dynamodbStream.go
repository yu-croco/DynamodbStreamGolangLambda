package model

type DynamodbStream struct {
	Records Records `json:"Records"`
}

type Records []Record

type Record struct {
	Dynamodb  DynamoDBImages `json:"dynamodb"`
	EventName string         `json:"eventName"`
}

type DynamoDBImages struct {
	NewImage RecordImage `json:"NewImage"`
	OldImage RecordImage `json:"OldImage"`
}

type RecordImage struct {
	UserId int `json:"userId" dynamodbav:"userId"`
	Age    int `json:"age" dynamodbav:"age"`
	Address string `json:"address" dynamodbav:"address"`
}
