--
-- PostgreSQL database dump
--

\restrict v89NdFjZrTamdiQZjbd2eaGa5RzI10z2HqNhusBvgzdSFl1YX2ulNh154vN43oZ

-- Dumped from database version 15.14 (Debian 15.14-1.pgdg13+1)
-- Dumped by pg_dump version 15.14 (Debian 15.14-1.pgdg13+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, image, created_at) FROM stdin;
bfd1dfdd-b3cc-4ef1-b782-22d00e5e313f	Breakfast	https://simply-delicious-food.com/wp-content/uploads/2019/07/Pancake-board-2.jpg	2025-09-15 10:34:31.440098+00
ac01077a-f72a-4e02-a57b-154b6e56cb86	Asian	https://easyfairsassets.com/sites/227/2023/09/AdobeStock_612129491.jpeg	2025-09-26 08:18:28.348468+00
28710db3-5cc1-4d05-b804-0dee73f40c8d	Vegetarian	https://images.everydayhealth.com/images/diet-nutrition/what-is-a-vegan-diet-benefits-food-list-beginners-guide-alt-1440x810.jpg?sfvrsn=1d260c85_5	2025-09-26 08:27:49.119452+00
95bf427a-cbfb-4a87-a2c0-fc7d46ccfcf3	Dinner	https://cdn.loveandlemons.com/wp-content/uploads/2019/09/dinner-ideas-2.jpg	2025-09-26 08:29:14.991519+00
247c24ca-ad3d-4eb2-b480-216e8d053ae6	Lunch	https://hips.hearstapps.com/hmg-prod/images/chipotle-shrimp-taco-bowls-lead-67dd8e0fedc8d.jpeg?crop=1xw:0.9998076553183305xh;center,top	2025-09-26 08:30:47.008351+00
0769238b-35ec-4d4c-a05f-c8c5d349ca6c	Snacks	https://stordfkenticomedia.blob.core.windows.net/df-us/rms/media/recipesmedia/recipes/retail/x17/2009/apr/2018_sweet-sallty-snack-mix_5817_600x600.jpg?ext=.jpg	2025-09-26 08:33:05.150439+00
1815fed3-cb49-4361-9e0f-8f0bda69eb79	Traditional Ethiopian	https://media.cnn.com/api/v1/images/stellar/prod/190205144959-shekla-tibs.jpg?q=w_1110,c_fill	2025-09-26 08:35:46.313833+00
c1f02777-87c0-465b-a878-151585d3c897	Seafood	https://www.foodandwine.com/thmb/ClPnka2WSnl5PtrMYOjlmXsXw1k=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/escovitch-fish-FT-RECIPE0920-8a733638c2ba4b72b48737782fa616c2.jpg	2025-09-26 08:36:37.03295+00
70d5fa26-76f9-429c-aeeb-5e28cb4ebf1a	Italian	https://hips.hearstapps.com/hmg-prod/images/frittata-affogato2-1662413900.jpg?crop=0.9979591836734694xw:1xh;center,top	2025-09-26 08:37:28.442354+00
f4e197f8-0702-49fc-9075-71a66d95fda1	Drinks	https://images.immediate.co.uk/production/volatile/sites/2/2021/04/Zombie-bbcc653.jpg?quality=90&resize=700,466	2025-09-26 08:38:32.587912+00
6226fdfb-945d-41ad-bbf9-6b35907f93f9	Bread	https://www.kingarthurbaking.com/sites/default/files/styles/featured_image/public/recipe_legacy/6079-3-large.jpg?itok=VWwKm7jt	2025-09-26 08:41:02.456365+00
18af902e-60d3-4128-8a7e-1914890b60a2	Cookies	https://inbloombakery.com/wp-content/uploads/2023/07/cinnamon-roll-cookies-15.jpg	2025-09-26 08:43:20.450251+00
3462a0b2-8e6b-4d41-8b07-be76f687460d	Soup and Stew	https://thewoksoflife.com/wp-content/uploads/2014/12/fish-tofu-soup-11-1.jpg	2025-09-26 08:47:12.902984+00
fc6254cd-41bc-43c7-805a-ca20a36a86da	Kid Friendly	https://thebakermama.com/wp-content/uploads/2020/04/fullsizeoutput_242f2-scaled.jpeg	2025-09-26 08:50:52.09849+00
cfb83a4d-3ac9-4813-ba7e-a377cba26fba	Appetizers	https://www.eatingwell.com/thmb/zp72KPMev-XvaIvE1VczYK6zuJM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/7947087-305293f23233424aa96ad21213d43b6f.jpg	2025-09-26 08:52:06.025684+00
43df6ee9-00c2-4a5e-8ef7-f4e8bdcb682b	Burgers	https://www.savingdessert.com/wp-content/uploads/2016/06/Marinated-Burgers-5.jpg	2025-09-26 08:54:01.059781+00
a5c1cdcd-0fb4-4acb-b730-6873d317ec1b	Desserts	https://www.iheartnaptime.net/wp-content/uploads/2022/07/I-Heart-Naptime-oreo-lasagna.jpg	2025-09-26 08:56:19.77517+00
99b75497-cf8c-4592-996d-f1b795315650	Ice Cream	https://img.freepik.com/premium-photo/three-desserts-with-chocolate-vanilla-ice-cream-tray_1276840-51653.jpg	2025-09-26 08:59:06.345292+00
c0caf9fc-562f-46c1-97cc-b7b9b7a8235e	Sauces	https://upload.wikimedia.org/wikipedia/commons/5/52/Fresh_Tomato_Sauce_%28Unsplash%29.jpg	2025-09-26 09:00:48.779022+00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password_hash, username, created_at, updated_at) FROM stdin;
37bccb1a-239d-4ef9-aa92-8c94eab57aba	tsion1@gmail.com	temp_hashed_password_123	chefTsi	2025-09-15 10:32:50.726845+00	2025-09-15 10:32:50.726845+00
03255642-0a64-447e-b70a-8d60fd3ab629	beza1@gmail.com	$2a$10$USQaFWPIIQZur84KkmtnmuPiAyYwPkZHz5vZxY7tB1sQksSq9YMJ6	bezthecook	2025-09-16 17:11:38.219893+00	2025-09-16 17:11:38.219893+00
74a0ab66-83c8-4d07-b13e-e40bb5e557b3	tsion2@gmail.com	$2a$10$uP/0/AHRZJDPqeFuxeKufea3MBVxq5EeUombQUFvFesi/zAFhRTdu	hackeduser4	2025-09-17 09:01:12.06834+00	2025-09-17 09:01:12.06834+00
4182aca0-7055-4e28-b85c-e3aff199b4ac	tsion3@gmail.com	$2a$10$Mop3MmPaAarAt0tAYI6ARe7Fzmjuuq3HZt7Y4OEKyGrBKpli/hxYy	hackeduser5	2025-09-17 11:32:05.332573+00	2025-09-17 11:32:05.332573+00
960d4be9-5371-4ac7-b831-e0f10cc343fd	daniel1@gmail.com	$2a$10$kMcZnmtnuw1joqbC.DYZee3NXMypV0LfYrFyxzYEdHwP5VUtPctWe	newname	2025-09-17 12:21:04.186802+00	2025-09-17 12:21:04.186802+00
8a459d07-b926-4f88-b460-832743501b35	beza@gmail.com	$2a$10$6KE3.X7lxEec6bUELpt40OhaRFN.fnj/CaXqYL7tL/49lGkExSn3.	nothacked	2025-09-16 13:00:57.869943+00	2025-09-16 13:00:57.869943+00
38816842-4bf8-4f0d-bd99-44dfd1ccb926	daniel2@gmail.com	$2a$10$..EjVm7z1W55wAz8WkiBEO/QJvxcwiZUFYlZvwujn8GBysEgrBQ7q	daniel2chef	2025-09-18 17:41:24.880571+00	2025-09-18 17:41:24.880571+00
eea6f74d-97e5-47b0-950f-9e6b931f3447	lola@gmail.com	$2a$10$Mka4bA9g7UB.bw5SXR1IUeOhMVwcg2HUSL4/kwKZ8CtAdcxWdtDM6	cheflola	2025-09-25 18:22:10.563476+00	2025-09-25 18:22:10.563476+00
f8b1e185-7710-4d97-b616-b8eed649f215	lola1@gmail.com	$2a$10$MwxQsmkrJjgztc5mBA3u4u3mA2jcbW48P2yd9/N7fGmCQF7FCt8L2	lola1chef	2025-09-25 21:05:11.187239+00	2025-09-25 21:05:11.187239+00
4bebabeb-7120-4d2b-9854-01ea52f66435	lola5@gmail.com	$2a$10$6/DI8GPtWylBQqcmV9wcjuKDZu5gzfOt56n0J9AY4hJ6ov1MtyBgq	lola5chef	2025-09-25 21:27:14.481613+00	2025-09-25 21:27:14.481613+00
c486e899-ae43-4163-a6a4-2b5a77944902	lola6@gmail.com	$2a$10$AIOwQxMvUh/3CtYMIasNcOY1joDzST9wmiiDOkbfmWcu1LYnK.3E2	lola6@chef	2025-09-25 21:29:32.284117+00	2025-09-25 21:29:32.284117+00
7ec35500-5247-4157-86ae-01207e24d8ff	lola7@gmail.com	$2a$10$6p0vVVNM6mu1TXVKNYcFIuHbF7uv6Ogwz.DWxVgxPOhDF4XxjWJc.	lola7chef	2025-09-25 21:39:05.738758+00	2025-09-25 21:39:05.738758+00
e3167a34-18cb-4ab5-8f71-4b1d8a3085a1	lola8@gmail.com	$2a$10$IgQuPqQ5j/d9yMxuQJFaBuAogOcLh/KSSpnbFs6D7jOxu3GLXLpJe	lola8chef	2025-09-25 21:42:34.279983+00	2025-09-25 21:42:34.279983+00
8a83f8db-7cdb-43cf-8fb2-4ecf3d4e69f6	abel1@gmail.com	$2a$10$pF17snjuGu0UxVeSPdPFeeux/ngA5HqyKAdkSeV9zFEhbGGpEl4.i	abel1chef	2025-09-25 21:44:08.642303+00	2025-09-25 21:44:08.642303+00
4b05efa0-16e3-438c-8047-db36174288d4	abel2@gmail.com	$2a$10$JgG20q6cR9VkE9eJobKXHuCYSTRqaGveGMEAt/Nu8KJkOANIpVquS	abel2chef	2025-09-25 21:55:31.649159+00	2025-09-25 21:55:31.649159+00
ea6a0f3c-d51b-4171-9604-eafae1b0eec6	abel3@g,ail.com	$2a$10$sdchmu.1ECOPxHb/u/r.auG74s1aX3gMXMuHFQiGux9jL1lekB/ze	abel3chef	2025-09-25 22:00:35.539069+00	2025-09-25 22:00:35.539069+00
884fc09f-8c34-4962-a7a1-c3aac8acc537	abel4@gmail.com	$2a$10$OVQL6GLKHtd7nlkAWlh42uZKH4G6/OkREaxbkm3h0ruZcpBGc6cEy	abel4	2025-09-25 22:06:57.075898+00	2025-09-25 22:06:57.075898+00
e4d64997-1c18-4960-9aa8-61fa549d7a50	please@gmail.com	$2a$10$arJcJVQtrvAKsXzInbeWke8uEtp.9epJj4gLhhjo7oormMKg8y6yu	please	2025-09-25 22:37:57.150845+00	2025-09-25 22:37:57.150845+00
5419e018-2e70-4646-9503-b41dcb88e01e	agneschef@gmail.com	$2a$10$XOrQ8RegpGxEm8tbSqtVU.5U5bOAaMTFwtGqrL5GDAw2TxMhBvHEG	AgnesChef	2025-09-26 08:04:13.621686+00	2025-09-26 08:04:13.621686+00
531cc3d5-65e9-4905-a56c-b34dc8fed2b6	ruya@gmail.com	$2a$10$CHoago1G5NcE1KYbkZMIje/2PNjETSsU9E07e1s4vUpH/cbOShZc6	ruya	2025-09-26 13:58:04.622129+00	2025-09-26 13:58:04.622129+00
25982b99-2d1b-403d-93dd-d51c1ae86549	ethel@gmail.com	$2a$10$ixOypLEC3o5vCIz71HV.yuOUGbb3O.ixB3ojw/Ol1Z7aqYnvWX08a	ethel	2025-09-28 00:42:39.28035+00	2025-09-28 00:42:39.28035+00
cbdc62ea-c979-4705-8a94-59c41001ea07	trying@gmail.com	$2a$10$d6rBVlVsIBQVDB/Mhx7/fuh.SjhHLxhL.RrZBJyZzwda3yu.VVGi6	trying	2025-09-28 00:43:29.786861+00	2025-09-28 00:43:29.786861+00
\.


--
-- Data for Name: recipes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.recipes (id, title, description, prep_time, featured_image, is_premium, price, user_id, category_id, created_at, updated_at) FROM stdin;
70521010-9f85-4fee-b4b0-ea6b60406f67	Scrambled Eggs with Toast	Fluffy scrambled eggs served with crispy buttered toast.	10	https://www.incredibleegg.org/wp-content/uploads/2015/06/basic-scrambled-mobile.jpeg	f	\N	37bccb1a-239d-4ef9-aa92-8c94eab57aba	bfd1dfdd-b3cc-4ef1-b782-22d00e5e313f	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
a119a17f-f672-433d-a459-47d833d4cd81	Spaghetti	Classic spaghetti recipe	20		f	\N	eea6f74d-97e5-47b0-950f-9e6b931f3447	bfd1dfdd-b3cc-4ef1-b782-22d00e5e313f	2025-09-25 19:03:36.386769+00	2025-09-25 19:03:36.386769+00
40156108-8a40-4782-9f2b-7b55ef8ac2cf	Chicken Teriyaki	A classic Japanese dish made with juicy chicken glazed in a sweet and savory teriyaki sauce.	20	https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSKCsXsEUcdrRJfYgF8ZpiNXhuDXU7SnM0q2kHjB-KTXpqARUbfy7ROrVspEXQ-rcwn_Rc&usqp=CAU	f	\N	5419e018-2e70-4646-9503-b41dcb88e01e	ac01077a-f72a-4e02-a57b-154b6e56cb86	2025-09-26 09:24:25.432338+00	2025-09-26 09:24:25.432338+00
299112fc-1778-4cd9-924a-46ba5af0a85e	cookies	cookies best	20	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758899346.jpeg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	18af902e-60d3-4128-8a7e-1914890b60a2	2025-09-26 15:09:10.191385+00	2025-09-26 15:09:10.191385+00
86b695f8-5322-470c-9940-d8f196547655	test recipe	test test	10	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758899773.jpg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	28710db3-5cc1-4d05-b804-0dee73f40c8d	2025-09-26 15:16:17.268923+00	2025-09-26 15:16:17.268923+00
2e3a8677-c69e-4c20-9671-332094d59e34	please work	please please	50	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758906219.jpg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	c1f02777-87c0-465b-a878-151585d3c897	2025-09-26 17:03:43.30571+00	2025-09-26 17:03:43.30571+00
cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	oh my	thank u bro	20	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758910625.jpg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	28710db3-5cc1-4d05-b804-0dee73f40c8d	2025-09-26 18:17:13.718673+00	2025-09-26 18:17:13.718673+00
be541a4c-89d3-4876-8563-823c7db016f1	milk	test milk	20	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758918527.jpeg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	f4e197f8-0702-49fc-9075-71a66d95fda1	2025-09-26 20:28:50.800719+00	2025-09-26 20:28:50.800719+00
128fe8c3-4af0-43d1-9023-228b5cd23305	testing images	image testt	20	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758920957.jpeg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	43df6ee9-00c2-4a5e-8ef7-f4e8bdcb682b	2025-09-26 21:09:18.80107+00	2025-09-26 21:09:18.80107+00
df757556-9312-4fa5-a92d-86b577739219	breakfast food	tryyyy	10	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758922770.jpeg	f	\N	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	bfd1dfdd-b3cc-4ef1-b782-22d00e5e313f	2025-09-26 21:39:32.2321+00	2025-09-26 21:39:32.2321+00
a9be63a1-159d-419c-bab0-04a2e0bfa50c	Doro Wot	Doro wat is a richly spiced, fragrant, and delicious chicken-and-egg stew from Ethiopia and Eritrea. The chicken cooks in a flavour-rich base of slow-cooked meal.	60	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1759002720.webp	t	5.00	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	1815fed3-cb49-4361-9e0f-8f0bda69eb79	2025-09-27 19:52:03.25826+00	2025-09-27 19:52:03.25826+00
b4ed09de-ce22-4620-9cda-1228856551d4	Fresh Garden Salad	A crisp, colorful salad packed with fresh vegetables and a light, tangy dressing. Perfect as a healthy side or a light meal.	40	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/cbdc62ea-c979-4705-8a94-59c41001ea07%2F1759020740.jpg	t	10.00	cbdc62ea-c979-4705-8a94-59c41001ea07	28710db3-5cc1-4d05-b804-0dee73f40c8d	2025-09-28 00:52:27.687071+00	2025-09-28 00:52:27.687071+00
22550c31-9cdd-46cf-b51d-5932cb274943	Mango Smoothie	A creamy, tropical smoothie that’s sweet, tangy, and perfect for a hot day. Packed with vitamins and easy to make in minutes.	15	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/cbdc62ea-c979-4705-8a94-59c41001ea07%2F1759021316.jpg	f	\N	cbdc62ea-c979-4705-8a94-59c41001ea07	f4e197f8-0702-49fc-9075-71a66d95fda1	2025-09-28 01:01:59.92792+00	2025-09-28 01:01:59.92792+00
\.


--
-- Data for Name: bookmarks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bookmarks (id, user_id, recipe_id, created_at) FROM stdin;
aa38b059-35a9-4e20-a380-8a5c62321dab	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	df757556-9312-4fa5-a92d-86b577739219	2025-09-27 17:05:18.307541+00
4e597842-a079-44b7-ac6f-91c3e3516a21	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	be541a4c-89d3-4876-8563-823c7db016f1	2025-09-27 17:05:41.898337+00
a749c0e8-2e8d-476e-9f8d-689867659ed2	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	2e3a8677-c69e-4c20-9671-332094d59e34	2025-09-27 17:05:45.201199+00
301f866f-474b-473a-be39-26399ded067c	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	299112fc-1778-4cd9-924a-46ba5af0a85e	2025-09-27 17:05:47.01354+00
40c6ca3b-821e-41bb-a31f-b5253da74ee1	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	70521010-9f85-4fee-b4b0-ea6b60406f67	2025-09-27 17:08:07.766645+00
eb4c0b66-2b66-41bd-9491-3b38d68e55d6	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	2025-09-27 18:19:14.16323+00
1484e9ff-10a2-4f7b-8462-75375f70e829	cbdc62ea-c979-4705-8a94-59c41001ea07	128fe8c3-4af0-43d1-9023-228b5cd23305	2025-09-28 01:23:57.113329+00
fe5c1fa5-d550-4a26-8008-b032b55d9d45	cbdc62ea-c979-4705-8a94-59c41001ea07	70521010-9f85-4fee-b4b0-ea6b60406f67	2025-09-28 01:24:29.344908+00
\.


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.comments (id, user_id, recipe_id, content, created_at, updated_at) FROM stdin;
ce2e2d8b-188d-4a27-aef6-373eff5109b8	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	70521010-9f85-4fee-b4b0-ea6b60406f67	This recipe is amazing!	2025-09-27 07:22:42.598036+00	2025-09-27 07:22:42.598036+00
74f8c332-c9b4-4889-870f-aa88fefcfadb	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	70521010-9f85-4fee-b4b0-ea6b60406f67	please work	2025-09-27 07:37:27.535915+00	2025-09-27 07:37:27.535915+00
e06cc43d-4366-4027-b45c-7ad5af430afd	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	70521010-9f85-4fee-b4b0-ea6b60406f67	hellooooo	2025-09-27 08:12:33.965753+00	2025-09-27 08:12:33.965753+00
5da92b29-0380-4498-8fd7-b5277db78816	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	be541a4c-89d3-4876-8563-823c7db016f1	great milk	2025-09-27 08:14:46.920996+00	2025-09-27 08:14:46.920996+00
fc3f4c72-b69e-467a-8b92-ca3ac77ce85c	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	128fe8c3-4af0-43d1-9023-228b5cd23305	niceeee recipe. yay	2025-09-27 19:07:13.587168+00	2025-09-27 19:07:13.587168+00
233a9c7e-da2e-4454-aded-58381b195eb1	cbdc62ea-c979-4705-8a94-59c41001ea07	128fe8c3-4af0-43d1-9023-228b5cd23305	good one ruya, thanks	2025-09-28 01:24:11.817718+00	2025-09-28 01:24:11.817718+00
\.


--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ingredients (id, recipe_id, name, quantity, unit, created_at, updated_at) FROM stdin;
36883c89-9039-4256-9bbb-3c445b799e29	70521010-9f85-4fee-b4b0-ea6b60406f67	Eggs	3	pieces	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
1b42da53-95c3-4aa9-a2c3-e8ae4e41f433	70521010-9f85-4fee-b4b0-ea6b60406f67	Butter	1	tbsp	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
f2bc1e3b-791d-4ace-84d6-53571fc6d1b6	70521010-9f85-4fee-b4b0-ea6b60406f67	Salt	0.25	tsp	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
f3835733-2492-469c-b7d8-13d9b6d5ad77	70521010-9f85-4fee-b4b0-ea6b60406f67	Black Pepper	0.25	tsp	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
f8c61e42-392e-45dc-a1cd-c9b60191d9cd	70521010-9f85-4fee-b4b0-ea6b60406f67	Bread Slices	2	pieces	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
53c68b69-0025-4183-9519-abc40a9c1ebc	a119a17f-f672-433d-a459-47d833d4cd81	Spaghetti	200	grams	2025-09-25 19:03:36.401836+00	2025-09-25 19:03:36.401836+00
e8354e8a-f8bb-4daf-b994-9f390c410301	a119a17f-f672-433d-a459-47d833d4cd81	Tomato Sauce	1	cup	2025-09-25 19:03:36.401836+00	2025-09-25 19:03:36.401836+00
e5fbc672-f6d7-4552-a2bf-df55fc4d4a21	40156108-8a40-4782-9f2b-7b55ef8ac2cf	Chicken thighs (boneless, skinless)	500	g	2025-09-26 09:28:04.488111+00	2025-09-26 09:28:04.488111+00
22bd7262-6cb0-4cb0-88a9-b14252205ffb	40156108-8a40-4782-9f2b-7b55ef8ac2cf	Ginger (grated)	1	tsp	2025-09-26 09:28:32.914076+00	2025-09-26 09:28:32.914076+00
c7d76ab4-4953-4fc0-903d-c862974e8aae	40156108-8a40-4782-9f2b-7b55ef8ac2cf	Green onions	2	stalks	2025-09-26 09:28:51.080068+00	2025-09-26 09:28:51.080068+00
ceb73d4d-30e3-455e-b0d8-afa151759360	40156108-8a40-4782-9f2b-7b55ef8ac2cf	Soy sauce	1	cup	2025-09-26 09:29:10.211678+00	2025-09-26 09:29:10.211678+00
418cbe6e-0bd6-463a-b1a1-8e373648fa69	299112fc-1778-4cd9-924a-46ba5af0a85e	milk	1	cup	2025-09-26 15:09:10.212232+00	2025-09-26 15:09:10.212232+00
bde5bad8-2128-4fd7-a15c-ea094b73bb0c	299112fc-1778-4cd9-924a-46ba5af0a85e	cookie dough	2	cup	2025-09-26 15:09:10.212232+00	2025-09-26 15:09:10.212232+00
3799bbe7-eedf-4073-9a8d-a3e2e9bd627b	86b695f8-5322-470c-9940-d8f196547655	text 1	3	tsp	2025-09-26 15:16:17.279119+00	2025-09-26 15:16:17.279119+00
c62f432a-7bfa-4f47-8502-89dcc9d135df	86b695f8-5322-470c-9940-d8f196547655	test 2	1	tbsp	2025-09-26 15:16:17.279119+00	2025-09-26 15:16:17.279119+00
812a3cae-0732-4e90-adf5-2f1f689c4e34	2e3a8677-c69e-4c20-9671-332094d59e34	work	1	cup	2025-09-26 17:03:43.321915+00	2025-09-26 17:03:43.321915+00
dab997e3-b2b1-4b4a-913a-5630fdf43152	2e3a8677-c69e-4c20-9671-332094d59e34	please bro	2	small	2025-09-26 17:03:43.321915+00	2025-09-26 17:03:43.321915+00
55aa4e5c-0acc-4607-a8f7-8bf59cc7e184	2e3a8677-c69e-4c20-9671-332094d59e34	wink wink	1	cup	2025-09-26 17:03:43.321915+00	2025-09-26 17:03:43.321915+00
c4853931-a992-4831-8e56-d7bc5054d41d	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	it workedddd	1	cup	2025-09-26 18:17:13.749329+00	2025-09-26 18:17:13.749329+00
8c1be5f4-c7a0-438d-a931-c4d38b67cbac	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	yayyy	2	tbsp	2025-09-26 18:17:13.749329+00	2025-09-26 18:17:13.749329+00
3dd80595-efe9-4ca8-8d44-44ed176a8210	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	huzzah	3	small	2025-09-26 18:17:13.749329+00	2025-09-26 18:17:13.749329+00
76d31c71-962f-44b1-83e0-01ad6defcd72	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	hurrah	1	tsp	2025-09-26 18:17:13.749329+00	2025-09-26 18:17:13.749329+00
12ca6b57-a9da-40ae-a1eb-3aaed21a0fbe	be541a4c-89d3-4876-8563-823c7db016f1	milks	1	cup	2025-09-26 20:28:50.810768+00	2025-09-26 20:28:50.810768+00
3951ab50-4ff8-4529-89ff-a663f8064b6b	be541a4c-89d3-4876-8563-823c7db016f1	banana	1	small	2025-09-26 20:28:50.810768+00	2025-09-26 20:28:50.810768+00
8fa73ba2-a605-4166-a407-b406c3d0d53f	128fe8c3-4af0-43d1-9023-228b5cd23305	test ing	1	tsp	2025-09-26 21:09:18.824228+00	2025-09-26 21:09:18.824228+00
ed22dcc2-df80-46d0-9c93-91899cfcfc47	128fe8c3-4af0-43d1-9023-228b5cd23305	test 2 ings	2	cups	2025-09-26 21:09:18.824228+00	2025-09-26 21:09:18.824228+00
bce44cd7-0134-4322-9ff1-380e5fdfd819	df757556-9312-4fa5-a92d-86b577739219	brekfast stuff	1	small	2025-09-26 21:39:32.240762+00	2025-09-26 21:39:32.240762+00
a8f7049b-1cc6-473f-800a-e3420d632a01	df757556-9312-4fa5-a92d-86b577739219	food	1	tsp	2025-09-26 21:39:32.240762+00	2025-09-26 21:39:32.240762+00
ef3fcec8-6721-4606-8b9e-9049cbf26bc5	a9be63a1-159d-419c-bab0-04a2e0bfa50c	doro	1	large	2025-09-27 19:52:03.279523+00	2025-09-27 19:52:03.279523+00
2faeb1a3-b38c-4a8c-a635-967001bb0368	a9be63a1-159d-419c-bab0-04a2e0bfa50c	onions	20	small	2025-09-27 19:52:03.279523+00	2025-09-27 19:52:03.279523+00
a419e64b-5df7-401f-9762-6439c9320554	a9be63a1-159d-419c-bab0-04a2e0bfa50c	butter	3	cups	2025-09-27 19:52:03.279523+00	2025-09-27 19:52:03.279523+00
ef719b8d-4231-4078-afb3-8ae8ce0a05de	a9be63a1-159d-419c-bab0-04a2e0bfa50c	spice	2	cup	2025-09-27 19:52:03.279523+00	2025-09-27 19:52:03.279523+00
04271ccf-5d28-4ca5-99f0-a64f3053fb46	b4ed09de-ce22-4620-9cda-1228856551d4	mixed lettuce	2	cups	2025-09-28 00:52:27.709865+00	2025-09-28 00:52:27.709865+00
76e15401-b8a2-4148-9311-bd54ddfde734	b4ed09de-ce22-4620-9cda-1228856551d4	cherry tomatoes	1	small	2025-09-28 00:52:27.709865+00	2025-09-28 00:52:27.709865+00
302fa194-a5c5-4003-8003-20c34e7650f9	b4ed09de-ce22-4620-9cda-1228856551d4	olive oil	2	tbsp	2025-09-28 00:52:27.709865+00	2025-09-28 00:52:27.709865+00
2d211ba9-17f1-46fb-8211-3635095ffde1	b4ed09de-ce22-4620-9cda-1228856551d4	shredded carrots	1	cup	2025-09-28 00:52:27.709865+00	2025-09-28 00:52:27.709865+00
cc88242d-a88a-47e7-a682-cf84c51e4865	b4ed09de-ce22-4620-9cda-1228856551d4	onion	1	small	2025-09-28 00:52:27.709865+00	2025-09-28 00:52:27.709865+00
dca6a04d-56f1-4fd0-86a1-34e131f668e4	22550c31-9cdd-46cf-b51d-5932cb274943	mango	1	large	2025-09-28 01:01:59.938721+00	2025-09-28 01:01:59.938721+00
a4ef4493-3146-4889-be25-1daa869c45ec	22550c31-9cdd-46cf-b51d-5932cb274943	pineapple juice	1	cup	2025-09-28 01:01:59.938721+00	2025-09-28 01:01:59.938721+00
abd78743-1790-4106-bfbd-ce887d5d7fcb	22550c31-9cdd-46cf-b51d-5932cb274943	Greek yogurt	1	cup	2025-09-28 01:01:59.938721+00	2025-09-28 01:01:59.938721+00
2f489420-3547-4ff3-a90c-3db102bbc8cb	22550c31-9cdd-46cf-b51d-5932cb274943	honey	2	tbsp	2025-09-28 01:01:59.938721+00	2025-09-28 01:01:59.938721+00
ca643a4b-05d2-430c-a1f9-84eeea74b7ae	22550c31-9cdd-46cf-b51d-5932cb274943	Ice cube	1	cup	2025-09-28 01:01:59.938721+00	2025-09-28 01:01:59.938721+00
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, user_id, recipe_id, created_at) FROM stdin;
47527a1c-7601-4e54-a97c-da564fee67ab	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	70521010-9f85-4fee-b4b0-ea6b60406f67	2025-09-27 17:08:02.847624+00
e459e3b8-855b-4aca-8703-6bccab16665f	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	40156108-8a40-4782-9f2b-7b55ef8ac2cf	2025-09-27 17:51:22.872666+00
bba3edf3-033f-4aef-aaf4-bb0d3367d863	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	be541a4c-89d3-4876-8563-823c7db016f1	2025-09-27 19:07:27.835169+00
a71b9c40-c35f-4eb3-b702-72134a969097	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	86b695f8-5322-470c-9940-d8f196547655	2025-09-27 19:07:32.590589+00
a8fc3abf-7632-4896-878a-20912ab3c4ba	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	df757556-9312-4fa5-a92d-86b577739219	2025-09-27 21:39:08.772335+00
a5db8e43-daa5-493f-9107-1633427515c0	cbdc62ea-c979-4705-8a94-59c41001ea07	128fe8c3-4af0-43d1-9023-228b5cd23305	2025-09-28 01:23:58.259724+00
51c29a2e-c34d-4436-a8b9-0d1703393fd1	cbdc62ea-c979-4705-8a94-59c41001ea07	70521010-9f85-4fee-b4b0-ea6b60406f67	2025-09-28 01:24:30.538041+00
\.


--
-- Data for Name: purchases; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.purchases (id, user_id, recipe_id, amount, status, chapa_transaction_id, created_at) FROM stdin;
1289c0e3-7c8f-4da1-b7cf-20fea63b61bc	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	a9be63a1-159d-419c-bab0-04a2e0bfa50c	5.00	pending	\N	2025-09-27 23:51:46.441539+00
\.


--
-- Data for Name: ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ratings (id, user_id, recipe_id, rating, created_at, updated_at) FROM stdin;
8c5e58c7-f2ff-46b0-bab8-54f30c093d1a	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	128fe8c3-4af0-43d1-9023-228b5cd23305	5.0	2025-09-27 11:15:56.775802+00	2025-09-27 11:15:56.775802+00
20cb91c1-5d2c-4011-8827-4c0adb5737e4	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	be541a4c-89d3-4876-8563-823c7db016f1	4.0	2025-09-27 11:16:09.071699+00	2025-09-27 11:16:09.071699+00
7006b573-9f34-413e-a9d9-2773f5f08fcf	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	70521010-9f85-4fee-b4b0-ea6b60406f67	3.0	2025-09-27 13:57:40.394585+00	2025-09-27 13:57:43.490774+00
7a48140d-3cb3-47bf-af0f-afd7f1b060ca	531cc3d5-65e9-4905-a56c-b34dc8fed2b6	df757556-9312-4fa5-a92d-86b577739219	5.0	2025-09-27 13:58:01.052945+00	2025-09-27 13:58:01.052945+00
7e454cd3-af02-41c6-9e27-190680a61324	cbdc62ea-c979-4705-8a94-59c41001ea07	128fe8c3-4af0-43d1-9023-228b5cd23305	5.0	2025-09-28 01:24:15.017424+00	2025-09-28 01:24:15.017424+00
\.


--
-- Data for Name: recipe_images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.recipe_images (id, recipe_id, image_url, is_featured, created_at) FROM stdin;
758f6eed-9518-452f-9ccb-b8a451425932	70521010-9f85-4fee-b4b0-ea6b60406f67	https://www.incredibleegg.org/wp-content/uploads/2015/06/basic-scrambled-mobile.jpeg	t	2025-09-15 11:29:05.6845+00
0f6473de-da04-49f3-8bc4-fb2ea2acc0be	40156108-8a40-4782-9f2b-7b55ef8ac2cf	https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSKCsXsEUcdrRJfYgF8ZpiNXhuDXU7SnM0q2kHjB-KTXpqARUbfy7ROrVspEXQ-rcwn_Rc&usqp=CAU	t	2025-09-26 09:24:25.432338+00
64b39357-9393-4633-ac13-daebc462f8bf	299112fc-1778-4cd9-924a-46ba5af0a85e	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758899346.jpeg	t	2025-09-26 15:09:10.191385+00
f40ee9f4-4896-40fd-bf2f-e32756261e31	86b695f8-5322-470c-9940-d8f196547655	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758899773.jpg	t	2025-09-26 15:16:17.268923+00
06bfeda2-1298-43b9-a7b4-721d519a042c	2e3a8677-c69e-4c20-9671-332094d59e34	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758906219.jpg	t	2025-09-26 17:03:43.34192+00
1ea8cbb9-9fbe-4871-aed6-aa2cfdb79f15	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758910625.jpg	t	2025-09-26 18:17:13.765209+00
56eccbb5-c629-4a8f-bd19-823fd18e859e	be541a4c-89d3-4876-8563-823c7db016f1	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758918527.jpeg	t	2025-09-26 20:28:50.82404+00
2ceebab5-7c0a-488f-a6a3-4d21075bc271	128fe8c3-4af0-43d1-9023-228b5cd23305	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758920957.jpeg	t	2025-09-26 21:09:18.837261+00
547804f8-1791-4870-82b9-94cc7bb45c00	128fe8c3-4af0-43d1-9023-228b5cd23305	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758920958.jpeg	f	2025-09-26 21:09:18.843731+00
6fea0705-6f3f-44ba-a855-8a04f1df6fd0	df757556-9312-4fa5-a92d-86b577739219	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758922770.jpeg	t	2025-09-26 21:39:32.255222+00
e2ec6fbf-c326-47ec-be43-fdfcd227cb62	df757556-9312-4fa5-a92d-86b577739219	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1758922771.jpeg	f	2025-09-26 21:39:32.261659+00
79b0b6f6-fc7d-4f20-b247-6294cc5e650f	a9be63a1-159d-419c-bab0-04a2e0bfa50c	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1759002720.webp	t	2025-09-27 19:52:03.294235+00
7be476c2-96aa-483e-98f1-4e270d7d19cc	a9be63a1-159d-419c-bab0-04a2e0bfa50c	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/531cc3d5-65e9-4905-a56c-b34dc8fed2b6%2F1759002722.webp	f	2025-09-27 19:52:03.301046+00
e17e8a22-e446-4d4f-9fa2-93a80ffd30d3	b4ed09de-ce22-4620-9cda-1228856551d4	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/cbdc62ea-c979-4705-8a94-59c41001ea07%2F1759020740.jpg	t	2025-09-28 00:52:27.735127+00
c48ab0a4-b4e5-48aa-aa9c-413473309bdb	b4ed09de-ce22-4620-9cda-1228856551d4	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/cbdc62ea-c979-4705-8a94-59c41001ea07%2F1759020745.jpg	f	2025-09-28 00:52:27.74775+00
f3d71ba8-55e5-47bd-881f-8ae930120b7b	22550c31-9cdd-46cf-b51d-5932cb274943	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/cbdc62ea-c979-4705-8a94-59c41001ea07%2F1759021316.jpg	t	2025-09-28 01:01:59.95373+00
c7244d9a-9a87-4601-a65f-44e8d53c1df6	22550c31-9cdd-46cf-b51d-5932cb274943	https://jlwyiaxjugfzuvpiscez.supabase.co/storage/v1/object/public/recipe-images/cbdc62ea-c979-4705-8a94-59c41001ea07%2F1759021318.jpg	f	2025-09-28 01:01:59.960386+00
\.


--
-- Data for Name: steps; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.steps (id, recipe_id, step_number, instruction, created_at, updated_at) FROM stdin;
0926263d-b7b4-4744-ab54-a57e2807a806	70521010-9f85-4fee-b4b0-ea6b60406f67	1	Crack eggs into a bowl, add salt and pepper, and whisk until smooth.	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
2890031a-da08-43e7-96ba-6cf03794c6b4	70521010-9f85-4fee-b4b0-ea6b60406f67	2	Melt butter in a pan over medium heat.	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
8a193706-21e1-4559-b196-e03e0ffc0985	70521010-9f85-4fee-b4b0-ea6b60406f67	3	Pour in eggs and stir gently until softly set.	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
390a5079-7f87-4741-a638-0d8ccc547619	70521010-9f85-4fee-b4b0-ea6b60406f67	4	Toast bread slices until golden brown.	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
dc32df73-b0ac-4652-9ec0-9649faffb269	70521010-9f85-4fee-b4b0-ea6b60406f67	5	Serve scrambled eggs with buttered toast on the side.	2025-09-15 11:29:05.6845+00	2025-09-15 11:29:05.6845+00
28b0ffb5-f160-41c9-8a3a-965a6fb5c4a5	40156108-8a40-4782-9f2b-7b55ef8ac2cf	1	In a bowl, mix soy sauce, mirin, sugar, garlic, and ginger	2025-09-26 09:25:46.469499+00	2025-09-26 09:25:46.469499+00
7cfd12c0-7687-4483-9f03-7920a6223231	40156108-8a40-4782-9f2b-7b55ef8ac2cf	2	Marinate chicken in the sauce for 15 minutes.	2025-09-26 09:26:05.088191+00	2025-09-26 09:26:05.088191+00
e8cab9e0-307d-45b2-b956-a2f8e29753cb	40156108-8a40-4782-9f2b-7b55ef8ac2cf	3	Heat a pan and cook the chicken until browned and cooked through	2025-09-26 09:26:19.172392+00	2025-09-26 09:26:19.172392+00
bc367846-eb3a-4785-8914-15fdc847d398	40156108-8a40-4782-9f2b-7b55ef8ac2cf	4	Garnish with chopped green onions and sesame seeds. Serve with rice	2025-09-26 09:26:34.875068+00	2025-09-26 09:26:34.875068+00
d2ade40f-1c4b-4c09-b4ba-7fbfde648d07	299112fc-1778-4cd9-924a-46ba5af0a85e	1	cookies are the best	2025-09-26 15:09:10.218896+00	2025-09-26 15:09:10.218896+00
aac88690-180e-44c3-8423-bf478f52902b	299112fc-1778-4cd9-924a-46ba5af0a85e	2	yess very much so	2025-09-26 15:09:10.218896+00	2025-09-26 15:09:10.218896+00
767a5ea4-84a2-4596-af24-9b8948a29809	86b695f8-5322-470c-9940-d8f196547655	1	best test	2025-09-26 15:16:17.286413+00	2025-09-26 15:16:17.286413+00
b754a4d9-e727-4c72-9a08-7db77c8f24b5	86b695f8-5322-470c-9940-d8f196547655	2	test 2	2025-09-26 15:16:17.286413+00	2025-09-26 15:16:17.286413+00
a80a1a68-b6b8-4457-a687-b7f61308b09f	2e3a8677-c69e-4c20-9671-332094d59e34	1	test recipe	2025-09-26 17:03:43.330855+00	2025-09-26 17:03:43.330855+00
5a74c13a-4cc2-4406-bc2d-7a5b480758f5	2e3a8677-c69e-4c20-9671-332094d59e34	2	i hope it works	2025-09-26 17:03:43.330855+00	2025-09-26 17:03:43.330855+00
fc1cc9de-e6fb-4af2-9570-e70aeb620697	2e3a8677-c69e-4c20-9671-332094d59e34	3	or wtv atp	2025-09-26 17:03:43.330855+00	2025-09-26 17:03:43.330855+00
5b25a6c3-a189-47e3-9607-8da5e26fd03c	2e3a8677-c69e-4c20-9671-332094d59e34	4	please im kidding	2025-09-26 17:03:43.330855+00	2025-09-26 17:03:43.330855+00
a610c1c2-ba08-403e-ba33-010df2610d82	cb2971e4-22fa-4f9e-ba3a-bdf8b374a101	1	receipeee works	2025-09-26 18:17:13.757615+00	2025-09-26 18:17:13.757615+00
0d40877d-802b-4d88-a1fe-fa4a8ce51dad	be541a4c-89d3-4876-8563-823c7db016f1	1	mix both in a shaker	2025-09-26 20:28:50.817542+00	2025-09-26 20:28:50.817542+00
55bad65d-7b59-4f9a-adf8-1e2a7b553eef	be541a4c-89d3-4876-8563-823c7db016f1	2	mix very well	2025-09-26 20:28:50.817542+00	2025-09-26 20:28:50.817542+00
c46bc211-0e56-431b-b031-4ca880bb15b2	be541a4c-89d3-4876-8563-823c7db016f1	3	drink	2025-09-26 20:28:50.817542+00	2025-09-26 20:28:50.817542+00
b9a9137d-0e0a-451c-89ad-d92609b2b8d2	128fe8c3-4af0-43d1-9023-228b5cd23305	1	test image	2025-09-26 21:09:18.830629+00	2025-09-26 21:09:18.830629+00
4b932391-7c6c-42ab-9ef0-8696bad6f73b	128fe8c3-4af0-43d1-9023-228b5cd23305	2	image test	2025-09-26 21:09:18.830629+00	2025-09-26 21:09:18.830629+00
adca2d42-ffbe-415c-93ee-4cdc1ed20c29	df757556-9312-4fa5-a92d-86b577739219	1	eat	2025-09-26 21:39:32.248531+00	2025-09-26 21:39:32.248531+00
0cf8fbba-131e-493a-b168-83ca182b3f94	df757556-9312-4fa5-a92d-86b577739219	2	die	2025-09-26 21:39:32.248531+00	2025-09-26 21:39:32.248531+00
8321c3ac-4287-4f80-9f19-d1cc26333e15	df757556-9312-4fa5-a92d-86b577739219	3	eat again	2025-09-26 21:39:32.248531+00	2025-09-26 21:39:32.248531+00
47242b57-c55a-4942-80c9-4eab1d7c741a	a9be63a1-159d-419c-bab0-04a2e0bfa50c	1	first kill the doro	2025-09-27 19:52:03.287278+00	2025-09-27 19:52:03.287278+00
1cc5c44c-6837-464a-a89f-fe1c916df9f3	a9be63a1-159d-419c-bab0-04a2e0bfa50c	2	second fillet it like a master	2025-09-27 19:52:03.287278+00	2025-09-27 19:52:03.287278+00
80ccf1a5-1c75-4f3e-9d9f-737d4c5d2345	a9be63a1-159d-419c-bab0-04a2e0bfa50c	3	boil boil boil	2025-09-27 19:52:03.287278+00	2025-09-27 19:52:03.287278+00
d27dc33c-55fa-4037-8e77-af8d902d5578	a9be63a1-159d-419c-bab0-04a2e0bfa50c	4	mix mix mix	2025-09-27 19:52:03.287278+00	2025-09-27 19:52:03.287278+00
5d82398e-9353-4fae-b515-680820cc6cfe	a9be63a1-159d-419c-bab0-04a2e0bfa50c	5	eat eat eat	2025-09-27 19:52:03.287278+00	2025-09-27 19:52:03.287278+00
5a1c3b50-d295-48d3-aa49-d2b6a7c46c8c	b4ed09de-ce22-4620-9cda-1228856551d4	1	Wash and dry all the vegetables thoroughly.	2025-09-28 00:52:27.723169+00	2025-09-28 00:52:27.723169+00
6f601dee-1d17-42e6-afc6-a1178a7875ae	b4ed09de-ce22-4620-9cda-1228856551d4	2	In a large salad bowl, combine lettuce, cherry tomatoes, cucumber, red onion, and carrots.	2025-09-28 00:52:27.723169+00	2025-09-28 00:52:27.723169+00
5b85d305-1280-4e72-ab8a-be69f431a619	b4ed09de-ce22-4620-9cda-1228856551d4	3	In a small bowl, whisk together olive oil, balsamic vinegar (or lemon juice), salt, and pepper.	2025-09-28 00:52:27.723169+00	2025-09-28 00:52:27.723169+00
3a891816-4258-45d5-a02a-440b28f53ff6	b4ed09de-ce22-4620-9cda-1228856551d4	4	Drizzle the dressing over the salad and toss gently to coat.	2025-09-28 00:52:27.723169+00	2025-09-28 00:52:27.723169+00
3737e9bc-e78a-436f-8772-1184480c488e	b4ed09de-ce22-4620-9cda-1228856551d4	5	Sprinkle feta cheese and nuts on top, if using.	2025-09-28 00:52:27.723169+00	2025-09-28 00:52:27.723169+00
08321ba5-b920-4085-8206-e41054bb92e7	b4ed09de-ce22-4620-9cda-1228856551d4	6	Serve immediately and enjoy your fresh garden salad!	2025-09-28 00:52:27.723169+00	2025-09-28 00:52:27.723169+00
ace16d3e-dbc5-4069-b45e-4d94c366d27b	22550c31-9cdd-46cf-b51d-5932cb274943	1	Add mango, pineapple, orange juice, yogurt, and honey into a blender.	2025-09-28 01:01:59.946781+00	2025-09-28 01:01:59.946781+00
3270b0a4-a5c1-4fab-8703-1d9200733ed7	22550c31-9cdd-46cf-b51d-5932cb274943	2	Blend until smooth and creamy.	2025-09-28 01:01:59.946781+00	2025-09-28 01:01:59.946781+00
ce6ca4b9-3189-4228-b16a-8aac37c84bae	22550c31-9cdd-46cf-b51d-5932cb274943	3	Add ice cubes and blend again until frothy.	2025-09-28 01:01:59.946781+00	2025-09-28 01:01:59.946781+00
f4ebae5a-8a9f-4d03-95a3-44d8c095f601	22550c31-9cdd-46cf-b51d-5932cb274943	4	Pour into a glass and garnish with a small pineapple wedge or mint leaves.	2025-09-28 01:01:59.946781+00	2025-09-28 01:01:59.946781+00
29fa5f8a-5176-4082-9863-71ac971a5fa9	22550c31-9cdd-46cf-b51d-5932cb274943	5	Serve immediately and enjoy your tropical drink!	2025-09-28 01:01:59.946781+00	2025-09-28 01:01:59.946781+00
\.


--
-- PostgreSQL database dump complete
--

\unrestrict v89NdFjZrTamdiQZjbd2eaGa5RzI10z2HqNhusBvgzdSFl1YX2ulNh154vN43oZ

