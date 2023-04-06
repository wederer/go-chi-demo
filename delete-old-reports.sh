#!/bin/bash

# mindepth 1 to get rid of "." and ".."
# -not -path to ignore hidden directories
# find ... | while read to make sure whitespace does not break loop
find . -mindepth 1 -maxdepth 1 -not -path '*/.*' -name 'report_*' -type d -print0 | while read -d $'\0' dir
do
       # dir is e.g. ./report_2023-04-06_4627658339
       stripped_dir=${dir#*/report_} # to remove prefix
       stripped_dir=${stripped_dir%%--*} # to remove suffix after date
       dir_date=$(date +%F -d $stripped_dir)
       retention_date=$(date +%F -d '1 day ago')
       echo "Parsing $stripped_dir -> $dir_date."
       echo "Deleting any reports older than $retention_date."
       if [[ $dir_date < $retention_date ]]; then
               echo "$dir is older than $retention_date -> deleting $dir..."
               rm -rf $dir
               echo "$dir was deleted."
       fi
done