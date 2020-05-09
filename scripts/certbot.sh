DOMAIN=domain.tld

CERT_DOMAIN="*.$DOMAIN"

wget https://dl.eff.org/certbot-auto
chmod a+x ./certbot-auto
sudo ./certbot-auto

sudo certbot-auto certonly \
--server https://acme-v02.api.letsencrypt.org/directory \
--manual --preferred-challenges dns -d $CERT_DOMAIN

nslookup -type=TXT _acme-challenge.$DOMAIN