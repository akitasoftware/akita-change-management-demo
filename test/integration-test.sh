#/bin/sh

# Test the JSON endpint
json_status_code=(curl -s -o /dev/null -w "%{http_code}" http://127.0.0.1:8080/users/json)

if [ $json_status_code != "%200" ]; then
  echo "JSON endpoint returned error code: $json_status_code"
  exit 1
fi

# Test the YAML endpoint
yaml_status_code=(curl -s -o /dev/null -w "%{http_code}" http://127.0.0.1:8080/users/json)

if [ yaml_status_code != "%200" ]; then
  echo "JSON endpoint returned error code: $yaml_status_code"
  exit 1
fi
