<p align="center">
  <img src="media/ahbcc-logo.png" width="100" alt="Repository logo" />
</p>
<p align="center">Adverse Human Behaviors Corpus Creator<p>
<p align="center">
    <img src="https://img.shields.io/github/repo-size/lhbelfanti/ahbcc?label=Repo%20size" alt="Repo size" />
    <img src="https://img.shields.io/github/license/lhbelfanti/ahbcc?label=License" alt="License" />
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
This application serves as the orchestrator, utilizing a docker-compose.yml file to **connect the other three applications among themselves and to the database it manages**.

The primary objective is to gather information from X (formerly Twitter) using [GoXCrap](https://github.com/lhbelfanti/goxcrap). Subsequently, each tweet is manually evaluated to determine if it discusses an Adverse Human Behavior using [Binarizer](https://github.com/lhbelfanti/binarizer). Finally, a balanced corpus is created from this data using Corpus Creator.

### Database

Tables: **Entity Relationship Diagram**

```mermaid
erDiagram
    tweets {
        INTEGER id PK
        TIMESTAMP tweet_timestamp
        BOOLEAN is_a_reply
        BOOLEAN has_text
        BOOLEAN has_images
        TEXT text
        TEXT[] images
        BOOLEAN has_quote
        INTEGER quote_id FK
    }
    tweets ||--o| tweets_quotes : ""
    tweets_quotes {
        INTEGER id PK
        BOOLEAN quote_is_a_reply
        BOOLEAN quote_has_text
        BOOLEAN quote_has_images
        TEXT quote_text
        TEXT[] quote_images
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
```