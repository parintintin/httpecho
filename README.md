# httpecho
httpecho is a HTTP server which responds with the requested HTTP status code.

## Usage
```
httpecho [options] 
options:
-host="localhost": Host for HTTP server
-port=80: Port to bind HTTP server to
```

## Client usage

### Single status code

URL path: */{[StatusCode](#status-codes)}*

**Request:**
```
curl --location http://localhost/200
```
**Response:**
````
< HTTP/1.1 200 OK
< Access-Control-Allow-Credentials: true
< Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token
< Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
< Access-Control-Allow-Origin: *
...
< Content-Type: text/plain; charset=utf-8
```
### Redirects
Redirect is supported for HTTP status codes 301 and 307. Additionally a second status code has to be specified.

URL path: */{301|307}/{[StatusCode](#status-codes)}*

**Request:**
```
curl --location http://localhost/301/200
```
**Response (redirect to /200):**
````
< HTTP/1.1 301 Moved Permanently
...
< Location: http://localhost/200
```
### Parameters
#### Delay
URL parameter `d` specifies a delay of the response in **seconds**.

Example for a 10 second delay:
```
url -s -w "%{time_total}s\n" -o /dev/null http://localhost/408?d=10
10.012s
```
## Supported HTTP status codes (RFC 2616)<a id="status-codes"></a>
* 100: Continue
* 101: Switching Protocols
* 200: OK
* 201: Created
* 202: Accepted
* 203: Non-Authoritative Information
* 204: No Content
* 205: Reset Content
* 206: Partial Content
* 300: Multiple Choices
* 301: Moved Permanently
* 302: Found
* 303: See Other
* 304: Not Modified
* 305: Use Proxy
* 307: Temporary Redirect
* 400: Bad Request
* 401: Unauthorized
* 402: Payment Required
* 403: Forbidden
* 404: Not Found
* 405: Method Not Allowed
* 406: Not Acceptable
* 407: Proxy Authentication Required
* 408: Request Timeout
* 409: Conflict
* 410: Gone
* 411: Length Required
* 412: Precondition Failed
* 413: Request Entity Too Large
* 414: Request URI Too Long
* 415: Unsupported Media Type
* 416: Requested Range Not Satisfiable
* 417: Expectation Failed
* 418: I'm a teapot
* 428: Precondition Required
* 429: Too Many Requests
* 431: Request Header Fields Too Large
* 451: Unavailable For Legal Reasons
* 500: Internal Server Error
* 501: Not Implemented
* 502: Bad Gateway
* 503: Service Unavailable
* 504: Gateway Timeout
* 505: HTTP Version Not Supported
* 511: Network Authentication Required
