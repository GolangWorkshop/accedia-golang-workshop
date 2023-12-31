INSERT INTO "public"."genres"(genre) 
VALUES 
  ('fiction'),
  ('criminal'),
  ('fantasy'),
  ('novel'),
  ('tragedy'),
  ('childrens literature'),
  ('biography'),
  ('self-help');

INSERT INTO "authors" ("first_name", "last_name") 
VALUES
  ('Agatha', 'Christie'),
  ('Joanne', 'Rowling'),
  ('Astrid', 'Lindgren'),
  ('Oscar', 'Wilde'),
  ('Mark', 'Twain'),
  ('Stephen', 'King'),
  ('Paulo', 'Coelho'),
  ('Leo', 'Tolstoy'),
  ('William', 'Shakespeare'),
  ('Aaron', 'Blabey'),
  ('Liz ', 'Wong'),
  ('Peppa', 'Pig'),
  ('Trent', 'Dalton'),
  ('Conrad', 'Mason'),
  ('Glennon', 'Doyle'),
  ('Jeff', 'Kinney'),
  ('Barack', 'Obama'),
  ('Julia', 'Donaldson'),
  ('John', 'Grisham'),
  ('Hanya', 'Yanagihara'),
  ('Nicola', 'Barker'),
  ('Jenny', 'Erpenbeck'),
  ('Arsene', 'Wenger'),
  ('Mary', 'McAleese'),
  ('Matthew', 'McConaughey'),
  ('Michelle', 'Obama'),
  ('Tara', 'Westover'),
  ('Natalie', 'Haynes'),
  ('Angela', 'Duckworth');

INSERT INTO "book_statuses" ("status") 
VALUES
  ('free'),
  ('borrowed'),
  ('unlisted');

INSERT INTO "books" ("name", "image", "description", "author_id", "genre_id", "status_id", "is_deleted") 
VALUES
  ('Pippi Longstocking', 'https://upload.wikimedia.org/wikipedia/en/7/78/L%C3%A5ngstrump_G%C3%A5r_Ombord.jpeg', 'The beloved story of a spunky young girl and her hilarious escapades. "A rollicking story."--The Horn Book Tommy and his sister Annika have a new neighbor, and her name is Pippi Longstocking. She has crazy red pigtails, no parents to tell her what to do, a horse that lives on her porch, and a flair for the outrageous that seems to lead to one adventure after another!', 3, 1, 1, false),
  ('Order of the Phoenix', 'https://upload.wikimedia.org/wikipedia/en/7/70/Harry_Potter_and_the_Order_of_the_Phoenix.jpg', 'Rowling and the fifth novel in the Harry Potter series. It follows Harry Potter''s struggles through his fifth year at Hogwarts School of Witchcraft and Wizardry, including the surreptitious return of the antagonist Lord Voldemort, O.W.L. exams, and an obstructive Ministry of Magic.', 2, 3, 1, false),
  ('The Mysterious Affair at Styles', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/0084/9780008400637.jpg', 'Agatha Christie’s first novel, The Mysterious Affair at Styles, was the result of a dare from her sister Madge who challenged her to write a story. The story begins when Hastings is sent back to England from the First World War due to injury and is invited to spend his sick leave at the beautiful Styles Court by his old friend John Cavendish. Here, Hastings meets John’s step-mother, Mrs Inglethorp, and her new husband, Alfred. Despite the tranquil surroundings Hastings begins to realise that all is not right. When Mrs Inglethorp is found poisoned, suspicion falls on the family, and another old friend, Hercule Poirot, is invited to investigate.', 1, 2, 2, false),
  ('Romeo and Juliet ', 'https://prodimage.images-bn.com/pimages/9780743477116_p0_v2_s1200x630.jpg', 'Romeo and Juliet is a tragedy written by William Shakespeare early in his career about two young star-crossed lovers whose deaths ultimately reconcile their feuding families. It was among Shakespeare''s most popular plays during his lifetime and, along with Hamlet, is one of his most frequently performed plays.', 9, 4, 1, false),
  ('Hamlet', 'https://i.pinimg.com/originals/a7/c2/5b/a7c25bb8144b76673d27593d8bf8081b.jpg', 'Hamlet is melancholy, bitter, and cynical, full of hatred for his uncle''s scheming and disgust for his mother''s sexuality. A reflective and thoughtful young man who has studied at the University of Wittenberg, Hamlet is often indecisive and hesitant, but at other times prone to rash and impulsive acts.\n', 9, 5, 1, false),
  ('The Brothers Lionheart', 'https://upload.wikimedia.org/wikipedia/en/5/5a/Lionheart_brothers.jpg', 'It is a story of optimism - a story that clearly tells that there is life beyond that, and a very interesting and eventful life at that. To cut a long story short, The Brothers Lionheart is a story about Karl, a 10 year old who is suffering from a terminal illness, and his 13 year old brother Jonatan.', 3, 1, 1, false),
  ('Christmas in Noisy Village', 'https://bethlehembooks.com/sites/default/files/HappyTimesInNoisyVillage.jpg', 'Let the beloved author of Pippi Longstocking take you on an adventure to Noisy Village! The noisy children of three neighboring families are celebrating the season by baking cookies, cutting and decorating trees, eating fruitcake and tarts, and opening Christmas gifts. With illustrations by Ilon Wikland, the master storyteller Astrid Lindgren takes us through Christmas in the Noisy Village!', 3, 1, 2, false),
  ('Murder on the Orient Express', 'https://images-na.ssl-images-amazon.com/images/I/51+2QZIRWfL.jpg', 'Just after midnight, a snowdrift stops the Orient Express in its tracks. The luxurious train is surprisingly full for the time of the year, but by the morning it is one passenger fewer. An American tycoon lies dead in his compartment, stabbed a dozen times, his door locked from the inside. Isolated and with a killer in their midst, detective Hercule Poirot must identify the murderer—in case he or she decides to strike again.', 1, 2, 1, false),
  ('Death on the Nile', 'https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcSDcfTPH66LY9LlgzSTdLfxvrb3beajgWfxJlf_PxMNCtCQZyRm', 'The tranquillity of a cruise along the Nile is shattered by the discovery that Linnet Ridgeway has been shot through the head. She was young, stylish and beautiful, a girl who had everything - until she lost her life. Hercule Poirot recalls an earlier outburst by a fellow passenger: "I would like to put my dear little pistol against her head and just press the trigger." Yet in this exotic setting, nothing is ever quite what it seems...', 1, 2, 1, false),
  ('The A.B.C. Murders', 'https://agathachristie.imgix.net/hcuk-paperback/The-ABC-Murders.JPG?auto=compress,format&fit=clip&q=65&w=400', 'When Alice Asher is murdered in Andover, Hercule Poirot is already on to the clues. Alphabetically speaking, it''s one down, twenty-five to go. There''s a serial killer on the loose. His macabre calling card is to leave the ABC Railway guide beside each victim''s body. But if A is for Alice Asher, bludgeoned to death in Andover; and B is for Betty Bernard, strangled with her belt on the beach at Bexhill; then who will Victim C be?', 1, 2, 1, false),
  ('The Adventures of Huckleberry Finn', 'https://m.media-amazon.com/images/I/51rt34sReUL.jpg', 'A nineteenth-century boy from a Mississippi River town recounts his adventures as he travels down the river with a runaway slave, encountering a family involved in a feud, two scoundrels pretending to be royalty, and Tom Sawyer''s aunt who mistakes him for Tom.', 5, 4, 1, false),
  ('Life on the Mississippi', 'https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1309286211l/99152.jpg', 'A stirring account of America''s vanished past... The book that earned Mark Twain his first recognition as a serious writer... Discover the magic of life on the Mississippi. At once a romantic history of a mighty river, an autobiographical account of Mark Twain''s early steamboat days, and a storehouse of humorous anecdotes and sketches, Life on the Mississippi is the raw material from which Twain wrote his finest novel: Adventures of Huckleberry Finn .\n', 5, 1, 1, false),
  ('The Tragedy of Pudd’nhead Wilson', 'https://kbimages1-a.akamaihd.net/12f0eeab-2378-4703-9f6e-a254be0a7336/353/569/90/False/the-tragedy-of-pudd-nhead-wilson-18.jpg', 'Two half brothers look so similar as infants that no one can tell them apart. One, the legitimate son of a rich man, is destined for a life of comfort, while the other is condemned to be a slave because he is part black. The mother of the would-be slave is also the nurse of the other; to give her son the best life possible, she switches the babies. Soon the boy who is given every advantage becomes spoiled and cruel. He takes sadistic pleasure in tormenting his half brother. As they grow older, the townspeople no longer notice that the boys look similar, and they readily accept that each is born to his station. A local lawyer, David Wilson, has had a similar experience. On his first day in the village, he made an odd remark about a dog, and the townspeople gave him the condescending name "Pudd''nhead." Although he was a young, intelligent lawyer, he is unable to live down this name, so he toils in obscurity for over twenty years. Finally, he is presented with a complex murder trial-a chance to prove himself to the townspeople and shake this unjust label.', 5, 4, 1, false),
  ('The Picture of Dorian Gray', 'https://www.prestwickhouse.com/ProductImages/ebeb01f5-41da-4f9b-a6cc-763c1610ccff/images/202121.jpg', 'Pudd''nhead Wilson (1894) is a novel by American writer Mark Twain. Its central intrigue revolves around two boys—one, born into slavery, with 1/32 black ancestry; the other, white, born to be the master of the house. The two boys, who look similar, are switched at infancy. Each grows into the other''s social role.', 4, 1, 1, false),
  ('It', 'https://images-na.ssl-images-amazon.com/images/I/61QiVXXKB5L.jpg', 'It is a 1986 horror novel by American author Stephen King. ... It was his 22nd book, and his 17th novel written under his own name. The story follows the experiences of seven children as they are terrorized by an evil entity that exploits the fears of its victims to disguise itself while hunting its prey.', 6, 4, 1, false),
  ('The Shining', 'https://i.pinimg.com/originals/e8/e7/04/e8e7046d4033157f27fe610045c60fcb.jpg', 'Jack Torrance''s new job at the Overlook Hotel is the perfect chance for a fresh start. As the off-season caretaker at the atmospheric old hotel, he''ll have plenty of time to spend reconnecting with his family and working on his writing. But as the harsh winter weather sets in, the idyllic location feels ever more remote...and more sinister. And the only one to notice the strange and terrible forces gathering around the Overlook is Danny Torrance, a uniquely gifted five-year-old. ', 6, 4, 1, false),
  ('Doctor Sleep', 'https://m.media-amazon.com/images/I/41biWTWGx+L.jpg', 'Stephen King returns to the characters and territory of one of his most popular novels ever, The Shining, in this instantly riveting novel about the now middle-aged Dan Torrance (the boy protagonist of The Shining) and the very special 12-year-old girl he must save from a tribe of murderous paranormals. On highways across America, a tribe of people called The True Knot travel in search of sustenance. They look harmless - mostly old, lots of polyester, and married to their RVs. But as Dan Torrance knows, and spunky 12-year-old Abra Stone learns, The True Knot are quasi-immortal, living off the "steam" that children with the "shining" produce when they are slowly tortured to death.', 6, 4, 1, false),
  ('The Alchemist', 'https://i.pinimg.com/originals/36/25/ae/3625aeda8106f77d2fdc859a09a25f0b.jpg', 'The Alchemist is the magical story of Santiago, an Andalusian shepherd boy who yearns to travel in search of a worldly treasure as extravagant as any ever found. From his home in Spain he journeys to the markets of Tangiers and across the Egyptian desert to a fateful encounter with the alchemist.', 7, 4, 1, false),
  ('Eleven Minutes', 'https://m.media-amazon.com/images/I/51USI3nvZ2L.jpg', 'Eleven Minutes is the story of Maria, a young girl from a Brazilian village, whose first innocent brushes with love leave her heartbroken. ... Maria''s despairing view of love is put to the test when she meets a handsome young painter', 7, 1, 1, false),
  ('Adultery', 'https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1397692452l/21939580.jpg', 'Adultery, the provocative new novel by Paulo Coelho, best-selling author of The Alchemist and Eleven Minutes, explores the question of what it means to live life fully and happily, finding the balance between life''s routine and the desire for something new.', 7, 4, 1, false),
  ('War and Peace', 'https://images.penguinrandomhouse.com/cover/9781400079988', 'In Russia''s struggle with Napoleon, Tolstoy saw a tragedy that involved all mankind. Greater than a historical chronicle, War and Peace is an affirmation of life itself, "a complete picture", as a contemporary reviewer put it, "of everything in which people find their happiness and greatness, their grief and humiliation". Tolstoy gave his personal approval to this translation, published here in a new single volume edition, which includes an introduction by Henry Gifford, and Tolstoy''s important essay "Some Words about War and Peace".', 8, 4, 1, false),
  ('The Death of Ivan Ilyich', 'https://images-na.ssl-images-amazon.com/images/I/5134GkH6CkL.jpg', 'Hailed as one of the world''s supreme masterpieces on the subject of death and dying, The Death of Ivan Ilyich is the story of a worldly careerist, a high court judge who has never given the inevitability of his death so much as a passing thought.\n', 8, 1, 1, false),
  ('Pig the Pug', 'https://c.booko.info/covers/a78b02febab074cb/v/600.jpeg', 'The story follows a day in the life of Pig (who is a rather selfish pug) and his play time with Trevor – the cutest sausage dog you''ll ever meet. All Trevor wants is one toy, but Pig refuses to share. As a poor Trevor''s luck would have it, Pig goes into a toy-snatching frenzy and makes a rather large pile of toys.', 10, 1, 2, false),
  ('I Need a Hug', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4071/9781407171586.jpg', '"I need a hug. Will you cuddle me, Lou?" "What? With those spikes? Get away from me! Shoo!" All this little porcupine wants is a hug. But with such prickly spikes, will she ever get the cuddle she craves?', 10, 1, 1, false),
  ('The Goose Egg', 'https://images-na.ssl-images-amazon.com/images/I/51A%2B8tZvLXL._SX218_BO1,204,203,200_QL40_ML2_.jpg', 'Henrietta likes her quiet life. A morning swim, a cup of tea--all is serene. But everything changes when she bumps her head and winds up with a goose egg--a REAL goose egg. Henrietta tries to return the baby goose to the nest, but her flock has flown. It''s up to Henrietta to raise her.', 11, 1, 1, false),
  ('Marvellous Magnet Book', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4093/9781409301769.jpg', 'Can you help create Peppa and George''s amazing adventures in this marvellous magnet book? A super-fun and interactive magnet book that lets little piglets use their imagination and storytelling skills again and again! Peppa and her little brother George are having fun doing their favourite things.They''re going to the museum, a birthday party and having a sunny summer holiday, and much more! But, they need you to help them finish their adventures! Choose the right magnets to finish the pictures, so you can be sure Peppa and George have a lovely time.', 12, 1, 1, false),
  ('Practise with Peppa: Wipe-Clean First Numbers', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/7232/9780723292111.jpg', 'Develop and practise first number skills with Peppa Pig and friends in this colourful wipe-clean activity book! Children will learn how to write the numbers 1-10 and begin to link number words with the numerals (one = 1) through a range of fun Peppa-themed activities. Ideal for young readers who are starting school and developing first number skills and pencil control. Children can wipe the page clean and practise again and again. Includes a free pen! For more confident learners, why not also try Practise with Peppa: Wipe-Clean First Counting to practise numbers up to 20?', 12, 1, 1, false),
  ('Practise with Peppa: Wipe-Clean First Letters', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/7232/9780723292081.jpg', 'Develop and practise first letter shapes with Peppa Pig and friends in this colourful wipe-clean activity book! Trace over the lowercase letters from a-z and learn new words through a range of fun Peppa-themed activities. Ideal for young readers who are starting school and learning to write first letter shapes, this book helps children form letters in the correct way with extra guidance for left-handers. Children can wipe the page clean each time and practise again and again. Also includes a free pen.', 12, 1, 1, true),
  ('All Our Shimmering Skies', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4607/9781460753903.jpg', 'The international bestselling author of Boy Swallows Universe, Trent Dalton returns with a glorious novel destined to become another Australian classic. "A glinting, big-hearted miracle of a book" Richard Glover.', 13, 4, 1, false),
  ('How Things Work : See Inside', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/7460/9780746098516.jpg', 'This is an amazing flap book packed with inventions, machines, gadgets and devices, and facts and information about how they work. Over 90 flaps reveal the insides of car engines, toilets, escalators, submarines and microwaves and many, many other machines. It includes internet links to websites with animations, games and experiments.', 14, 6, 1, false),
  ('Untamed', 'https://images-na.ssl-images-amazon.com/images/I/51m7MVU4OWL._SX329_BO1,204,203,200_.jpg', 'For many years, Glennon Doyle denied her own discontent. Then, while speaking at a conference, she looked at a woman across the room and fell instantly in love. Three words flooded her mind: There She Is. At first, Glennon assumed these words came to her from on high. But she soon realized they had come to her from within. This was her own voice—the one she had buried beneath decades of numbing addictions, cultural conditioning, and institutional allegiances. This was the voice of the girl she had been before the world told her who to be. Glennon decided to quit abandoning herself and to instead abandon the world’s expectations of her. She quit being good so she could be free. She quit pleasing and started living.', 15, 4, 1, false),
  ('The Deep End', 'https://images-na.ssl-images-amazon.com/images/I/51x8WaTt84L._SX339_BO1,204,203,200_.jpg', 'But things take an unexpected turn, and they find themselves stranded at an RV park that’s not exactly a summertime paradise. When the skies open up and the water starts to rise, the Heffleys wonder if they can save their vacation—or if they’re already in too deep.', 16, 6, 1, false),
  ('A Promised Land', 'https://images-na.ssl-images-amazon.com/images/I/41L5qgUW2fL._SX327_BO1,204,203,200_.jpg', 'In the stirring, highly anticipated first volume of his presidential memoirs, Barack Obama tells the story of his improbable odyssey from young man searching for his identity to leader of the free world, describing in strikingly personal detail both his political education and the landmark moments of the first term of his historic presidency—a time of dramatic transformation and turmoil.', 17, 7, 1, false),
  ('Night Monkey, Day Monkey', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4052/9781405283342.jpg', 'Night Monkey and Day Monkey don''t think they have much in common. But when they each spend time in the other''s opposite worlds, they learn a lot - and they also learn to be the best of friends. Little hands will love turning the pages of this beautiful board book. The perfect storytelling experience for babies and toddlers.', 18, 1, 1, false),
  ('A Time for Mercy', 'https://images-na.ssl-images-amazon.com/images/I/51Q4FT2AIfL._SX327_BO1,204,203,200_.jpg', 'Clanton, Mississippi. 1990. Jake Brigance finds himself embroiled in a deeply divisive trial when the court appoints him attorney for Drew Gamble, a timid sixteen-year-old boy accused of murdering a local deputy. Many in Clanton want a swift trial and the death penalty, but Brigance digs in and discovers that there is more to the story than meets the eye. Jake’s fierce commitment to saving Drew from the gas chamber puts his career, his financial security, and the safety of his family on the line.', 19, 4, 1, false),
  ('A Little Life', 'https://images-na.ssl-images-amazon.com/images/I/81eBp+Fp-jL.jpg', 'This operatically harrowing American gay melodrama became an unlikely bestseller, and one of the most divisive novels of the century so far. One man’s life is blighted by abuse and its aftermath, but also illuminated by love and friendship. Some readers wept all night, some condemned it as titillating and exploitative, but no one could deny its power.', 20, 4, 1, false),
  ('Darkmans', 'https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1356131015l/1139452.jpg', 'British fiction’s most anarchic author is as prolific as she is playful, but this freewheeling, visionary epic set around the Thames Gateway is her magnum opus. Barker brings her customary linguistic invention and wild humour to a tale about history’s hold on the present, as contemporary Ashford is haunted by the spirit of a medieval jester.', 21, 1, 1, false),
  ('Visitation', 'https://images-na.ssl-images-amazon.com/images/I/61qlVcRC9sL.jpg', 'A grand house by a lake in the east of Germany is both the setting and main character of Erpenbeck’s third novel. The turbulent waves of 20th-century history crash over it as the house is sold by a Jewish family fleeing the Third Reich, requisitioned by the Russian army, reclaimed by exiles returning from Siberia, and sold again.', 22, 4, 1, true),
  ('My Life in Red and White', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4746/9781474618243.jpg', 'One of the most influential figures in world football, Wenger won multiple Premier League titles, a record number of FA Cups, and masterminded Arsenal''s historic "Invincibles" season of 2003-2004 and 49-match unbeaten run. He changed the game in England forever, popularising an attacking approach and changing attitudes towards nutrition, fitness and coaching methods - and towards foreign managers.', 23, 7, 1, false),
  ('Here''s the Story', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/8448/9781844884704.jpg', 'When a young Mary McAleese told a priest that she planned to become a lawyer, the priest dismissed the idea: she knew no one in the law, and she was female. The reality of what she went on to achieve - despite those obstacles, and despite a sectarian attack that forced her family to flee their home - is even more improbable.', 24, 7, 1, false),
  ('Greenlights', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4722/9781472280831.jpg', 'I''ve been in this life for fifty years, been trying to work out its riddle for forty-two, and been keeping diaries of clues to that riddle for the last thirty-five. Notes about successes and failures, joys and sorrows, things that made me marvel, and things that made me laugh out loud. How to be fair. How to have less stress. How to have fun. How to hurt people less. How to get hurt less. How to be a good man. How to have meaning in life. How to be more me.', 25, 7, 1, false),
  ('Becoming', 'https://images-na.ssl-images-amazon.com/images/I/414JfiBCutL._SX327_BO1,204,203,200_.jpg', 'n a life filled with meaning and accomplishment, Michelle Obama has emerged as one of the most iconic and compelling women of our era. As First Lady of the United States of America—the first African American to serve in that role—she helped create the most welcoming and inclusive White House in history, while also establishing herself as a powerful advocate for women and girls in the U.S. and around the world, dramatically changing the ways that families pursue healthier and more active lives, and standing with her husband as he led America through some of its most harrowing moments. Along the way, she showed us a few dance moves, crushed Carpool Karaoke, and raised two down-to-earth daughters under an unforgiving media glare.', 26, 7, 1, false),
  ('Educated: A Memoir', 'https://images-na.ssl-images-amazon.com/images/I/417zjFiUwEL._SX327_BO1,204,203,200_.jpg', 'Born to survivalists in the mountains of Idaho, Tara Westover was seventeen the first time she set foot in a classroom. Her family was so isolated from mainstream society that there was no one to ensure the children received an education, and no one to intervene when one of Tara’s older brothers became violent. When another brother got himself into college, Tara decided to try a new kind of life. Her quest for knowledge transformed her, taking her over oceans and across continents, to Harvard and to Cambridge University. Only then would she wonder if she’d traveled too far, if there was still a way home.', 27, 7, 1, false),
  ('Pandora''s Jar', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/5098/9781509873128.jpg', 'Stories of gods and monsters are the mainstay of epic poetry and Greek tragedy, from Homer to Aeschylus, Sophocles and Euripides, from the Trojan War to Jason and the Argonauts. And still, today, a wealth of novels, plays and films draw their inspiration from stories first told almost three thousand years ago. But modern tellers of Greek myth have usually been men, and have routinely shown little interest in telling women''s stories. And when they do, those women are often painted as monstrous, vengeful or just plain evil. But Pandora - the first woman, who according to legend unloosed chaos upon the world - was not a villain, and even Medea and Phaedra have more nuanced stories than generations of retellings might indicate.', 28, 4, 1, false),
  ('Grit', 'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/7850/9781785040207.jpg', 'In this must-read for anyone seeking to succeed, pioneering psychologist Angela Duckworth takes us on an eye-opening journey to discover the true qualities that lead to outstanding achievement. Winningly personal, insightful and powerful, Grit is a book about what goes through your head when you fall down, and how that - not talent or luck - makes all the difference.', 29, 8, 1, false);

INSERT INTO "users" ("username", "password") 
VALUES
  ('admin', '$2b$10$Ub9V/Qx/jJn7F38HTRtwR.XlPsQ2Cjm.pOnJ15eKmYMJ20GsCyFaK'),
  ('nadya', '$2b$10$QCGTAt8Y5mczHPd8Ng7GSOuh8qdkZY/g56tuMKoX/dy.xivu0GnAW'),
  ('mayya', '$2b$10$MLTcAM/6VxgquI8uksDdpeHMQGhG3lYKWwwI0ehy4t3uGg4XXUQES'),
  ('pesho', '$2b$10$.YGx5e/EdM9JaakfEWULl.EYGVbZnQ/9Z.CfrPXyfWmLpA3F4xqEe'),
  ('gosho', '$2b$10$tjIawuxpcyDgVMU0YffQjuYXfRddezj8aL0dvOej6GZLee9U7mGg6'),
  ('sasho', '$2b$10$8pk.8mJRFC58Efy0rIMvruUlET1udOWbsh9EYtOONysb1blwWK/dK'),
  ('tosho', '$2b$10$av2bp2ieM7WQfluBp8o19OPOzhUxXYkgKfpurLbqkRVGScgUjNzlq'),
  ('ico', '$2b$10$WSglGSlNm7XEszvZ9RpL4um.vLiUMs5t.Ow26JGZANNDKgBtaJfD2'),   
  ('krasi', '$2b$10$Jlojv69t7jYQZHlc.b5/fOT9EzZdCMoe5GNWNQtwC7.Sr8OqKniZy'),
  ('venci', '$2b$10$Q2p3XnWCHe/.f0cUOH71a.oObOpLmRE3nMyV/cdiyZIt6Ye2r75dO');

INSERT INTO "reviews" ("book_id", "user_id", "text", "is_deleted") 
VALUES
  (2, 4, 'Awesome book!', true),
  (10, 3, 'Scary book!', false),
  (22, 2, 'Liked it!', false),
  (1, 1, 'This is my favourite book!', false),
  (2, 1, 'Not bad!!', false),
  (3, 1, 'I love this book!', true),
  (8, 1, 'Like!', false),
  (5, 1, 'Nice..', true),
  (9, 1, 'Cool!', false),
  (10, 1, 'Can''t wait to read it!', false),
  (14, 1, 'Wow!', false),
  (13, 1, 'Nice', false),
  (31, 4, 'Awesome!', false),
  (1, 4, 'Nice book!', false),
  (30, 1, 'Amaizing!', false),
  (32, 1, 'Me gusta!', false),
  (29, 1, 'Nice!', false),
  (31, 1, 'Wooooooooooow!', false),
  (6, 1, 'Awesome!', false),
  (7, 1, 'The best book ever!', false),
  (35, 1, 'Cool!', false),
  (3, 4, 'Nice!', true),
  (4, 1, 'I cried while reading!', false),
  (4, 3, 'When can I borrow it?', false),
  (1, 4, 'I didn''t like it!', true);

INSERT INTO "votes" ("vote") 
VALUES
  ('like'),
  ('dislike');

INSERT INTO "user_voted_reviews" ("user_id", "review_id", "vote_id") 
VALUES
  (3, 1, 1),
  (2, 1, 1),
  (2, 2, 2),
  (4, 1, 1),
  (5, 1, 1),
  (3, 1, 2),
  (1, 14, 2),
  (5, 4, 1),
  (1, 1, 1),
  (1, 2, 1),
  (4, 4, 1),
  (1, 22, 1),
  (1, 25, 2),
  (6, 14, 1),
  (6, 4, 2);

INSERT INTO "history" ("user_id", "book_id", "borrowed", "returned") 
VALUES
  (4, 2, '2020-09-07 18:46:17', '2020-10-07 18:46:40'),
  (4, 13, '2020-10-07 18:48:20', '2020-10-22 22:24:21'),
  (3, 9, '2020-08-27 18:50:23', '2020-09-12 18:50:33'),
  (3, 7, '2020-10-01 18:50:43', '2020-10-07 18:50:49'),
  (3, 23, '2020-09-15 18:51:05', NULL),
  (2, 22, '2020-09-10 18:53:43', '2020-10-07 18:53:53'),
  (4, 4, '2020-10-22 15:58:33', '2020-10-24 13:38:17'),
  (4, 4, '2020-10-24 13:38:31', '2020-10-26 13:47:34'),
  (1, 1, '2020-10-24 14:44:52', '2020-10-24 14:44:57'),
  (1, 8, '2020-10-24 14:55:06', '2020-10-24 14:55:10'),
  (1, 29, '2020-10-24 14:55:54', '2020-10-24 14:55:59'),
  (1, 31, '2020-10-24 14:56:55', '2020-10-24 14:56:59'),
  (1, 5, '2020-10-24 19:16:21', '2020-10-24 19:16:29'),
  (1, 3, '2020-10-24 19:19:32', '2020-10-24 19:19:38'),
  (1, 1, '2020-10-25 20:36:39', '2020-10-25 20:36:56'),
  (1, 5, '2020-10-25 20:37:17', '2020-10-25 20:38:01'),
  (1, 1, '2020-10-25 20:37:45', '2020-10-25 20:37:57'),
  (1, 1, '2020-10-26 09:40:51', '2020-10-26 09:41:01'),
  (1, 4, '2020-10-26 13:53:56', '2020-10-26 13:54:02'),
  (1, 1, '2020-10-26 16:10:09', '2020-10-26 16:10:15'),
  (1, 35, '2020-10-26 16:55:09', '2020-10-26 16:55:50'),
  (4, 3, '2020-10-26 17:21:50', '2020-10-26 17:21:57'),
  (1, 2, '2020-10-26 20:06:57', '2020-10-26 20:07:02'),
  (1, 4, '2020-10-27 13:37:54', '2020-10-27 13:38:01'),
  (1, 27, '2020-10-27 14:19:30', '2020-10-27 14:19:52'),
  (1, 27, '2020-10-27 14:20:04', '2020-10-27 14:21:13'),
  (1, 27, '2020-10-27 14:21:31', '2020-10-27 14:25:09'),
  (1, 1, '2020-10-27 14:24:55', '2020-10-27 14:25:06'),
  (1, 3, '2020-10-28 15:30:56', '2020-10-28 15:31:07'),
  (1, 4, '2020-10-30 09:20:52', '2020-10-30 09:20:58'),
  (1, 1, '2020-10-30 10:02:06', '2020-10-30 10:02:23'),
  (1, 1, '2020-10-30 16:16:32', '2020-10-30 16:16:41'),
  (9, 3, '2020-10-30 17:17:19', NULL),
  (9, 7, '2020-10-30 17:17:31', NULL);

INSERT INTO "book_rates" ("book_id", "user_id", "rate") 
VALUES
  (2, 4, 5),
  (32, 1, 5),
  (1, 1, 4),
  (8, 1, 5),
  (29, 1, 3),
  (31, 1, 3),
  (3, 1, 5),
  (35, 1, 3),
  (3, 4, 3),
  (2, 1, 5),
  (4, 1, 4);