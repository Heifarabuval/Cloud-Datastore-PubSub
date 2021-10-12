
.PHONY: deploy

deploy:
	@python main.py --streaming --runner DataflowRunner \
	--project calicut-315113 --region europe-west1 \
	--job_name "processed-tasks-`date +"%Y%m%d"`-`uuid | cut -d- -f1`" --requirements_file requirements.txt \