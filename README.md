# YouTube Fetch

This project fetches the latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

## Features
- **YouTube API Fetcher**: Continuously fetches the latest YouTube videos for a predefined search query, updating every minute.
- **Data Storage**: Stores video data including title, description, publishing date-time, and thumbnail URLs in a PostgreSQL database.
- **Paginated API**: Provides a GET API to retrieve stored video data in reverse chronological order.
- **Search API**: Search API to search stored videos by their title and description.

## Directory Structure
Manuelmastro-youtube-fetch/ ├── Dockerfile ├── go.mod ├── go.sum ├── main.go ├── config/ │ └── models.go ├── handlers/ │ └── handlers.go ├── models/ │ └── model.go └── yt_api/ ├── api.go ├── background.go └── store.go


## Setup Instructions

### Clone the Repository

```bash
git clone https://github.com/your-username/Manuelmastro-youtube-fetch.git
cd Manuelmastro-youtube-fetch
```

### Install Dependencies

```bash
go mod tidy
```





### Setup environment variables

Create an `.env` file in the root directory with the following parameters:

```env
YOUTUBE_API_KEYS=<api_key_1>,<api_key_2>,<api_key_3>
DB_CONNECTION=<psql_connection_string>
QUERY=<query_parameter>
```
YOUTUBE_API_KEYS: Comma-separated list of YouTube API keys (to rotate between them in case of quota limitations).
DB_CONNECTION: PostgreSQL connection string to connect to the database.
QUERY: The search query to fetch YouTube videos for (e.g., "football").

### Run the services:

```bash
go run main.go
```






