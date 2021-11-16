module github.com/lilien1010/taskq/example/api_worker/sqs_api_worker

go 1.15

require (
	github.com/aws/aws-sdk-go v1.40.25
	github.com/go-redis/redis/v8 v8.11.4
	github.com/lilien1010/taskq v3.2.7
)

replace github.com/lilien1010/taskq => ../..
