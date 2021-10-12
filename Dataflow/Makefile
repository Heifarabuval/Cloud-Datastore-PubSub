
.PHONY: deploy

deploy:
	@python3 dataflow.py --streaming --runner DirectRunner \
	--project heifara-test --region europe-west1 \
	--output gs://dataflow-data