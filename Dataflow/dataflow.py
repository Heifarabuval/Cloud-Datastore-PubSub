import json
import logging
import apache_beam as beam
from apache_beam import Map
from apache_beam.io import ReadFromPubSub
from apache_beam.io.gcp.datastore.v1new.datastoreio import WriteToDatastore
from apache_beam.options.pipeline_options import PipelineOptions
from apache_beam.io.gcp.datastore.v1new.types import Entity, Key

PROJECT_NAME = "heifara-test"
subscription = "projects/{}/subscriptions/{}".format(PROJECT_NAME, "getCompute-sub")


def create(data):
    entity = Entity(key=Key(
        path_elements=['Computation', data['computation_id']],
        project=PROJECT_NAME
    ))

    entity.set_properties({
        "computed": True,
        "result": data['result'],
        "values": data["values"],
        "webhookId": data["webhook_id"]
    })
    return entity


def run():
    pipeline_options = PipelineOptions(streaming=True, save_main_session=True)

    with beam.Pipeline(options=pipeline_options) as pipeline:
        (
                pipeline

                | "Read from PubSub"
                >> ReadFromPubSub(subscription=subscription)
                .with_output_types(bytes)

                | "Data to utf8"
                >> Map(lambda x: x.decode("utf-8"))

                | "Load json data"
                >> Map(json.loads)

                | "Get entity"
                >> Map(create)

                | "Write to datastore"
                >> WriteToDatastore(PROJECT_NAME)
        )


if __name__ == '__main__':
    logging.getLogger().setLevel(logging.INFO)
    run()
