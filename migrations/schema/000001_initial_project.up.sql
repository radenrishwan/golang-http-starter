CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists articles (
  id uuid primary key default uuid_generate_v4(),
  title varchar(255) not null,
  slug varchar(255) not null unique,
  body text not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);