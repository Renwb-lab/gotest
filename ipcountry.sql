create table IpCountryTab (
    id bigint not null primary key auto_increment,
    ip unsigned int(20) not null default 0, 
    country char(255) not null default '',
    ctime unsigned int(20) not null default 0,
    primary key (id),
    unique index ip_country_idx (ip, country)
) engine=innoDB;
