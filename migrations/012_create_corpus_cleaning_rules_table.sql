-- Create the enum type for rules types
SELECT create_enum_type_if_not_exists('cleaning_rules', ARRAY['BAD_WORD', 'REPLACE', 'DELETE']);

-- Create corpus_cleaning_rules table
CREATE TABLE IF NOT EXISTS corpus_cleaning_rules (
    id              SERIAL PRIMARY KEY,
    rule_type       cleaning_rules NOT NULL,
    source_text     TEXT NOT NULL,
    target_text     TEXT,
    priority        INTEGER DEFAULT 10 CHECK (priority >= 1 AND priority <= 10),
    description     TEXT,

    CONSTRAINT uq_rule_type_source_text_priority UNIQUE (rule_type, source_text, priority)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_corpus_cleaning_rules_priority ON corpus_cleaning_rules(priority);

-- Table comments
COMMENT ON TABLE corpus_cleaning_rules                IS 'Records the corpus generated after the tweets categorization';
COMMENT ON COLUMN corpus_cleaning_rules.id            IS 'Auto-incrementing id of the entry of the corpus, agnostic to business logic';
COMMENT ON COLUMN corpus_cleaning_rules.rule_type     IS 'The type of rule that will be applied to the text';
COMMENT ON COLUMN corpus_cleaning_rules.source_text   IS 'The regex that the rule will identify for cleaning';
COMMENT ON COLUMN corpus_cleaning_rules.target_text   IS 'The replacement text for the source_text, if not already given by the rule_type';
COMMENT ON COLUMN corpus_cleaning_rules.priority      IS 'The priority of the rule, used to order the rules applied to the text. Lower values are applied first.';
COMMENT ON COLUMN corpus_cleaning_rules.description   IS 'A description of the rule, used for debugging purposes';
