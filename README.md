# Jasfer Inventory Software

### How to run it

#### Run the server with docker
Soon available...

#### Run the server on local
First you need to have exported the `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS` and `DB_NAME` variables.
Then run the `build` recipe with `make build` and the executable will be created in `bin/`
After that you just need have to run it `./bin/app`

### Changelog

#### All CRUD operations can be performed
- Delete products `DELETE /products/{id}` 
- Update products `PUT /products/{id}`
- Add products `POST /products`
- Retrieve products `GET /products`

#### Other features
- More error handling
- New graceful shutdown method
- Added reusing deleted IDs
- Use a database

 ## Note for my self: Clean the code before merge it into main 