# Metrics

The API exposes a variety of metrics in standard Prometheus format.  These metrics are exposed at the `/metrics` endpoint.

Name  | Type | Labels | Description
----  | ---- | ------ | -----------
api_processing_ops_total | gauge | | The number of events processing
api_responses_total | counter |  - HTTP Status<br/>- method<br>- endpoint  | The number of responses by endpoint and status
