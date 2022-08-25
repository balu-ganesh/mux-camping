insert into CAMPS (id,name, description, location, active) values(100,"pacific new found","new land found", "Pacific Ocean",1);

insert into SITES (id,name, description, camp_id, active) values(500,"north-1", "1 - north",1,1);
insert into SITES (id,name, description, camp_id, active) values(501,"north-2", "1 - north",1,1);

-- one site has multiple slots for people
insert into SLOTS (id,site_id, active) values(1001,500, 1);
insert into SLOTS (id,site_id, active) values(1002,500, 1);
insert into SLOTS (id,site_id, active) values(1003,500, 1);

insert into SLOTS (id,site_id, active) values(2001,501, 1);
insert into SLOTS (id,site_id, active) values(2002,501, 1);
insert into SLOTS (id,site_id, active) values(2003,501, 1);

insert into USERS (id, first_name, last_name, email) values(10001, "Jess", "Sheine", "jess.sheine@gmail.com");
insert into USERS (id, first_name, last_name, email) values(10002, "Dave", "Mcq", "dave.mncq@gmail.com");
insert into USERS (id, first_name, last_name, email) values(10003, "Jesse", "Ezze", "jesse.ezze@gmail.com");
insert into USERS (id, first_name, last_name, email) values(10004, "Sree", "Sreekumar", "sree.sree@gmail.com");
insert into USERS (id, first_name, last_name, email) values(10005, "Tarun", "Nair", "tarun.nair@gmail.com");
insert into USERS (id, first_name, last_name, email) values(10006, "Balu", "Nair", "balu.nair@gmail.com");

insert into BOOKINGS (id, slot_id, user_id, start_date, end_date,  description) VALUES (20001, 1001,10001, 2022-08-24 12:00, '2022-08-25 12:00', "booking for one day");
insert into BOOKINGS (id, slot_id, user_id, start_date, end_date,  description) VALUES (20002, 1002,10002, '2022-08-24 12:00', '2022-08-26 12:00', "booking for two day");
insert into BOOKINGS (id, slot_id, user_id, start_date, end_date,  description) VALUES (20003, 1003,10003, '2022-08-24 12:00', '2022-08-27 12:00', "booking for three day");

insert into BOOKINGS (id, slot_id, user_id, start_date, end_date,  description) VALUES (20004, 1001,10004, '2022-08-25 12:00', '2022-08-26 12:00', "booking for one day from second day");
insert into BOOKINGS (id, slot_id, user_id, start_date, end_date,  description) VALUES (20005, 1002,10005, '2022-08-26 12:00', '2022-08-27 12:00', "booking for one day from second day");

insert into BOOKINGS (id, slot_id, user_id, start_date, end_date,  description) VALUES (20006, 2001,10006, '2022-08-25 12:00', '2022-08-26 12:00', "booking for one day from second day");

