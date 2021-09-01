create table "users" (
    username varchar primary key,
    hashed_password varchar not null,
    full_name varchar not null,
    email varchar unique not null,
    password_changed_at  timestamptz not null default '0001-01-01 00:00:00+00Z',
    "create_at" timestamptz not null default now()
);

alter table "users" add foreign key ("username") references accounts (owner);
create unique index on "accounts" ("owner", "currency");

create table "accounts" (
  "id" bigserial primary key,
  "owner" varchar not null,
  "balance" bigint not null,
  "currency" varchar not null,
  "create_at" timestamptz not null default now()
);

create table "entries" (
  "id" bigserial primary key,
  "account_id" bigint not null,
  "amount" bigint not null,
  "create_at" timestamptz not null default now()
);

create table "transfers"(
  "id" bigserial primary key,
  "from_account_id" bigint not null,
  "to_account_id" bigint not null,
  "amount" bigint not null,
  "create_at" timestamptz not null default now()
);

alter table "entries" add foreign key ("account_id") references accounts (id);
alter table "transfers" add foreign key ("from_account_id") references accounts (id);
alter table "transfers" add foreign key ("to_account_id") references accounts (id);