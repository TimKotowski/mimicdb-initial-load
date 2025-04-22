#!/usr/bin/env bash
set -xeuo pipefail

export PGPASSWORD=12345

psql -d postgres_source -U postgres_source -h localhost << EOF
CREATE DATABASE postgres_target;
\c postgres_target

CREATE TABLE IF NOT EXISTS test (
    id SERIAL PRIMARY KEY,
    name varchar,
    create_at timestamp with time zone
);

\c postgres_source
\q
EOF
