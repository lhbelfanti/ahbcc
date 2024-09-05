-- Insert initial Search Criteria
INSERT INTO search_criteria (name, all_of_these_words, this_exact_phrase, any_of_these_words, none_of_these_words, language, since_date, until_date)
VALUES
    -- Cocaine Search Criteria
    ('Cocaine Search Criteria A', NULL, 'tomarme una linea', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Cocaine Search Criteria B', NULL, 'tomar una linea', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Cocaine Search Criteria C', NULL, 'tomar una raya', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Cocaine Search Criteria D', NULL, 'tomar merca', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Cocaine Search Criteria E', NULL, 'necesito un saque', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Cocaine Search Criteria F', NULL, 'volver a tomar merca', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    -- Marijuana Search Criteria
    ('Marijuana Search Criteria A', NULL, 'fumar un porro', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Marijuana Search Criteria B', NULL, 'necesito fumar porro', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('Marijuana Search Criteria C', NULL, 'un porro', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    -- Heroin Search Criteria
    ('Heroin Search Criteria A', NULL, 'inyectarme heroina', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    -- Ecstasy Search Criteria
    ('Ecstasy Search Criteria A', NULL, 'Necesito éxtasis', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    -- LSD Search Criteria
    ('LSD Search Criteria A', NULL, 'lsd', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE),
    ('LSD Search Criteria B', NULL, 'trip de ácido', NULL, NULL, 'es', '2006-01-01', CURRENT_DATE);


