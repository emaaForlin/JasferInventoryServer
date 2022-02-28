# Jasfer Inventory Software

## How to run it

### Run the server with docker
You can run the app inside a docker container (which is the best way).
1. You just need to modify the `docker-compose.yml.example` at your liking.
2. Remove the `.example` at the end of the file name.
3. Use `docker-compose up -d` or just `docker-compose up`
4. Enjoy it.

### Run the server on local
1. First you need to export the `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS` and `DB_NAME` variables.
2. Then run the `build` recipe with `make build` and the executable will be created in `bin/`
3. After that you just need have to run it `./bin/app`

### Changelog
- Added some CI workflow
- Added Dockerfile
- Error handling in all operations
- Now it creates the table if not exists 
- New graceful shutdown method
- Added reusing deleted IDs
- Use a database

#### All CRUD operations can be performed
- Delete products `DELETE /products/{id}` 
- Update products `PUT /products/{id}`
- Add products `POST /products`
- Retrieve products `GET /products`
#### Docs
- API Documentation in `GET /docs`



 ## Note for my self: Clean the code before merge it into main 