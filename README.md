# fampay-youtube

## Goto `localhost:50051` after running to test the API

1. Pagination commands  
2. Debounced Search
3. Multiple API-key Support
4. Dockerized
5. For searching on a bigger dataset change the `max_results` parameter or the `query` parameter in `youtube-video.util.go`

## ✨ How to run

1. Rename `.env.example` to `.env`
2. Add Google API Keys in the `.envfile`

Run with `docker-compose`

    ```bash
    docker-compose up --build
    ```

Run with `go`

    ```bash
    go run main.go
    ```
