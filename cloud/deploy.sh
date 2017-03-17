pushd ..
GOOS=linux go build
cp wright cloud
popd
gcloud beta functions deploy wright --stage-bucket tp-wright --trigger-topic wright
