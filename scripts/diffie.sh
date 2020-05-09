# exit script on error
set -e

echo "Creating Diffie-Hellman params..."
openssl dhparam -out dhparam.pem 4096
echo "creating self signed certificate"
openssl req -new -newkey rsa:4096 -x509 -sha512 -days 3650 -nodes -out status.lab.crt -keyout status.lab.key
echo "setting right permissions..."
chmod 400 status.lab.crt
chmod 400 status.lab.key
