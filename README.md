# SIGN ME

# Endpoints
### Sign Up
```shell
curl --request POST \
  --url http://localhost:8000/v1/auth/sign-up \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '{
    "email": "test@test.com",
    "name": "John Doe",
    "password": "qwerty123",
    "passwordConfirmation": "qwerty123"
}'
```

# Todo's
* [ ] Sign In endpoint
* [ ] Sign out endpoint
* [ ] Create credential keys
* [ ] Upload pdf
* [ ] Sign pdf
* [ ] Verify pdf
* [ ] Share signed pdf
