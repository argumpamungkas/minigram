--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

-- Started on 2025-05-30 14:24:42

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
-- TOC entry 220 (class 1259 OID 16417)
-- Name: postings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.postings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    photo text NOT NULL,
    caption character varying(250),
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.postings OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16416)
-- Name: postings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.postings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.postings_id_seq OWNER TO postgres;

--
-- TOC entry 4912 (class 0 OID 0)
-- Dependencies: 219
-- Name: postings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.postings_id_seq OWNED BY public.postings.id;


--
-- TOC entry 218 (class 1259 OID 16404)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    full_name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password text NOT NULL,
    avatar text,
    bio character varying(200),
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16403)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4913 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4748 (class 2604 OID 16420)
-- Name: postings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.postings ALTER COLUMN id SET DEFAULT nextval('public.postings_id_seq'::regclass);


--
-- TOC entry 4747 (class 2604 OID 16407)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4906 (class 0 OID 16417)
-- Dependencies: 220
-- Data for Name: postings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.postings (id, user_id, photo, caption, created_date, updated_date) FROM stdin;
3	1	6YaGHUYxgrLlg0JZofr8xwRwQF7.jpeg	larilari	2025-05-30 09:36:25.022286	\N
\.


--
-- TOC entry 4904 (class 0 OID 16404)
-- Dependencies: 218
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, full_name, email, password, avatar, bio, created_date, updated_date) FROM stdin;
1	argum1	Argumelar	argum1@gmail.com	$2a$08$TsSpbYTje3DR1hfiA3jUeO/9GcmY6/HTBqABqiCnPwz4ZnSdh.GDq	\N	\N	2025-05-30 09:20:37.505618	\N
2	argum2	Argumelar	argum2@gmail.com	$2a$08$ZRlgGYmAYckORzMPLFkdhemgdtTQFOYMicojRCxd67eCugD5sApJe	\N	\N	2025-05-30 09:21:24.116706	\N
\.


--
-- TOC entry 4914 (class 0 OID 0)
-- Dependencies: 219
-- Name: postings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.postings_id_seq', 3, true);


--
-- TOC entry 4915 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 4756 (class 2606 OID 16424)
-- Name: postings postings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.postings
    ADD CONSTRAINT postings_pkey PRIMARY KEY (id);


--
-- TOC entry 4750 (class 2606 OID 16415)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4752 (class 2606 OID 16411)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4754 (class 2606 OID 16413)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 4757 (class 2606 OID 16425)
-- Name: postings fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.postings
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2025-05-30 14:24:43

--
-- PostgreSQL database dump complete
--

