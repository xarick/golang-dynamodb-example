package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/xarick/golang-dynamodb-example/models"
)

// CreateUser - foydalanuvchi qo'shish
func CreateUser(user *models.User) error {
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	_, err = DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      item,
	})

	return err
}

// GetAllUsers - barcha foydalanuvchilarni olish
func GetAllUsers() ([]models.User, error) {
	output, err := DynamoDB.Scan(&dynamodb.ScanInput{
		TableName: aws.String("users"),
	})
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &users)

	return users, err
}

// GetUserByID - foydalanuvchini ID orqali olish
func GetUserByID(id string) (*models.User, error) {
	output, err := DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	})
	if err != nil || output.Item == nil {
		return nil, err
	}

	var user models.User
	err = dynamodbattribute.UnmarshalMap(output.Item, &user)
	return &user, err
}

// UpdateUser - foydalanuvchini yangilash
func UpdateUser(id string, user *models.User) error {
	_, err := DynamoDB.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("name"), // 'name' atributini #name bilan almashtirish
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name":  {S: aws.String(user.Name)},
			":email": {S: aws.String(user.Email)},
		},
		UpdateExpression:    aws.String("SET #name = :name, email = :email"),
		ConditionExpression: aws.String("attribute_exists(ID)"), // Bu faqat mavjud yozuvlarni yangilashga imkon beradi (yangi qo'shmaydi)
	})
	return err
}

// DeleteUser - foydalanuvchini o'chirish
func DeleteUser(id string) error {
	_, err := DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	})
	return err
}
