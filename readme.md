# Documentation

## Design decision

### Architecture & Folder layout
The project layout followed golang standard project layout. I decided not to move handler func to its own folder "api" because
it contain only 2 endpoints, had it grow bigger, it needs to be moved to its own folder and module.

### Dependency injection for store
Storage is abstracted using interfaces, to anticipate the future need of going into persistence. Had it needed to go for persistence, memory store can be swapped out
with persistence store, without breaking its method.

## API Docs

### POST /shorten
This endpoint will convert given URL to shorten 6 alphanumeric URL

Form Data
> url:

Response
> {
"shortened_url": "http://{host}/{alphanumeric}"
}

### GET /{6 alphanumeric}
This endpoint will redirect given shortened URL to the destination URL
