package dbase

const createMenu = `create table if not exists sl_menu(
    id serial primary key not null,
    name varchar(31) not null,
    entry_ids text not null default '1');`

var schema = []string{createMenu}
