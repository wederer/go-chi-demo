#!/bin/bash

# mindepth 1 to get rid of "." and ".."
# -not -path to ignore hidden directories
# find ... | while read to make sure whitespace does not break loop

retention_date=$(date +%F -d '1 day ago')
echo "Deleting any reports older than $retention_date."

find . -mindepth 1 -maxdepth 1 -not -path '*/.*' -name 'report_*' -type d -print0 | while read -d $'\0' dir
do
       # dir is e.g. ./report_2023-04-06_4627658339
       stripped_dir=${dir#*/report_} # to remove "report_" prefix
       stripped_dir=${stripped_dir%%_*} # to remove "_<run-id>" suffix after date
       dir_date=$(date +%F -d $stripped_dir)
       if [[ $dir_date < $retention_date ]]; then
               echo "$dir is older than $retention_date -> deleting $dir..."
               rm -rf $dir
               echo "$dir was deleted."
       fi
done
