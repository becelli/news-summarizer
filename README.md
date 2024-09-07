# News Summarizer

News Summarizer is a Go application that fetches, summarizes, evaluates, and translates news articles for a podcast tailored to tech professionals and developers. It leverages a Language Model (LLM) server to generate concise summaries, assess the relevance of articles, and provide translations in Brazilian Portuguese.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/news-summarizer.git
    cd news-summarizer
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Ensure you have a running LLM server at `http://localhost:4000/v1/chat/completions`.

## Usage

To run the application:
```sh
go run main.go
```
The app will fetch articles from RSS feeds, generate summaries, evaluate their importance, translate them, and save the top articles to a markdown file.

## Project Structure
```
├── .gitignore
├── go.mod
├── go.sum
├── internal/
│   ├── articles/
│   │   ├── fetch.go
│   │   ├── summarize.go
│   │   ├── evaluate.go
│   │   ├── translate.go
│   │   ├── save.go
│   │   ├── sort.go
│   ├── llm/
│   │   ├── client.go
│   │   ├── request.go
│   ├── models/
│   │   ├── article.go
│   │   ├── llm.go
├── main.go
```

### Key Files
- **main.go**: Entry point of the application.
- **internal/articles**: Manages article fetching, summarizing, evaluating, translating, and saving.
- **internal/llm**: Handles communication with the LLM server.
- **internal/models**: Defines the data models for articles and LLM requests.
