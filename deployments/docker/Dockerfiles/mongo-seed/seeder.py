#!/usr/bin/env python3

import subprocess
import uuid
import sys
import shlex
import time

#from pymongo import MongoClient


flag = True

print(len(sys.argv))

if len(sys.argv) <= 1:
    print("Require mongoDB Host with argument")
    sys.exit(0)

# client = MongoClient("mongodb://%s:27017" % sys.argv[1])
# db = client.csapi
# db.people.remove({})

print("Reading titanic.csv..")

with open('/tmp/titanic.csv', 'r') as file:
    titanic = open("/tmp/data.csv", "w")
    for line in file.readlines()[:2]:
        line = line.strip()

        if flag:
            line = "uuid,survived,passengerclass,name,sex,age,siblingsorspousesaboard,parentsorchildrenaboard,fare\n"
            flag = False
        else:    
            id = str(uuid.uuid4())
            line = id + "," + line
            values = line.split(",")

            if (values[1] == "0"):
                values[1] = "False"
            else:
                values[1] = "True"

            person = {
                "uuid": values[0],
                "survived": values[1],
                "passengerclass": int(values[2]),
                "name": values[3],
                "sex": values[4],
                "age": int(values[5]),
                "siblingsorspousesaboard": int(values[6]),
                "parentsorchildrenaboard": int(values[7]),
                "fare": float(values[8]),
            }
            print(person)
            #result = db.people.insert_one(person)
            print(values)
            line = ",".join(values) + "\n"
            titanic.write(line)
            # print("Created {0} register".format(result.inserted_id))
    titanic.close()

#print("Populate MongoDB..")

while not flag:
    try:
        cmd = '/usr/bin/mongoimport --type csv -d csapi -c people --columnsHaveTypes  \
                --fields="uuid.string(),survived.boolean(),passengerclass.int32(),name.string(),sex.string(),age.int32(),siblingsorspousesaboard.int32(),parentsorchildrenaboard.int32(),fare.double()" \
                --drop --host %s:27017 /tmp/data.csv' % sys.argv[1]
        args = shlex.split(cmd)
        # print("Executing: " + cmd)
        #subprocess.run([args])
        #process.wait()
        process  = subprocess.check_output(args)
        print("** Finished seed Database: \n", process)
        flag = True
    except:
        print("Error import data to Database and retry in 5 sec ...")
        time.sleep(5)

