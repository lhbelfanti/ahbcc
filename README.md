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

### Database

Tables: **Entity Relationship Diagram**

```mermaid
erDiagram
    tweets ||--o| tweets_quotes : ""
    tweets }|--|{ search_criteria : ""
    tweets {
        TEXT hash "CK"
        TIMESTAMP posted_at
        BOOLEAN is_a_reply
        TEXT text_content
        TEXT[] images
        INTEGER quote_id FK
        INTEGER search_criteria_id FK "CK"
    }
    tweets_quotes {
        INTEGER id PK
        BOOLEAN is_a_reply
        TEXT text_content
        TEXT[] images
    }
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

To connect to the database we need to define a `.env` file in the root of the project. It should contain the following environment variables

```
DB_NAME=<Database name>
DB_USER=<Database username>
DB_PASS=<Database password>
```

Replace the `< ... >` by the correct value. For example: `DB_NAME=<Database name>` --> `DB_NAME=ahbcc`.
