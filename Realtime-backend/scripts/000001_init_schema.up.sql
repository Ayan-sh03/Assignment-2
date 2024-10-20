CREATE TABLE IF NOT EXISTS weather_summaries (
    city TEXT NOT NULL,
    date DATE NOT NULL,
    data JSONB NOT NULL,
    PRIMARY KEY (city, date)
);