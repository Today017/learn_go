insert into articles (title, contents, username, nice, created_at) values
    ('firstPost', 'This is my first blog', 'soma', 2, now());

insert into articles (title, contents, username, nice) values
    ('2nd', 'Second blog post', 'soma', 4);

insert into comments (article_id, message, created_at) values
    (1, '1st commnet yeah', now());

insert into comments (article_id, message) values
    (1, 'welcome');
