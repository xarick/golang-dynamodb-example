# golang-dynamodb-example

- Running DynamoDB
```bash
docker run -d -p 8000:8000 amazon/dynamodb-local
```

- Setting Up AWS CLI Configuration
```bash
aws configure

AWS Access Key ID [None]: test
AWS Secret Access Key [None]: test
Default region name [None]: us-east-1
Default output format [None]: json
```

- View list of tables
```bash
aws dynamodb list-tables --endpoint-url http://127.0.0.1:8000 --region us-east-1
```

- Creating a table
```sql
aws dynamodb create-table \
    --table-name users \
    --attribute-definitions AttributeName=ID,AttributeType=S \
    --key-schema AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url http://127.0.0.1:8000 \
    --region us-east-1
```
