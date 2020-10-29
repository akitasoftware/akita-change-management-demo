Example for Blog Post 6

# Set Up

We have an API that returns information about users. By default, the endpoint
returns JSON and does not include user phone numbers in response.

```shell
$ curl localhost:8080/users/json
[{"id":"usr_295oDMFK8b1yS5dwlSTdgP"},{"id":"usr_6NiejyYEVpWfziUXJgovV6"}]
```

One day, an engineer added an endpoint to return YAML instead. Now phone numbers
are leaked!

```shell
$ curl localhost:8080/users/yaml
- id: usr_295oDMFK8b1yS5dwlSTdgP
  phone: (123) 456-7890
- id: usr_6NiejyYEVpWfziUXJgovV6
  phone: (777) 888-9999

```

# How to Run

The HTTP server listens on port 8080.

```shell
go run .
```
