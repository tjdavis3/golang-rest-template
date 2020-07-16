# Response Codes

The standard  API returns HTTP status codes in addition to JSON-based error codes and messages.


## HTTP Status Codes

The API attempts to return appropriate HTTP status codes for every request.

Code |	Text |	Description |
--- | --- | ---
200 |	OK |	Success!
201 | Created | The requested object has been created
202 | Accepted | The request has been accepted for processing, but the processing has not been completed. The request might or might not eventually be acted upon, as it might be disallowed when processing actually takes place.
304 |	Not Modified |	There was no new data to return.
400 |	Bad Request	 |The request was invalid or cannot be otherwise served. An accompanying error message will explain further. Requests without authentication are considered invalid and will yield this response.
401 |	Unauthorized |	Missing or incorrect authentication credentials. This may also returned in other undefined circumstances.
403 |	Forbidden |	The request is understood, but it has been refused or access is not allowed. An accompanying error message will explain why. This code is used when requests are being denied due to update limits . Other reasons for this status being returned are listed alongside the error codes in the table below.
404 |	Not Found |	The URI requested is invalid or the resource requested, such as a user, does not exist.
405 | Method not allowed | An invalid HTTP verb was specified for the given endpoint
406 |	Not Acceptable |	Returned when an invalid format is specified in the request.
410 |	Gone |	This resource is gone. Used to indicate that an API endpoint has been turned off and is no longer available.
420 |	Enhance Your Calm |	Returned when an app is being rate limited for making too many requests.
422 |	Unprocessable Entity |	Returned when the data is unable to be processed (for example, if an image uploaded to POST account / update_profile_banner is not valid, or the JSON body of a request is badly-formed).
429 |	Too Many Requests |	Returned when a request cannot be served due to the app's rate limit having been exhausted for the resource. See Rate Limiting.
500 |	Internal Server Error |	Something is broken. This is usually a temporary error, for example in a high load situation or if an endpoint is temporarily having issues. Check in the developer forums in case others are having similar issues,  or try again later.
501 | Not Implemented | The requested function has not yet been implemented.
502 |	Bad Gateway |	The service is down, or being upgraded.
503 |	Service Unavailable |	The servers are up, but overloaded with requests. Try again later.
504 |	Gateway timeout |	The  servers are up, but the request couldn’t be serviced due to some failure within the internal stack. Try again later.

## Error Messages

The API error messages are returned in JSON format.  An example of an error might look like this:

```json
{
  "code": "999",
  "message": "An error has occurred"
}
```
