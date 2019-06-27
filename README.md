# Password API

> A simple RESTful API that generates passwords, stores and publishes them in
> an encrypted format.

###### Note
This project is part of the hiring process of METRONOM.

## API documentation 
The API is documented using the
[OpenAPI](https://swagger.io/docs/specification/about/) specification standard.
A rendered version of the documentation can be found
[here](http://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/mariuskiessling/password-api/master/specifications/openapi.yaml).

## Running the project
1. Download the project.
2. Modify the `config.json` file if you need to change the port or environment.
3. Run
  ```go run main.go```
4. Visit [http://localhost:8082](http://localhost:8082) in your favourite REST
   API client. Alternatively you can load the example script's functions using
   ```source example/testing.sh```.
   This is only possible in a bash compatible shell. You can now access the
   functions `generate_passwords` and `decrypt_passwords`.
   ```
   # Create a password that is 16 characters long and contains 4 numbers and 6
   # special characters. It will be stored under the tag "example".
   generate_passwords example 16 4 6 

   # Fetch an decrypt the generated password using the private key that is
   # also located inside the example directory.
   generate_passwords example
   ```

## Config
On startup the file `config.json` is loaded and parsed. You can tweak the
following paramters:
- **`port`**: The port the API listens on.
- **`env`**: The environment the application is executed in. Only *development*
  and *production* are valid values. In the production environment, all output
  is suppressed. Inside the development environment, the mutation steps for
  each vowel replacement is shown as well as the generated password both in its
  unencrypted and encrypted form.
