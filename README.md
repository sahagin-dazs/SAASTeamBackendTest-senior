# SaasTeamBackendTest
SAAS Team takehome backend test

## Required software

Install the following dependencies:

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/) or equivalent make processor

## System Setup

Run `make build` before you begin, to build the docker images. 

Run `make dev` to stand up the service.

If it is working, running 

```bash
curl 'http://0.0.0.0:8000/health'
```
should return **OK** and running 

```bash
curl 'http://0.0.0.0:8000/products'

```
should return a JSON object containing the 4 products in the system.

Run `make test_before` and all tests should pass.

## Tasks

#### 1) Tests
* Fix any bugs that prevent `make test_after` from passing successfully. Note this will cause `make test_before` to fail, this is expected.
* Note that in order to do all of this, the following tasks will need to be completed.

#### 2) ProductType
* Implement a new field on products in the system called ProductType.
* Product 1 and 2 are a "food", Product 3 and 4 are "sporting_good".

#### 3) Hidden Fields
* The Product model contains hidden fields that we don't want to expose through the `GET` endpoints, but want to keep internally.
* Both `product_discount_price` and `coupon_code` should be hidden fields and should _not_ be exposed through the `GET /products` or `GET /products/{product_id}` endpoint.
* Fix the system such that `product_discount_price` and `coupon_code` are still fields that can be set with `POST /products`, but are not returned to the user through the `GET` endpoints.

#### 4) Calculate Price
* Implement the `/calculate-price` API endpoint.
* This API should take a JSON payload containing an array (`cart`) of objects containing a `product_id`, `quantity`, and an optional field `coupon_code`. Eg:
```
{
  "cart": [
    {
      "product_id": "1",
      "quantity": 2,
      "coupon_code": "food50"
    },
    {
      "product_id": "2",
      "quantity": 4      
    },
  ]
}
```
* `/calculate-price` should return a JSON consisting of the `total_objects` and `total_cost` for the entire cart. The response should be in the form:
```
{
	total_objects: int,
	total_cost: int
}
```
* It is up to you to define the request payload and implement it. Add any code at any level to return the correct responses.
* If the cart consists of some valid items and some invalid items (`product_id` does not exist), only tally the valid items.
* If a product in the cart has a valid coupon code, the new price of the product should be the product discount price.
* If the cart is empty or consists of _only_ invalid items, return an `HTTP 200 OK` with the following body:
```
{
	total_objects: 0,
	total_cost: 0
}
```
#### 5) Error Handling
* Currently any and all errors in the system will throw a `HTTP 500 Internal Server Error`, we don't want this. Fix the handler responses with the following rules:
	* `400 Bad Request`
		* If the request payload is malformed.
		* This applies to the `POST /products` and `POST /calculate-price` endpoints.
	* `404 Not Found`
		* If the request's ID is not found.
		* This applies to the `GET /products/{product_id}` endpoint.

#### 6) Add a store
Right now the GET products endpoint returns an hardcoded list of products and the POST to create a new product only returns a payload, but no new product is actually created. Define an interface to store and retrieve products from storage and add at least one implementation backed by any storage method you prefer (database, JSON file, in-memory data structure). It doesn't matter if the stored data won't survive a reboot of the Docker container, as long as products can be added and fetched while the process lives. The tests should be fixed accordingly to take into account the new changes. Everything needs to keep working out of the box with `make dev`, so make sure that new dependencies (if any) are added to the docker-compose file.
# SAASTeamBackendTest-senior
