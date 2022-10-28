create table products
(
    id                  serial primary key,
    title               varchar(100)                               not null,
    content             text                                       not null,
    count               int       default 0                        not null,
    cost                int check ( cost > 0)                      not null,
    date_of_publication timestamp default now()                    not null,
    rating              int check ( rating >= 0 and rating <= 10 ) not null
);

create table users
(
    id            serial primary key,
    first_name    varchar(150)                                             not null,
    second_name   varchar(150)                                             not null,
    username      varchar(150) check ( char_length(username) >= 5 ) unique not null,
    sex           varchar(10) check ( sex = 'Male' or sex = 'Female' )     not null,
    password_hash varchar(150)                                             not null
);

create table orders
(
    id                  serial primary key,
    status              varchar(50) default 'Created'              not null,
    comment             varchar(250)                               not null,
    date_of_publication timestamp   default now()                  not null,
    date_of_completion  timestamp,
    user_id             int references users (id) on delete cascade not null
);

create table reviews
(
    id                  serial primary key,
    content             text                                          not null,
    rating              int check ( rating >= 0 and rating <= 10 )    not null,
    date_of_publication timestamp default now()                       not null,
    user_id             int references users (id) on delete cascade    not null,
    product_id          int references products (id) on delete cascade not null
);

create table likes_product
(
    id                  serial primary key,
    mark                int check ( mark = -1 or mark = 1 )           not null,
    date_of_publication timestamp default now()                       not null,
    product_id          int references products (id) on delete cascade not null,
    user_id             int references users (id) on delete cascade    not null
);

create table likes_review
(
    id                  serial primary key,
    mark                int check ( mark = -1 or mark = 1 )          not null,
    date_of_publication timestamp default now()                      not null,
    review_id           int references reviews (id) on delete cascade not null,
    user_id             int references users (id) on delete cascade   not null
);

create table categories
(
    id   serial primary key,
    name varchar(100) unique not null
);

create table products_categories
(
    product_id  int references products (id) on delete cascade  not null,
    category_id int references categories (id) on delete cascade not null,
    primary key (product_id, category_id)
);

create table orders_products
(
    count      int default 1 check ( count >= 0 )            not null,
    order_id   int references orders (id) on delete cascade   not null,
    product_id int references products (id) on delete cascade not null,
    primary key (order_id, product_id)
);