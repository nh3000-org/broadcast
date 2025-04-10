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
  id varchar(32),
  description text not null
);
insert into categories (id,description) values('VIDEOSTATIONID', 'VIDEO Station ID');
insert into categories (id,description) values('VIDEOIMAGINGID', 'VIDEO Imaging ID');
insert into categories (id,description)  values('VIDEONEXT', 'VIDEO Play Next');
insert into categories (id,description) values('VIDEOADDS', 'VIDEO  Advertising');
insert into categories (id,description) values('VIDEO', 'VIDEO Library');


insert into categories (id,description) values('STATIONID', 'Station ID');
insert into categories (id,description) values('IMAGINGID', 'Imaging ID');
insert into categories (id,description)  values('NEXT', 'Play Next');
insert into categories (id,description) values('LIVE', 'Live');
insert into categories (id,description) values('ADDSTOH', 'ADDS - Advertising Top Of Hour');
insert into categories (id,description) values('ADDSDRIVETIME', 'ADDS - Advertising Drive Time');
insert into categories (id,description) values('ADDS', 'ADDS - Advertising');
insert into categories (id,description) values('TOP40', 'Top 40 MUSIC');
insert into categories (id,description) values('ROOTS', 'Roots MUSIC');
insert into categories (id,description) values('MUSIC', 'Music Library');
insert into categories (id,description) values('FILLTOTOH', 'Fill To TOH Schedule');
insert into categories (id,description) values('NWS', 'News Weather Sports');
create index categoriesindex on categories (id);

create table schedule (
  rowid serial primary key,
  days varchar(3), 
  hours char(2), 
  position char(2),
  categories varchar(32), 
  spinstoplay integer
);

insert into schedule (days,hours,position,categories,spinstoplay) values('VID', '00','01','VIDEOSTATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('VID', '00','02','VIDEOIMAGINGID',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('VID', '00','03','VIDEOADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('VID', '00','04','VIDEO',1);


insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '00','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '01','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '02','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '03','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '04','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '05','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '06','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '07','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '08','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '09','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','09','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '10','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '11','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '12','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '13','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '14','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '15','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '16','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '17','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','02','ADDSDRIVETIME',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '18','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '19','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '20','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '21','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '22','15','FILLTOTOH',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','01','STATIONID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','02','ADDSTOH',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','03','NWS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','04','TOP40',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','05','MUSIC',2);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','06','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','07','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','08','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','09','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','10','IMAGINGID',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','11','TOP40',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','12','ROOTS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','13','ADDS',1);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','14','MUSIC',3);
insert into schedule (days,hours,position,categories,spinstoplay) values('MON', '23','15','FILLTOTOH',1);
create index scheduleindex on schedule (days,hours,position);
create table inventory (
  rowid serial primary key,
  category varchar(32) not null,
  artist text not null,
  song   text not null,
  album  text,
  songlength integer,
  rndorder  text,
  startson  text,
  expireson text,
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
  artist text not null,
  song   text not null,
  album  text,
  playedon text
);
create index trafficbyartist on traffic (artist,song,album);
