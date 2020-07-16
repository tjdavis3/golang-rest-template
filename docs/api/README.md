# API Overview

This API is based on simple REST principles.   The API endpoints return JSON responses and are defined using the OpenAPI specification which makes the generation of client SDK much easier.

## Requests

The API is based on REST principles. Data resources are accessed via standard HTTPS requests in UTF-8 format to an API endpoint. Where possible, the API uses appropriate HTTP verbs for each action:

METHOD |	ACTION
------ | ---------
GET |	Retrieves resources
POST |	Creates resources
PUT |	Changes and/or replaces resources or collections in their entirety
PATCH |  Changes a portion of a resource
DELETE	 |Deletes resources

## Responses

The API returns all response data as a JSON object. See the Web API Object Model for a description of all the retrievable objects.

## Timestamps

Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.
