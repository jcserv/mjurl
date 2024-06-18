# mjurl

mjurl ("Mike and Jarrod's URL Shortening Service") is a simple Go API that allows users to shorten URLs.
Based on ByteByteGo's URL Shortener problem in System Design Interview Vol. 1

## Getting Started

### Starting the Server with Docker
1. Ensure you have [Docker Engine](https://docs.docker.com/engine/install/) installed.
2. Run `make run`
3. Execute requests against `http://localhost:8080` to your heart's content

## API Routes

### POST /v1/api/url
Request Body:
```
{
    "url": "https://github.com"
}
```

Response:
```
200 OK
"<shortened URL hash>"
```

### GET /v1/api/url/{shortURL}
Response:
```
308 PERMANENT REDIRECT
Location: <longURL>
```

## Contributors
- [@jcserv](https://github.com/jcserv)
- [@imphungky](https://github.com/imphungky)

## License

Distributed under the GNU-GPL License. See `LICENSE` for more information.