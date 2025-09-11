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
3	MON	00	03	CURRENTS	1
4	MON	00	04	ADS	2
5	MON	00	05	IMAGINGID	1
6	MON	00	06	CURRENTS	1
7	MON	00	07	RECURRENTS	1
8	MON	00	08	CURRENTS	1
9	MON	00	09	IMAGINGID	1
10	MON	00	10	RECURRENTS	2
11	MON	00	11	IMAGINGID	1
12	MON	00	12	RECURRENTS	2
13	MON	00	13	IMAGINGID	1
14	MON	00	14	CURRENTS	1
15	MON	00	15	RECURRENTS	1
16	MON	00	16	CURRENTS	1
17	MON	00	17	ADS	2
18	MON	00	18	RECURRENTS	2
19	MON	00	19	IMAGINGID	1
20	MON	00	20	RECURRENTS	2
21	MON	00	21	IMAGINGID	1
22	MON	00	22	FILLTOTOH	1
23	MON	01	01	STATIONID	1
24	MON	01	02	PROMOS	1
25	MON	01	03	CURRENTS	1
26	MON	01	04	ADS	2
27	MON	01	05	IMAGINGID	1
28	MON	01	06	CURRENTS	1
29	MON	01	07	RECURRENTS	1
30	MON	01	08	CURRENTS	1
31	MON	01	09	IMAGINGID	1
32	MON	01	10	RECURRENTS	2
33	MON	01	11	IMAGINGID	1
34	MON	01	12	RECURRENTS	2
35	MON	01	13	IMAGINGID	1
36	MON	01	14	CURRENTS	1
37	MON	01	15	RECURRENTS	1
38	MON	01	16	CURRENTS	1
39	MON	01	17	ADS	2
40	MON	01	18	RECURRENTS	2
41	MON	01	19	IMAGINGID	1
42	MON	01	20	RECURRENTS	2
43	MON	01	21	IMAGINGID	1
44	MON	01	22	FILLTOTOH	1
45	MON	02	01	STATIONID	1
46	MON	02	02	PROMOS	1
47	MON	02	03	CURRENTS	1
48	MON	02	04	ADS	2
49	MON	02	05	IMAGINGID	1
50	MON	02	06	CURRENTS	1
51	MON	02	07	RECURRENTS	1
52	MON	02	08	CURRENTS	1
53	MON	02	09	IMAGINGID	1
54	MON	02	10	RECURRENTS	2
55	MON	02	11	IMAGINGID	1
56	MON	02	12	RECURRENTS	2
57	MON	02	13	IMAGINGID	1
58	MON	02	14	CURRENTS	1
59	MON	02	15	RECURRENTS	1
60	MON	02	16	CURRENTS	1
61	MON	02	17	ADS	2
62	MON	02	18	RECURRENTS	2
63	MON	02	19	IMAGINGID	1
64	MON	02	20	RECURRENTS	2
65	MON	02	21	IMAGINGID	1
66	MON	02	22	FILLTOTOH	1
67	MON	03	01	STATIONID	1
68	MON	03	02	PROMOS	1
69	MON	03	03	CURRENTS	1
70	MON	03	04	ADS	2
71	MON	03	05	IMAGINGID	1
72	MON	03	06	CURRENTS	1
73	MON	03	07	RECURRENTS	1
74	MON	03	08	CURRENTS	1
75	MON	03	09	IMAGINGID	1
76	MON	03	10	RECURRENTS	2
77	MON	03	11	IMAGINGID	1
78	MON	03	12	RECURRENTS	2
79	MON	03	13	IMAGINGID	1
80	MON	03	14	CURRENTS	1
81	MON	03	15	RECURRENTS	1
82	MON	03	16	CURRENTS	1
83	MON	03	17	ADS	2
84	MON	03	18	RECURRENTS	2
85	MON	03	19	IMAGINGID	1
86	MON	03	20	RECURRENTS	2
87	MON	03	21	IMAGINGID	1
88	MON	03	22	FILLTOTOH	1
89	MON	04	01	STATIONID	1
90	MON	04	02	PROMOS	1
91	MON	04	03	CURRENTS	1
92	MON	04	04	ADS	2
93	MON	04	05	IMAGINGID	1
94	MON	04	06	CURRENTS	1
95	MON	04	07	RECURRENTS	1
96	MON	04	08	CURRENTS	1
97	MON	04	09	IMAGINGID	1
98	MON	04	10	RECURRENTS	2
99	MON	04	11	IMAGINGID	1
100	MON	04	12	RECURRENTS	2
101	MON	04	13	IMAGINGID	1
102	MON	04	14	CURRENTS	1
103	MON	04	15	RECURRENTS	1
104	MON	04	16	CURRENTS	1
105	MON	04	17	ADS	2
106	MON	04	18	RECURRENTS	2
107	MON	04	19	IMAGINGID	1
108	MON	04	20	RECURRENTS	2
109	MON	04	21	IMAGINGID	1
110	MON	04	22	FILLTOTOH	1
111	MON	05	01	STATIONID	1
112	MON	05	02	PROMOS	1
113	MON	05	03	CURRENTS	1
114	MON	05	04	ADS	2
115	MON	05	05	IMAGINGID	1
116	MON	05	06	CURRENTS	1
117	MON	05	07	RECURRENTS	1
118	MON	05	08	CURRENTS	1
119	MON	05	09	IMAGINGID	1
120	MON	05	10	RECURRENTS	2
121	MON	05	11	IMAGINGID	1
122	MON	05	12	RECURRENTS	2
123	MON	05	13	IMAGINGID	1
124	MON	05	14	CURRENTS	1
125	MON	05	15	RECURRENTS	1
126	MON	05	16	CURRENTS	1
127	MON	05	17	ADS	2
128	MON	05	18	RECURRENTS	2
129	MON	05	19	IMAGINGID	1
130	MON	05	20	RECURRENTS	2
131	MON	05	21	IMAGINGID	1
132	MON	05	22	FILLTOTOH	1
133	MON	06	01	STATIONID	1
134	MON	06	02	PROMOS	1
135	MON	06	03	CURRENTS	2
136	MON	06	04	ADS	3
137	MON	06	05	CURRENTS	1
138	MON	06	06	RECURRENTS	2
139	MON	06	07	IMAGINGID	1
140	MON	06	08	CURRENTS	1
141	MON	06	09	RECURRENTS	2
142	MON	06	10	IMAGINGID	1
143	MON	06	11	ADS	5
144	MON	06	12	NWS-1-PLAYONCE	1
145	MON	06	13	CURRENTS	2
146	MON	06	14	IMAGINGID	1
147	MON	06	15	RECURRENTS	1
148	MON	06	16	IMAGINGID	1
149	MON	06	17	FILLTOTOH	1
150	MON	07	01	STATIONID	1
151	MON	07	02	PROMOS	1
152	MON	07	03	CURRENTS	2
153	MON	07	04	DJAM-71-PLAYONCE	1
154	MON	07	05	ADS	4
155	MON	07	06	CURRENTS	1
156	MON	07	07	NWS-2-PLAYONCE	1
157	MON	07	08	RECURRENTS	3
158	MON	07	09	IMAGINGID	1
159	MON	07	10	CURRENTS	1
160	MON	07	11	DJAM-72-PLAYONCE	1
161	MON	07	12	ADS	4
162	MON	07	13	CURRENTS	1
163	MON	07	14	NWS-3-PLAYONCE	1
164	MON	07	15	CURRENTS	2
165	MON	07	16	IMAGINGID	1
166	MON	07	17	FILLTOTOH	1
167	MON	08	01	STATIONID	1
168	MON	08	02	PROMOS	1
169	MON	08	03	CURRENTS	2
170	MON	08	04	DJAM-81-PLAYONCE	1
171	MON	08	05	ADS	4
172	MON	08	06	CURRENTS	1
173	MON	08	07	NWS-4-PLAYONCE	1
174	MON	08	08	RECURRENTS	2
175	MON	08	09	IMAGINGID	1
176	MON	08	10	CURRENTS	1
177	MON	08	11	DJAM-82-PLAYONCE	1
178	MON	08	12	ADS	4
179	MON	08	13	CURRENTS	1
180	MON	08	14	NWS-5-PLAYONCE	1
181	MON	08	15	CURRENTS	2
182	MON	08	16	IMAGINGID	1
183	MON	08	17	FILLTOTOH	1
184	MON	09	01	STATIONID	1
185	MON	09	02	PROMOS	1
186	MON	09	03	CURRENTS	2
187	MON	09	04	DJAM-91-PLAYONCE	1
188	MON	09	05	ADS	4
189	MON	09	06	CURRENTS	1
190	MON	09	07	IMAGINGID	1
191	MON	09	08	RECURRENTS	2
192	MON	09	09	IMAGINGID	1
193	MON	09	10	CURRENTS	2
194	MON	09	11	DJAM-92-PLAYONCE	1
195	MON	09	12	RECURRENTS	2
196	MON	09	13	ADS	4
197	MON	09	14	CURRENTS	1
198	MON	09	15	IMAGINGID	1
199	MON	09	16	FILLTOTOH	1
200	MON	10	01	STATIONID	1
201	MON	10	02	PROMOS	1
202	MON	10	03	CURRENTS	2
203	MON	10	04	DJAM-101-PLAYONCE	1
204	MON	10	05	ADS	4
205	MON	10	06	CURRENTS	1
206	MON	10	06	IMAGINGID	1
207	MON	10	07	RECURRENTS	2
208	MON	10	08	IMAGINGID	1
209	MON	10	09	CURRENTS	2
210	MON	10	10	DJAM-102-PLAYONCE	1
211	MON	10	11	CURRENTS	1
212	MON	10	12	ADS	4
213	MON	10	14	IMAGINGID	1
214	MON	10	14	CURRENTS	1
215	MON	10	15	IMAGINGID	1
216	MON	10	16	FILLTOTOH	1
217	MON	11	01	STATIONID	1
218	MON	11	02	PROMOS	1
219	MON	11	03	CURRENTS	2
220	MON	11	04	DJAM-111-PLAYONCE	1
221	MON	11	05	ADS	4
222	MON	11	06	CURRENTS	1
223	MON	11	07	IMAGINGID	1
224	MON	11	08	RECURRENTS	2
225	MON	11	09	IMAGINGID	1
226	MON	11	10	CURRENTS	2
227	MON	11	11	DJAM-112-PLAYONCE	1
228	MON	11	12	CURRENTS	1
229	MON	11	13	RECURRENTS	2
230	MON	11	14	ADS	4
231	MON	11	15	CURRENTS	1
232	MON	11	16	IMAGINGID	1
233	MON	11	17	FILLTOTOH	1
234	MON	12	01	STATIONID	1
235	MON	12	02	PROMOS	1
236	MON	12	03	CURRENTS	2
237	MON	12	04	DJPM-121-PLAYONCE	1
238	MON	12	05	ADS	4
239	MON	12	06	CURRENTS	1
240	MON	12	07	IMAGINGID	1
241	MON	12	08	RECURRENTS	2
242	MON	12	09	IMAGINGID	1
243	MON	12	10	CURRENTS	2
244	MON	12	11	DJPM-122-PLAYONCE	1
245	MON	12	12	CURRENTS	1
246	MON	12	13	RECURRENTS	2
247	MON	12	14	ADS	4
248	MON	12	15	CURRENTS	2
249	MON	12	16	IMAGINGID	1
250	MON	12	17	FILLTOTOH	1
251	MON	13	01	STATIONID	1
252	MON	13	02	PROMOS	1
253	MON	13	03	CURRENTS	2
254	MON	13	04	DJPM-121-PLAYONCE	1
255	MON	13	05	ADS	4
256	MON	13	06	CURRENTS	1
257	MON	13	07	IMAGINGID	1
258	MON	13	08	RECURRENTS	2
259	MON	13	09	IMAGINGID	1
260	MON	13	10	CURRENTS	2
261	MON	13	11	DJPM-122-PLAYONCE	1
262	MON	13	12	CURRENTS	1
263	MON	13	13	RECURRENTS	2
264	MON	13	14	ADS	4
265	MON	13	15	CURRENTS	2
266	MON	13	16	IMAGINGID	1
267	MON	13	17	FILLTOTOH	1
268	MON	14	01	STATIONID	1
269	MON	14	02	PROMOS	1
270	MON	14	03	CURRENTS	2
271	MON	14	04	DJPM-141-PLAYONCE	1
272	MON	14	05	ADS	4
273	MON	14	06	CURRENTS	1
274	MON	14	07	IMAGINGID	1
275	MON	14	08	RECURRENTS	2
276	MON	14	09	IMAGINGID	1
277	MON	14	10	CURRENTS	2
278	MON	14	11	DJPM-142-PLAYONCE	1
279	MON	14	12	CURRENTS	1
280	MON	14	13	RECURRENTS	2
281	MON	14	14	ADS	4
282	MON	14	15	CURRENTS	2
283	MON	14	16	IMAGINGID	1
284	MON	14	17	FILLTOTOH	1
285	MON	15	01	STATIONID	1
286	MON	15	02	PROMOS	1
287	MON	15	03	CURRENTS	2
288	MON	15	04	DJPM-151-PLAYONCE	1
289	MON	15	05	ADS	4
290	MON	15	06	CURRENTS	1
291	MON	15	07	IMAGINGID	1
292	MON	15	08	RECURRENTS	2
293	MON	15	09	IMAGINGID	1
294	MON	15	10	CURRENTS	2
295	MON	15	11	DJPM-152-PLAYONCE	1
296	MON	15	12	CURRENTS	1
297	MON	15	13	RECURRENTS	2
298	MON	15	14	ADS	4
299	MON	15	15	CURRENTS	2
300	MON	15	16	IMAGINGID	1
301	MON	15	17	FILLTOTOH	1
302	MON	16	01	STATIONID	1
303	MON	16	02	PROMOS	1
304	MON	16	03	CURRENTS	2
305	MON	16	04	DJPM-161-PLAYONCE	1
306	MON	16	05	ADS	4
307	MON	16	06	CURRENTS	1
308	MON	16	07	IMAGINGID	1
309	MON	16	08	RECURRENTS	2
310	MON	16	09	IMAGINGID	1
311	MON	16	10	CURRENTS	2
312	MON	16	11	DJPM-162-PLAYONCE	1
313	MON	16	12	CURRENTS	1
314	MON	16	13	RECURRENTS	2
315	MON	16	14	ADS	4
316	MON	16	15	CURRENTS	2
317	MON	16	16	IMAGINGID	1
318	MON	16	17	FILLTOTOH	1
319	MON	17	01	STATIONID	1
320	MON	17	02	PROMOS	1
321	MON	17	03	CURRENTS	2
322	MON	17	04	DJPM-171-PLAYONCE	1
323	MON	17	05	ADS	4
324	MON	17	06	CURRENTS	1
325	MON	17	07	IMAGINGID	1
326	MON	17	08	RECURRENTS	2
327	MON	17	09	IMAGINGID	1
328	MON	17	10	CURRENTS	1
329	MON	17	11	DJPM-172-PLAYONCE	1
330	MON	17	12	CURRENTS	1
331	MON	17	13	RECURRENTS	2
332	MON	17	14	ADS	4
333	MON	17	15	CURRENTS	2
334	MON	17	16	IMAGINGID	1
335	MON	17	17	FILLTOTOH	1
336	MON	18	01	STATIONID	1
337	MON	18	02	PROMOS	1
338	MON	18	03	CURRENTS	2
339	MON	18	04	DJPM-181-PLAYONCE	1
340	MON	18	05	ADS	4
341	MON	18	06	CURRENTS	1
342	MON	18	07	IMAGINGID	1
343	MON	18	08	RECURRENTS	3
344	MON	18	09	IMAGINGID	1
345	MON	18	10	CURRENTS	2
346	MON	18	11	DJPM-182-PLAYONCE	1
347	MON	18	12	CURRENTS	1
348	MON	18	13	RECURRENTS	2
349	MON	18	14	ADS	4
350	MON	18	15	CURRENTS	2
351	MON	18	16	IMAGINGID	1
352	MON	18	17	FILLTOTOH	1
353	MON	19	01	STATIONID	1
354	MON	19	02	PROMOS	1
355	MON	19	03	CURRENTS	1
356	MON	19	04	ADS	2
357	MON	19	05	IMAGINGID	1
358	MON	19	06	CURRENTS	1
359	MON	19	07	RECURRENTS	1
360	MON	19	08	CURRENTS	1
361	MON	19	09	IMAGINGID	1
362	MON	19	10	RECURRENTS	2
363	MON	19	11	IMAGINGID	1
364	MON	19	12	RECURRENTS	2
365	MON	19	13	IMAGINGID	1
366	MON	19	14	CURRENTS	1
367	MON	19	15	RECURRENTS	1
368	MON	19	16	CURRENTS	1
369	MON	19	17	ADS	2
370	MON	19	18	RECURRENTS	2
371	MON	19	19	IMAGINGID	1
372	MON	19	20	RECURRENTS	2
373	MON	19	21	IMAGINGID	1
374	MON	19	22	FILLTOTOH	1
375	MON	20	01	STATIONID	1
376	MON	20	02	PROMOS	1
377	MON	20	03	CURRENTS	1
378	MON	20	04	ADS	2
379	MON	20	05	IMAGINGID	1
380	MON	20	06	CURRENTS	1
381	MON	20	07	RECURRENTS	1
382	MON	20	08	CURRENTS	1
383	MON	20	09	IMAGINGID	1
384	MON	20	10	RECURRENTS	2
385	MON	20	11	IMAGINGID	1
386	MON	20	12	RECURRENTS	2
387	MON	20	13	IMAGINGID	1
388	MON	20	14	CURRENTS	1
389	MON	20	15	RECURRENTS	1
390	MON	20	16	CURRENTS	1
391	MON	20	17	ADS	2
392	MON	20	18	RECURRENTS	2
393	MON	20	19	IMAGINGID	1
394	MON	20	20	RECURRENTS	2
395	MON	20	21	IMAGINGID	1
396	MON	20	22	FILLTOTOH	1
397	MON	21	01	STATIONID	1
398	MON	21	02	PROMOS	1
399	MON	21	03	CURRENTS	1
400	MON	21	04	ADS	2
401	MON	21	05	IMAGINGID	1
402	MON	21	06	CURRENTS	1
403	MON	21	07	RECURRENTS	1
404	MON	21	08	CURRENTS	1
405	MON	21	09	IMAGINGID	1
406	MON	21	10	RECURRENTS	2
407	MON	21	11	IMAGINGID	1
408	MON	21	12	RECURRENTS	2
409	MON	21	13	IMAGINGID	1
410	MON	21	14	CURRENTS	1
411	MON	21	15	RECURRENTS	1
412	MON	21	16	CURRENTS	1
413	MON	21	17	ADS	2
414	MON	21	18	RECURRENTS	2
415	MON	21	19	IMAGINGID	1
416	MON	21	20	RECURRENTS	2
417	MON	21	21	IMAGINGID	1
418	MON	21	22	FILLTOTOH	1
419	MON	22	01	STATIONID	1
420	MON	22	02	PROMOS	1
421	MON	22	03	CURRENTS	1
422	MON	22	04	ADS	2
423	MON	22	05	IMAGINGID	1
424	MON	22	06	CURRENTS	1
425	MON	22	07	RECURRENTS	1
426	MON	22	08	CURRENTS	1
427	MON	22	09	IMAGINGID	1
428	MON	22	10	RECURRENTS	2
429	MON	22	11	IMAGINGID	1
430	MON	22	12	RECURRENTS	2
431	MON	22	13	IMAGINGID	1
432	MON	22	14	CURRENTS	1
433	MON	22	15	RECURRENTS	1
434	MON	22	16	CURRENTS	1
435	MON	22	17	ADS	2
436	MON	22	18	RECURRENTS	2
437	MON	22	19	IMAGINGID	1
438	MON	22	20	RECURRENTS	2
439	MON	22	21	IMAGINGID	1
440	MON	22	22	FILLTOTOH	1
441	MON	23	01	STATIONID	1
442	MON	23	02	PROMOS	1
443	MON	23	03	CURRENTS	1
444	MON	23	04	ADS	2
445	MON	23	05	IMAGINGID	1
446	MON	23	06	CURRENTS	1
447	MON	23	07	RECURRENTS	1
448	MON	23	08	CURRENTS	1
449	MON	23	09	IMAGINGID	1
450	MON	23	10	RECURRENTS	2
451	MON	23	11	IMAGINGID	1
452	MON	23	12	RECURRENTS	2
453	MON	23	13	IMAGINGID	1
454	MON	23	14	CURRENTS	1
455	MON	23	15	RECURRENTS	1
456	MON	23	16	CURRENTS	1
457	MON	23	17	ADS	2
458	MON	23	18	RECURRENTS	2
459	MON	23	19	IMAGINGID	1
460	MON	23	20	RECURRENTS	2
461	MON	23	21	IMAGINGID	1
462	MON	23	22	FILLTOTOH	1
\.


--
-- Data for Name: traffic; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.traffic (rowid, artist, song, album, playedon) FROM stdin;
\.


--
-- Name: categories_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_rowid_seq', 37, true);


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

SELECT pg_catalog.setval('public.schedule_rowid_seq', 462, true);


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

