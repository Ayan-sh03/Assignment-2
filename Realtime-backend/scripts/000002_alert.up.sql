create table alerts (
  id integer primary key,
  city_name varchar(255) not null,
  alert_count integer not null,
  alert_date date not null
  );

create index alert_city_name_alert_date_idx on alerts (city_name, alert_date);
