create database mygram;
create extension if not exists "uuid-ossp";
create table users(
    id UUID primary key default uuid_generate_v4(),
    username varchar(255) not null unique, 
    email varchar(255) not null unique,  
    password varchar(255) not null,
    dob date, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp 
);

create table photos(
    id UUID primary key default uuid_generate_v4(),
    title varchar(255) not null, 
    url text not null,  
    caption text,
    user_id  UUID not null, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp, 
    constraint fk_photo_user_id 
        foreign key (user_id) 
        references users(id)
);

create table comments(
    id UUID primary key default uuid_generate_v4(),
    message text not null, 
    user_id  UUID not null, 
    photo_id  UUID not null, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp, 
    constraint fk_comments_photo_id 
        foreign key (photo_id) 
        references photos(id),
    constraint fk_comments_user_id 
        foreign key (user_id) 
        references users(id)
);

create table social_medias(
    id UUID primary key default uuid_generate_v4(),
    name varchar(255) not null, 
    url text not null, 
    user_id  UUID not null, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp, 
    constraint fk_social_medias_user_id 
        foreign key (user_id) 
        references users(id)
);

-- check data types
SELECT column_name, data_type
FROM information_schema.columns
WHERE table_name = 'comments';