-- migrate:up
create table cargo (
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    created_by int not null,
    updated_at timestamp not null default CURRENT_TIMESTAMP,
    updated_by int not null
);

-- migrate:down
drop table cargo;
