create table if not exists articles (
    article_id integer unsigned auto_increment primary key,
    title varchar(100) not null,
    contents text not null,
    username varchar(100) not null,
    nice integer not null default 0,
    created_at datetime not null default current_timestamp
);

create table if not exists comments (
    comment_id integer unsigned auto_increment primary key,
    article_id integer unsigned not null,
    message text not null,
    created_at datetime,
    foreign key (article_id) references articles(article_id)
);