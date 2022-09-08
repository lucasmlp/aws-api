aws ecr create-repository \
    --repository-name aws-api \
    --image-scanning-configuration scanOnPush=true \
    --region us-west-2