CREATE TABLE moviments (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    date        TEXT NOT NULL,
    month       TEXT,
    description TEXT,
    category    TEXT,
    bank        TEXT,
    method      TEXT,
    value       REAL NOT NULL,
    tags        TEXT,
    content     TEXT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE purchases (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,
    date                TEXT NOT NULL,
    month               TEXT,
    description         TEXT,
    current_installment INTEGER,
    total_installment   INTEGER,
    bank                TEXT,
    number              TEXT,
    category            TEXT,
    value               REAL NOT NULL,
    tags                TEXT,
    content             TEXT,
    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE companies (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,
    name                TEXT,
    category            TEXT,
    tags                TEXT,
    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP
);
