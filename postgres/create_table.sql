CREATE TABLE summary (
    summary_id VARCHAR (64) PRIMARY KEY,
    summary TEXT NOT NULL,
    original_text TEXT NOT NULL,
    temperature FLOAT(3) NOT NULL,
    top_p FLOAT(3) NOT NULL,
    engine VARCHAR(32) NOT NULL
);