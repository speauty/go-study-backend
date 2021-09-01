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