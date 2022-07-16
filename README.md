# Description

The repository is the simplified implementation of star wars movies and comments API. In consists of 
- `app` - golang application;
- `redis` - cache layer to avoid http requests to external system;
- `postgres` - db to persist comments.

## Important to mention

For simplicity as it is a test task:

- No migrations added - db reset each time with orm model.
- No config - only hardcoded values.
- Models and repository implemented coupled, no separation.
- No input validation - simplicity.
- No makefile committed.

## Run

To run the app you need docker-compose and docker installed. Run:
```bash
docker-compose up -d
```

After that the app will be available on port `:80`.

## API 

There are multiple endpoints in the API:

- GET `/api/v1/movies` - list all movies

```bash
curl http://localhost:80/api/v1/movies

[{"episode_id":"4","title":"A New Hope","opening_crawl":"It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....","release_date":"1977-05-25T00:00:00Z","comments_count":0},{"episode_id":"5","title":"The Empire Strikes Back","opening_crawl":"It is a dark time for the\r\nRebellion. Although the Death\r\nStar has been destroyed,\r\nImperial troops have driven the\r\nRebel forces from their hidden\r\nbase and pursued them across\r\nthe galaxy.\r\n\r\nEvading the dreaded Imperial\r\nStarfleet, a group of freedom\r\nfighters led by Luke Skywalker\r\nhas established a new secret\r\nbase on the remote ice world\r\nof Hoth.\r\n\r\nThe evil lord Darth Vader,\r\nobsessed with finding young\r\nSkywalker, has dispatched\r\nthousands of remote probes into\r\nthe far reaches of space....","release_date":"1980-05-17T00:00:00Z","comments_count":0},{"episode_id":"6","title":"Return of the Jedi","opening_crawl":"Luke Skywalker has returned to\r\nhis home planet of Tatooine in\r\nan attempt to rescue his\r\nfriend Han Solo from the\r\nclutches of the vile gangster\r\nJabba the Hutt.\r\n\r\nLittle does Luke know that the\r\nGALACTIC EMPIRE has secretly\r\nbegun construction on a new\r\narmored space station even\r\nmore powerful than the first\r\ndreaded Death Star.\r\n\r\nWhen completed, this ultimate\r\nweapon will spell certain doom\r\nfor the small band of rebels\r\nstruggling to restore freedom\r\nto the galaxy...","release_date":"1983-05-25T00:00:00Z","comments_count":0},{"episode_id":"1","title":"The Phantom Menace","opening_crawl":"Turmoil has engulfed the\r\nGalactic Republic. The taxation\r\nof trade routes to outlying star\r\nsystems is in dispute.\r\n\r\nHoping to resolve the matter\r\nwith a blockade of deadly\r\nbattleships, the greedy Trade\r\nFederation has stopped all\r\nshipping to the small planet\r\nof Naboo.\r\n\r\nWhile the Congress of the\r\nRepublic endlessly debates\r\nthis alarming chain of events,\r\nthe Supreme Chancellor has\r\nsecretly dispatched two Jedi\r\nKnights, the guardians of\r\npeace and justice in the\r\ngalaxy, to settle the conflict....","release_date":"1999-05-19T00:00:00Z","comments_count":0},{"episode_id":"2","title":"Attack of the Clones","opening_crawl":"There is unrest in the Galactic\r\nSenate. Several thousand solar\r\nsystems have declared their\r\nintentions to leave the Republic.\r\n\r\nThis separatist movement,\r\nunder the leadership of the\r\nmysterious Count Dooku, has\r\nmade it difficult for the limited\r\nnumber of Jedi Knights to maintain \r\npeace and order in the galaxy.\r\n\r\nSenator Amidala, the former\r\nQueen of Naboo, is returning\r\nto the Galactic Senate to vote\r\non the critical issue of creating\r\nan ARMY OF THE REPUBLIC\r\nto assist the overwhelmed\r\nJedi....","release_date":"2002-05-16T00:00:00Z","comments_count":0},{"episode_id":"3","title":"Revenge of the Sith","opening_crawl":"War! The Republic is crumbling\r\nunder attacks by the ruthless\r\nSith Lord, Count Dooku.\r\nThere are heroes on both sides.\r\nEvil is everywhere.\r\n\r\nIn a stunning move, the\r\nfiendish droid leader, General\r\nGrievous, has swept into the\r\nRepublic capital and kidnapped\r\nChancellor Palpatine, leader of\r\nthe Galactic Senate.\r\n\r\nAs the Separatist Droid Army\r\nattempts to flee the besieged\r\ncapital with their valuable\r\nhostage, two Jedi Knights lead a\r\ndesperate mission to rescue the\r\ncaptive Chancellor....","release_date":"2005-05-19T00:00:00Z","comments_count":0}] 
```

- GET `/api/v1/movies/:id/comments` - get all comments for movie with id `:id`

```bash
curl http://localhost:80/api/v1/movies/1/comments

[{"id":1,"movie_id":"1","text":"some comment 1","ip":"::1","created_at":"2022-07-16T03:15:19.474797Z"}]
```

- POST `/api/v1/movies/:id/comments` - create new comment for the movie with id `:id`

```bash
curl -X POST http://localhost:80/api/v1/movies/1/comments -d '{"text": "some comment"}'

{}
```

- GET `/api/v1/movies/:id/characters` - get all comments for the movie with id `:id`. Accepts query params: 
    - `sort` - that can be `name`, `gender`, `height`.
    - `order` - `asc` or `desc`. Default `asc`. 
    - `gender` - string to filter by `gender`.

```bash
curl "http://localhost:80/api/v1/movies/1/characters?sort=name&order=desc&gender=male"

{"characters":[{"name":"Wilhuff Tarkin","gender":"male","height":180,"height_feet":"5 10.85"},{"name":"Wedge Antilles","gender":"male","height":170,"height_feet":"5 6.91"},{"name":"Raymus Antilles","gender":"male","height":188,"height_feet":"6 2.00"},{"name":"Owen Lars","gender":"male","height":178,"height_feet":"5 10.06"},{"name":"Obi-Wan Kenobi","gender":"male","height":182,"height_feet":"5 11.64"},{"name":"Luke Skywalker","gender":"male","height":172,"height_feet":"5 7.70"},{"name":"Jek Tono Porkins","gender":"male","height":180,"height_feet":"5 10.85"},{"name":"Han Solo","gender":"male","height":180,"height_feet":"5 10.85"},{"name":"Greedo","gender":"male","height":173,"height_feet":"5 8.09"},{"name":"Darth Vader","gender":"male","height":202,"height_feet":"6 7.51"},{"name":"Chewbacca","gender":"male","height":228,"height_feet":"7 5.74"},{"name":"Biggs Darklighter","gender":"male","height":183,"height_feet":"6 0.03"}],"count":12,"height":2216}
```
