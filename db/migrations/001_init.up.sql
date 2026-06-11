-- ENUMS
CREATE TYPE order_creator_type AS ENUM ('user', 'system');
CREATE TYPE order_type AS ENUM ('market', 'limit');
CREATE TYPE spot_order_side AS ENUM ('buy', 'sell');
CREATE TYPE perp_order_side AS ENUM ('long', 'short');
CREATE TYPE order_status AS ENUM ('created', 'partially_filled', 'filled', 'canceled');
CREATE TYPE transaction_side AS ENUM ('debit', 'credit');
CREATE TYPE transaction_type AS ENUM ('money_added', 'money-withdrawn', 'affiliate_bonus', 'spot_sell', 'spot_buy');

-- User Deatils
CREATE TABLE IF NOT EXISTS users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    password_hash TEXT NOT NULL,
    affiliate_code TEXT NOT NULL,
    referrer_id INT,
    available_balance BIGINT NOT NULL DEFAULT 0,
    locked_balance BIGINT NOT NULL DEFAULT 0,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


-- Markets (Spot and Perps will have seperate entry)
CREATE TABLE IF NOT EXISTS market (
    id  INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    symbol TEXT NOT NULL UNIQUE,
    UNIQUE(name, symbol),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()

);

-- User Assets
CREATE TABLE IF NOT EXISTS asset (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT  NOT NULL,
    market_id INT NOT NULL,
    available_asset BIGINT NOT NULL DEFAULT 0,
    locked_asset BIGINT NOT NULL DEFAULT 0,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, market_id), -- Single user can have only one record for the market
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_market
        FOREIGN KEY (market_id)
        REFERENCES market(id)
        ON DELETE CASCADE
);

-- User Positions(Will be implemented later for PERPS)


-- Spot Order of All the Markets
CREATE TABLE IF NOT EXISTS spot_order (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    market_id INT NOT NULL,
    user_id INT NOT NULL,
    type order_type NOT NULL,
    side spot_order_side NOT NULL,
    locked_balance BIGINT NOT NULL,
    quantityStep BIGINT NOT NULL,
    quantityStepFilled BIGINT NOT NULL,
    priceTick BIGINT NOT NULL,
    status order_status NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_market
        FOREIGN KEY (market_id)
        REFERENCES market(id),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id),

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Spot Fills
CREATE TABLE IF NOT EXISTS spot_fill (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    market_id INT NOT NULL,
    maker_id INT NOT NULL,
    maker_order_id UUID NOT NULL,
    taker_id  INT NOT NULL,
    taker_order_id UUID NOT NULL,
    priceTick  BIGINT NOT NULL,
    stepSize BIGINT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_maker
        FOREIGN KEY (maker_id)
        REFERENCES users(id),
    CONSTRAINT fk_maker_order
        FOREIGN KEY (maker_order_id)
        REFERENCES spot_order(id),
    CONSTRAINT fk_taker
        FOREIGN KEY (taker_id)
        REFERENCES users(id),
    CONSTRAINT fk_taker_order
        FOREIGN KEY (taker_order_id)
        REFERENCES spot_order(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Perp Order for All the Markets(Will be implemented later for PERPS)
-- CREATE TABLE IF NOT EXISTS perp_order (
--     id UUID PRIMARY KEY DEFAULT uuidv7(),

--     is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
--     created_at TIMESTAMP DEFAULT NOW(),
--     updated_at TIMESTAMP DEFAULT NOW()
-- );

-- Perp Fills(Will be implemented later for PERPS)


-- Account Transactions
CREATE TABLE IF NOT EXISTS account_transaction (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    amount BIGINT NOT NULL,
    user_id INT NOT NULL,
    side transaction_side NOT NULL,
    type transaction_type NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);
