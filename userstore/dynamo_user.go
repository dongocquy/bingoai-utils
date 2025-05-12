package userstore

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// 🔎 Lấy user config theo ID (email)
func GetUserFromDynamo(client *dynamodb.Client, id string) (*UserConfig, error) {
	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: id},
	}

	result, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TableUserConfig),
		Key:       key,
	})
	if err != nil {
		return nil, fmt.Errorf("❌ DynamoDB GetItem lỗi: %v", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("❌ Không tìm thấy user có id = %s", id)
	}

	var user UserConfig
	if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
		return nil, fmt.Errorf("❌ Unmarshal lỗi: %v", err)
	}

	return &user, nil
}

// 🔄 Truy ngược từ user_email → client_email
func GetClientEmailFromUser(userEmail string, db *dynamodb.Client) (string, error) {
	out, err := db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TableClientUsers),
		Key: map[string]types.AttributeValue{
			"user_email": &types.AttributeValueMemberS{Value: userEmail},
		},
	})
	if err != nil {
		return "", fmt.Errorf("❌ DynamoDB GetItem lỗi: %w", err)
	}
	if out.Item == nil {
		return "", fmt.Errorf("⛔ Không tìm thấy client cho user: %s", userEmail)
	}

	var result ClientUser
	if err := attributevalue.UnmarshalMap(out.Item, &result); err != nil {
		return "", fmt.Errorf("❌ Unmarshal lỗi: %w", err)
	}
	return result.ClientEmail, nil
}

// 📋 Lấy danh sách user thuộc 1 client
func ListUsersByClientEmail(clientEmail string, db *dynamodb.Client) ([]ClientUser, error) {
	out, err := db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(TableClientUsers),
		IndexName:              aws.String("ClientEmailGSI"),
		KeyConditionExpression: aws.String("client_email = :val"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": &types.AttributeValueMemberS{Value: clientEmail},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("query lỗi: %w", err)
	}

	var users []ClientUser
	err = attributevalue.UnmarshalListOfMaps(out.Items, &users)
	return users, err
}

// 🔎 Lấy ProjectID từ client email
func GetProjectIDFromClient(clientEmail string, db *dynamodb.Client) (string, error) {
	out, err := db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TableUserConfig),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: clientEmail},
		},
	})
	if err != nil {
		return "", fmt.Errorf("❌ Lỗi GetItem: %w", err)
	}
	if out.Item == nil {
		return "", fmt.Errorf("⛔ Không tìm thấy client: %s", clientEmail)
	}

	var config UserConfig
	if err := attributevalue.UnmarshalMap(out.Item, &config); err != nil {
		return "", fmt.Errorf("❌ Lỗi giải mã: %w", err)
	}
	return config.ProjectID, nil
}

// 🔄 Lấy client từ ProjectID (truy ngược bằng GSI)
func GetClientFromProjectID(projectID string, db *dynamodb.Client) (*UserConfig, error) {
	out, err := db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(TableUserConfig),
		IndexName:              aws.String("ProjectIDIndex"),
		KeyConditionExpression: aws.String("project_id = :val"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": &types.AttributeValueMemberS{Value: projectID},
		},
		Limit: aws.Int32(1),
	})
	if err != nil {
		return nil, fmt.Errorf("❌ Query lỗi: %w", err)
	}
	if len(out.Items) == 0 {
		return nil, fmt.Errorf("⛔ Không tìm thấy client với project_id: %s", projectID)
	}

	var config UserConfig
	if err := attributevalue.UnmarshalMap(out.Items[0], &config); err != nil {
		return nil, fmt.Errorf("❌ Lỗi giải mã: %w", err)
	}
	return &config, nil
}
