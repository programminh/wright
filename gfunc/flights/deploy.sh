GOOS=linux go build
gcloud beta functions deploy flights --stage-bucket wright-functions --trigger-topic flights
rm flights
