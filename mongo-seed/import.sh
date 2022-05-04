#! /bin/bash

mongoimport --host database --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection users --type json --file /mongo-seed/users.json --jsonArray
mongoimport --host database --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection stages --type json --file /mongo-seed/stages.json --jsonArray
mongoimport --host database --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection languages --type json --file /mongo-seed/languages.json --jsonArray
mongoimport --host database --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection groups --type json --file /mongo-seed/groups.json --jsonArray
mongoimport --host database --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection baseStrings --type json --file /mongo-seed/baseStrings.json --jsonArray
