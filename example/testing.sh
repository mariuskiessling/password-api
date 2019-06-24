# Load this file using `source testing.sh`

generate_passwords() {
  curl --silent -X "POST" "http://localhost:8082/password" \
       -H 'Content-Type: application/json; charset=utf-8' \
       -d $'{
    "alternatives": 0,
    "public_key": "-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7cnq0P7dtzpY3mHGk59N\\nUBkBVDydid3WN0bV4Z4wZMxcjbE3zYPWRFXcU07l5F6Kkp0CeXnXegKxhCdqPEJc\\noeHHcxR0cN7aeyEYf6vQ/r8atbdXrh0ciYTac36zE4h6iJXgd5Gauf55VmpeTRF8\\nn2CY7xqEIk14fNG6DSevl+idkncSiKTgG6KNFQDsTEdTL+edaEkqJDk8px4JdpHL\\n63nDw9+tqQ6H4lkct5Na3+f9zYE9DpjE/KQGM3TDfx3vsUWLlFTWxQkObgPUqY2W\\nXQ6qP4rz7MdIEAqjHuAkcVXK27YJB3f2ZFf00L+s8YCVVVPuK/HOCVLgsmGZSFkt\\nSwIDAQAB\\n-----END PUBLIC KEY-----\\n",
    "public_key_fingerprint": "ddf1d3fb4f581a043dacea4e67eb87f8886190e391861731ce9955e933c49392",
    "options": {
      "special_characters": '"$4"',
      "numbers": '"$3"',
      "length": '"$2"'
    },
    "tag": "'"$1"'"
  }'
}

decrypt_passwords() {
  PUBLIC_KEY_FINGERPRINT="ddf1d3fb4f581a043dacea4e67eb87f8886190e391861731ce9955e933c49392"
  for pw in $(curl "http://localhost:8082/password/$PUBLIC_KEY_FINGERPRINT?tag=$1" | jq -c -r ".[] | .[]"); do
    echo ${pw} | base64 -D | openssl pkeyutl -inkey private-pkcs1.pem -keyform PEM -decrypt -pkeyopt rsa_padding_mode:oaep -pkeyopt rsa_oaep_md:sha256
    echo "\n"
  done
}
