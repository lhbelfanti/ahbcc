<p align="center">
  <img src="media/ahbcc-logo.png" width="100" alt="Repository logo" />
</p>
<p align="center">Adverse Human Behaviors Corpus Creator<p>
<p align="center">
    <img src="https://img.shields.io/github/repo-size/lhbelfanti/ahbcc?label=Repo%20size" alt="Repo size" />
    <img src="https://img.shields.io/github/license/lhbelfanti/ahbcc?label=License" alt="License" />
    <img src="https://codecov.io/gh/lhbelfanti/ahbcc/graph/badge.svg?token=69LLNMKXRU" alt="Coverage" />
</p>

---


# AHBCC: Adverse Human Behaviors Corpus Creator

Adverse Human Behaviors is a term created to encompass all types of human behaviors that affect one or more individuals in physical, psychological, or emotional ways.

There are four main categories:
- Hate speech
- Depression and/or suicidal attempt
- Eating disorders
- Illicit drug use

## Application
This application serves as the orchestrator, utilizing a docker-compose.yml file to **connect the other two applications with the database managed by [AHBCC](https://github.com/lhbelfanti/ahbcc)**.

The primary objective is to gather information from X (formerly Twitter) using [GoXCrap](https://github.com/lhbelfanti/goxcrap). Subsequently, each tweet is manually evaluated to determine if it discusses an Adverse Human Behavior using [Binarizer](https://github.com/lhbelfanti/binarizer). Finally, [AHBCC](https://github.com/lhbelfanti/ahbcc) is in charge of creating a balanced corpus from the retrieved and categorized tweets.

### Endpoints

To allow [GoXCrap](https://github.com/lhbelfanti/goxcrap) to save the tweets into the database and then retrieve them using [Binarizer](https://github.com/lhbelfanti/binarizer), this application exposes different endpoints, encapsulating the access to the database in one place (this app).

#### Network
This app calls an endpoint defined by the env variable `ENQUEUE_CRITERIA_API_URL`. To ensure proper communication, the app that owns this endpoint must be on the same network (named shared), which is defined in the [compose.yml](compose.yml) as follows:
```
networks:
  shared_network:
    driver: bridge
    name: shared_network
```

To join the same network, the corresponding `compose.yml` for the other app should include the following configuration:
```
networks:
  shared_network:
    external: true
```

### Database

Tables: **Entity Relationship Diagram**

```mermaid
erDiagram
    tweets ||--o| tweets_quotes : ""
    tweets }|--|{ search_criteria : ""
    tweets {
        INTEGER uuid PK
        TEXT id
        TEXT author
        TEXT avatar
        TIMESTAMP posted_at
        BOOLEAN is_a_reply
        TEXT text_content
        TEXT[] images
        INTEGER quote_id FK
        INTEGER search_criteria_id FK
    }
    tweets_quotes {
        INTEGER id PK
        TEXT author
        TEXT avatar
        TIMESTAMP posted_at
        BOOLEAN is_a_reply
        TEXT text_content
        TEXT[] images
    }
    search_criteria ||--o{ search_criteria_executions : ""
    search_criteria {
        INTEGER id PK
        TEXT name
        TEXT[] all_of_these_words
        TEXT this_exact_phrase
        TEXT[] any_of_these_words
        TEXT[] none_of_these_words
        TEXT[] these_hashtags
        TEXT language
        DATE since_date
        DATE until_date
    }
    search_criteria_executions ||--o{ search_criteria_execution_days : ""
    search_criteria_executions {
        INTEGER id PK
        ENUM status
        INTEGER search_criteria_id FK
    }
    search_criteria_execution_days {
        INTEGER id PK
        DATE execution_date
        INTEGER tweets_quantity
        TEXT error_reason
        INTEGER search_criteria_execution_id FK
    }
    users {
        INTEGER id PK
        TEXT name
    }
    categorized_tweets ||--|{ tweets : ""
    categorized_tweets ||--|{ users : ""
    categorized_tweets {
        INTEGER id PK
        INTEGER tweet_id FK
        INTEGER user_id FK
        BOOLEAN adverse_behavior
    }
```

#### Necessary files to start the database

To connect to the database create a `.env` file in the root of the project or rename the provided [.env.example](.env.example). 

This file should contain the following environment variables:
```
# App settings
APP_EXPOSED_PORT=<AHBCC Host Port>
APP_INTERNAL_PORT=<AHBCC Container Port>

# Database
DB_NAME=<Database name>
DB_USER=<Database username>
DB_PASS=<Database password>
DB_PORT=<Database port>

# External APIs URLs
ENQUEUE_CRITERIA_API_URL=<Domain of the application with the endpoint /criteria/enqueue/v1> --> Example: the URL to the GoXCrap API
```

Replace the `< ... >` by the correct value. For example: `DB_NAME=<Database name>` --> `DB_NAME=ahbcc`.
