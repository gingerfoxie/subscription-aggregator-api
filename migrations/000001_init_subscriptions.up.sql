-- migrations/000001_init_subscriptions.up.sql
CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date VARCHAR(7) NOT NULL,
    end_date VARCHAR(7),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_service_name ON subscriptions(service_name);