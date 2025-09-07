--
-- PostgreSQL database dump
--

-- Dumped from database version 17.1 (Ubuntu 17.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 17.1 (Ubuntu 17.1-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    rowid integer NOT NULL,
    id character varying(64),
    description text NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_rowid_seq OWNER TO postgres;

--
-- Name: categories_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_rowid_seq OWNED BY public.categories.rowid;


--
-- Name: days; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.days (
    rowid integer NOT NULL,
    id character(3),
    description text NOT NULL,
    dayofweek integer
);


ALTER TABLE public.days OWNER TO postgres;

--
-- Name: days_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.days_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.days_rowid_seq OWNER TO postgres;

--
-- Name: days_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.days_rowid_seq OWNED BY public.days.rowid;


--
-- Name: hours; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.hours (
    rowid integer NOT NULL,
    id character(2),
    description text NOT NULL
);


ALTER TABLE public.hours OWNER TO postgres;

--
-- Name: hours_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.hours_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.hours_rowid_seq OWNER TO postgres;

--
-- Name: hours_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.hours_rowid_seq OWNED BY public.hours.rowid;


--
-- Name: inventory; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inventory (
    rowid integer NOT NULL,
    category character varying(64) NOT NULL,
    artist text NOT NULL,
    song text NOT NULL,
    album text NOT NULL,
    songlength integer,
    rndorder text,
    startson text,
    expireson text,
    adstimeslots text[],
    adsdayslots text[],
    adsmaxspins integer,
    adsmaxspinsperhour integer,
    lastplayed text,
    dateadded text,
    spinstoday integer,
    spinsweek integer,
    spinstotal integer,
    sourcelink text
);


ALTER TABLE public.inventory OWNER TO postgres;

--
-- Name: inventory_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.inventory_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.inventory_rowid_seq OWNER TO postgres;

--
-- Name: inventory_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.inventory_rowid_seq OWNED BY public.inventory.rowid;


--
-- Name: schedule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedule (
    rowid integer NOT NULL,
    days character varying(3),
    hours character(2),
    "position" character(2),
    categories character varying(64),
    spinstoplay integer
);


ALTER TABLE public.schedule OWNER TO postgres;

--
-- Name: schedule_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedule_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.schedule_rowid_seq OWNER TO postgres;

--
-- Name: schedule_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedule_rowid_seq OWNED BY public.schedule.rowid;


--
-- Name: traffic; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.traffic (
    rowid integer NOT NULL,
    artist text NOT NULL,
    song text NOT NULL,
    album text,
    playedon text
);


ALTER TABLE public.traffic OWNER TO postgres;

--
-- Name: traffic_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.traffic_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.traffic_rowid_seq OWNER TO postgres;

--
-- Name: traffic_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.traffic_rowid_seq OWNED BY public.traffic.rowid;


--
-- Name: categories rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN rowid SET DEFAULT nextval('public.categories_rowid_seq'::regclass);


--
-- Name: days rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.days ALTER COLUMN rowid SET DEFAULT nextval('public.days_rowid_seq'::regclass);


--
-- Name: hours rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hours ALTER COLUMN rowid SET DEFAULT nextval('public.hours_rowid_seq'::regclass);


--
-- Name: inventory rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory ALTER COLUMN rowid SET DEFAULT nextval('public.inventory_rowid_seq'::regclass);


--
-- Name: schedule rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule ALTER COLUMN rowid SET DEFAULT nextval('public.schedule_rowid_seq'::regclass);


--
-- Name: traffic rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.traffic ALTER COLUMN rowid SET DEFAULT nextval('public.traffic_rowid_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (rowid, id, description) FROM stdin;
1	STATIONID	Station ID
2	IMAGINGID	Imaging ID
3	PROMOS	Promotions
4	NEXT	Play Next
5	ADS	ADS - Advertising Top Of Hour
6	CURRENTS	Top 40 Currants
7	RECURRENTS	Recurrants Library
8	NWS-1-PLAYONCE	NWS Spots 6 30 Bot AM Play Once
9	NWS-2-PLAYONCE	NWS Spots 7 00 Bot AM Play Once
10	NWS-3-PLAYONCE	NWS Spots 7 30 Bot AM Play Once
11	NWS-4-PLAYONCE	NWS Spots 8 00 Bot AM Play Once
12	NWS-5-PLAYONCE	NWS Spots 8 30 Bot AM Play Once
13	DJAM-71-PLAYONCE	DJ Morning Spots 7 Top AM Play Once
14	DJAM-72-PLAYONCE	DJ Morning Spots 7 Bot AM Play Once
15	DJAM-81-PLAYONCE	DJ Morning Spots 8 Top AM Play Once
16	DJAM-82-PLAYONCE	DJ Morning Spots 8 Bot AM Play Once
17	DJAM-91-PLAYONCE	DJ Morning Spots 9 TopAM Play Once
18	DJAM-92-PLAYONCE	DJ Morning Spots 9 Bot AM Play Once
19	DJAM-101-PLAYONCE	DJ Morning Spots 10 Top AM Play Once
20	DJAM-102-PLAYONCE	DJ Morning Spots 10 Bot AM Play Once
21	DJAM-111-PLAYONCE	DJ Morning Spots 11 Top AM Play Once
22	DJAM-112-PLAYONCE	DJ Morning Spots 11 Bot AM Play Once
23	DJPM-121-PLAYONCE	DJ Afternoon Spots 12 Top AM Play Once
24	DJPM-122-PLAYONCE	DJ Afternoon Spots 12 Bot AM Play Once
25	DJPM-131-PLAYONCE	DJ Afternoon Spots 13 Top PM Play Once
26	DJPM-132-PLAYONCE	DJ Afternoon Spots 13 Bot PM Play Once
27	DJPM-141-PLAYONCE	DJ Afternoon Spots 14 Top PM Play Once
28	DJPM-142-PLAYONCE	DJ Afternoon Spots 14 Bot PM Play Once
29	DJPM-151-PLAYONCE	DJ Afternoon Spots 15 Top PM Play Once
30	DJPM-152-PLAYONCE	DJ Afternoon Spots 15 Bot PM Play Once
31	DJPM-161-PLAYONCE	DJ Afternoon Spots 16 Top PM Play Once
32	DJPM-162-PLAYONCE	DJ Afternoon Spots 16 Bot PM Play Once
33	DJPM-141-PLAYONCE	DJ Afternoon Spots 17 Top PM Play Once
34	DJPM-172-PLAYONCE	DJ Afternoon Spots 17 Bot PM Play Once
35	DJPM-181-PLAYONCE	DJ Afternoon Spots 18 Top PM Play Once
36	DJPM-182-PLAYONCE	DJ Afternoon Spots 18 Bot PM Play Once
37	FILLTOTOH	Fill To TOH Schedule
38	NWS	News Weather Sports
\.


--
-- Data for Name: days; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.days (rowid, id, description, dayofweek) FROM stdin;
1	MON	Monday	1
2	TUE	Tuesday	2
3	WED	Wednesday	3
4	THU	Thursday	4
5	FRI	Friday	5
6	SAT	Saturday	6
7	SUN	Sunday	7
\.


--
-- Data for Name: hours; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.hours (rowid, id, description) FROM stdin;
1	00	Hour Part 00
2	01	Hour Part 01
3	02	Hour Part 02
4	03	Hour Part 03
5	04	Hour Part 04
6	05	Hour Part 05
7	06	Hour Part 06
8	07	Hour Part 07
9	08	Hour Part 08
10	09	Hour Part 09
11	10	Hour Part 10
12	11	Hour Part 11
13	12	Hour Part 12
14	13	Hour Part 13
15	14	Hour Part 14
16	15	Hour Part 15
17	16	Hour Part 16
18	17	Hour Part 17
19	18	Hour Part 18
20	19	Hour Part 19
21	20	Hour Part 20
22	21	Hour Part 21
23	22	Hour Part 22
24	23	Hour Part 23
\.


--
-- Data for Name: inventory; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.inventory (rowid, category, artist, song, album, songlength, rndorder, startson, expireson, adstimeslots, adsdayslots, adsmaxspins, adsmaxspinsperhour, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink) FROM stdin;
\.


--
-- Data for Name: schedule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedule (rowid, days, hours, "position", categories, spinstoplay) FROM stdin;
1	MON	00	01	STATIONID	1
2	MON	00	02	PROMOS	1
3	MON	00	03	ADS	1
4	MON	00	04	CURRENTS	3
5	MON	00	05	IMAGINGID	1
6	MON	00	06	RECURRENTS	4
7	MON	00	07	IMAGINGID	1
8	MON	00	08	CURRENTS	3
9	MON	00	09	ADS	2
10	MON	00	10	RECURRENTS	4
11	MON	00	11	FILLTOTOH	1
12	MON	01	01	STATIONID	1
13	MON	01	02	PROMOS	1
14	MON	00	03	ADS	2
15	MON	01	04	CURRENTS	3
16	MON	01	05	IMAGINGID	1
17	MON	01	06	RECURRENTS	4
18	MON	01	07	IMAGINGID	1
19	MON	01	08	ADS	2
20	MON	01	09	CURRENTS	3
21	MON	01	10	RECURRENTS	4
22	MON	01	11	FILLTOTOH	1
23	MON	02	01	STATIONID	1
24	MON	02	02	PROMOS	1
25	MON	00	03	ADS	2
26	MON	02	04	CURRENTS	3
27	MON	02	05	IMAGINGID	1
28	MON	02	06	RECURRENTS	4
29	MON	02	07	IMAGINGID	1
30	MON	02	08	ADS	2
31	MON	02	09	CURRENTS	3
32	MON	02	10	RECURRENTS	4
33	MON	02	11	FILLTOTOH	1
34	MON	03	01	STATIONID	1
35	MON	03	02	PROMOS	1
36	MON	00	03	ADS	2
37	MON	03	04	CURRENTS	3
38	MON	03	05	IMAGINGID	1
39	MON	03	06	RECURRENTS	4
40	MON	03	07	IMAGINGID	1
41	MON	03	08	ADS	2
42	MON	03	09	CURRENTS	3
43	MON	03	10	RECURRENTS	4
44	MON	03	11	FILLTOTOH	1
45	MON	04	01	STATIONID	1
46	MON	04	02	PROMOS	1
47	MON	00	03	ADS	2
48	MON	04	04	CURRENTS	3
49	MON	04	05	IMAGINGID	1
50	MON	04	06	RECURRENTS	4
51	MON	04	07	IMAGINGID	1
52	MON	04	08	ADS	2
53	MON	04	09	CURRENTS	3
54	MON	04	10	RECURRENTS	4
55	MON	04	11	FILLTOTOH	1
56	MON	05	01	STATIONID	1
57	MON	05	02	PROMOS	1
58	MON	00	03	ADS	2
59	MON	05	04	CURRENTS	3
60	MON	05	05	IMAGINGID	1
61	MON	05	06	RECURRENTS	4
62	MON	05	07	IMAGINGID	1
63	MON	05	08	ADS	2
64	MON	05	09	CURRENTS	3
65	MON	05	10	RECURRENTS	4
66	MON	05	11	FILLTOTOH	1
67	MON	06	01	STATIONID	1
68	MON	06	02	PROMOS	1
69	MON	06	03	CURRENTS	3
70	MON	06	04	ADS	5
71	MON	06	05	IMAGINGID	1
72	MON	06	06	RECURRENTS	4
73	MON	06	07	IMAGINGID	1
74	MON	06	08	ADS	5
75	MON	06	09	NWS-1-PLAYONCE	1
76	MON	06	10	CURRENTS	3
77	MON	06	11	FILLTOTOH	1
78	MON	07	01	STATIONID	1
79	MON	07	02	PROMOS	1
80	MON	07	03	CURRENTS	3
81	MON	07	04	DJAM-71-PLAYONCE	1
82	MON	07	05	ADS	5
83	MON	07	06	NWS-2-PLAYONCE	1
84	MON	07	07	RECURRENTS	4
85	MON	07	08	DJAM-72-PLAYONCE	1
86	MON	07	09	ADS	5
87	MON	07	10	NWS-3-PLAYONCE	1
88	MON	07	11	CURRENTS	3
89	MON	07	12	FILLTOTOH	1
90	MON	08	01	STATIONID	1
91	MON	08	02	PROMOS	1
92	MON	08	03	CURRENTS	3
93	MON	08	04	DJAM-81-PLAYONCE	1
94	MON	08	05	ADS	5
95	MON	08	06	NWS-4-PLAYONCE	1
96	MON	08	07	RECURRENTS	4
97	MON	08	08	DJAM-82-PLAYONCE	1
98	MON	08	09	ADS	5
99	MON	08	10	NWS-5-PLAYONCE	1
100	MON	08	11	CURRENTS	3
101	MON	08	12	FILLTOTOH	1
102	MON	09	01	STATIONID	1
103	MON	09	02	PROMOS	1
104	MON	09	03	CURRENTS	3
105	MON	09	04	DJAM-91-PLAYONCE	1
106	MON	09	05	ADS	5
107	MON	09	06	CURRENTS	3
108	MON	09	07	DJAM-92-PLAYONCE	1
109	MON	09	08	RECURRENTS	4
110	MON	09	09	ADS	5
111	MON	09	10	CURRENTS	3
112	MON	09	11	FILLTOTOH	1
113	MON	10	01	STATIONID	1
114	MON	10	02	PROMOS	1
115	MON	10	03	CURRENTS	3
116	MON	10	04	DJAM-101-PLAYONCE	1
117	MON	10	05	ADS	5
118	MON	10	06	CURRENTS	3
119	MON	10	07	DJAM-102-PLAYONCE	1
120	MON	10	08	RECURRENTS	4
121	MON	10	09	ADS	5
122	MON	10	10	CURRENTS	3
123	MON	10	11	FILLTOTOH	1
124	MON	11	01	STATIONID	1
125	MON	11	02	PROMOS	1
126	MON	11	03	CURRENTS	3
127	MON	11	04	DJAM-111-PLAYONCE	1
128	MON	11	05	ADS	5
129	MON	11	06	CURRENTS	3
130	MON	11	07	DJAM-112-PLAYONCE	1
131	MON	12	08	RECURRENTS	4
132	MON	11	09	ADS	5
133	MON	11	10	CURRENTS	3
134	MON	11	11	FILLTOTOH	1
135	MON	12	01	STATIONID	1
136	MON	12	02	PROMOS	1
137	MON	12	03	CURRENTS	3
138	MON	12	04	DJPM-121-PLAYONCE	1
139	MON	12	05	ADS	5
140	MON	12	06	CURRENTS	3
141	MON	12	07	DJPMH-122-PLAYONCE	1
142	MON	12	08	RECURRENTS	4
143	MON	12	09	ADS	5
144	MON	12	10	CURRENTS	3
145	MON	12	11	FILLTOTOH	1
146	MON	13	01	STATIONID	1
147	MON	13	02	PROMOS	1
148	MON	13	03	CURRENTS	3
149	MON	13	04	DJPM-121-PLAYONCE	1
150	MON	13	05	ADS	5
151	MON	13	06	CURRENTS	3
152	MON	13	07	DJPMH-122-PLAYONCE	1
153	MON	13	08	RECURRENTS	4
154	MON	13	09	ADS	5
155	MON	13	10	CURRENTS	3
156	MON	13	11	FILLTOTOH	1
157	MON	14	01	STATIONID	1
158	MON	14	02	PROMOS	1
159	MON	14	03	CURRENTS	3
160	MON	14	04	DJPM-141-PLAYONCE	1
161	MON	14	05	ADS	5
162	MON	14	06	CURRENTS	3
163	MON	14	07	DJPMH-142-PLAYONCE	1
164	MON	14	08	RECURRENTS	4
165	MON	14	09	ADS	5
166	MON	14	10	CURRENTS	3
167	MON	14	11	FILLTOTOH	1
168	MON	15	01	STATIONID	1
169	MON	15	02	PROMOS	1
170	MON	15	03	CURRENTS	3
171	MON	15	04	DJPM-151-PLAYONCE	1
172	MON	15	05	ADS	5
173	MON	15	06	CURRENTS	3
174	MON	15	07	DJPMH-152-PLAYONCE	1
175	MON	15	08	RECURRENTS	4
176	MON	15	09	ADS	5
177	MON	15	10	CURRENTS	3
178	MON	15	11	FILLTOTOH	1
179	MON	16	01	STATIONID	1
180	MON	16	02	PROMOS	1
181	MON	16	03	CURRENTS	3
182	MON	16	04	DJPM-161-PLAYONCE	1
183	MON	16	05	ADS	5
184	MON	16	06	CURRENTS	3
185	MON	16	07	DJPMH-162-PLAYONCE	1
186	MON	16	08	RECURRENTS	4
187	MON	16	09	ADS	5
188	MON	16	10	CURRENTS	3
189	MON	16	11	FILLTOTOH	1
190	MON	17	01	STATIONID	1
191	MON	17	02	PROMOS	1
192	MON	17	03	CURRENTS	3
193	MON	17	04	DJPM-171-PLAYONCE	1
194	MON	17	05	ADS	5
195	MON	17	06	CURRENTS	3
196	MON	17	07	DJPMH-172-PLAYONCE	1
197	MON	17	08	RECURRENTS	4
198	MON	17	09	ADS	5
199	MON	17	10	CURRENTS	3
200	MON	17	11	FILLTOTOH	1
201	MON	18	01	STATIONID	1
202	MON	18	02	PROMOS	1
203	MON	18	03	CURRENTS	3
204	MON	18	04	DJPM-181-PLAYONCE	1
205	MON	18	05	ADS	5
206	MON	18	06	CURRENTS	3
207	MON	18	07	DJPMH-182-PLAYONCE	1
208	MON	18	08	RECURRENTS	4
209	MON	18	09	ADS	5
210	MON	18	10	CURRENTS	3
211	MON	18	11	FILLTOTOH	1
212	MON	19	01	STATIONID	1
213	MON	19	02	PROMOS	1
214	MON	19	03	CURRENTS	3
215	MON	19	04	IMAGINGID	1
216	MON	19	05	ADS	2
217	MON	19	06	RECURRENTS	4
218	MON	19	07	IMAGINGID	1
219	MON	19	08	CURRENTS	3
220	MON	19	09	RECURRENTS	3
221	MON	19	10	IMAGINGID	1
222	MON	19	11	ADS	2
223	MON	19	12	FILLTOTOH	1
224	MON	20	01	STATIONID	1
225	MON	20	02	PROMOS	1
226	MON	20	03	CURRENTS	3
227	MON	20	04	IMAGINGID	1
228	MON	20	05	ADS	2
229	MON	20	06	RECURRENTS	4
230	MON	20	07	IMAGINGID	1
231	MON	20	08	CURRENTS	3
232	MON	20	09	RECURRENTS	3
233	MON	20	10	IMAGINGID	1
234	MON	20	11	ADS	2
235	MON	20	12	FILLTOTOH	1
236	MON	21	01	STATIONID	1
237	MON	21	02	PROMOS	1
238	MON	21	03	CURRENTS	4
239	MON	21	04	IMAGINGID	1
240	MON	21	05	ADS	2
241	MON	21	06	RECURRENTS	4
242	MON	21	07	IMAGINGID	1
243	MON	21	08	CURRENTS	3
244	MON	21	09	RECURRENTS	3
245	MON	21	10	IMAGINGID	1
246	MON	21	11	ADS	2
247	MON	21	12	FILLTOTOH	1
248	MON	22	01	STATIONID	1
249	MON	22	02	PROMOS	1
250	MON	22	03	CURRENTS	4
251	MON	22	04	IMAGINGID	1
252	MON	22	05	ADS	2
253	MON	22	06	RECURRENTS	4
254	MON	22	07	IMAGINGID	1
255	MON	22	08	CURRENTS	3
256	MON	22	09	RECURRENTS	3
257	MON	22	10	IMAGINGID	1
258	MON	22	11	ADS	2
259	MON	22	12	FILLTOTOH	1
260	MON	23	01	STATIONID	1
261	MON	23	02	PROMOS	1
262	MON	23	03	CURRENTS	4
263	MON	23	04	IMAGINGID	1
264	MON	23	05	ADS	2
265	MON	23	06	RECURRENTS	4
266	MON	23	07	IMAGINGID	1
267	MON	23	08	CURRENTS	3
268	MON	23	09	RECURRENTS	3
269	MON	23	10	IMAGINGID	1
270	MON	23	11	ADS	2
271	MON	23	12	FILLTOTOH	1
\.


--
-- Data for Name: traffic; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.traffic (rowid, artist, song, album, playedon) FROM stdin;
\.


--
-- Name: categories_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_rowid_seq', 38, true);


--
-- Name: days_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.days_rowid_seq', 7, true);


--
-- Name: hours_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.hours_rowid_seq', 24, true);


--
-- Name: inventory_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.inventory_rowid_seq', 1, false);


--
-- Name: schedule_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedule_rowid_seq', 271, true);


--
-- Name: traffic_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.traffic_rowid_seq', 1, false);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (rowid);


--
-- Name: days days_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.days
    ADD CONSTRAINT days_pkey PRIMARY KEY (rowid);


--
-- Name: hours hours_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hours
    ADD CONSTRAINT hours_pkey PRIMARY KEY (rowid);


--
-- Name: inventory inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_pkey PRIMARY KEY (rowid);


--
-- Name: schedule schedule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule
    ADD CONSTRAINT schedule_pkey PRIMARY KEY (rowid);


--
-- Name: traffic traffic_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.traffic
    ADD CONSTRAINT traffic_pkey PRIMARY KEY (rowid);


--
-- Name: categoriesindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX categoriesindex ON public.categories USING btree (id);


--
-- Name: dayindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dayindex ON public.days USING btree (dayofweek);


--
-- Name: hoursindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX hoursindex ON public.hours USING btree (id);


--
-- Name: inventorybyartist; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX inventorybyartist ON public.inventory USING btree (artist, song);


--
-- Name: inventorybycategorysong; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX inventorybycategorysong ON public.inventory USING btree (category, song);


--
-- Name: inventoryplayget; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX inventoryplayget ON public.inventory USING btree (category, lastplayed, rndorder);


--
-- Name: scheduleindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX scheduleindex ON public.schedule USING btree (days, hours, "position");


--
-- Name: trafficbyartist; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX trafficbyartist ON public.traffic USING btree (artist, song, album);


--
-- PostgreSQL database dump complete
--

