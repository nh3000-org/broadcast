create table  days  (
  rowid serial primary key,
  id char(3),
  description text not null,
  dayofweek integer
);
insert into days (id,description,dayofweek) values('MON', 'Monday',1);
insert into days (id,description,dayofweek) values('TUE', 'Tuesday',2);
insert into days (id,description,dayofweek) values('WED', 'Wednesday',3);
insert into days (id,description,dayofweek) values('THU', 'Thursday',4);
insert into days (id,description,dayofweek) values('FRI', 'Friday',5);
insert into days (id,description,dayofweek) values('SAT', 'Saturday',6);
insert into days (id,description,dayofweek) values('SUN', 'Sunday',7);

create index dayindex on days (dayofweek);
create table hours (
  rowid serial primary key,
  id char(2),
  description text not null
);
insert into hours (id,description) values('00', 'Hour Part 00');
insert into hours (id,description) values('01', 'Hour Part 01');
insert into hours (id,description) values('02', 'Hour Part 02');
insert into hours (id,description) values('03', 'Hour Part 03');
insert into hours (id,description) values('04', 'Hour Part 04');
insert into hours (id,description) values('05', 'Hour Part 05');
insert into hours (id,description) values('06', 'Hour Part 06');
insert into hours (id,description) values('07', 'Hour Part 07');
insert into hours (id,description) values('08', 'Hour Part 08');
insert into hours (id,description) values('09', 'Hour Part 09');
insert into hours (id,description) values('10', 'Hour Part 10');
insert into hours (id,description) values('11', 'Hour Part 11');
insert into hours (id,description) values('12', 'Hour Part 12');
insert into hours (id,description) values('13', 'Hour Part 13');
insert into hours (id,description) values('14', 'Hour Part 14');
insert into hours (id,description) values('15', 'Hour Part 15');
insert into hours (id,description) values('16', 'Hour Part 16');
insert into hours (id,description) values('17', 'Hour Part 17');
insert into hours (id,description) values('18', 'Hour Part 18');
insert into hours (id,description) values('19', 'Hour Part 19');
insert into hours (id,description) values('20', 'Hour Part 20');
insert into hours (id,description) values('21', 'Hour Part 21');
insert into hours (id,description) values('22', 'Hour Part 22');
insert into hours (id,description) values('23', 'Hour Part 23');

create index hoursindex on hours (id);
create table categories (
  rowid serial primary key,
  id varchar(64),
  description text not null
);


insert into categories (id,description) values('STATIONID', 'Station ID');
insert into categories (id,description) values('IMAGINGID', 'Imaging ID');
insert into categories (id,description) values('PROMOS', 'Promotions');
insert into categories (id,description)  values('NEXT', 'Play Next');
insert into categories (id,description) values('ADS', 'ADS - Advertising Top Of Hour');
insert into categories (id,description) values('CURRENTS', 'Top 40 Currants');
insert into categories (id,description) values('RECURRENTS', 'Recurrants Library');
insert into categories (id,description) values('NWS-1-PLAYONCE', 'NWS Spots 6 30 Bot AM Play Once');
insert into categories (id,description) values('NWS-2-PLAYONCE', 'NWS Spots 7 00 Bot AM Play Once');
insert into categories (id,description) values('NWS-3-PLAYONCE', 'NWS Spots 7 30 Bot AM Play Once');
insert into categories (id,description) values('NWS-4-PLAYONCE', 'NWS Spots 8 00 Bot AM Play Once');
insert into categories (id,description) values('NWS-5-PLAYONCE', 'NWS Spots 8 30 Bot AM Play Once');

insert into categories (id,description) values('DJAM-71-PLAYONCE', 'DJ Morning Spots 7 Top AM Play Once');
insert into categories (id,description) values('DJAM-72-PLAYONCE', 'DJ Morning Spots 7 Bot AM Play Once');
insert into categories (id,description) values('DJAM-81-PLAYONCE', 'DJ Morning Spots 8 Top AM Play Once');
insert into categories (id,description) values('DJAM-82-PLAYONCE', 'DJ Morning Spots 8 Bot AM Play Once');
insert into categories (id,description) values('DJAM-91-PLAYONCE', 'DJ Morning Spots 9 TopAM Play Once');
insert into categories (id,description) values('DJAM-92-PLAYONCE', 'DJ Morning Spots 9 Bot AM Play Once');
insert into categories (id,description) values('DJAM-101-PLAYONCE', 'DJ Morning Spots 10 Top AM Play Once');
insert into categories (id,description) values('DJAM-102-PLAYONCE', 'DJ Morning Spots 10 Bot AM Play Once');
insert into categories (id,description) values('DJAM-111-PLAYONCE', 'DJ Morning Spots 11 Top AM Play Once');
insert into categories (id,description) values('DJAM-112-PLAYONCE', 'DJ Morning Spots 11 Bot AM Play Once');

insert into categories (id,description) values('DJPM-121-PLAYONCE', 'DJ Afternoon Spots 12 Top AM Play Once');
insert into categories (id,description) values('DJPM-122-PLAYONCE', 'DJ Afternoon Spots 12 Bot AM Play Once');
insert into categories (id,description) values('DJPM-131-PLAYONCE', 'DJ Afternoon Spots 13 Top PM Play Once');
insert into categories (id,description) values('DJPM-132-PLAYONCE', 'DJ Afternoon Spots 13 Bot PM Play Once');

insert into categories (id,description) values('DJPM-141-PLAYONCE', 'DJ Afternoon Spots 14 Top PM Play Once');
insert into categories (id,description) values('DJPM-142-PLAYONCE', 'DJ Afternoon Spots 14 Bot PM Play Once');

insert into categories (id,description) values('DJPM-151-PLAYONCE', 'DJ Afternoon Spots 15 Top PM Play Once');
insert into categories (id,description) values('DJPM-152-PLAYONCE', 'DJ Afternoon Spots 15 Bot PM Play Once');

insert into categories (id,description) values('DJPM-161-PLAYONCE', 'DJ Afternoon Spots 16 Top PM Play Once');
insert into categories (id,description) values('DJPM-162-PLAYONCE', 'DJ Afternoon Spots 16 Bot PM Play Once');

insert into categories (id,description) values('DJPM-141-PLAYONCE', 'DJ Afternoon Spots 17 Top PM Play Once');
insert into categories (id,description) values('DJPM-172-PLAYONCE', 'DJ Afternoon Spots 17 Bot PM Play Once');

insert into categories (id,description) values('DJPM-181-PLAYONCE', 'DJ Afternoon Spots 18 Top PM Play Once');
insert into categories (id,description) values('DJPM-182-PLAYONCE', 'DJ Afternoon Spots 18 Bot PM Play Once');


insert into categories (id,description) values('FILLTOTOH', 'Fill To TOH Schedule');

create index categoriesindex on categories (id);

create table schedule (
  rowid serial primary key,
  days varchar(3), 
  hours char(2), 
  position char(2),
  categories varchar(64), 
  spinstoplay integer
);


insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','04','ADS',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','05','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','09','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','11','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','12','NWS-1-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','13','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','14','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','15','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','04','DJAM-71-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','07','NWS-2-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','08','RECURRENTS',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','10','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','11','DJAM-72-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','12','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','13','RECURRENTS',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','14','NWS-3-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','04','DJAM-81-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','07','NWS-4-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','10','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','11','DJAM-82-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','12','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','13','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','14','NWS-5-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','15','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','04','DJAM-91-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','10','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','11','DJAM-92-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','13','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','14','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','15','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','16','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','04','DJAM-101-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','07','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','08','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','09','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','10','DJAM-102-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','11','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','12','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','14','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','14','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','15','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','16','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','04','DJAM-111-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','10','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','11','DJAM-112-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','04','DJPM-121-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','08','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','11','DJPM-122-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','04','DJPM-121-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','10','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','11','DJPM-122-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','17','FILLTOTOH',1);


insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','04','DJPM-141-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','10','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','11','DJPM-142-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','04','DJPM-151-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','10','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','11','DJPM-152-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','12','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','04','DJPM-161-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','10','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','11','DJPM-162-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','04','DJPM-171-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','08','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','10','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','11','DJPM-172-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','03','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','04','DJPM-181-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','05','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','06','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','07','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','08','RECURRENTS',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','10','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','11','DJPM-182-PLAYONCE',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','14','ADS',4);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','15','CURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','16','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','17','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','22','FILLTOTOH',1);

insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','02','PROMOS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','03','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','04','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','05','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','06','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','07','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','08','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','09','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','10','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','11','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','12','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','13','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','14','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','15','RECURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','16','CURRENTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','17','ADS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','18','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','19','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','20','RECURRENTS',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','21','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','22','FILLTOTOH',1);

create index scheduleindex on schedule (days,hours,position);
create table inventory (
  rowid serial primary key,
  category varchar(64) not null,
  artist text not null,
  song   text not null,
  album  text not null,
  songlength integer,
  rndorder  text,
  startson  text,
  expireson text,
  adstimeslots text array[23],
  adsdayslots text array[7],
  adsmaxspins int,
  adsmaxspinsperhour int,
  lastplayed text,
  dateadded text,
  spinstoday integer,
  spinsweek  integer,
  spinstotal integer,
  sourcelink text
);
create index inventorybyartist on inventory (artist,song);
create index inventorybycategorysong on inventory (category,song);
create index inventoryplayget on inventory (category,lastplayed,rndorder);
create table traffic (
  rowid serial primary key,
  category varchar(64) not null,
  artist text not null,
  song   text not null,
  album  text,
  playedon text
);
create index trafficbyartist on traffic (artist,song,album);
create index trafficbycategory on traffic (category,album);
