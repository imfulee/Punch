#!/bin/sh

status=$1

if [ -z $status ]; then
    echo "No argument supplied"
    exit 1
fi

if [ $status != "In" ] && [ $status != "Out" ]; then
    echo "Status could only be \"In\" or \"Out\""
    exit 1
fi

if [ -z $USERNAME ] || [ -z $PASSWORD ] || [ -z $COMPANY ]; then
    echo "No fields supplied"
    exit 1
fi

/app/punch $status \
    --username=$USERNAME \
    --password=$PASSWORD \
    --company=$COMPANY
