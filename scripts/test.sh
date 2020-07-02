aws lambda invoke \
  --function-name attachTag \
  --cli-binary-format raw-in-base64-out \
  --payload fileb://tests/test1.json output.txt \
  --log-type Tail \
  --query 'LogResult' | tr -d '"' | base64 -D
