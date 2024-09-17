create table vehicles(
    id   serial primary key,
    year int not null,
    model varchar(100) not null,
    plate varchar(30) not null
);

create table accessory (
  id serial primary key,
  vehicle_id int not null,
  name varchar(50),
  CONSTRAINT accessory_vehicle
      FOREIGN KEY(vehicle_id)
          REFERENCES vehicles(id)
);


