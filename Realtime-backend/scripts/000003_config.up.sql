CREATE TABLE weather_configs (
    id SERIAL PRIMARY KEY,
    cities TEXT[] NOT NULL,
    threshold_temperature DOUBLE PRECISION NOT NULL,
    email VARCHAR(255) NOT NULL,
    consecutive_alert_threshold INTEGER NOT NULL
);
