#! /bin/bash

mongoimport --host database --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection users --type json --file /mongo-seed/users.json --jsonArray