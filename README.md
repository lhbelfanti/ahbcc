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
This application serves as the orchestrator, using a docker-compose.yml file to **connect the other two applications with the database managed by [AHBCC](https://github.com/lhbelfanti/ahbcc)**.

The primary goal is to gather information from X (formerly Twitter) using [GoXCrap](https://github.com/lhbelfanti/goxcrap). Subsequently, each tweet is manually evaluated to determine if it discusses an Adverse Human Behavior using [Binarizer](https://github.com/lhbelfanti/binarizer). Finally, [AHBCC](https://github.com/lhbelfanti/ahbcc) is in charge of creating a corpus from the retrieved and categorized tweets.

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
        ENUM status "'PENDING', 'IN PROGRESS', 'DONE'"
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
        TEXT username
        TEXT password_hash
        TIMESTAMP created_at
    }
    categorized_tweets ||--|{ search_criteria : ""
    categorized_tweets ||--|{ tweets : ""
    categorized_tweets ||--|{ users : ""
    categorized_tweets {
        INTEGER id PK
        INTEGER search_criteria_id FK "Intentional redundancy"
        INTEGER tweet_id FK
        INTEGER tweet_year "Intentional redundancy"
        INTEGER tweet_month "Intentional redundancy"
        INTEGER user_id FK
        ENUM categorization "'POSITIVE', 'INDETERMINATE', 'NEGATIVE'"
    }
    users_sessions ||--|{ users : ""
    users_sessions {
        INTEGER id PK
        INTEGER user_id FK
        TEXT token
        TIMESTAMP expires_at
        TIMESTAMP created_at
    }
    search_criteria_executions_summary ||--|{ search_criteria : ""
    search_criteria_executions_summary {
        INTEGER id PK
        INTEGER search_criteria_id FK
        INTEGER tweets_year
        INTEGER tweets_month
        INTEGER total_tweets
    }
    
    corpus { 
        INTEGER id PK 
        TEXT tweet_author
        TEXT tweet_avatar
        TEXT tweet_text
        TEXT[] tweet_images
        BOOLEAN is_tweet_a_reply
        TEXT quote_author
        TEXT quote_avatar
        TEXT quote_text
        TEXT[] quote_images
        BOOLEAN is_quote_a_reply
    }

    corpus_cleaning_rules {
        INTEGER id PK
        ENUM cleaning_rules "'BAD_WORD', 'PERSON', etc."
        TEXT source_text
        TEXT target_text
        INTEGER priority
    }
```

> The corpus table is designed as a denormalized structure that consolidates information from both tweets and their 
> corresponding quoted tweets into a single record. This schema is intended to optimize read performance, particularly 
> during data extraction processes such as exporting to CSV or JSON formats. By avoiding JOIN operations between the 
> tweets and quotes tables, the denormalized design reduces query complexity and improves efficiency at the cost of 
> some data redundancy, which is considered acceptable in this context due to the read-heavy nature of the task.


## Setup

To run the application, you must define specific environment variables. 
You can create a .env file in the root directory of the project or rename the provided example file, .env.example.

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

# Session
SESSION_SECRET_KEY=<Secret key used for signing and verifying HMAC-based tokens>

# External APIs URLs
ENQUEUE_CRITERIA_API_URL=<Domain of the application with the endpoint /criteria/enqueue/v1> --> Example: the URL to the GoXCrap API
```

Replace the `< ... >` by the correct value. For example: `DB_NAME=<Database name>` --> `DB_NAME=ahbcc`.

#### Session secret key

The value of the session secret key should be a random and long byte sequence that isn't easily guessable.

An example of how to generate this key is by using a tool like `OpenSSL` or a `password generator`.
```bash
openssl rand -base64 32
```
This generates a 256-bit (32-byte) key encoded in Base64, which is suitable for HMAC-SHA256.


