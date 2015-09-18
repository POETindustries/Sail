package dbase

const createWidget = `create table if not exists sl_widget(
    id serial primary key not null,
    type varchar(31) not null,
    type_id integer not null);`

const createMenu = `create table if not exists sl_widget_menu(
    id serial primary key not null,
    name varchar(31) not null,
    entry_ids text not null default '1');`

const addWidgetColumn = `do $$ begin if not exists (
    select column_name
    from information_schema.columns
    where table_name='sl_domain' and column_name='plugin_sail_widgets' )
    then
    alter table sl_domain
    add column plugin_sail_widgets text default '';
    end if; end $$`

var schema = []string{createWidget, createMenu, addWidgetColumn}
